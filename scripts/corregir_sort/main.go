// Command corregir_sort corrige el ordenamiento (sortby/order) de los CRUD generados con
// beego. Lo aplica a cualquier repositorio que siga la forma de bee: modelos con una
// función GetAll<Modelo> y controladores con un método GetAll. No depende de nombres del
// repositorio donde se ejecuta.
//
// Cambios sobre models/*.go, en cada GetAll<Modelo> con el bloque de ordenamiento
// reconocido:
//
//   - sortby sin order se ordena ascendente (asc por omisión).
//   - Cada campo de sortby se valida contra el modelo antes de consultar la base de
//     datos. Un campo inexistente, un order distinto de asc o desc, o cantidades
//     incompatibles devuelven un error que envuelve ErrSortInvalido.
//
// Cambios sobre controllers/*.go, en el método GetAll:
//
//   - Un error que envuelve models.ErrSortInvalido responde 400. Los demás errores
//     conservan su respuesta (404).
//
// El script crea models/orden.go con ErrSortInvalido y construirOrden si no existe. Si
// ya existe, no lo sobrescribe. Un GetAll o un controlador que no tiene la forma
// esperada se advierte y se omite sin modificarlo. Ejecutarlo de nuevo sobre código ya
// corregido no produce cambios.
//
// Uso (desde la raíz del proyecto beego):
//
//	go run ./scripts/corregir_sort            # o: go run ./scripts/corregir_sort -raiz ruta
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	dirModelos          = "models"
	dirControladores    = "controllers"
	archivoOrden        = "orden.go"
	mensajeSortInvalido = "Error service GetAll: The request contains an invalid sortby or order parameter"
)

// edicion reemplaza el texto entre los desplazamientos inicio y fin del archivo fuente.
type edicion struct {
	inicio, fin int
	texto       string
}

// transformacion corrige el código fuente de un archivo. Retorna el código nuevo, que
// puede ser igual al original si el archivo no requiere cambios.
type transformacion func(ruta string, fuente []byte, advertir func(string, ...interface{})) ([]byte, error)

func main() {
	raiz := flag.String("raiz", ".", "directorio raíz del proyecto beego (contiene models/ y controllers/)")
	flag.Parse()
	if err := corregirRaiz(*raiz, advertir); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
	fmt.Println("Corrección de sortby/order completada")
}

func advertir(formato string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "ADVERTENCIA: "+formato+"\n", args...)
}

// corregirRaiz aplica la corrección a los modelos y controladores del directorio raíz.
// Las advertencias se reportan con advertir y no detienen la ejecución.
func corregirRaiz(raiz string, advertir func(string, ...interface{})) error {
	dirM := filepath.Join(raiz, dirModelos)
	listo, err := asegurarOrden(dirM, advertir)
	if err != nil {
		return err
	}
	if !listo {
		advertir("%s no está listo para la corrección; se omiten modelos y controladores", dirM)
		return nil
	}
	modelos, err := archivosGo(dirM)
	if err != nil {
		return err
	}
	for _, r := range modelos {
		if err := corregirArchivo(r, corregirModelo, advertir); err != nil {
			return err
		}
	}
	controladores, err := archivosGo(filepath.Join(raiz, dirControladores))
	if err != nil {
		return err
	}
	for _, r := range controladores {
		if err := corregirArchivo(r, corregirControlador, advertir); err != nil {
			return err
		}
	}
	return nil
}

// corregirArchivo lee ruta, aplica la transformación y escribe el resultado solo si cambió.
func corregirArchivo(ruta string, t transformacion, advertir func(string, ...interface{})) error {
	fuente, err := os.ReadFile(ruta)
	if err != nil {
		return err
	}
	nuevo, err := t(ruta, fuente, advertir)
	if err != nil {
		return fmt.Errorf("%s: %v", ruta, err)
	}
	if bytes.Equal(nuevo, fuente) {
		return nil
	}
	return os.WriteFile(ruta, nuevo, 0o644)
}

