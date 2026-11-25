// Command generar_pruebas genera las pruebas unitarias de los controladores a
// partir del código fuente del proyecto. Procesa como texto:
//
//   - models/*.go: structs, tipos de campo y etiquetas orm (pk, auto, rel(fk), null, column, type)
//   - routers/router.go: ruta de cada controlador
//   - controllers/*.go: archivo de cada controlador (para nombrar su _test.go)
//   - database/agora_cotizaciones.sql: columnas NOT NULL, longitudes y llaves foráneas
//
// y genera:
//
//   - tests/helpers/registros_gen.go: registros de ejemplo válidos por modelo
//   - tests/controllers/<controlador>_test.go: pruebas de cada controlador
//
// Uso (desde la raíz del proyecto):
//
//	go run ./scripts/generar_pruebas
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

const (
	archivoSQL      = "database/agora_cotizaciones.sql"
	archivoRouter   = "routers/router.go"
	dirModelos      = "models"
	dirControllers  = "controllers"
	dirPruebas      = "tests/controllers"
	archivoRegistro = "tests/helpers/registros_gen.go"
	uuidEjemplo     = "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b"
	cabecera        = "// Código generado por scripts/generar_pruebas. NO EDITAR.\n// Para regenerar: go run ./scripts/generar_pruebas\n\n"
)

// Campos administrados por el servicio: no se incluyen en los datos de ejemplo.
var camposServicio = map[string]bool{"Id": true, "Activo": true, "FechaCreacion": true, "FechaModificacion": true}

// columnaSQL describe una columna de database/agora_cotizaciones.sql.
type columnaSQL struct {
	Tipo     string
	NotNull  bool
	Longitud int
}

// tablaSQL describe una tabla de database/agora_cotizaciones.sql.
type tablaSQL struct {
	Columnas map[string]columnaSQL
	FKs      map[string]string // columna -> tabla referenciada
}

// Campo describe un campo de un modelo.
type Campo struct {
	Nombre    string
	Tipo      string // string, int, float64, bool, time.Time o el modelo relacionado
	Columna   string
	Relacion  string // modelo relacionado cuando es rel(fk)
	Nulo      bool   // etiqueta orm null
	UUID      bool   // etiqueta orm type(uuid)
	Ejemplo   string // literal Go del valor de ejemplo
	Requerido bool   // llave foránea obligatoria
}

// Entidad describe un modelo, su tabla y su controlador.
type Entidad struct {
	Nombre        string
	Tabla         string
	Ruta          string
	Archivo       string // nombre del archivo del controlador sin .go
	Campos        []Campo
	IdCompartido  string // modelo cuyo Id comparte (relación 1:1), si aplica
	CampoPut      Campo  // campo que se modifica en la prueba de PUT
	ValorPut      string // literal Go del nuevo valor
	Relaciones    []Campo
	RelacionesReq []Campo
	RelacionesOpc []Campo
}

func main() {
	tablas, err := leerSQL(archivoSQL)
	salirSiError(err)
	rutas, err := leerRutas(archivoRouter)
	salirSiError(err)
	archivos, err := leerControllers(dirControllers)
	salirSiError(err)
	entidades, err := leerModelos(dirModelos)
	salirSiError(err)

	porTabla := map[string]*Entidad{}
	for _, e := range entidades {
		porTabla[e.Tabla] = e
	}
	var lista []*Entidad
	for _, e := range entidades {
		e.Ruta, e.Archivo = rutas[e.Nombre], archivos[e.Nombre]
		if e.Ruta == "" || e.Archivo == "" {
			advertir("el modelo %s no tiene controlador o ruta; se omite", e.Nombre)
			continue
		}
		completar(e, tablas[e.Tabla], porTabla)
		lista = append(lista, e)
	}
	sort.Slice(lista, func(i, j int) bool { return lista[i].Nombre < lista[j].Nombre })

	salirSiError(escribir(archivoRegistro, plantillaRegistros, lista))
	for _, e := range lista {
		salirSiError(escribir(filepath.Join(dirPruebas, e.Archivo+"_test.go"), plantillaPruebas, e))
	}
	fmt.Printf("Pruebas generadas para %d controladores en %s\n", len(lista), dirPruebas)
}

