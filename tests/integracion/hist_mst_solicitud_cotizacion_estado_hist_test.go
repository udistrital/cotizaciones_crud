// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package integracion

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas del controlador HistMstSolicitudCotizacionEstadoHistController (/v1/hist_mst_solicitud_cotizacion_estado_hist).

const (
	rutaHistMstSolicitudCotizacionEstadoHist  = "/v1/hist_mst_solicitud_cotizacion_estado_hist"
	tablaHistMstSolicitudCotizacionEstadoHist = "hist_mst_solicitud_cotizacion_estado_hist"
)

// cuerpoPostHistMstSolicitudCotizacionEstadoHist retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostHistMstSolicitudCotizacionEstadoHist(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoHistMstSolicitudCotizacionEstadoHist(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestHistMstSolicitudCotizacionEstadoHistPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacionEstadoHist, tablaHistMstSolicitudCotizacionEstadoHist, cuerpoPostHistMstSolicitudCotizacionEstadoHist(t))
	id2 := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacionEstadoHist, tablaHistMstSolicitudCotizacionEstadoHist, cuerpoPostHistMstSolicitudCotizacionEstadoHist(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaHistMstSolicitudCotizacionEstadoHist, id)

	// La llave foránea HistMstSolicitudCotizacionId es obligatoria
	sinHistMstSolicitudCotizacionId := cuerpoPostHistMstSolicitudCotizacionEstadoHist(t)
	sinHistMstSolicitudCotizacionId["HistMstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaHistMstSolicitudCotizacionEstadoHist+"/", helpers.JSON(t, sinHistMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaHistMstSolicitudCotizacionEstadoHist)
	t.Log("TestHistMstSolicitudCotizacionEstadoHistPost Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionEstadoHistGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacionEstadoHist(t)
	dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacionEstadoHist, creado.Id)

	// La llave foránea HistMstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetHistMstSolicitudCotizacionById(creado.HistMstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "HistMstSolicitudCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaHistMstSolicitudCotizacionEstadoHist)
	t.Log("TestHistMstSolicitudCotizacionEstadoHistGetOne Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionEstadoHistGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacionEstadoHist(t)
	dato := helpers.ProbarGetAll(t, rutaHistMstSolicitudCotizacionEstadoHist, tablaHistMstSolicitudCotizacionEstadoHist, creado.Id)

	// La llave foránea HistMstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetHistMstSolicitudCotizacionById(creado.HistMstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "HistMstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaHistMstSolicitudCotizacionEstadoHist, fmt.Sprintf("query=HistMstSolicitudCotizacionId.Id:%d", creado.HistMstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por HistMstSolicitudCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestHistMstSolicitudCotizacionEstadoHistGetAll Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionEstadoHistPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacionEstadoHist(t)
	helpers.ProbarPut(t, rutaHistMstSolicitudCotizacionEstadoHist, tablaHistMstSolicitudCotizacionEstadoHist, creado.Id, "Justificacion", "Prueba-2", "justificacion")
	t.Log("TestHistMstSolicitudCotizacionEstadoHistPut Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionEstadoHistDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacionEstadoHist(t)
	helpers.ProbarDelete(t, rutaHistMstSolicitudCotizacionEstadoHist, tablaHistMstSolicitudCotizacionEstadoHist, creado.Id)
	t.Log("TestHistMstSolicitudCotizacionEstadoHistDelete Finalizado Correctamente (OK)")
}

// La llave foránea HistMstSolicitudCotizacionId debe referenciar un registro existente: POST y PUT
// responden 400 y no crean ni modifican registros.
func TestHistMstSolicitudCotizacionEstadoHistHistMstSolicitudCotizacionIdInexistente(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	helpers.ProbarPostFKInexistente(t, rutaHistMstSolicitudCotizacionEstadoHist, tablaHistMstSolicitudCotizacionEstadoHist, cuerpoPostHistMstSolicitudCotizacionEstadoHist(t), "HistMstSolicitudCotizacionId")
	creado := helpers.CrearHistMstSolicitudCotizacionEstadoHist(t)
	helpers.ProbarPutFKInexistente(t, rutaHistMstSolicitudCotizacionEstadoHist, tablaHistMstSolicitudCotizacionEstadoHist, creado.Id, "HistMstSolicitudCotizacionId")
	t.Log("TestHistMstSolicitudCotizacionEstadoHistHistMstSolicitudCotizacionIdInexistente Finalizado Correctamente (OK)")
}