// archivosGo retorna los archivos de código del directorio, sin pruebas ni orden.go.
func archivosGo(dir string) ([]string, error) {
	rutas, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return nil, err
	}
	var salida []string
	for _, r := range rutas {
		base := filepath.Base(r)
		if strings.HasSuffix(base, "_test.go") || base == archivoOrden {
			continue
		}
		salida = append(salida, r)
	}
	return salida, nil
}

// asegurarOrden crea models/orden.go si no existe. Retorna false si el paquete no puede
// usar la corrección: orden.go existe sin construirOrden, o ErrSortInvalido ya está
// declarado en otro archivo.
func asegurarOrden(dir string, advertir func(string, ...interface{})) (bool, error) {
	ruta := filepath.Join(dir, archivoOrden)
	if b, err := os.ReadFile(ruta); err == nil {
		if !bytes.Contains(b, []byte("func construirOrden(")) {
			advertir("%s existe y no define construirOrden; no se sobrescribe", ruta)
			return false, nil
		}
		return true, nil
	}
	rutas, err := archivosGo(dir)
	if err != nil {
		return false, err
	}
	declaracion := regexp.MustCompile(`var ErrSortInvalido\b|ErrSortInvalido\s*=`)
	pkg := ""
	for _, r := range rutas {
		b, err := os.ReadFile(r)
		if err != nil {
			return false, err
		}
		if declaracion.Match(b) {
			advertir("ErrSortInvalido ya está declarado en %s; no se crea %s", r, ruta)
			return false, nil
		}
		if pkg == "" {
			f, err := parser.ParseFile(token.NewFileSet(), r, b, parser.PackageClauseOnly)
			if err != nil {
				return false, err
			}
			pkg = f.Name.Name
		}
	}
	if pkg == "" {
		return false, fmt.Errorf("no hay archivos de modelo en %s para determinar el paquete", dir)
	}
	codigo := strings.Replace(plantillaOrden, "package models\n", "package "+pkg+"\n", 1)
	return true, os.WriteFile(ruta, []byte(codigo), 0o644)
}

// corregirModelo reemplaza el bloque de ordenamiento de cada GetAll<Modelo> reconocido.
func corregirModelo(ruta string, fuente []byte, advertir func(string, ...interface{})) ([]byte, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, ruta, fuente, parser.ParseComments)
	if err != nil {
		advertir("%s no se puede analizar: %v; se omite", ruta, err)
		return fuente, nil
	}
	tipos := map[string]bool{}
	for _, d := range f.Decls {
		gd, ok := d.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, s := range gd.Specs {
			if ts, ok := s.(*ast.TypeSpec); ok {
				if _, esStruct := ts.Type.(*ast.StructType); esStruct {
					tipos[ts.Name.Name] = true
				}
			}
		}
	}
	var ediciones []edicion
	for _, d := range f.Decls {
		fn, ok := d.(*ast.FuncDecl)
		if !ok || fn.Recv != nil || fn.Body == nil || !strings.HasPrefix(fn.Name.Name, "GetAll") {
			continue
		}
		modelo := strings.TrimPrefix(fn.Name.Name, "GetAll")
		if modelo == "" || llamaA(fn, "construirOrden") {
			continue
		}
		e, motivo := edicionOrden(fset, fuente, fn, modelo, tipos)
		if motivo != "" {
			advertir("%s: %s %s; se omite", ruta, fn.Name.Name, motivo)
			continue
		}
		ediciones = append(ediciones, e)
	}
	if len(ediciones) == 0 {
		return fuente, nil
	}
	antes := nombresUsados(fuente)
	nuevo := aplicarEdiciones(fuente, ediciones)
	nuevo = asegurarImport(nuevo, "reflect")
	nuevo, err = quitarImportsSinUso(nuevo, antes)
	if err != nil {
		return nil, err
	}
	return format.Source(nuevo)
}