// leerSQL extrae tablas, columnas y llaves foráneas del script SQL.
func leerSQL(ruta string) (map[string]*tablaSQL, error) {
	b, err := os.ReadFile(ruta)
	if err != nil {
		return nil, err
	}
	texto := string(b)
	tablas := map[string]*tablaSQL{}
	reTabla := regexp.MustCompile(`(?s)CREATE TABLE "\w+"\."(\w+)" \((.*?)\n\);`)
	reColumna := regexp.MustCompile(`^\s*"(\w+)" (\w+)(?:\((\d+)(?:,\d+)?\))?(.*?),?$`)
	for _, m := range reTabla.FindAllStringSubmatch(texto, -1) {
		t := &tablaSQL{Columnas: map[string]columnaSQL{}, FKs: map[string]string{}}
		for _, linea := range strings.Split(m[2], "\n") {
			if c := reColumna.FindStringSubmatch(linea); c != nil {
				longitud, _ := strconv.Atoi(c[3])
				t.Columnas[c[1]] = columnaSQL{Tipo: c[2], Longitud: longitud,
					NotNull: strings.Contains(c[4], "NOT NULL") || strings.Contains(c[4], "PRIMARY KEY")}
			}
		}
		tablas[m[1]] = t
	}
	reFK := regexp.MustCompile(`ALTER TABLE "\w+"\."(\w+)" ADD CONSTRAINT "\w+" FOREIGN KEY \("(\w+)"\) REFERENCES "\w+"\."(\w+)"`)
	for _, m := range reFK.FindAllStringSubmatch(texto, -1) {
		if t, ok := tablas[m[1]]; ok {
			t.FKs[m[2]] = m[3]
		}
	}
	return tablas, nil
}

// leerRutas asocia cada controlador (por nombre de modelo) con su ruta.
func leerRutas(ruta string) (map[string]string, error) {
	b, err := os.ReadFile(ruta)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`NSNamespace\("(/\w+)",\s*beego\.NSInclude\(\s*&controllers\.(\w+)Controller\{\}`)
	rutas := map[string]string{}
	for _, m := range re.FindAllStringSubmatch(string(b), -1) {
		rutas[m[2]] = "/v1" + m[1]
	}
	return rutas, nil
}

// leerControllers asocia cada controlador (por nombre de modelo) con su archivo.
func leerControllers(dir string) (map[string]string, error) {
	rutas, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`type (\w+)Controller struct`)
	archivos := map[string]string{}
	for _, r := range rutas {
		b, err := os.ReadFile(r)
		if err != nil {
			return nil, err
		}
		for _, m := range re.FindAllStringSubmatch(string(b), -1) {
			archivos[m[1]] = strings.TrimSuffix(filepath.Base(r), ".go")
		}
	}
	return archivos, nil
}

