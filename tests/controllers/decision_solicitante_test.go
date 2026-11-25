// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas del controlador DecisionSolicitanteController (/v1/decision_solicitante).

const (
	rutaDecisionSolicitante  = "/v1/decision_solicitante"
	tablaDecisionSolicitante = "decision_solicitante"
)

// cuerpoPostDecisionSolicitante retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostDecisionSolicitante(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoDecisionSolicitante(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestDecisionSolicitantePost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaDecisionSolicitante, tablaDecisionSolicitante, cuerpoPostDecisionSolicitante(t))
	id2 := helpers.CrearPorPost(t, rutaDecisionSolicitante, tablaDecisionSolicitante, cuerpoPostDecisionSolicitante(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaDecisionSolicitante, id)

	// La llave foránea RelSolicitudCotizacionProveedorId es obligatoria
	sinRelSolicitudCotizacionProveedorId := cuerpoPostDecisionSolicitante(t)
	sinRelSolicitudCotizacionProveedorId["RelSolicitudCotizacionProveedorId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaDecisionSolicitante+"/", helpers.JSON(t, sinRelSolicitudCotizacionProveedorId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaDecisionSolicitante)
	t.Log("TestDecisionSolicitantePost Finalizado Correctamente (OK)")
}

func TestDecisionSolicitanteGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearDecisionSolicitante(t)
	dato := helpers.ProbarGetOne(t, rutaDecisionSolicitante, creado.Id)

	// La llave foránea RelSolicitudCotizacionProveedorId viene expandida
	if relacionado, err := models.GetRelSolicitudCotizacionProveedorById(creado.RelSolicitudCotizacionProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar RelSolicitudCotizacionProveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "RelSolicitudCotizacionProveedorId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaDecisionSolicitante)
	t.Log("TestDecisionSolicitanteGetOne Finalizado Correctamente (OK)")
}

func TestDecisionSolicitanteGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearDecisionSolicitante(t)
	dato := helpers.ProbarGetAll(t, rutaDecisionSolicitante, tablaDecisionSolicitante, creado.Id)

	// La llave foránea RelSolicitudCotizacionProveedorId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetRelSolicitudCotizacionProveedorById(creado.RelSolicitudCotizacionProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar RelSolicitudCotizacionProveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "RelSolicitudCotizacionProveedorId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaDecisionSolicitante, fmt.Sprintf("query=RelSolicitudCotizacionProveedorId.Id:%d", creado.RelSolicitudCotizacionProveedorId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por RelSolicitudCotizacionProveedorId.Id y se obtuvo", l)
	}
	t.Log("TestDecisionSolicitanteGetAll Finalizado Correctamente (OK)")
}

func TestDecisionSolicitantePut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearDecisionSolicitante(t)
	helpers.ProbarPut(t, rutaDecisionSolicitante, tablaDecisionSolicitante, creado.Id, "MensajeAProveedor", "Prueba-2", "mensaje_a_proveedor")
	t.Log("TestDecisionSolicitantePut Finalizado Correctamente (OK)")
}

func TestDecisionSolicitanteDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearDecisionSolicitante(t)
	helpers.ProbarDelete(t, rutaDecisionSolicitante, tablaDecisionSolicitante, creado.Id)
	t.Log("TestDecisionSolicitanteDelete Finalizado Correctamente (OK)")
}