// edicionOrden calcula el reemplazo del bloque var sortFields + if err del GetAll. Si el
// GetAll no tiene esa forma, retorna el motivo.
func edicionOrden(fset *token.FileSet, fuente []byte, fn *ast.FuncDecl, modelo string, tipos map[string]bool) (edicion, string) {
	if !tipos[modelo] {
		return edicion{}, "no tiene el tipo " + modelo
	}
	if contarResultados(fn) != 3 {
		return edicion{}, "no devuelve tres valores"
	}
	stmts := fn.Body.List
	iDecl := -1
	for i, s := range stmts {
		if declaraVariable(s, "sortFields") {
			iDecl = i
			break
		}
	}
	if iDecl < 0 {
		return edicion{}, "no tiene el bloque de ordenamiento reconocido (var sortFields)"
	}
	iCount := -1
	for i := iDecl + 1; i < len(stmts); i++ {
		if llamaA(stmts[i], "Count") {
			iCount = i
			break
		}
	}
	if iCount < 0 {
		return edicion{}, "no tiene la consulta Count después del ordenamiento"
	}
	intermedio := stmts[iDecl+1 : iCount]
	ifOrden, ok := singularIf(intermedio)
	if !ok {
		return edicion{}, "el bloque de ordenamiento no tiene la forma esperada"
	}
	inicio := inicioConComentario(fuente, fset.Position(stmts[iDecl].Pos()).Offset)
	fin := fset.Position(ifOrden.End()).Offset
	return edicion{inicio: inicio, fin: fin, texto: bloqueOrden(modelo)}, ""
}

// singularIf retorna el único statement de la lista si es un if.
func singularIf(stmts []ast.Stmt) (*ast.IfStmt, bool) {
	if len(stmts) != 1 {
		return nil, false
	}
	ifStmt, ok := stmts[0].(*ast.IfStmt)
	return ifStmt, ok
}

// inicioConComentario retorna el inicio de la línea de pos, o el de la línea anterior si
// esta es el comentario "// order by" que acompaña al bloque.
func inicioConComentario(fuente []byte, pos int) int {
	inicio := bytes.LastIndexByte(fuente[:pos], '\n') + 1
	if inicio == 0 {
		return inicio
	}
	inicioAnterior := bytes.LastIndexByte(fuente[:inicio-1], '\n') + 1
	anterior := strings.TrimSpace(string(fuente[inicioAnterior : inicio-1]))
	if strings.HasPrefix(anterior, "// order by") {
		return inicioAnterior
	}
	return inicio
}

// bloqueOrden es el código que reemplaza el bloque de ordenamiento en el modelo.
func bloqueOrden(modelo string) string {
	return "\t// order by: sortby sin order se ordena ascendente; los campos inválidos y el order\n" +
		"\t// incorrecto se rechazan con ErrSortInvalido antes de consultar la base de datos\n" +
		"\tsortFields, err := construirOrden(reflect.TypeOf(" + modelo + "{}), sortby, order)\n" +
		"\tif err != nil {\n" +
		"\t\treturn nil, 0, err\n" +
		"\t}"
}

// corregirControlador agrega el caso 400 al manejo de error del método GetAll.
func corregirControlador(ruta string, fuente []byte, advertir func(string, ...interface{})) ([]byte, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, ruta, fuente, parser.ParseComments)
	if err != nil {
		advertir("%s no se puede analizar: %v; se omite", ruta, err)
		return fuente, nil
	}
	for _, d := range f.Decls {
		fn, ok := d.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || fn.Name.Name != "GetAll" || fn.Body == nil {
			continue
		}
		if len(fn.Recv.List) == 0 || len(fn.Recv.List[0].Names) == 0 {
			continue
		}
		if bytes.Contains(fuente, []byte("errors.Is(")) && bytes.Contains(fuente, []byte("ErrSortInvalido)")) {
			return fuente, nil
		}
		alias, ok := aliasDeModelos(fn)
		if !ok {
			advertir("%s: GetAll no llama a un GetAll de modelos; se omite", ruta)
			return fuente, nil
		}
		receptor := fn.Recv.List[0].Names[0].Name
		ifErr, abortIdx := ifDeError404(fn.Body.List)
		if ifErr == nil {
			advertir("%s: GetAll no tiene el manejo de error 404 reconocido; se omite", ruta)
			return fuente, nil
		}
		s := abortIdx
		if s > 0 {
			if _, ok := ifErr.Body.List[s-1].(*ast.AssignStmt); ok {
				s--
			}
		}
		inicio := fset.Position(ifErr.Body.List[s].Pos()).Offset
		fin := fset.Position(ifErr.Body.List[abortIdx].End()).Offset
		prefijo := fmt.Sprintf("if errors.Is(err, %s.ErrSortInvalido) {\n%s.Data[\"message\"] = %s\n%s.Abort(\"400\")\n} else {\n",
			alias, receptor, strconv.Quote(mensajeSortInvalido), receptor)
		e := edicion{inicio: inicio, fin: fin, texto: prefijo + string(fuente[inicio:fin]) + "\n}"}
		nuevo := aplicarEdiciones(fuente, []edicion{e})
		nuevo = asegurarImport(nuevo, "errors")
		return format.Source(nuevo)
	}
	return fuente, nil
}