// leerModelos analiza los structs de los modelos y sus etiquetas orm.
func leerModelos(dir string) ([]*Entidad, error) {
	rutas, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return nil, err
	}
	var entidades []*Entidad
	fset := token.NewFileSet()
	for _, r := range rutas {
		f, err := parser.ParseFile(fset, r, nil, 0)
		if err != nil {
			return nil, err
		}
		tablas := map[string]string{}
		ast.Inspect(f, func(n ast.Node) bool {
			if fn, ok := n.(*ast.FuncDecl); ok && fn.Name.Name == "TableName" && fn.Recv != nil {
				receptor := strings.TrimPrefix(tipoTexto(fn.Recv.List[0].Type), "*")
				ast.Inspect(fn.Body, func(n ast.Node) bool {
					if l, ok := n.(*ast.BasicLit); ok && l.Kind == token.STRING {
						tablas[receptor], _ = strconv.Unquote(l.Value)
					}
					return true
				})
			}
			return true
		})
		for _, d := range f.Decls {
			gd, ok := d.(*ast.GenDecl)
			if !ok || gd.Tok != token.TYPE {
				continue
			}
			for _, s := range gd.Specs {
				ts := s.(*ast.TypeSpec)
				st, ok := ts.Type.(*ast.StructType)
				if !ok {
					continue
				}
				e := &Entidad{Nombre: ts.Name.Name, Tabla: tablas[ts.Name.Name]}
				for _, fl := range st.Fields.List {
					tag := ""
					if fl.Tag != nil {
						v, _ := strconv.Unquote(fl.Tag.Value)
						tag = reflect.StructTag(v).Get("orm")
					}
					c := Campo{Nombre: fl.Names[0].Name, Tipo: tipoTexto(fl.Type), Columna: opcionORM(tag, "column"),
						Nulo: tieneOpcion(tag, "null"), UUID: opcionORM(tag, "type") == "uuid"}
					if tieneOpcion(tag, "rel(fk)") {
						c.Relacion = strings.TrimPrefix(c.Tipo, "*")
					}
					e.Campos = append(e.Campos, c)
				}
				entidades = append(entidades, e)
			}
		}
	}
	return entidades, nil
}

// completar cruza el modelo con el SQL, define los valores de ejemplo y el
// campo a modificar en la prueba de PUT.
func completar(e *Entidad, t *tablaSQL, porTabla map[string]*Entidad) {
	if t == nil {
		advertir("la tabla %s del modelo %s no está en %s", e.Tabla, e.Nombre, archivoSQL)
		t = &tablaSQL{Columnas: map[string]columnaSQL{}, FKs: map[string]string{}}
	}
	if padre, ok := t.FKs["id"]; ok && porTabla[padre] != nil {
		e.IdCompartido = porTabla[padre].Nombre
	}
	for i := range e.Campos {
		c := &e.Campos[i]
		col, existe := t.Columnas[c.Columna]
		if !existe {
			advertir("%s.%s: la columna %s no existe en la tabla %s", e.Nombre, c.Nombre, c.Columna, e.Tabla)
		} else if c.Relacion != "" && c.Nulo == col.NotNull {
			advertir("%s.%s: el modelo indica null=%v pero en el SQL NOT NULL=%v", e.Nombre, c.Nombre, c.Nulo, col.NotNull)
		}
		switch {
		case camposServicio[c.Nombre]:
		case c.Relacion != "":
			c.Ejemplo = fmt.Sprintf("&models.%s{Id: Crear%s(t).Id}", c.Relacion, c.Relacion)
			c.Requerido = !c.Nulo
			e.Relaciones = append(e.Relaciones, *c)
			if c.Requerido {
				e.RelacionesReq = append(e.RelacionesReq, *c)
			} else {
				e.RelacionesOpc = append(e.RelacionesOpc, *c)
			}
		case c.Tipo == "string" && c.UUID:
			c.Ejemplo = strconv.Quote(uuidEjemplo)
		case c.Tipo == "string" && (col.Tipo == "json" || col.Tipo == "jsonb"):
			c.Ejemplo = strconv.Quote(`{"prueba": true}`)
		case c.Tipo == "string":
			c.Ejemplo = strconv.Quote(recortar("Prueba", col.Longitud))
		case c.Tipo == "int":
			c.Ejemplo = "1"
		case c.Tipo == "int16":
			c.Ejemplo = "2026"
		case c.Tipo == "time.Time":
			c.Ejemplo = "time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)"
		case c.Tipo == "float64":
			c.Ejemplo = "1.5"
		case c.Tipo == "bool":
			c.Ejemplo = "true"
		default:
			advertir("%s.%s: tipo %s sin valor de ejemplo", e.Nombre, c.Nombre, c.Tipo)
		}
	}
	// Campo para la prueba de PUT: primer texto editable o, si no hay, primer entero
	// (usuario_creacion_id no se modifica en PUT y usuario_modificacion_id es el último recurso)
	for _, tipo := range []string{"string", "int"} {
		for _, c := range e.Campos {
			if e.CampoPut.Nombre == "" && c.Tipo == tipo && c.Ejemplo != "" && c.Relacion == "" && !c.UUID && c.Nombre != "UsuarioModificacionId" && c.Nombre != "UsuarioCreacionId" {
				e.CampoPut = c
				if tipo == "string" {
					e.ValorPut = strconv.Quote(recortar("Prueba-2", t.Columnas[c.Columna].Longitud))
				} else {
					e.ValorPut = "2"
				}
			}
		}
	}
	// Tablas con solo llaves foráneas y auditoría: se modifica usuario_modificacion_id
	for _, c := range e.Campos {
		if e.CampoPut.Nombre == "" && c.Nombre == "UsuarioModificacionId" && c.Tipo == "int" {
			e.CampoPut = c
			e.ValorPut = "2"
		}
	}
	if e.CampoPut.Nombre == "" {
		advertir("%s: no hay un campo de texto o entero para la prueba de PUT", e.Nombre)
	}
}

