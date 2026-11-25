// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package modelos

import (
	"log"
	"os"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// TestMain prepara el entorno de las pruebas de esta carpeta.
func TestMain(m *testing.M) {
	if err := helpers.ConectarFalso(); err != nil {
		log.Fatal("No se pudo registrar la base de datos falsa: ", err)
	}
	os.Exit(m.Run())
}
