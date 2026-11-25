package main

// Helpers de prueba. Las fuentes de tests/helpers se emiten desde este archivo:
// los archivos escritos a mano hasta ahora (basedatos, http y verificaciones) se
// copian tal cual, y los generados (registros, punteros y controlador falso) se
// construyen con las plantillas de abajo.

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"text/template"
)

// Rutas de salida de los helpers.
const (
	archivoRegistro       = "tests/helpers/registros_gen.go"
	archivoPunteros       = "tests/helpers/punteros_gen.go"
	archivoFalso          = "tests/helpers/falso_gen.go"
	archivoBaseDatos      = "tests/helpers/basedatos.go"
	archivoHTTP           = "tests/helpers/http.go"
	archivoVerificaciones = "tests/helpers/verificaciones.go"
)

// generarHelpers escribe todos los archivos de tests/helpers a partir de sus fuentes.
func generarHelpers(lista []*Entidad) error {
	if err := escribirCrudo(archivoBaseDatos, fuenteBaseDatos); err != nil {
		return err
	}
	if err := escribirCrudo(archivoHTTP, fuenteHTTP); err != nil {
		return err
	}
	if err := escribirCrudo(archivoVerificaciones, fuenteVerificaciones); err != nil {
		return err
	}
	if err := escribir(archivoPunteros, plantillaPunteros, nil); err != nil {
		return err
	}
	if err := escribir(archivoFalso, plantillaFalso, nil, importesFalso...); err != nil {
		return err
	}
	return escribir(archivoRegistro, plantillaRegistros, lista)
}

// escribirBytes formatea el código y lo escribe en ruta, creando las carpetas necesarias.
func escribirBytes(ruta string, codigo []byte) error {
	formateado, err := format.Source(codigo)
	if err != nil {
		return fmt.Errorf("%s: %v\n%s", ruta, err, codigo)
	}
	if err := os.MkdirAll(filepath.Dir(ruta), 0o755); err != nil {
		return err
	}
	return os.WriteFile(ruta, formateado, 0o644)
}

// escribirCrudo escribe el contenido tal cual, con la cabecera de generación, y lo
// formatea. Se usa para helpers que declaran su propio paquete e imports.
func escribirCrudo(ruta, contenido string) error {
	return escribirBytes(ruta, []byte(cabecera+contenido))
}

// importesFalso son los imports del controlador falso que no se detectan por uso.
var importesFalso = []string{
	`"context"`, `"database/sql"`, `"database/sql/driver"`, `"encoding/json"`, `"errors"`, `"io"`,
	`"reflect"`, `"strings"`, `"github.com/beego/beego/v2/client/orm"`,
}

