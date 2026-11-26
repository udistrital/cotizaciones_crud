package main

import (
	"path/filepath"
	"text/template"
)

// Carpetas y archivos de soporte de las pruebas por capa. Todo se emite desde aquí.
const (
	dirUnitControladores = "tests/unitarias/controladores"
	dirUnitModelos       = "tests/unitarias/modelos"
)

// pruebasMain indica qué configuración necesita el TestMain de cada carpeta.
type pruebasMain struct {
	Integracion   bool // conecta la base de datos PostgreSQL de las variables COTIZACIONES_CRUD_PG*
	Controladores bool // configura beego y registra las rutas (requiere el router)
}

// generarCapas escribe el controlador falso, los TestMain de cada carpeta y, por
// entidad con campos escalares opcionales, las pruebas unitarias del controlador y
// del modelo.
func generarCapas(lista []*Entidad) error {
	mainIntegracion := pruebasMain{Integracion: true, Controladores: true}
	if err := escribir(filepath.Join(dirPruebas, "main_gen_test.go"), plantillaMainPruebas, mainIntegracion, importesMain(true)...); err != nil {
		return err
	}
	mainControladores := pruebasMain{Controladores: true}
	if err := escribir(filepath.Join(dirUnitControladores, "main_gen_test.go"), plantillaMainPruebas, mainControladores, importesMain(true)...); err != nil {
		return err
	}
	mainModelos := pruebasMain{}
	if err := escribir(filepath.Join(dirUnitModelos, "main_gen_test.go"), plantillaMainPruebas, mainModelos, importesMain(false)...); err != nil {
		return err
	}
	for _, e := range lista {
		if len(e.EscalaresOpc) == 0 {
			continue
		}
		if err := escribir(filepath.Join(dirUnitControladores, e.Archivo+"_test.go"), plantillaUnitControlador, e); err != nil {
			return err
		}
		if err := escribir(filepath.Join(dirUnitModelos, e.Archivo+"_test.go"), plantillaUnitModelo, e); err != nil {
			return err
		}
	}
	// Pruebas de GetAll por capa y de validación del Id de ruta para todas las entidades
	for _, e := range lista {
		if err := escribir(filepath.Join(dirUnitControladores, e.Archivo+"_orden_test.go"), plantillaOrdenControlador, e); err != nil {
			return err
		}
		if err := escribir(filepath.Join(dirUnitModelos, e.Archivo+"_orden_test.go"), plantillaOrdenModelo, e); err != nil {
			return err
		}
		// Formato de las columnas date en el JSON: solo el día
		if len(e.Fechas) > 0 {
			if err := escribir(filepath.Join(dirUnitControladores, e.Archivo+"_fechas_test.go"), plantillaFechasControlador, e); err != nil {
				return err
			}
		}
		// Validación del Id de ruta: GetOne, PUT y DELETE rechazan un Id no numérico
		if err := escribir(filepath.Join(dirUnitControladores, e.Archivo+"_validacion_test.go"), plantillaValidacionControlador, e); err != nil {
			return err
		}
	}
	return nil
}

// importesMain son los imports del TestMain. Con controladores también se
// registran las rutas y la configuración de beego.
func importesMain(controladores bool) []string {
	imp := []string{`"log"`, `"os"`}
	if controladores {
		imp = append(imp,
			`beego "github.com/beego/beego/v2/server/web"`,
			`"github.com/udistrital/utils_oas/v2/customerror"`,
			`_ "github.com/udistrital/cotizaciones_crud/routers"`,
		)
	}
	return imp
}

// plantillaMainPruebas emite el TestMain de una carpeta de pruebas.
var plantillaMainPruebas = template.Must(template.New("main").Parse(`
// TestMain prepara el entorno de las pruebas de esta carpeta.
func TestMain(m *testing.M) {
{{- if .Integracion}}
	if err := helpers.IniciarBaseDatos(); err != nil {
		log.Fatal("No se pudo conectar a la base de datos de pruebas: ", err)
	}
	if !helpers.BaseDisponible() {
		log.Println("AVISO: variables COTIZACIONES_CRUD_PG* no definidas; las pruebas de integración se omiten")
	}
{{- else}}
	if err := helpers.ConectarFalso(); err != nil {
		log.Fatal("No se pudo registrar la base de datos falsa: ", err)
	}
{{- end}}
{{- if .Controladores}}
	// Equivalen a copyrequestbody = true y autorender = false de conf/app.conf
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	// Mismas respuestas de error (Success/Status/Message/Data) que en main.go
	beego.ErrorController(&customerror.CustomErrorController{})
{{- end}}
	os.Exit(m.Run())
}
`))