// ifDeError404 retorna el if "err != nil" de nivel superior cuyo cuerpo termina en
// Abort("404"), y el índice de la sentencia Abort dentro de su cuerpo.
func ifDeError404(stmts []ast.Stmt) (*ast.IfStmt, int) {
	for _, s := range stmts {
		ifStmt, ok := s.(*ast.IfStmt)
		if !ok || !esErrNoNulo(ifStmt.Cond) {
			continue
		}
		for i, c := range ifStmt.Body.List {
			if llamaAbort(c, "404") {
				return ifStmt, i
			}
		}
	}
	return nil, -1
}

// esErrNoNulo indica si la condición es err != nil.
func esErrNoNulo(e ast.Expr) bool {
	b, ok := e.(*ast.BinaryExpr)
	if !ok || b.Op != token.NEQ {
		return false
	}
	x, okX := b.X.(*ast.Ident)
	y, okY := b.Y.(*ast.Ident)
	return okX && okY && x.Name == "err" && y.Name == "nil"
}

// llamaAbort indica si la sentencia es receptor.Abort(codigo).
func llamaAbort(s ast.Stmt, codigo string) bool {
	es, ok := s.(*ast.ExprStmt)
	if !ok {
		return false
	}
	call, ok := es.X.(*ast.CallExpr)
	if !ok || len(call.Args) != 1 {
		return false
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Abort" {
		return false
	}
	lit, ok := call.Args[0].(*ast.BasicLit)
	return ok && lit.Value == strconv.Quote(codigo)
}

// aliasDeModelos retorna el nombre con que el controlador importa el paquete de modelos,
// tomado de la llamada a GetAll<Modelo> del método.
func aliasDeModelos(fn *ast.FuncDecl) (string, bool) {
	alias := ""
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok || alias != "" {
			return alias == ""
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok || !strings.HasPrefix(sel.Sel.Name, "GetAll") {
			return true
		}
		if id, ok := sel.X.(*ast.Ident); ok {
			alias = id.Name
		}
		return true
	})
	return alias, alias != ""
}

// llamaA indica si el nodo contiene una llamada a una función o método llamado nombre.
func llamaA(n ast.Node, nombre string) bool {
	encontrado := false
	ast.Inspect(n, func(x ast.Node) bool {
		if encontrado {
			return false
		}
		if call, ok := x.(*ast.CallExpr); ok {
			switch fun := call.Fun.(type) {
			case *ast.Ident:
				encontrado = fun.Name == nombre
			case *ast.SelectorExpr:
				encontrado = fun.Sel.Name == nombre
			}
		}
		return !encontrado
	})
	return encontrado
}

// declaraVariable indica si la sentencia es una declaración var con la variable indicada.
func declaraVariable(s ast.Stmt, nombre string) bool {
	ds, ok := s.(*ast.DeclStmt)
	if !ok {
		return false
	}
	gd, ok := ds.Decl.(*ast.GenDecl)
	if !ok || gd.Tok != token.VAR {
		return false
	}
	for _, spec := range gd.Specs {
		if vs, ok := spec.(*ast.ValueSpec); ok {
			for _, id := range vs.Names {
				if id.Name == nombre {
					return true
				}
			}
		}
	}
	return false
}

