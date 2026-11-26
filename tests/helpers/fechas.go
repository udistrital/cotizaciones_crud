// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package helpers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
)

// FechaGuardada es una columna date como la entrega el driver lib/pq: la medianoche del
// día guardado, marcada como UTC. FechaJSON es como debe salir en el JSON: solo el día.
var FechaGuardada = time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)

const FechaJSON = "2026-01-15"

var reFecha = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

// ProbarGetFechas fija una fila en la base de datos falsa, hace GET del registro 1 y
// comprueba el texto de cada campo en el JSON.
func ProbarGetFechas(t *testing.T, ruta string, fila map[string]interface{}, esperado map[string]string) {
	t.Helper()
	DefinirFila(fila)
	dato := VerificarExito(t, Peticion(http.MethodGet, ruta+"/1", ""), http.StatusOK).Dato(t)
	for campo, valor := range esperado {
		VerificarIgual(t, "Campo "+campo+" en el JSON", dato[campo], valor)
	}
}

// VerificarFechaJSON comprueba que un campo date venga en el JSON solo con el día
// (2006-01-02) y que sea el día guardado en la columna.
func VerificarFechaJSON(t *testing.T, dato map[string]interface{}, campo string, fila map[string]interface{}, columna string) {
	t.Helper()
	valor, _ := dato[campo].(string)
	if !reFecha.MatchString(valor) {
		t.Errorf("Se esperaba %s solo con el día (%s) y se obtuvo %v", campo, models.FormatoFecha, dato[campo])
	}
	VerificarIgual(t, "Campo "+campo+" igual a la columna "+columna, valor, fila[columna])
}

// ProbarFechasEstables envía dos veces por PUT el registro tal como lo devuelve GET y
// comprueba que las columnas indicadas no cambien (no se corren por la zona horaria).
func ProbarFechasEstables(t *testing.T, ruta, tabla string, id int, columnas ...string) {
	t.Helper()
	antes := Fila(t, tabla, id)
	for i := 1; i <= 2; i++ {
		cuerpo := ProbarGetOne(t, ruta, id)
		VerificarExito(t, Peticion(http.MethodPut, ruta+"/"+strconv.Itoa(id), JSON(t, cuerpo)), http.StatusOK)
	}
	despues := Fila(t, tabla, id)
	for _, c := range columnas {
		VerificarIgual(t, fmt.Sprintf("Columna %s tras dos GET + PUT", c), despues[c], antes[c])
	}
}