const fuenteBaseDatos = `// Package helpers contiene utilidades compartidas por las pruebas unitarias
// y de integración.
package helpers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"

	// Registra los modelos en el ORM
	_ "github.com/udistrital/cotizaciones_crud/models"
)

// IdInexistente es un Id que no existe en ninguna tabla.
const IdInexistente = 999999999

// Variables de entorno con la conexión a la base de datos (las mismas de
// environment.env), según el lineamiento de paso de parámetros de la OAS.
const (
	envUsuario    = "COTIZACIONES_CRUD_PGUSER"
	envClave      = "COTIZACIONES_CRUD_PGPASS"
	envServidor   = "COTIZACIONES_CRUD_PGHOST"
	envPuerto     = "COTIZACIONES_CRUD_PGPORT"
	envBaseDatos  = "COTIZACIONES_CRUD_PGDB"
	envEsquema    = "COTIZACIONES_CRUD_PGSCHEMA"
	motivoOmision = "Base de datos no configurada: cargue las variables COTIZACIONES_CRUD_PG* (environment.env) para ejecutar esta prueba"
)

var baseDisponible bool

// IniciarBaseDatos registra la base de datos PostgreSQL indicada en las
// variables de entorno. Si no están definidas, las pruebas se omiten (por
// ejemplo en el pipeline de IC, que no cuenta con base de datos).
func IniciarBaseDatos() error {
	usuario, servidor, base := os.Getenv(envUsuario), os.Getenv(envServidor), os.Getenv(envBaseDatos)
	if usuario == "" || servidor == "" || base == "" {
		return nil
	}
	conexion := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=disable",
		usuario, url.QueryEscape(os.Getenv(envClave)), servidor, os.Getenv(envPuerto), base, os.Getenv(envEsquema))
	if err := orm.RegisterDataBase("default", "postgres", conexion); err != nil {
		return err
	}
	baseDisponible = true
	return nil
}

// BaseDisponible indica si las pruebas cuentan con base de datos.
func BaseDisponible() bool {
	return baseDisponible
}

// RequerirBaseDatos omite la prueba cuando no hay base de datos configurada.
func RequerirBaseDatos(t *testing.T) {
	t.Helper()
	if !baseDisponible {
		t.Skip(motivoOmision)
	}
}

// EliminarAlFinal borra físicamente el registro al terminar la prueba, para
// no dejar datos de prueba en la base. Los registros se eliminan en orden
// inverso a su creación (primero los hijos y luego sus llaves foráneas).
func EliminarAlFinal(t *testing.T, tabla string, id int) {
	t.Helper()
	t.Cleanup(func() {
		if _, err := orm.NewOrm().Raw("DELETE FROM "+tabla+" WHERE id = ?", id).Exec(); err != nil {
			t.Errorf("No se pudo eliminar el registro de prueba %s %d: %v", tabla, id, err)
		}
	})
}

// Fila retorna el registro tal como está guardado en la base de datos.
func Fila(t *testing.T, tabla string, id int) map[string]interface{} {
	t.Helper()
	var texto string
	if err := orm.NewOrm().Raw("SELECT row_to_json(t)::text FROM "+tabla+" t WHERE id = ?", id).QueryRow(&texto); err != nil {
		t.Fatalf("No se pudo leer %s %d de la base de datos: %v", tabla, id, err)
	}
	fila := map[string]interface{}{}
	if err := json.Unmarshal([]byte(texto), &fila); err != nil {
		t.Fatal("No se pudo decodificar el registro:", err)
	}
	return fila
}

// SinColumnas retorna una copia de la fila sin las columnas indicadas.
func SinColumnas(fila map[string]interface{}, columnas ...string) map[string]interface{} {
	copia := map[string]interface{}{}
	for k, v := range fila {
		copia[k] = v
	}
	for _, c := range columnas {
		delete(copia, c)
	}
	return copia
}

// Contar retorna la cantidad de registros de la tabla.
func Contar(t *testing.T, tabla string) int {
	t.Helper()
	var total int
	if err := orm.NewOrm().Raw("SELECT count(*) FROM " + tabla).QueryRow(&total); err != nil {
		t.Fatalf("No se pudo contar los registros de %s: %v", tabla, err)
	}
	return total
}
`