// contarResultados retorna la cantidad de valores que devuelve la función.
func contarResultados(fn *ast.FuncDecl) int {
	n := 0
	if fn.Type.Results == nil {
		return 0
	}
	for _, r := range fn.Type.Results.List {
		if len(r.Names) == 0 {
			n++
		} else {
			n += len(r.Names)
		}
	}
	return n
}

// aplicarEdiciones aplica las ediciones de atrás hacia adelante para que los
// desplazamientos pendientes sigan siendo válidos.
func aplicarEdiciones(fuente []byte, ediciones []edicion) []byte {
	sort.Slice(ediciones, func(i, j int) bool { return ediciones[i].inicio > ediciones[j].inicio })
	salida := append([]byte{}, fuente...)
	for _, e := range ediciones {
		nueva := append([]byte{}, salida[:e.inicio]...)
		nueva = append(nueva, e.texto...)
		salida = append(nueva, salida[e.fin:]...)
	}
	return salida
}

var regexImportAbierto = regexp.MustCompile(`(?m)^import \(\n`)
var regexImportUnico = regexp.MustCompile(`(?m)^import ("[^"]+")\n`)
var regexPaquete = regexp.MustCompile(`(?m)^package \w+\n`)

// asegurarImport agrega el import de la ruta si el archivo no lo tiene. Formatear después
// ordena el bloque de imports.
func asegurarImport(fuente []byte, ruta string) []byte {
	if bytes.Contains(fuente, []byte("\""+ruta+"\"")) {
		return fuente
	}
	if loc := regexImportAbierto.FindIndex(fuente); loc != nil {
		return aplicarEdiciones(fuente, []edicion{{loc[1], loc[1], "\t\"" + ruta + "\"\n"}})
	}
	if loc := regexImportUnico.FindSubmatchIndex(fuente); loc != nil {
		unico := string(fuente[loc[2]:loc[3]])
		return aplicarEdiciones(fuente, []edicion{{loc[0], loc[1], "import (\n\t\"" + ruta + "\"\n\t" + unico + "\n)\n"}})
	}
	if loc := regexPaquete.FindIndex(fuente); loc != nil {
		return aplicarEdiciones(fuente, []edicion{{loc[1], loc[1], "\nimport \"" + ruta + "\"\n"}})
	}
	return fuente
}

// nombresUsados retorna los nombres de paquete que el código referencia (p. ej. errors
// en errors.New). Si el código no analiza, retorna un conjunto vacío.
func nombresUsados(fuente []byte) map[string]bool {
	usados := map[string]bool{}
	f, err := parser.ParseFile(token.NewFileSet(), "", fuente, 0)
	if err != nil {
		return usados
	}
	ast.Inspect(f, func(n ast.Node) bool {
		if sel, ok := n.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok {
				usados[id.Name] = true
			}
		}
		return true
	})
	return usados
}

// quitarImportsSinUso elimina los imports que el código usaba antes de la corrección y que
// ya no usa. Los imports que nunca estuvieron en uso no se tocan.
func quitarImportsSinUso(fuente []byte, usadosAntes map[string]bool) ([]byte, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", fuente, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	usadosAhora := nombresUsados(fuente)
	var ediciones []edicion
	for _, imp := range f.Imports {
		ruta, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			return nil, err
		}
		nombre := filepath.Base(ruta)
		if imp.Name != nil {
			nombre = imp.Name.Name
		}
		if nombre == "_" || nombre == "." || !usadosAntes[nombre] || usadosAhora[nombre] {
			continue
		}
		inicio := bytes.LastIndexByte(fuente[:fset.Position(imp.Pos()).Offset], '\n') + 1
		fin := fset.Position(imp.End()).Offset
		if fin < len(fuente) && fuente[fin] == '\n' {
			fin++
		}
		ediciones = append(ediciones, edicion{inicio: inicio, fin: fin})
	}
	return aplicarEdiciones(fuente, ediciones), nil
}

