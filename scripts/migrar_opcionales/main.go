// Command migrar_opcionales convierte en punteros los campos escalares opcionales
// de los modelos Beego, según la nulabilidad de sus columnas en el script SQL.
//
// Regla para cada campo escalar (los campos rel(fk) no se modifican):
//
//   - Columna nullable en el SQL y campo no listado en ignorar.conf: tipo puntero (*T).
//   - Columna NOT NULL, o campo listado en ignorar.conf: tipo primitivo (T).
//
// Solo se reescribe el tipo del campo. Las etiquetas orm se conservan, y ejecutar
// el comando de nuevo sobre modelos ya migrados no produce cambios.
//
// Entradas (rutas relativas a la raíz del proyecto):
//
//   - database/agora_cotizaciones.sql: columnas y su nulabilidad.
//   - models/*.go: structs, TableName() y etiquetas orm.
//   - scripts/migrar_opcionales/ignorar.conf: patrones tabla.Campo con comodín *.
//
// El contraste entre etiquetas null y el SQL se advierte en el log sin modificar
// el archivo. Un patrón de ignorar.conf que no coincide con ningún campo también
// se advierte, sin detener la ejecución.
//
// Uso (desde la raíz del proyecto):
//
//	go run ./scripts/migrar_opcionales
package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	archivoSQL     = "database/agora_cotizaciones.sql"
	dirModelos     = "models"
	archivoIgnorar = "scripts/migrar_opcionales/ignorar.conf"
)

// tiposEscalares son los tipos de campos escalares que el migrador convierte a puntero.
var tiposEscalares = map[string]bool{
	"string": true, "int": true, "int16": true, "int32": true, "int64": true,
	"float64": true, "bool": true, "time.Time": true,
}

// columnaSQL describe una columna del script SQL.
type columnaSQL struct {
	NotNull bool // true para NOT NULL y PRIMARY KEY
}

// campo describe un campo de un modelo y la posición de su tipo en el archivo fuente.
type campo struct {
	Modelo  string
	Nombre  string
	Tabla   string
	Columna string
	Tipo    string // tipo tal como aparece en el código, por ejemplo "*string"
	Inicio  int    // desplazamiento en bytes del inicio del tipo dentro del archivo
	Nulo    bool   // etiqueta orm null
	FK      bool   // etiqueta orm rel(fk)
}

// archivoModelo agrupa el código fuente de un archivo de modelos y sus campos.
type archivoModelo struct {
	ruta   string
	fuente []byte
	campos []campo
}

// cambio es la edición del tipo de un campo: agregar '*' o quitarlo.
type cambio struct {
	inicio  int
	agregar bool
}

func main() {
	tablas, err := leerSQL(archivoSQL)
	salirSiError(err)
	patrones, err := leerIgnorar(archivoIgnorar)
	salirSiError(err)
	archivos, err := leerModelos(dirModelos)
	salirSiError(err)

	coincidencias := make([]bool, len(patrones))
	for _, a := range archivos {
		var cambios []cambio
		for _, c := range a.campos {
			ignorado := marcarCoincidencias(patrones, coincidencias, c.Tabla, c.Nombre)
			col, existe := tablas[c.Tabla][c.Columna]
			contrastar(c, col, existe)
			if c.FK || !existe {
				continue
			}
			base := strings.TrimPrefix(c.Tipo, "*")
			if !tiposEscalares[base] {
				advertir("%s.%s: tipo %s no soportado por el migrador; se omite", c.Modelo, c.Nombre, c.Tipo)
				continue
			}
			debePuntero := !col.NotNull && !ignorado
			if debePuntero != strings.HasPrefix(c.Tipo, "*") {
				cambios = append(cambios, cambio{inicio: c.Inicio, agregar: debePuntero})
				nuevo := base
				if debePuntero {
					nuevo = "*" + base
				}
				fmt.Printf("%s.%s: %s -> %s\n", c.Modelo, c.Nombre, c.Tipo, nuevo)
			}
		}
		if len(cambios) == 0 {
			continue
		}
		nuevo, err := aplicar(a.fuente, cambios)
		salirSiError(err)
		salirSiError(os.WriteFile(a.ruta, nuevo, 0o644))
	}
	for i, p := range patrones {
		if !coincidencias[i] {
			advertir("el patrón %s de %s no coincide con ningún campo", p, archivoIgnorar)
		}
	}
	fmt.Println("Migración de opcionales completada")
}

// leerSQL extrae, por tabla, las columnas y si son NOT NULL.
func leerSQL(ruta string) (map[string]map[string]columnaSQL, error) {
	b, err := os.ReadFile(ruta)
	if err != nil {
		return nil, err
	}
	reTabla := regexp.MustCompile(`(?s)CREATE TABLE "\w+"\."(\w+)" \((.*?)\n\);`)
	reColumna := regexp.MustCompile(`^\s*"(\w+)" (\w+)(?:\((\d+)(?:,\d+)?\))?(.*?),?$`)
	tablas := map[string]map[string]columnaSQL{}
	for _, m := range reTabla.FindAllStringSubmatch(string(b), -1) {
		columnas := map[string]columnaSQL{}
		for _, linea := range strings.Split(m[2], "\n") {
			if c := reColumna.FindStringSubmatch(linea); c != nil {
				columnas[c[1]] = columnaSQL{NotNull: strings.Contains(c[4], "NOT NULL") || strings.Contains(c[4], "PRIMARY KEY")}
			}
		}
		tablas[m[1]] = columnas
	}
	return tablas, nil
}

