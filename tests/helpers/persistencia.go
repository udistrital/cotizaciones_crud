// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package helpers

import (
	"net/http"
	"strings"
	"testing"

	"github.com/lib/pq"
)

// Errores de persistencia simulados, como los que devuelve el driver lib/pq.
var (
	// ErrBaseNoDisponible: la base de datos no está disponible (SQLSTATE 08006, connection_failure).
	ErrBaseNoDisponible = &pq.Error{Code: "08006", Message: "connection failure"}
	// ErrRestriccion: la base rechaza los datos (SQLSTATE 23503, foreign_key_violation).
	ErrRestriccion = &pq.Error{Code: "23503", Message: "insert or update violates foreign key constraint"}
)

// ProbarErrorPersistencia simula err en la siguiente sentencia que empiece por prefijo,
// envía la petición y comprueba la respuesta: el estado indicado con la estructura
// institucional (Success=false, Data nulo), un mensaje del servicio indicado y sin
// detalles internos del error de la base de datos. Falla si la petición no llegó a
// ejecutar la sentencia que debía fallar.
func ProbarErrorPersistencia(t *testing.T, metodo, ruta, cuerpo, prefijo string, err *pq.Error, estado int, servicio string) {
	t.Helper()
	LimpiarSentencia()
	DefinirFila(nil)
	SimularError(prefijo, err)
	w := Peticion(metodo, ruta, cuerpo)
	pendiente := ErrorSimuladoPendiente()
	SimularError("", nil)
	if pendiente {
		t.Fatalf("La petición no ejecutó ninguna sentencia %s: el error simulado no se aplicó (respuesta %d: %s)", prefijo, w.Code, w.Body.String())
	}
	r := VerificarError(t, w, estado)
	if r.Data != nil {
		t.Error("Se esperaba Data nulo y se obtuvo", r.Data)
	}
	if !strings.Contains(r.Message, "Error service "+servicio+":") {
		t.Errorf("Se esperaba un mensaje del servicio %s y se obtuvo: %s", servicio, r.Message)
	}
	if estado == http.StatusInternalServerError && !strings.Contains(r.Message, "The database is not available") {
		t.Error("Se esperaba un mensaje de base de datos no disponible y se obtuvo:", r.Message)
	}
	for _, interno := range []string{"pq:", err.Message, string(err.Code), "SELECT", "INSERT", "UPDATE"} {
		if strings.Contains(r.Message, interno) {
			t.Errorf("El mensaje expone detalles internos de la base de datos (%q): %s", interno, r.Message)
		}
	}
}