// plantillaUnitControlador emite las pruebas unitarias del POST de un controlador.
// Cada escenario aplica a todos los campos opcionales de la entidad.
var plantillaUnitControlador = template.Must(template.New("unitControlador").Parse(`
// Pruebas unitarias del POST de {{.Nombre}}Controller ({{.Ruta}}), sin base de datos.
// Cada escenario aplica a todos los campos escalares opcionales de la entidad.

const ruta{{.Nombre}} = "{{.Ruta}}"

// ejemplo{{.Nombre}} retorna la entidad con todos sus campos de ejemplo.
func ejemplo{{.Nombre}}() *models.{{.Nombre}} {
	return &models.{{.Nombre}}{
	{{- range .Campos}}{{if .Externo}}
		{{.Nombre}}: {{.Externo}},
	{{- end}}{{end}}
	}
}

func cuerpo{{.Nombre}}(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, ejemplo{{.Nombre}}())
}

// TestPost{{.Nombre}}OpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPost{{.Nombre}}OpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpo{{.Nombre}}(t)
	{{- range .EscalaresOpc}}
	delete(cuerpo, "{{.Nombre}}")
	{{- end}}
	helpers.ProbarPostOpcionales(t, ruta{{.Nombre}}, cuerpo, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: nil},
	{{- end}}
	})
}

// TestPost{{.Nombre}}OpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPost{{.Nombre}}OpcionalesNull(t *testing.T) {
	cuerpo := cuerpo{{.Nombre}}(t)
	{{- range .EscalaresOpc}}
	cuerpo["{{.Nombre}}"] = nil
	{{- end}}
	helpers.ProbarPostOpcionales(t, ruta{{.Nombre}}, cuerpo, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: nil},
	{{- end}}
	})
}

// TestPost{{.Nombre}}OpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPost{{.Nombre}}OpcionalesConValor(t *testing.T) {
	ejemplo := ejemplo{{.Nombre}}()
	helpers.ProbarPostOpcionales(t, ruta{{.Nombre}}, cuerpo{{.Nombre}}(t), []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: helpers.Desreferenciar(ejemplo.{{.Nombre}})},
	{{- end}}
	})
}

// TestGet{{.Nombre}}OpcionalesNull: la base de datos devuelve NULL en cada campo opcional.
func TestGet{{.Nombre}}OpcionalesNull(t *testing.T) {
	helpers.ProbarGetOpcionales(t, ruta{{.Nombre}}, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: nil},
	{{- end}}
	})
}

// TestGet{{.Nombre}}OpcionalesConValor: la base de datos devuelve el valor de cada campo opcional.
func TestGet{{.Nombre}}OpcionalesConValor(t *testing.T) {
	ejemplo := ejemplo{{.Nombre}}()
	helpers.ProbarGetOpcionales(t, ruta{{.Nombre}}, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: helpers.Desreferenciar(ejemplo.{{.Nombre}})},
	{{- end}}
	})
}

// TestPut{{.Nombre}}OpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPut{{.Nombre}}OpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpo{{.Nombre}}(t)
	{{- range .EscalaresOpc}}
	delete(cuerpo, "{{.Nombre}}")
	{{- end}}
	helpers.ProbarPutOpcionales(t, ruta{{.Nombre}}, cuerpo, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: nil},
	{{- end}}
	})
}

// TestPut{{.Nombre}}OpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPut{{.Nombre}}OpcionalesNull(t *testing.T) {
	cuerpo := cuerpo{{.Nombre}}(t)
	{{- range .EscalaresOpc}}
	cuerpo["{{.Nombre}}"] = nil
	{{- end}}
	helpers.ProbarPutOpcionales(t, ruta{{.Nombre}}, cuerpo, []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: nil},
	{{- end}}
	})
}

// TestPut{{.Nombre}}OpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPut{{.Nombre}}OpcionalesConValor(t *testing.T) {
	ejemplo := ejemplo{{.Nombre}}()
	helpers.ProbarPutOpcionales(t, ruta{{.Nombre}}, cuerpo{{.Nombre}}(t), []helpers.CampoOpcional{
	{{- range .EscalaresOpc}}
		{Campo: "{{.Nombre}}", Columna: "{{.Columna}}", Valor: helpers.Desreferenciar(ejemplo.{{.Nombre}})},
	{{- end}}
	})
}
`))

