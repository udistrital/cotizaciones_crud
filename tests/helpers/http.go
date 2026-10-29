package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
)

// Respuesta representa el cuerpo JSON que retornan los controladores, tanto
// en las respuestas exitosas como en las de error (customerror).
type Respuesta struct {
	Success bool
	Status  string
	Message string
	Total   *int64 // solo en las respuestas de GetAll
	Data    interface{}
}

// Peticion envía una petición HTTP al enrutador de beego y retorna la
// respuesta. Requiere que el paquete routers haya sido importado.
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

// decodificar verifica el código HTTP y decodifica el cuerpo de la respuesta.
func decodificar(t *testing.T, w *httptest.ResponseRecorder, estadoEsperado int) *Respuesta {
	t.Helper()
	if w.Code != estadoEsperado {
		t.Fatalf("Se esperaba el estado %d y se obtuvo %d: %s", estadoEsperado, w.Code, w.Body.String())
	}
	var r Respuesta
	if err := json.Unmarshal(w.Body.Bytes(), &r); err != nil {
		t.Fatal("La respuesta no es un JSON válido:", err, w.Body.String())
	}
	if r.Status != strconv.Itoa(estadoEsperado) {
		t.Error("Se esperaba Status", estadoEsperado, "y se obtuvo", r.Status)
	}
	return &r
}

// VerificarRespuesta comprueba que la respuesta sea exitosa con el código
// HTTP esperado y retorna su cuerpo.
func VerificarRespuesta(t *testing.T, w *httptest.ResponseRecorder, estadoEsperado int) *Respuesta {
	t.Helper()
	r := decodificar(t, w, estadoEsperado)
	if !r.Success {
		t.Error("Se esperaba Success=true y se obtuvo", r.Success)
	}
	return r
}

// VerificarError comprueba que la respuesta sea un error con el código HTTP
// esperado, con la estructura de customerror y con un mensaje que contenga
// mensajeEsperado.
func VerificarError(t *testing.T, w *httptest.ResponseRecorder, estadoEsperado int, mensajeEsperado string) {
	t.Helper()
	if estadoEsperado < http.StatusBadRequest {
		t.Fatal("VerificarError solo aplica a códigos de error, se recibió", estadoEsperado)
	}
	r := decodificar(t, w, estadoEsperado)
	if r.Success || r.Data != nil {
		t.Error("Se esperaba Success=false y Data=null y se obtuvo", r.Success, r.Data)
	}
	if !strings.Contains(r.Message, mensajeEsperado) {
		t.Errorf("Se esperaba un mensaje que contenga %q y se obtuvo %q", mensajeEsperado, r.Message)
	}
}

// Dato retorna el campo Data de la respuesta como un mapa.
func (r *Respuesta) Dato(t *testing.T) map[string]interface{} {
	t.Helper()
	m, ok := r.Data.(map[string]interface{})
	if !ok {
		t.Fatal("Se esperaba un objeto en Data y se obtuvo", r.Data)
	}
	return m
}

// Lista retorna el campo Data de la respuesta como una lista.
func (r *Respuesta) Lista(t *testing.T) []interface{} {
	t.Helper()
	l, ok := r.Data.([]interface{})
	if !ok {
		t.Fatal("Se esperaba una lista en Data y se obtuvo", r.Data)
	}
	return l
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