func recortar(s string, longitud int) string {
	if longitud > 0 && len(s) > longitud {
		return s[:longitud]
	}
	return s
}

// opcionORM retorna el valor de opcion(valor) en una etiqueta orm.
func opcionORM(tag, opcion string) string {
	for _, p := range strings.Split(tag, ";") {
		if strings.HasPrefix(p, opcion+"(") && strings.HasSuffix(p, ")") {
			return p[len(opcion)+1 : len(p)-1]
		}
	}
	return ""
}

func tieneOpcion(tag, opcion string) bool {
	for _, p := range strings.Split(tag, ";") {
		if p == opcion {
			return true
		}
	}
	return false
}

func tipoTexto(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + tipoTexto(t.X)
	case *ast.SelectorExpr:
		return tipoTexto(t.X) + "." + t.Sel.Name
	}
	return fmt.Sprintf("%T", e)
}

// escribir ejecuta la plantilla, agrega los imports usados y formatea el código.
func escribir(ruta string, plantilla *template.Template, datos interface{}) error {
	var cuerpo bytes.Buffer
	if err := plantilla.Execute(&cuerpo, datos); err != nil {
		return err
	}
	paquete := "controllers"
	if strings.Contains(ruta, "helpers") {
		paquete = "helpers"
	}
	imports := []string{}
	for _, imp := range [][2]string{
		{"fmt.", `"fmt"`},
		{"http.", `"net/http"`},
		{"testing.", `"testing"`},
		{"time.", `"time"`},
		{"models.", `"github.com/udistrital/cotizaciones_crud/models"`},
		{"helpers.", `"github.com/udistrital/cotizaciones_crud/tests/helpers"`},
	} {
		if strings.Contains(cuerpo.String(), imp[0]) {
			imports = append(imports, imp[1])
		}
	}
	// Imports de la biblioteca estándar primero y luego los del proyecto
	var estandar, proyecto []string
	for _, imp := range imports {
		if strings.Contains(imp, ".") {
			proyecto = append(proyecto, imp)
		} else {
			estandar = append(estandar, imp)
		}
	}
	codigo := cabecera + "package " + paquete + "\n\nimport (\n" + strings.Join(estandar, "\n") + "\n\n" +
		strings.Join(proyecto, "\n") + "\n)\n" + cuerpo.String()
	formateado, err := format.Source([]byte(codigo))
	if err != nil {
		return fmt.Errorf("%s: %v\n%s", ruta, err, codigo)
	}
	if err := os.MkdirAll(filepath.Dir(ruta), 0o755); err != nil {
		return err
	}
	return os.WriteFile(ruta, formateado, 0o644)
}