// plantillaUnitModelo emite, por campo opcional, las pruebas unitarias de
// models.Add<Modelo> con tres subpruebas: omitido (campo no asignado), null (campo
// asignado a nil) y con valor.
var plantillaUnitModelo = template.Must(template.New("unitModelo").Parse(`
// Pruebas unitarias de models.Add{{.Nombre}} para los campos escalares opcionales, sin base de datos.
{{range $f := .EscalaresOpc}}
// Test{{$.Nombre}}{{$f.Nombre}} verifica el valor de {{$f.Columna}} en el INSERT para
// los tres casos del campo {{$f.Nombre}}.
func Test{{$.Nombre}}{{$f.Nombre}}(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.{{$.Nombre}}{
		{{- range $c := $.Campos}}{{if and $c.Externo (ne $c.Nombre $f.Nombre)}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
		}
		if _, err := models.Add{{$.Nombre}}(m); err != nil {
			t.Fatal("No se pudo agregar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.{{$.Nombre}}{
		{{- range $c := $.Campos}}{{if and $c.Externo (ne $c.Nombre $f.Nombre)}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
			{{$f.Nombre}}: nil,
		}
		if _, err := models.Add{{$.Nombre}}(m); err != nil {
			t.Fatal("No se pudo agregar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.{{$.Nombre}}{
		{{- range $c := $.Campos}}{{if $c.Externo}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
		}
		esperado := helpers.Desreferenciar(m.{{$f.Nombre}})
		if _, err := models.Add{{$.Nombre}}(m); err != nil {
			t.Fatal("No se pudo agregar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", esperado)
	})
}

// Test{{$.Nombre}}{{$f.Nombre}}Get verifica la lectura de {{$f.Columna}}: NULL en la base
// de datos llega como nil y un valor llega igual.
func Test{{$.Nombre}}{{$f.Nombre}}Get(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"{{$f.Columna}}": nil})
		v, err := models.Get{{$.Nombre}}ById(1)
		if err != nil {
			t.Fatal("No se pudo consultar {{$.Nombre}}:", err)
		}
		helpers.VerificarValorLeido(t, "{{$f.Nombre}}", nil, helpers.Desreferenciar(v.{{$f.Nombre}}))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar({{$f.Externo}})
		helpers.DefinirFila(map[string]interface{}{"{{$f.Columna}}": esperado})
		v, err := models.Get{{$.Nombre}}ById(1)
		if err != nil {
			t.Fatal("No se pudo consultar {{$.Nombre}}:", err)
		}
		helpers.VerificarValorLeido(t, "{{$f.Nombre}}", esperado, helpers.Desreferenciar(v.{{$f.Nombre}}))
	})
}

// Test{{$.Nombre}}{{$f.Nombre}}Update verifica el valor de {{$f.Columna}} en el UPDATE
// para los tres casos del campo {{$f.Nombre}}.
func Test{{$.Nombre}}{{$f.Nombre}}Update(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.{{$.Nombre}}{
			Id: 1,
		{{- range $c := $.Campos}}{{if and $c.Externo (ne $c.Nombre $f.Nombre)}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
		}
		if err := models.Update{{$.Nombre}}ById(m); err != nil {
			t.Fatal("No se pudo actualizar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.{{$.Nombre}}{
			Id: 1,
		{{- range $c := $.Campos}}{{if and $c.Externo (ne $c.Nombre $f.Nombre)}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
			{{$f.Nombre}}: nil,
		}
		if err := models.Update{{$.Nombre}}ById(m); err != nil {
			t.Fatal("No se pudo actualizar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.{{$.Nombre}}{
			Id: 1,
		{{- range $c := $.Campos}}{{if $c.Externo}}
			{{$c.Nombre}}: {{$c.Externo}},
		{{- end}}{{end}}
		}
		esperado := helpers.Desreferenciar(m.{{$f.Nombre}})
		if err := models.Update{{$.Nombre}}ById(m); err != nil {
			t.Fatal("No se pudo actualizar {{$.Nombre}}:", err)
		}
		helpers.VerificarColumnaEnviada(t, "{{$f.Columna}}", esperado)
	})
}
{{end}}`))

