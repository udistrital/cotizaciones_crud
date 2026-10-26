package models

import (
	"log"
	"os"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestMain(m *testing.M) {
	limpiar, err := helpers.IniciarBaseDatos()
	if err != nil {
		log.Fatal("No se pudo iniciar la base de datos de pruebas: ", err)
	}
	codigo := m.Run()
	limpiar()
	os.Exit(codigo)
}