const fuenteHTTP = `package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
)

// Respuesta representa la estructura institucional que retornan los servicios.
type Respuesta struct {
	Success bool
	Status  string
	Message string
	Data    interface{}
	Total   interface{} // solo en GetAll
}

// Peticion envía una petición HTTP al enrutador de beego (sin levantar el
// servidor) y retorna la respuesta. Requiere importar el paquete routers.
func Peticion(metodo, ruta, cuerpo string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(metodo, ruta, strings.NewReader(cuerpo))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	return w
}

// JSON serializa v para usarlo como cuerpo de una petición.
func JSON(t *testing.T, v interface{}) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal("No se pudo serializar el cuerpo de la petición:", err)
	}
	return string(b)
}

// AMapa convierte v (p. ej. un modelo) en un mapa con la misma forma que
// tiene en las respuestas JSON del API.
func AMapa(t *testing.T, v interface{}) map[string]interface{} {
	t.Helper()
	m := map[string]interface{}{}
	if err := json.Unmarshal([]byte(JSON(t, v)), &m); err != nil {
		t.Fatal("No se pudo convertir a mapa:", err)
	}
	return m
}

// leer decodifica la respuesta con la estructura institucional.
func leer(t *testing.T, w *httptest.ResponseRecorder) *Respuesta {
	t.Helper()
	var r Respuesta
	if err := json.Unmarshal(w.Body.Bytes(), &r); err != nil {
		t.Fatalf("La respuesta no tiene la estructura institucional (JSON inválido): %v\n%s", err, w.Body.String())
	}
	return &r
}

// VerificarExito comprueba el código HTTP de una respuesta exitosa y que el
// cuerpo tenga Success=true y Status igual al código.
func VerificarExito(t *testing.T, w *httptest.ResponseRecorder, estado int) *Respuesta {
	t.Helper()
	if w.Code != estado {
		t.Fatalf("Se esperaba el estado %d y se obtuvo %d: %s", estado, w.Code, w.Body.String())
	}
	r := leer(t, w)
	if !r.Success || r.Status != strconv.Itoa(estado) {
		t.Errorf("Se esperaba Success=true y Status=%d y se obtuvo Success=%v Status=%s", estado, r.Success, r.Status)
	}
	return r
}

// VerificarError comprueba el código HTTP de una respuesta de error y que el
// cuerpo tenga la estructura institucional con Success=false.
func VerificarError(t *testing.T, w *httptest.ResponseRecorder, estado int) *Respuesta {
	t.Helper()
	if w.Code != estado {
		t.Fatalf("Se esperaba el estado %d y se obtuvo %d: %s", estado, w.Code, w.Body.String())
	}
	r := leer(t, w)
	if r.Success || r.Status != strconv.Itoa(estado) {
		t.Errorf("Se esperaba Success=false y Status=%d y se obtuvo Success=%v Status=%s", estado, r.Success, r.Status)
	}
	return r
}

// VerificarFiltroInvalido comprueba la respuesta de GetAll ante un parámetro
// query mal formado: HTTP 400 con la estructura institucional, Data nulo y
// un mensaje que explica el formato esperado.
func VerificarFiltroInvalido(t *testing.T, ruta string) {
	t.Helper()
	w := Peticion(http.MethodGet, ruta, "")
	r := VerificarError(t, w, http.StatusBadRequest)
	if r.Data != nil {
		t.Error("Se esperaba Data nulo y se obtuvo", r.Data)
	}
	if !strings.Contains(r.Message, "query=campo:valor") {
		t.Error("Se esperaba un mensaje con el formato esperado (query=campo:valor) y se obtuvo:", r.Message)
	}
}

// Dato retorna el campo Data de la respuesta como un objeto.
func (r *Respuesta) Dato(t *testing.T) map[string]interface{} {
	t.Helper()
	m, ok := r.Data.(map[string]interface{})
	if !ok {
		t.Fatal("Se esperaba un objeto en Data y se obtuvo", r.Data)
	}
	return m
}

// Lista retorna el campo Data de la respuesta como una lista de objetos.
func (r *Respuesta) Lista(t *testing.T) []map[string]interface{} {
	t.Helper()
	l, ok := r.Data.([]interface{})
	if !ok {
		t.Fatal("Se esperaba una lista en Data y se obtuvo", r.Data)
	}
	lista := make([]map[string]interface{}, 0, len(l))
	for _, e := range l {
		m, ok := e.(map[string]interface{})
		if !ok {
			t.Fatal("Se esperaba una lista de objetos y se obtuvo", r.Data)
		}
		lista = append(lista, m)
	}
	return lista
}

// IdDato retorna el Id del objeto contenido en Data.
func (r *Respuesta) IdDato(t *testing.T) int {
	t.Helper()
	id, ok := r.Dato(t)["Id"].(float64)
	if !ok {
		t.Fatal("La respuesta no contiene Id:", r.Data)
	}
	return int(id)
}

// VerificarIgual compara dos valores y reporta la diferencia.
func VerificarIgual(t *testing.T, descripcion string, obtenido, esperado interface{}) {
	t.Helper()
	if !reflect.DeepEqual(obtenido, esperado) {
		t.Errorf("%s\n  esperado: %v\n  obtenido: %v", descripcion, esperado, obtenido)
	}
}

// TotalLista retorna el campo Total de una respuesta de GetAll. Falla si no
// existe o no es numérico.
func (r *Respuesta) TotalLista(t *testing.T) int {
	t.Helper()
	total, ok := r.Total.(float64)
	if !ok {
		t.Fatal("Se esperaba el campo Total numérico en la respuesta de GetAll y se obtuvo", r.Total)
	}
	return int(total)
}
`