// plantillaOrdenModelo emite, para cada entidad, las pruebas de la capa de modelo del
// ordenamiento de GetAll: sortby sin order ascendente, order desc, campo inexistente,
// order inválido, y limit/offset/total. Verifican la sentencia SQL capturada y el error
// devuelto, sin base de datos.
var plantillaOrdenModelo = template.Must(template.New("ordenModelo").Parse(`
// Pruebas de la capa de modelo del ordenamiento de GetAll{{.Nombre}}, sin base de datos.

func Test{{.Nombre}}GetAllOrdenSinOrderEsAsc(t *testing.T) {
	helpers.LimpiarSentencia()
	helpers.DefinirFilas(map[string]interface{}{"id": 1}, map[string]interface{}{"id": 2})
	if _, _, err := models.GetAll{{.Nombre}}(nil, nil, []string{"Id"}, nil, 0, 10); err != nil {
		t.Fatal("Se esperaba sortby sin order sin error y se obtuvo:", err)
	}
	helpers.VerificarSQL(t, "ORDER BY T0.\"id\" ASC")
}

func Test{{.Nombre}}GetAllOrdenDesc(t *testing.T) {
	helpers.LimpiarSentencia()
	if _, _, err := models.GetAll{{.Nombre}}(nil, nil, []string{"Id"}, []string{"desc"}, 0, 10); err != nil {
		t.Fatal("Se esperaba order desc sin error y se obtuvo:", err)
	}
	helpers.VerificarSQL(t, "ORDER BY T0.\"id\" DESC")
}

func Test{{.Nombre}}GetAllOrdenCampoInexistente(t *testing.T) {
	helpers.LimpiarSentencia()
	err := helpers.LlamarSinPanico(t, func() error {
		_, _, err := models.GetAll{{.Nombre}}(nil, nil, []string{"campo_falso"}, []string{"asc"}, 0, 10)
		return err
	})
	helpers.VerificarErrorSortInvalido(t, err, "campo_falso")
	helpers.VerificarSinConsulta(t)
}

func Test{{.Nombre}}GetAllOrdenOrderInvalido(t *testing.T) {
	helpers.LimpiarSentencia()
	_, _, err := models.GetAll{{.Nombre}}(nil, nil, []string{"Id"}, []string{"ascending"}, 0, 10)
	helpers.VerificarErrorSortInvalido(t, err, "")
	helpers.VerificarSinConsulta(t)
}

func Test{{.Nombre}}GetAllPaginacionYTotal(t *testing.T) {
	helpers.LimpiarSentencia()
	helpers.DefinirFilas(map[string]interface{}{"id": 1}, map[string]interface{}{"id": 2}, map[string]interface{}{"id": 3})
	_, total, err := models.GetAll{{.Nombre}}(nil, nil, []string{"Id"}, []string{"asc"}, 1, 1)
	if err != nil {
		t.Fatal("Se esperaba limit y offset sin error y se obtuvo:", err)
	}
	helpers.VerificarIgual(t, "Total con limit=1 y offset=1", total, int64(3))
	helpers.VerificarSQL(t, "LIMIT 1 OFFSET 1")
}
`))

