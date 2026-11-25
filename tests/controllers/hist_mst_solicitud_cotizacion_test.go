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

// Pruebas del controlador HistMstSolicitudCotizacionController (/v1/hist_mst_solicitud_cotizacion).

const (
	rutaHistMstSolicitudCotizacion  = "/v1/hist_mst_solicitud_cotizacion"
	tablaHistMstSolicitudCotizacion = "hist_mst_solicitud_cotizacion"
)

// cuerpoPostHistMstSolicitudCotizacion retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostHistMstSolicitudCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoHistMstSolicitudCotizacion(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestHistMstSolicitudCotizacionPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpoPostHistMstSolicitudCotizacion(t))
	id2 := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpoPostHistMstSolicitudCotizacion(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaHistMstSolicitudCotizacion, id)

	// La llave foránea MstSolicitudCotizacionId es obligatoria
	sinMstSolicitudCotizacionId := cuerpoPostHistMstSolicitudCotizacion(t)
	sinMstSolicitudCotizacionId["MstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaHistMstSolicitudCotizacion+"/", helpers.JSON(t, sinMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaHistMstSolicitudCotizacion)
	t.Log("TestHistMstSolicitudCotizacionPost Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacion(t)
	dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaHistMstSolicitudCotizacion)
	t.Log("TestHistMstSolicitudCotizacionGetOne Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacion(t)
	dato := helpers.ProbarGetAll(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaHistMstSolicitudCotizacion, fmt.Sprintf("query=MstSolicitudCotizacionId.Id:%d", creado.MstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MstSolicitudCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestHistMstSolicitudCotizacionGetAll Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacion(t)
	helpers.ProbarPut(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, creado.Id, "NumeroSolicitud", "Prueba-2", "numero_solicitud")
	t.Log("TestHistMstSolicitudCotizacionPut Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacion(t)
	helpers.ProbarDelete(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, creado.Id)
	t.Log("TestHistMstSolicitudCotizacionDelete Finalizado Correctamente (OK)")
}