// plantillaOrden es models/orden.go: el centinela ErrSortInvalido y la validación de
// sortby/order. corregir_sort lo escribe con el nombre del paquete del repositorio.
const plantillaOrden = `package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// ErrSortInvalido indica que los parámetros sortby u order de una consulta GetAll son
// inválidos: un campo que no existe en el modelo, un order distinto de asc o desc, o
// cantidades incompatibles. Los controladores lo traducen a HTTP 400.
var ErrSortInvalido = errors.New("parámetro de ordenamiento inválido")

// construirOrden retorna las expresiones ORDER BY de Beego para los campos de sortby.
// Si order está vacío, cada campo se ordena ascendente. Si order tiene un solo valor,
// aplica a todos los campos de sortby. Sin sortby no se ordena. Todo error de
// validación envuelve ErrSortInvalido.
func construirOrden(modelo reflect.Type, sortby, order []string) ([]string, error) {
	if len(sortby) == 0 {
		if len(order) != 0 {
			return nil, fmt.Errorf("%w: se envió order sin sortby", ErrSortInvalido)
		}
		return nil, nil
	}
	if len(order) == 0 {
		order = []string{"asc"}
	}
	if len(order) != 1 && len(order) != len(sortby) {
		return nil, fmt.Errorf("%w: sortby y order deben tener la misma cantidad, o order un solo valor", ErrSortInvalido)
	}
	var expresiones []string
	for i, campo := range sortby {
		if err := validarCampoOrden(modelo, campo); err != nil {
			return nil, err
		}
		sentido := order[0]
		if len(order) > 1 {
			sentido = order[i]
		}
		switch sentido {
		case "asc":
			expresiones = append(expresiones, campo)
		case "desc":
			expresiones = append(expresiones, "-"+campo)
		default:
			return nil, fmt.Errorf("%w: order %q debe ser asc o desc", ErrSortInvalido, sentido)
		}
	}
	return expresiones, nil
}

// validarCampoOrden verifica que campo sea un campo del modelo, por nombre de campo o de
// columna. Un campo de relación se recorre con "__" (por ejemplo MstCotizacionId__Id) y
// cada segmento se valida contra el modelo relacionado.
func validarCampoOrden(modelo reflect.Type, campo string) error {
	if modelo.Kind() != reflect.Struct {
		return fmt.Errorf("%w: %q no se puede ordenar", ErrSortInvalido, campo)
	}
	nombre, resto, anidado := strings.Cut(campo, "__")
	f, ok := buscarCampoOrden(modelo, nombre)
	if !ok {
		return fmt.Errorf("%w: %q no es un campo de %s", ErrSortInvalido, campo, modelo.Name())
	}
	if !anidado {
		return nil
	}
	tipo := f.Type
	for tipo.Kind() == reflect.Ptr {
		tipo = tipo.Elem()
	}
	if tipo.Kind() != reflect.Struct {
		return fmt.Errorf("%w: %q no es una relación", ErrSortInvalido, nombre)
	}
	return validarCampoOrden(tipo, resto)
}

// buscarCampoOrden retorna el campo del struct cuyo nombre o columna orm es nombre.
func buscarCampoOrden(modelo reflect.Type, nombre string) (reflect.StructField, bool) {
	if nombre == "" {
		return reflect.StructField{}, false
	}
	for i := 0; i < modelo.NumField(); i++ {
		f := modelo.Field(i)
		if f.Name == nombre || columnaORM(f.Tag.Get("orm")) == nombre {
			return f, true
		}
	}
	return reflect.StructField{}, false
}

// columnaORM retorna el valor de column(...) en una etiqueta orm, o "" si no lo tiene.
func columnaORM(etiqueta string) string {
	for _, p := range strings.Split(etiqueta, ";") {
		if strings.HasPrefix(p, "column(") && strings.HasSuffix(p, ")") {
			return p[len("column(") : len(p)-1]
		}
	}
	return ""
}
`