// plantillaOrdenControlador emite, para cada entidad, las pruebas de la capa de
// controlador del GetAll: sortby sin order responde 200 ascendente, order desc, campo
// inexistente y order inválido responden 400 sin consultar la base, y limit/offset/total.
var plantillaOrdenControlador = template.Must(template.New("ordenControlador").Parse(`
// Pruebas de la capa de controlador del GetAll de {{.Nombre}}Controller ({{.Ruta}}), sin base de datos.

const rutaOrden{{.Nombre}} = "{{.Ruta}}"

func Test{{.Nombre}}GetAllOrdenSortbySinOrderEsAsc(t *testing.T) {
	helpers.DefinirFilas(map[string]interface{}{"id": 1}, map[string]interface{}{"id": 2}, map[string]interface{}{"id": 3})
	lista, total := helpers.ConsultarOrdenado(t, rutaOrden{{.Nombre}}, "sortby=Id", "ORDER BY T0.\"id\" ASC")
	helpers.VerificarIgual(t, "Registros devueltos", len(lista), 3)
	helpers.VerificarIgual(t, "Total", total, 3)
}

func Test{{.Nombre}}GetAllOrdenOrderDesc(t *testing.T) {
	helpers.DefinirFilas(map[string]interface{}{"id": 3}, map[string]interface{}{"id": 2}, map[string]interface{}{"id": 1})
	helpers.ConsultarOrdenado(t, rutaOrden{{.Nombre}}, "sortby=Id&order=desc", "ORDER BY T0.\"id\" DESC")
}

func Test{{.Nombre}}GetAllOrdenCampoInexistente(t *testing.T) {
	helpers.ProbarSortInvalido(t, rutaOrden{{.Nombre}}, "sortby=campo_falso&order=asc")
	helpers.ProbarSortInvalido(t, rutaOrden{{.Nombre}}, "sortby=campo_falso")
}

func Test{{.Nombre}}GetAllOrdenOrderInvalido(t *testing.T) {
	helpers.ProbarSortInvalido(t, rutaOrden{{.Nombre}}, "sortby=Id&order=ascending")
}

func Test{{.Nombre}}GetAllPaginacionYTotal(t *testing.T) {
	helpers.DefinirFilas(map[string]interface{}{"id": 1}, map[string]interface{}{"id": 2}, map[string]interface{}{"id": 3})
	_, total := helpers.ConsultarOrdenado(t, rutaOrden{{.Nombre}}, "sortby=Id&order=asc&limit=1&offset=1", "LIMIT 1 OFFSET 1")
	helpers.VerificarIgual(t, "Total con limit=1 y offset=1", total, 3)
}
`))

// plantillaValidacionControlador emite, para cada entidad, las pruebas de la capa de
// controlador de la validación de entradas: GetOne, PUT y DELETE con un Id no numérico
// responden 404 sin consultar la base (el Id no se convierte en 0), y POST y PUT sin una
// columna NOT NULL que no es llave foránea (ausente, null o texto vacío) responden 400
// sin consultar la base.
var plantillaValidacionControlador = template.Must(template.New("validacionControlador").Parse(`
// Pruebas de la capa de controlador de la validación de entradas de
// {{.Nombre}}Controller ({{.Ruta}}), sin base de datos.

const rutaValidacion{{.Nombre}} = "{{.Ruta}}"

// cuerpoValidacion{{.Nombre}} retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacion{{.Nombre}}(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.{{.Nombre}}{
	{{- range .Campos}}{{if .Externo}}
		{{.Nombre}}: {{.Externo}},
	{{- end}}{{end}}
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func Test{{.Nombre}}GetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacion{{.Nombre}}, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func Test{{.Nombre}}PutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacion{{.Nombre}}, helpers.JSON(t, cuerpoValidacion{{.Nombre}}(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func Test{{.Nombre}}DeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacion{{.Nombre}}, "")
}
{{- range .EscalaresReq}}

// La columna {{.Columna}} es NOT NULL: POST sin {{.Nombre}}, con null{{if eq .Tipo "string"}} o vacío{{end}}
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func Test{{$.Nombre}}PostSin{{.Nombre}}(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "{{.Nombre}}") },
		"null":    func(c map[string]interface{}) { c["{{.Nombre}}"] = nil },
	{{- if eq .Tipo "string"}}
		"vacio":    func(c map[string]interface{}) { c["{{.Nombre}}"] = "" },
		"espacios": func(c map[string]interface{}) { c["{{.Nombre}}"] = "   " },
	{{- end}}
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacion{{$.Nombre}}(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacion{{$.Nombre}}, cuerpo)
		})
	}
}

// La columna {{.Columna}} es NOT NULL: PUT sin {{.Nombre}}, con null{{if eq .Tipo "string"}} o vacío{{end}}
// responde 400 sin actualizar.
func Test{{$.Nombre}}PutSin{{.Nombre}}(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "{{.Nombre}}") },
		"null":    func(c map[string]interface{}) { c["{{.Nombre}}"] = nil },
	{{- if eq .Tipo "string"}}
		"vacio":    func(c map[string]interface{}) { c["{{.Nombre}}"] = "" },
		"espacios": func(c map[string]interface{}) { c["{{.Nombre}}"] = "   " },
	{{- end}}
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacion{{$.Nombre}}(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacion{{$.Nombre}}, cuerpo)
		})
	}
}
{{- end}}
{{- if .CerosReq}}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func Test{{.Nombre}}PostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacion{{.Nombre}}(t)
	{{- range .EscalaresReq}}
	{{- if or (eq .Tipo "int") (eq .Tipo "int16") (eq .Tipo "float64")}}
	cuerpo["{{.Nombre}}"] = 0
	{{- else if eq .Tipo "bool"}}
	cuerpo["{{.Nombre}}"] = false
	{{- end}}
	{{- end}}
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacion{{.Nombre}}+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
{{- end}}
`))