func advertir(formato string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "ADVERTENCIA: "+formato+"\n", args...)
}

func salirSiError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}

var plantillaRegistros = template.Must(template.New("registros").Parse(`
// Las funciones Nuevo<Modelo> retornan un registro de ejemplo válido sin
// guardar (creando en la base de datos sus llaves foráneas). Las funciones
// Crear<Modelo> lo guardan directamente con el modelo y programan su
// eliminación al terminar la prueba.
{{range .}}

func Nuevo{{.Nombre}}(t *testing.T) *models.{{.Nombre}} {
	t.Helper()
	return &models.{{.Nombre}}{
		{{- if .IdCompartido}}
		Id: Crear{{.IdCompartido}}(t).Id, // relación 1:1: comparte el Id de {{.IdCompartido}}
		{{- end}}
		{{- range .Campos}}{{if .Ejemplo}}
		{{.Nombre}}: {{.Ejemplo}},
		{{- end}}{{end}}
	}
}

func Crear{{.Nombre}}(t *testing.T) *models.{{.Nombre}} {
	t.Helper()
	v := Nuevo{{.Nombre}}(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.Add{{.Nombre}}(v); err != nil {
		t.Fatal("No se pudo crear {{.Nombre}}:", err)
	}
	EliminarAlFinal(t, "{{.Tabla}}", v.Id)
	return v
}
{{end}}`))

