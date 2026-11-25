package helpers

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