// leerIgnorar retorna los patrones tabla.Campo de ignorar.conf. Las líneas vacías
// y los comentarios (#) se omiten. Una línea sin punto se advierte y se omite.
func leerIgnorar(ruta string) ([]string, error) {
	b, err := os.ReadFile(ruta)
	if err != nil {
		return nil, err
	}
	var patrones []string
	for _, linea := range strings.Split(string(b), "\n") {
		linea = strings.TrimSpace(linea)
		if linea == "" || strings.HasPrefix(linea, "#") {
			continue
		}
		if !strings.Contains(linea, ".") {
			advertir("%s: la línea %q no tiene formato tabla.Campo; se omite", ruta, linea)
			continue
		}
		patrones = append(patrones, linea)
	}
	return patrones, nil
}

// leerModelos analiza cada archivo de modelos y retorna sus campos con la posición
// de su tipo. Los archivos sin campos con etiqueta column se retornan igualmente.
func leerModelos(dir string) ([]*archivoModelo, error) {
	rutas, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return nil, err
	}
	var archivos []*archivoModelo
	fset := token.NewFileSet()
	for _, r := range rutas {
		fuente, err := os.ReadFile(r)
		if err != nil {
			return nil, err
		}
		f, err := parser.ParseFile(fset, r, fuente, 0)
		if err != nil {
			return nil, err
		}
		tablas := tablasDeModelos(f)
		a := &archivoModelo{ruta: r, fuente: fuente}
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
				for _, fl := range st.Fields.List {
					if len(fl.Names) == 0 {
						continue
					}
					tag := ""
					if fl.Tag != nil {
						v, _ := strconv.Unquote(fl.Tag.Value)
						tag = reflect.StructTag(v).Get("orm")
					}
					columna := opcionORM(tag, "column")
					if columna == "" {
						continue
					}
					inicio := fset.Position(fl.Type.Pos()).Offset
					fin := fset.Position(fl.Type.End()).Offset
					a.campos = append(a.campos, campo{
						Modelo:  ts.Name.Name,
						Nombre:  fl.Names[0].Name,
						Tabla:   tablas[ts.Name.Name],
						Columna: columna,
						Tipo:    string(fuente[inicio:fin]),
						Inicio:  inicio,
						Nulo:    tieneOpcion(tag, "null"),
						FK:      tieneOpcion(tag, "rel(fk)"),
					})
				}
			}
		}
		archivos = append(archivos, a)
	}
	return archivos, nil
}

// tablasDeModelos asocia cada modelo con la tabla que retorna su método TableName.
func tablasDeModelos(f *ast.File) map[string]string {
	tablas := map[string]string{}
	for _, d := range f.Decls {
		fn, ok := d.(*ast.FuncDecl)
		if !ok || fn.Name.Name != "TableName" || fn.Recv == nil || fn.Body == nil {
			continue
		}
		tipo := fn.Recv.List[0].Type
		if estrella, ok := tipo.(*ast.StarExpr); ok {
			tipo = estrella.X
		}
		ident, ok := tipo.(*ast.Ident)
		if !ok {
			continue
		}
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			if l, ok := n.(*ast.BasicLit); ok && l.Kind == token.STRING {
				tablas[ident.Name], _ = strconv.Unquote(l.Value)
			}
			return true
		})
	}
	return tablas
}

// marcarCoincidencias indica si algún patrón de ignorar.conf coincide con tabla.campo.
// Marca en coincidencias los patrones que casaron con cualquier campo.
func marcarCoincidencias(patrones []string, coincidencias []bool, tabla, campo string) bool {
	ignorado := false
	for i, p := range patrones {
		ok, err := path.Match(p, tabla+"."+campo)
		if err != nil {
			advertir("patrón %s inválido: %v", p, err)
			continue
		}
		if ok {
			coincidencias[i] = true
			ignorado = true
		}
	}
	return ignorado
}

// contrastar advierte cuando la etiqueta null del campo no coincide con la
// nulabilidad de su columna, o cuando la columna no existe en el SQL.
func contrastar(c campo, col columnaSQL, existe bool) {
	if !existe {
		advertir("%s.%s: la columna %s no existe en la tabla %s", c.Modelo, c.Nombre, c.Columna, c.Tabla)
		return
	}
	if c.Nulo == col.NotNull {
		if c.Nulo {
			advertir("%s.%s: tiene etiqueta null pero la columna %s es NOT NULL", c.Modelo, c.Nombre, c.Columna)
		} else {
			advertir("%s.%s: no tiene etiqueta null pero la columna %s es nullable", c.Modelo, c.Nombre, c.Columna)
		}
	}
}

// aplicar retorna el código fuente con los cambios de tipo aplicados y formateado.
// Los cambios se aplican de atrás hacia adelante para que los desplazamientos
// pendientes sigan siendo válidos.
func aplicar(fuente []byte, cambios []cambio) ([]byte, error) {
	sort.Slice(cambios, func(i, j int) bool { return cambios[i].inicio > cambios[j].inicio })
	salida := append([]byte{}, fuente...)
	for _, c := range cambios {
		if c.agregar {
			salida = append(append(append([]byte{}, salida[:c.inicio]...), '*'), salida[c.inicio:]...)
			continue
		}
		if salida[c.inicio] != '*' {
			return nil, fmt.Errorf("se esperaba '*' en el desplazamiento %d", c.inicio)
		}
		salida = append(append([]byte{}, salida[:c.inicio]...), salida[c.inicio+1:]...)
	}
	return format.Source(salida)
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

// tieneOpcion indica si la etiqueta orm contiene la opción indicada como bandera.
func tieneOpcion(tag, opcion string) bool {
	for _, p := range strings.Split(tag, ";") {
		if p == opcion {
			return true
		}
	}
	return false
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