var plantillaPruebas = template.Must(template.New("pruebas").Parse(`
// Pruebas del controlador {{.Nombre}}Controller ({{.Ruta}}).

const (
	ruta{{.Nombre}}  = "{{.Ruta}}"
	tabla{{.Nombre}} = "{{.Tabla}}"
)

// cuerpoPost{{.Nombre}} retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente
{{- if not .IdCompartido}} y un Id (lo asigna la base de datos){{end}}.
func cuerpoPost{{.Nombre}}(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.Nuevo{{.Nombre}}(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	{{- if not .IdCompartido}}
	cuerpo["Id"] = helpers.IdInexistente
	{{- end}}
	return cuerpo
}

func Test{{.Nombre}}Post(t *testing.T) {
	helpers.RequerirBaseDatos(t)
{{- if .IdCompartido}}

	// El Id es obligatorio: es el Id de {{.IdCompartido}} (relación 1:1)
	sinId := cuerpoPost{{.Nombre}}(t)
	delete(sinId, "Id")
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, ruta{{.Nombre}}+"/", helpers.JSON(t, sinId)), http.StatusBadRequest)
	idInexistente := cuerpoPost{{.Nombre}}(t)
	idInexistente["Id"] = helpers.IdInexistente
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, ruta{{.Nombre}}+"/", helpers.JSON(t, idInexistente)), http.StatusBadRequest)

	cuerpo := cuerpoPost{{.Nombre}}(t)
	id := helpers.CrearPorPost(t, ruta{{.Nombre}}, tabla{{.Nombre}}, cuerpo)
	helpers.VerificarIgual(t, "Id igual al de {{.IdCompartido}}", float64(id), cuerpo["Id"])
{{- else}}

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, ruta{{.Nombre}}, tabla{{.Nombre}}, cuerpoPost{{.Nombre}}(t))
	id2 := helpers.CrearPorPost(t, ruta{{.Nombre}}, tabla{{.Nombre}}, cuerpoPost{{.Nombre}}(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}
{{- end}}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tabla{{.Nombre}}, id)
{{- range .RelacionesReq}}

	// La llave foránea {{.Nombre}} es obligatoria
	sin{{.Nombre}} := cuerpoPost{{$.Nombre}}(t)
	sin{{.Nombre}}["{{.Nombre}}"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, ruta{{$.Nombre}}+"/", helpers.JSON(t, sin{{.Nombre}})), http.StatusBadRequest)
{{- end}}

	helpers.ProbarPostInvalido(t, ruta{{.Nombre}})
	t.Log("Test{{.Nombre}}Post Finalizado Correctamente (OK)")
}

func Test{{.Nombre}}GetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.Crear{{.Nombre}}(t)
	{{if .Relaciones}}dato := {{end}}helpers.ProbarGetOne(t, ruta{{.Nombre}}, creado.Id)
{{- range .Relaciones}}

	// La llave foránea {{.Nombre}} viene expandida
	if relacionado, err := models.Get{{.Relacion}}ById(creado.{{.Nombre}}.Id); err != nil {
		t.Fatal("No se pudo consultar {{.Relacion}}:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "{{.Nombre}}", relacionado)
	}
{{- end}}

	helpers.ProbarGetOneInexistente(t, ruta{{.Nombre}})
	t.Log("Test{{.Nombre}}GetOne Finalizado Correctamente (OK)")
}

func Test{{.Nombre}}GetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.Crear{{.Nombre}}(t)
	{{if .Relaciones}}dato := {{end}}helpers.ProbarGetAll(t, ruta{{.Nombre}}, tabla{{.Nombre}}, creado.Id)
{{- range .Relaciones}}

	// La llave foránea {{.Nombre}} viene expandida y se puede filtrar por ella
	if relacionado, err := models.Get{{.Relacion}}ById(creado.{{.Nombre}}.Id); err != nil {
		t.Fatal("No se pudo consultar {{.Relacion}}:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "{{.Nombre}}", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, ruta{{$.Nombre}}, fmt.Sprintf("query={{.Nombre}}.Id:%d", creado.{{.Nombre}}.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por {{.Nombre}}.Id y se obtuvo", l)
	}
{{- end}}
	t.Log("Test{{.Nombre}}GetAll Finalizado Correctamente (OK)")
}

func Test{{.Nombre}}Put(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.Crear{{.Nombre}}(t)
	helpers.ProbarPut(t, ruta{{.Nombre}}, tabla{{.Nombre}}, creado.Id, "{{.CampoPut.Nombre}}", {{.ValorPut}}, "{{.CampoPut.Columna}}")
	t.Log("Test{{.Nombre}}Put Finalizado Correctamente (OK)")
}

func Test{{.Nombre}}Delete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.Crear{{.Nombre}}(t)
	helpers.ProbarDelete(t, ruta{{.Nombre}}, tabla{{.Nombre}}, creado.Id)
	t.Log("Test{{.Nombre}}Delete Finalizado Correctamente (OK)")
}
{{- range .RelacionesOpc}}

// La llave foránea {{.Nombre}} es opcional (null): el registro sin ella se
// puede crear, consultar, actualizar y borrar.
func Test{{$.Nombre}}{{.Nombre}}Opcional(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPost{{$.Nombre}}(t)
	cuerpo["{{.Nombre}}"] = nil
	id := helpers.CrearPorPost(t, ruta{{$.Nombre}}, tabla{{$.Nombre}}, cuerpo)

	if dato := helpers.ProbarGetOne(t, ruta{{$.Nombre}}, id); dato["{{.Nombre}}"] != nil {
		t.Error("Se esperaba {{.Nombre}} nulo y se obtuvo", dato["{{.Nombre}}"])
	}
	if l := helpers.ConsultarGetAll(t, ruta{{$.Nombre}}, fmt.Sprintf("query=Id:%d", id)); !helpers.ContieneId(l, id) {
		t.Error("Se esperaba el registro sin {{.Nombre}} en GetAll y se obtuvo", l)
	}
	helpers.ProbarPut(t, ruta{{$.Nombre}}, tabla{{$.Nombre}}, id, "{{$.CampoPut.Nombre}}", {{$.ValorPut}}, "{{$.CampoPut.Columna}}")
	helpers.ProbarDelete(t, ruta{{$.Nombre}}, tabla{{$.Nombre}}, id)
	t.Log("Test{{$.Nombre}}{{.Nombre}}Opcional Finalizado Correctamente (OK)")
}
{{- end}}
`))