const fuenteVerificaciones = `package helpers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

// FechaCliente es una fecha que el cliente envía en los campos administrados
// por el servicio (FechaCreacion/FechaModificacion) y que debe ignorarse.
const FechaCliente = "2000-01-01T00:00:00Z"

// CrearPorPost crea un registro mediante el servicio POST, verifica la
// respuesta 201 y programa su eliminación al terminar la prueba.
func CrearPorPost(t *testing.T, ruta, tabla string, cuerpo interface{}) int {
	t.Helper()
	w := Peticion(http.MethodPost, ruta+"/", JSON(t, cuerpo))
	id := VerificarExito(t, w, http.StatusCreated).IdDato(t)
	EliminarAlFinal(t, tabla, id)
	return id
}

// VerificarCamposAdministrados comprueba que, al crear, el servicio asigna
// Activo=true y las fechas de creación y modificación, ignorando los valores
// enviados por el cliente.
func VerificarCamposAdministrados(t *testing.T, tabla string, id int) {
	t.Helper()
	fila := Fila(t, tabla, id)
	if fila["activo"] != true {
		t.Error("Se esperaba activo=true al crear (aunque el cliente envíe false) y se obtuvo", fila["activo"])
	}
	for _, columna := range []string{"fecha_creacion", "fecha_modificacion"} {
		valor, _ := fila[columna].(string)
		if valor == "" || strings.HasPrefix(valor, "2000-") {
			t.Errorf("Se esperaba %s asignada por el servicio y se obtuvo %v", columna, fila[columna])
		}
	}
}

// ProbarPostInvalido comprueba que POST responda 400 ante un JSON mal formado
// o un cuerpo vacío.
func ProbarPostInvalido(t *testing.T, ruta string) {
	t.Helper()
	VerificarError(t, Peticion(http.MethodPost, ruta+"/", ` + "`" + `{"Id":` + "`" + `), http.StatusBadRequest)
	VerificarError(t, Peticion(http.MethodPost, ruta+"/", ""), http.StatusBadRequest)
}

// ProbarGetOne consulta el registro por Id, verifica la respuesta 200 y
// retorna el objeto obtenido.
func ProbarGetOne(t *testing.T, ruta string, id int) map[string]interface{} {
	t.Helper()
	r := VerificarExito(t, Peticion(http.MethodGet, ruta+"/"+strconv.Itoa(id), ""), http.StatusOK)
	VerificarIgual(t, "Id del registro consultado", r.IdDato(t), id)
	return r.Dato(t)
}

// ProbarGetOneInexistente comprueba que GetOne responda 404 ante un Id que no
// existe o que no es numérico.
func ProbarGetOneInexistente(t *testing.T, ruta string) {
	t.Helper()
	VerificarError(t, Peticion(http.MethodGet, ruta+"/"+strconv.Itoa(IdInexistente), ""), http.StatusNotFound)
	VerificarError(t, Peticion(http.MethodGet, ruta+"/abc", ""), http.StatusNotFound)
}

// ConsultarGetAll ejecuta GetAll con los parámetros indicados y retorna la
// lista obtenida.
func ConsultarGetAll(t *testing.T, ruta, parametros string) []map[string]interface{} {
	t.Helper()
	return VerificarExito(t, Peticion(http.MethodGet, ruta+"/?"+parametros, ""), http.StatusOK).Lista(t)
}

// ContieneId indica si la lista contiene un objeto con el Id indicado.
func ContieneId(lista []map[string]interface{}, id int) bool {
	for _, m := range lista {
		if m["Id"] == float64(id) {
			return true
		}
	}
	return false
}

// ConsultarPagina ejecuta GetAll con los parámetros indicados y retorna la
// lista obtenida y el campo Total.
func ConsultarPagina(t *testing.T, ruta, parametros string) ([]map[string]interface{}, int) {
	t.Helper()
	r := VerificarExito(t, Peticion(http.MethodGet, ruta+"/?"+parametros, ""), http.StatusOK)
	return r.Lista(t), r.TotalLista(t)
}

// ProbarGetAll verifica el servicio GetAll: filtro por Id, lista vacía sin
// resultados, Total (registros que cumplen el filtro, no solo los de la
// página), limit/offset, orden asc/desc, fields y HTTP 400 ante filtros
// inválidos. Retorna el registro obtenido con el filtro por Id.
func ProbarGetAll(t *testing.T, ruta, tabla string, id int) map[string]interface{} {
	t.Helper()
	lista, total := ConsultarPagina(t, ruta, fmt.Sprintf("query=Id:%d", id))
	if len(lista) != 1 || !ContieneId(lista, id) {
		t.Fatalf("Se esperaba solo el registro %d y se obtuvo %v", id, lista)
	}
	VerificarIgual(t, "Total con filtro por Id", total, 1)

	l, total := ConsultarPagina(t, ruta, fmt.Sprintf("query=Id:%d", IdInexistente))
	if len(l) != 0 {
		t.Error("Se esperaba una lista vacía (Data: []) y se obtuvo", l)
	}
	VerificarIgual(t, "Total sin resultados", total, 0)

	// Paginación: Total es la cantidad de registros de la tabla aunque la
	// página traiga uno solo, y offset salta los primeros registros
	registros := Contar(t, tabla)
	asc, total := ConsultarPagina(t, ruta, "sortby=Id&order=asc&limit=1")
	if len(asc) != 1 {
		t.Fatal("Se esperaba 1 registro con limit=1 y se obtuvo", len(asc))
	}
	VerificarIgual(t, "Total con limit=1", total, registros)
	desc, _ := ConsultarPagina(t, ruta, "sortby=Id&order=desc&limit=1")
	if len(desc) != 1 {
		t.Fatal("Se esperaba 1 registro con limit=1 y se obtuvo", len(desc))
	}
	primero, ultimo := asc[0]["Id"].(float64), desc[0]["Id"].(float64)
	if registros > 1 && primero >= ultimo {
		t.Errorf("Se esperaba el menor Id con order=asc y el mayor con order=desc y se obtuvo %v y %v", primero, ultimo)
	}
	siguiente, total := ConsultarPagina(t, ruta, "sortby=Id&order=asc&limit=1&offset=1")
	VerificarIgual(t, "Total con offset=1", total, registros)
	if registros > 1 && (len(siguiente) != 1 || siguiente[0]["Id"].(float64) <= primero) {
		t.Error("Se esperaba el segundo registro con offset=1 y se obtuvo", siguiente)
	} else if registros == 1 && len(siguiente) != 0 {
		t.Error("Se esperaba una lista vacía con offset=1 y se obtuvo", siguiente)
	}

	l = ConsultarGetAll(t, ruta, fmt.Sprintf("query=Id:%d&fields=Id", id))
	if len(l) != 1 || len(l[0]) != 1 || l[0]["Id"] != float64(id) {
		t.Error("Se esperaba solo el campo Id con fields=Id y se obtuvo", l)
	}

	// Filtros inválidos: HTTP 400 con la estructura institucional
	VerificarFiltroInvalido(t, ruta+"/?query=Id")
	VerificarFiltroInvalido(t, fmt.Sprintf("%s/?query=Id:%d,Activo", ruta, id))
	return lista[0]
}

// VerificarFKExpandida comprueba que la llave foránea campo venga expandida
// con los mismos datos del registro relacionado. Las llaves foráneas del
// registro relacionado solo se comparan por Id: el ORM expande hasta dos
// niveles (RelatedSel), por lo que en cadenas más largas el último nivel
// llega solo con su Id.
func VerificarFKExpandida(t *testing.T, dato map[string]interface{}, campo string, relacionado interface{}) {
	t.Helper()
	obtenido, ok := dato[campo].(map[string]interface{})
	if !ok {
		t.Errorf("Se esperaba la llave foránea %s expandida como objeto y se obtuvo %v", campo, dato[campo])
		return
	}
	VerificarIgual(t, "Llave foránea "+campo+" expandida", soloIdAnidados(obtenido), soloIdAnidados(AMapa(t, relacionado)))
}

// soloIdAnidados retorna una copia del objeto en la que cada objeto anidado
// (llave foránea) se reduce a su Id.
func soloIdAnidados(m map[string]interface{}) map[string]interface{} {
	copia := map[string]interface{}{}
	for k, v := range m {
		if anidado, ok := v.(map[string]interface{}); ok {
			copia[k] = anidado["Id"]
		} else {
			copia[k] = v
		}
	}
	return copia
}

// ProbarPut verifica el servicio PUT: actualiza el campo indicado, conserva
// fecha_creacion y usuario_creacion_id (aunque el cliente envíe otro valor)
// tras dos actualizaciones, actualiza fecha_modificacion, responde 400 ante
// Id distinto o JSON inválido y 404 ante un Id inexistente.
func ProbarPut(t *testing.T, ruta, tabla string, id int, campo string, valor interface{}, columna string) {
	t.Helper()
	rutaId := ruta + "/" + strconv.Itoa(id)
	antes := Fila(t, tabla, id)
	cuerpo := ProbarGetOne(t, ruta, id)
	delete(cuerpo, "FechaCreacion")
	delete(cuerpo, "FechaModificacion")
	cuerpo[campo] = valor
	if campo != "UsuarioCreacionId" {
		cuerpo["UsuarioCreacionId"] = IdInexistente
	}

	for i := 0; i < 2; i++ {
		VerificarExito(t, Peticion(http.MethodPut, rutaId, JSON(t, cuerpo)), http.StatusOK)
	}
	despues := Fila(t, tabla, id)
	VerificarIgual(t, "Columna "+columna+" actualizada", fmt.Sprint(despues[columna]), fmt.Sprint(valor))
	VerificarIgual(t, "fecha_creacion se conserva tras el PUT", despues["fecha_creacion"], antes["fecha_creacion"])
	VerificarIgual(t, "usuario_creacion_id se conserva tras el PUT", despues["usuario_creacion_id"], antes["usuario_creacion_id"])
	if despues["fecha_modificacion"] == nil || despues["fecha_modificacion"] == antes["fecha_modificacion"] {
		t.Error("Se esperaba fecha_modificacion actualizada y se obtuvo", despues["fecha_modificacion"])
	}

	cuerpo["Id"] = id + 1
	VerificarError(t, Peticion(http.MethodPut, rutaId, JSON(t, cuerpo)), http.StatusBadRequest)
	delete(cuerpo, "Id")
	VerificarError(t, Peticion(http.MethodPut, ruta+"/"+strconv.Itoa(IdInexistente), JSON(t, cuerpo)), http.StatusNotFound)
	VerificarError(t, Peticion(http.MethodPut, rutaId, ` + "`" + `{"Id":"texto"}` + "`" + `), http.StatusBadRequest)
}

// ProbarDelete verifica el borrado lógico: responde 200, deja activo=false,
// actualiza fecha_modificacion, no modifica las demás columnas, el registro
// sigue consultable y un Id inexistente responde 404.
func ProbarDelete(t *testing.T, ruta, tabla string, id int) {
	t.Helper()
	antes := Fila(t, tabla, id)
	VerificarExito(t, Peticion(http.MethodDelete, ruta+"/"+strconv.Itoa(id), ""), http.StatusOK)
	despues := Fila(t, tabla, id)

	if despues["activo"] != false {
		t.Error("Se esperaba activo=false (borrado lógico) y se obtuvo", despues["activo"])
	}
	if despues["fecha_modificacion"] == nil || despues["fecha_modificacion"] == antes["fecha_modificacion"] {
		t.Error("Se esperaba fecha_modificacion actualizada y se obtuvo", despues["fecha_modificacion"])
	}
	VerificarIgual(t, "Las demás columnas no cambian con el borrado lógico",
		SinColumnas(despues, "activo", "fecha_modificacion"), SinColumnas(antes, "activo", "fecha_modificacion"))

	if dato := ProbarGetOne(t, ruta, id); dato["Activo"] != false {
		t.Error("Se esperaba el registro consultable con Activo=false y se obtuvo", dato["Activo"])
	}
	VerificarError(t, Peticion(http.MethodDelete, ruta+"/"+strconv.Itoa(IdInexistente), ""), http.StatusNotFound)
}
`

