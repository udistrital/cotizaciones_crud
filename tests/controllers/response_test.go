package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// TestRespuestaEstructura verifica que las respuestas exitosas y de error
// tengan la estructura del lineamiento: Success, Status (texto), Message y Data.
func TestRespuestaEstructura(t *testing.T) {
	creado := helpers.CrearSolicitudCiiu(t)
	casos := []struct {
		nombre string
		metodo string
		ruta   string
		estado int
	}{
		{"Exito", "GET", fmt.Sprintf("%s/%d", rutaSolicitudCiiu, creado.Id), http.StatusOK},
		{"Error", "GET", fmt.Sprintf("%s/%d", rutaSolicitudCiiu, helpers.IdInexistente), http.StatusNotFound},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			w := helpers.Peticion(c.metodo, c.ruta, "")
			if w.Code != c.estado {
				t.Fatalf("Se esperaba el estado %d y se obtuvo %d", c.estado, w.Code)
			}
			var cuerpo map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &cuerpo); err != nil {
				t.Fatal("La respuesta no es un JSON válido:", err)
			}
			for _, campo := range []string{"Success", "Status", "Message", "Data"} {
				if _, ok := cuerpo[campo]; !ok {
					t.Error("La respuesta no contiene el campo", campo)
				}
			}
			if len(cuerpo) != 4 {
				t.Error("Se esperaban 4 campos en la respuesta y se obtuvo", cuerpo)
			}
			if _, ok := cuerpo["Status"].(string); !ok {
				t.Error("Se esperaba Status como texto y se obtuvo", cuerpo["Status"])
			}
		})
	}
	t.Log("TestRespuestaEstructura Finalizado Correctamente (OK)")
}