// plantillaFechasControlador emite, para cada entidad con columnas date (models.Fecha), las
// pruebas de la capa de controlador de su formato: al leer salen solo con el día guardado
// y al escribir se guarda el día enviado (sin correrse por la zona horaria); un formato
// inválido responde 400.
var plantillaFechasControlador = template.Must(template.New("fechasControlador").Parse(`
// Pruebas de la capa de controlador del formato de las columnas date de
// {{.Nombre}}Controller ({{.Ruta}}), sin base de datos.

// La base devuelve las columnas date como las entrega lib/pq (medianoche marcada como
// UTC): el JSON trae solo el día guardado.
func Test{{.Nombre}}GetFechasSoloDia(t *testing.T) {
	helpers.ProbarGetFechas(t, rutaValidacion{{.Nombre}}, map[string]interface{}{
	{{- range .Fechas}}
		"{{.Columna}}": helpers.FechaGuardada,
	{{- end}}
	}, map[string]string{
	{{- range .Fechas}}
		"{{.Nombre}}": helpers.FechaJSON,
	{{- end}}
	})
}
{{- range .Fechas}}

// {{.Nombre}} es una columna date: POST acepta solo el día o una fecha y hora RFC 3339 (de la
// que toma el día escrito), guarda ese día sin correrlo por la zona horaria, lo devuelve
// solo con el día y rechaza otros formatos.
func Test{{$.Nombre}}Post{{.Nombre}}SoloDia(t *testing.T) {
	for enviado, guardado := range map[string]string{
		"2026-01-15":                "2026-01-15",
		"2026-01-15T00:00:00Z":      "2026-01-15",
		"2026-01-15T23:30:00-05:00": "2026-01-15",
	} {
		t.Run(enviado, func(t *testing.T) {
			cuerpo := cuerpoValidacion{{$.Nombre}}(t)
			cuerpo["{{.Nombre}}"] = enviado
			helpers.LimpiarSentencia()
			r := helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacion{{$.Nombre}}+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
			helpers.VerificarColumnaEnviada(t, "{{.Columna}}", guardado)
			helpers.VerificarIgual(t, "Campo {{.Nombre}} en la respuesta", r.Dato(t)["{{.Nombre}}"], guardado)
		})
	}
	for _, invalido := range []string{"15/01/2026", "2026-13-45", "texto"} {
		t.Run("invalido "+invalido, func(t *testing.T) {
			cuerpo := cuerpoValidacion{{$.Nombre}}(t)
			cuerpo["{{.Nombre}}"] = invalido
			helpers.ProbarPostRechazado(t, rutaValidacion{{$.Nombre}}, cuerpo)
		})
	}
}
{{- end}}
`))