// plantillaFalso emite el controlador falso de base de datos. Registra cada sentencia
// SQL con sus argumentos y responde con una fila cuyo id es 1. Permite pruebas sin
// servidor de base de datos.
var plantillaFalso = template.Must(template.New("falso").Parse(`
// Sentencia es una sentencia SQL recibida por el controlador falso, con sus argumentos.
type Sentencia struct {
	SQL  string
	Args []driver.Value
}

// CampoOpcional describe un campo escalar opcional y el valor esperado. Valor nil
// indica que se espera NULL en la base de datos y null en la respuesta.
type CampoOpcional struct {
	Campo   string      // nombre del campo en el modelo y en el JSON
	Columna string      // nombre de la columna en la base de datos
	Valor   interface{} // valor esperado, o nil
}

const nombreFalso = "cotizaciones_falso"

var (
	ultima          Sentencia
	filaSiguiente   map[string]interface{}
	falsoRegistrado bool
)

// DefinirFila fija los valores de la siguiente fila que devuelve el controlador falso,
// por nombre de columna. nil es NULL. Las columnas no indicadas devuelven 1.
func DefinirFila(valores map[string]interface{}) {
	filaSiguiente = valores
}

// ConectarFalso registra el alias default del ORM con una base de datos que no se
// conecta a ningún servidor. Debe llamarse una vez por proceso, desde TestMain.
func ConectarFalso() error {
	if !falsoRegistrado {
		sql.Register(nombreFalso, falsoDriver{})
		falsoRegistrado = true
	}
	db, err := sql.Open(nombreFalso, "")
	if err != nil {
		return err
	}
	return orm.AddAliasWthDB("default", "postgres", db)
}

// UltimaSentencia retorna la última sentencia SQL que recibió el controlador falso.
func UltimaSentencia() Sentencia {
	return ultima
}

// ValorColumna retorna el valor que la sentencia envía para la columna indicada, en un
// INSERT o en un UPDATE. El segundo valor es false si la sentencia no la incluye.
func ValorColumna(s Sentencia, columna string) (driver.Value, bool) {
	if strings.HasPrefix(s.SQL, "UPDATE ") {
		marca := "\"" + columna + "\" = $"
		i := strings.Index(s.SQL, marca)
		if i < 0 {
			return nil, false
		}
		n := 0
		for _, c := range s.SQL[i+len(marca):] {
			if c < '0' || c > '9' {
				break
			}
			n = n*10 + int(c-'0')
		}
		if n < 1 || n > len(s.Args) {
			return nil, false
		}
		return s.Args[n-1], true
	}
	ini := strings.Index(s.SQL, "(")
	fin := strings.Index(s.SQL, ") VALUES")
	if ini < 0 || fin < ini {
		return nil, false
	}
	for i, c := range strings.Split(s.SQL[ini+1:fin], ",") {
		if strings.Trim(strings.TrimSpace(c), "\"") == columna && i < len(s.Args) {
			return s.Args[i], true
		}
	}
	return nil, false
}

// Desreferenciar retorna el valor apuntado por p, o nil si p es nil.
func Desreferenciar[T any](p *T) interface{} {
	if p == nil {
		return nil
	}
	return *p
}

// VerificarColumnaEnviada comprueba el valor que la última sentencia (INSERT o UPDATE)
// envía para la columna. Un esperado nil exige NULL.
func VerificarColumnaEnviada(t *testing.T, columna string, esperado interface{}) {
	t.Helper()
	s := UltimaSentencia()
	valor, ok := ValorColumna(s, columna)
	if !ok {
		t.Fatalf("La sentencia no incluye la columna %s: %s", columna, s.SQL)
	}
	if esperado == nil {
		if valor != nil {
			t.Errorf("Se esperaba NULL en %s y se envió %v", columna, valor)
		}
		return
	}
	if valor == nil {
		t.Errorf("Se esperaba %v en %s y se envió NULL", esperado, columna)
		return
	}
	if fmt.Sprint(valor) != fmt.Sprint(esperado) {
		t.Errorf("Se esperaba %v en %s y se envió %v", esperado, columna, valor)
	}
}

// VerificarValorLeido compara el valor leído de un campo con el esperado. Un esperado
// nil exige nil. Las fechas se comparan como instantes: beego las entrega en la zona
// horaria de la sesión.
func VerificarValorLeido(t *testing.T, campo string, esperado, obtenido interface{}) {
	t.Helper()
	if e, ok := esperado.(time.Time); ok {
		if o, ok := obtenido.(time.Time); ok && e.Equal(o) {
			return
		}
	}
	if !reflect.DeepEqual(esperado, obtenido) {
		t.Errorf("Campo %s: se esperaba %v y se obtuvo %v", campo, esperado, obtenido)
	}
}

// verificarRespuesta compara el campo opcional de la respuesta JSON. Los valores
// esperados se normalizan con JSON para que coincidan con lo que decodifica el test.
// Las fechas se comparan como instantes.
func verificarRespuesta(t *testing.T, c CampoOpcional, obtenido interface{}) {
	t.Helper()
	if c.Valor == nil {
		if obtenido != nil {
			t.Errorf("Campo %s en la respuesta: se esperaba null y se obtuvo %v", c.Campo, obtenido)
		}
		return
	}
	if e, ok := c.Valor.(time.Time); ok {
		cadena, _ := obtenido.(string)
		if o, err := time.Parse(time.RFC3339Nano, cadena); err == nil && e.Equal(o) {
			return
		}
		t.Errorf("Campo %s en la respuesta: se esperaba %v y se obtuvo %v", c.Campo, e, obtenido)
		return
	}
	b, err := json.Marshal(c.Valor)
	if err != nil {
		t.Fatal("No se pudo serializar el valor esperado:", err)
	}
	var esperado interface{}
	if err := json.Unmarshal(b, &esperado); err != nil {
		t.Fatal("No se pudo decodificar el valor esperado:", err)
	}
	VerificarValorLeido(t, c.Campo, esperado, obtenido)
}

// ProbarPostOpcionales envía el cuerpo por POST y verifica la respuesta 201. Comprueba
// el valor que llega a la base de datos y el que devuelve la respuesta.
func ProbarPostOpcionales(t *testing.T, ruta string, cuerpo map[string]interface{}, campos []CampoOpcional) {
	t.Helper()
	r := VerificarExito(t, Peticion(http.MethodPost, ruta+"/", JSON(t, cuerpo)), http.StatusCreated)
	for _, c := range campos {
		VerificarColumnaEnviada(t, c.Columna, c.Valor)
	}
	dato := r.Dato(t)
	for _, c := range campos {
		verificarRespuesta(t, c, dato[c.Campo])
	}
}

// ProbarGetOpcionales define en la base de datos falsa una fila con el valor de cada
// campo opcional, consulta el registro 1 y verifica que la respuesta lo devuelva igual.
func ProbarGetOpcionales(t *testing.T, ruta string, campos []CampoOpcional) {
	t.Helper()
	fila := map[string]interface{}{}
	for _, c := range campos {
		fila[c.Columna] = c.Valor
	}
	DefinirFila(fila)
	r := VerificarExito(t, Peticion(http.MethodGet, ruta+"/1", ""), http.StatusOK)
	dato := r.Dato(t)
	for _, c := range campos {
		verificarRespuesta(t, c, dato[c.Campo])
	}
}

// ProbarPutOpcionales envía el cuerpo por PUT al registro 1 y verifica la respuesta 200.
// Comprueba el valor que llega en el UPDATE y el que devuelve la respuesta.
func ProbarPutOpcionales(t *testing.T, ruta string, cuerpo map[string]interface{}, campos []CampoOpcional) {
	t.Helper()
	DefinirFila(nil)
	cuerpo["Id"] = 1
	r := VerificarExito(t, Peticion(http.MethodPut, ruta+"/1", JSON(t, cuerpo)), http.StatusOK)
	for _, c := range campos {
		VerificarColumnaEnviada(t, c.Columna, c.Valor)
	}
	dato := r.Dato(t)
	for _, c := range campos {
		verificarRespuesta(t, c, dato[c.Campo])
	}
}

// falsoDriver es el driver de database/sql del controlador falso.
type falsoDriver struct{}

func (falsoDriver) Open(string) (driver.Conn, error) { return falsoConexion{}, nil }

// falsoConexion registra cada sentencia y no admite preparación ni transacciones.
type falsoConexion struct{}

func (falsoConexion) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("el controlador falso no prepara sentencias")
}
func (falsoConexion) Close() error { return nil }
func (falsoConexion) Begin() (driver.Tx, error) {
	return nil, errors.New("el controlador falso no maneja transacciones")
}

func (falsoConexion) ExecContext(_ context.Context, q string, a []driver.NamedValue) (driver.Result, error) {
	registrar(q, a)
	return falsoResultado{}, nil
}

func (falsoConexion) QueryContext(_ context.Context, q string, a []driver.NamedValue) (driver.Rows, error) {
	registrar(q, a)
	nombres := []string{"id"}
	if i := strings.Index(q, " FROM "); strings.HasPrefix(q, "SELECT ") && i > 0 {
		nombres = nil
		for _, p := range strings.Split(q[len("SELECT "):i], ", ") {
			partes := strings.Split(p, ".")
			nombres = append(nombres, strings.Trim(partes[len(partes)-1], "\""))
		}
	}
	return &falsoFilas{nombres: nombres, valores: filaSiguiente}, nil
}

func registrar(q string, a []driver.NamedValue) {
	args := make([]driver.Value, len(a))
	for i, v := range a {
		args[i] = v.Value
	}
	ultima = Sentencia{SQL: q, Args: args}
}

type falsoResultado struct{}

func (falsoResultado) LastInsertId() (int64, error) { return 1, nil }
func (falsoResultado) RowsAffected() (int64, error) { return 1, nil }

// falsoFilas entrega una única fila. Las columnas que no están en valores devuelven
// 1; las que sí, el valor indicado (nil es NULL).
type falsoFilas struct {
	nombres   []string
	valores   map[string]interface{}
	entregada bool
}

func (f *falsoFilas) Columns() []string { return f.nombres }
func (*falsoFilas) Close() error        { return nil }
func (f *falsoFilas) Next(d []driver.Value) error {
	if f.entregada {
		return io.EOF
	}
	f.entregada = true
	for i, n := range f.nombres {
		d[i] = int64(1)
		if v, ok := f.valores[n]; ok {
			d[i] = v
		}
	}
	return nil
}
`))

// plantillaPunteros emite el helper genérico Ptr. Los campos escalares opcionales
// de los modelos son punteros y Go no permite tomar la dirección de un literal,
// así que los registros de ejemplo usan Ptr(valor) para construirlos.
var plantillaPunteros = template.Must(template.New("punteros").Parse(`
// Ptr retorna un puntero al valor recibido. Construye los valores de ejemplo de
// los campos escalares opcionales (punteros) en los registros generados.
func Ptr[T any](v T) *T {
	return &v
}
`))

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
