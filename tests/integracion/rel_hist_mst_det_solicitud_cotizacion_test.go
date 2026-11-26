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

// Pruebas del controlador RelHistMstDetSolicitudCotizacionController (/v1/rel_hist_mst_det_solicitud_cotizacion).

const (
	rutaRelHistMstDetSolicitudCotizacion  = "/v1/rel_hist_mst_det_solicitud_cotizacion"
	tablaRelHistMstDetSolicitudCotizacion = "rel_hist_mst_det_solicitud_cotizacion"
)

// cuerpoPostRelHistMstDetSolicitudCotizacion retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostRelHistMstDetSolicitudCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoRelHistMstDetSolicitudCotizacion(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestRelHistMstDetSolicitudCotizacionPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaRelHistMstDetSolicitudCotizacion, tablaRelHistMstDetSolicitudCotizacion, cuerpoPostRelHistMstDetSolicitudCotizacion(t))
	id2 := helpers.CrearPorPost(t, rutaRelHistMstDetSolicitudCotizacion, tablaRelHistMstDetSolicitudCotizacion, cuerpoPostRelHistMstDetSolicitudCotizacion(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaRelHistMstDetSolicitudCotizacion, id)

	// La llave foránea MaestroId es obligatoria
	sinMaestroId := cuerpoPostRelHistMstDetSolicitudCotizacion(t)
	sinMaestroId["MaestroId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaRelHistMstDetSolicitudCotizacion+"/", helpers.JSON(t, sinMaestroId)), http.StatusBadRequest)

	// La llave foránea DetalleId es obligatoria
	sinDetalleId := cuerpoPostRelHistMstDetSolicitudCotizacion(t)
	sinDetalleId["DetalleId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaRelHistMstDetSolicitudCotizacion+"/", helpers.JSON(t, sinDetalleId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaRelHistMstDetSolicitudCotizacion)
	t.Log("TestRelHistMstDetSolicitudCotizacionPost Finalizado Correctamente (OK)")
}

func TestRelHistMstDetSolicitudCotizacionGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelHistMstDetSolicitudCotizacion(t)
	dato := helpers.ProbarGetOne(t, rutaRelHistMstDetSolicitudCotizacion, creado.Id)

	// La llave foránea MaestroId viene expandida
	if relacionado, err := models.GetHistMstSolicitudCotizacionById(creado.MaestroId.Id); err != nil {
		t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MaestroId", relacionado)
	}

	// La llave foránea DetalleId viene expandida
	if relacionado, err := models.GetHistDetSolicitudCotizacionById(creado.DetalleId.Id); err != nil {
		t.Fatal("No se pudo consultar HistDetSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "DetalleId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaRelHistMstDetSolicitudCotizacion)
	t.Log("TestRelHistMstDetSolicitudCotizacionGetOne Finalizado Correctamente (OK)")
}

func TestRelHistMstDetSolicitudCotizacionGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelHistMstDetSolicitudCotizacion(t)
	dato := helpers.ProbarGetAll(t, rutaRelHistMstDetSolicitudCotizacion, tablaRelHistMstDetSolicitudCotizacion, creado.Id)

	// La llave foránea MaestroId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetHistMstSolicitudCotizacionById(creado.MaestroId.Id); err != nil {
		t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MaestroId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaRelHistMstDetSolicitudCotizacion, fmt.Sprintf("query=MaestroId.Id:%d", creado.MaestroId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MaestroId.Id y se obtuvo", l)
	}

	// La llave foránea DetalleId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetHistDetSolicitudCotizacionById(creado.DetalleId.Id); err != nil {
		t.Fatal("No se pudo consultar HistDetSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "DetalleId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaRelHistMstDetSolicitudCotizacion, fmt.Sprintf("query=DetalleId.Id:%d", creado.DetalleId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por DetalleId.Id y se obtuvo", l)
	}
	t.Log("TestRelHistMstDetSolicitudCotizacionGetAll Finalizado Correctamente (OK)")
}

func TestRelHistMstDetSolicitudCotizacionPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelHistMstDetSolicitudCotizacion(t)
	helpers.ProbarPut(t, rutaRelHistMstDetSolicitudCotizacion, tablaRelHistMstDetSolicitudCotizacion, creado.Id, "UsuarioModificacionId", 2, "usuario_modificacion_id")
	t.Log("TestRelHistMstDetSolicitudCotizacionPut Finalizado Correctamente (OK)")
}

func TestRelHistMstDetSolicitudCotizacionDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelHistMstDetSolicitudCotizacion(t)
	helpers.ProbarDelete(t, rutaRelHistMstDetSolicitudCotizacion, tablaRelHistMstDetSolicitudCotizacion, creado.Id)
	t.Log("TestRelHistMstDetSolicitudCotizacionDelete Finalizado Correctamente (OK)")
}

// La llave foránea MaestroId debe referenciar un registro existente: POST y PUT
// responden 400 y no crean ni modifican registros.
func TestRelHistMstDetSolicitudCotizacionMaestroIdInexistente(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	helpers.ProbarPostFKInexistente(t, rutaRelHistMstDetSolicitudCotizacion, tablaRelHistMstDetSolicitudCotizacion, cuerpoPostRelHistMstDetSolicitudCotizacion(t), "MaestroId")
	creado := helpers.CrearRelHistMstDetSolicitudCotizacion(t)
	helpers.ProbarPutFKInexistente(t, rutaRelHistMstDetSolicitudCotizacion, tablaRelHistMstDetSolicitudCotizacion, creado.Id, "MaestroId")
	t.Log("TestRelHistMstDetSolicitudCotizacionMaestroIdInexistente Finalizado Correctamente (OK)")
}

// La llave foránea DetalleId debe referenciar un registro existente: POST y PUT
// responden 400 y no crean ni modifican registros.
func TestRelHistMstDetSolicitudCotizacionDetalleIdInexistente(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	helpers.ProbarPostFKInexistente(t, rutaRelHistMstDetSolicitudCotizacion, tablaRelHistMstDetSolicitudCotizacion, cuerpoPostRelHistMstDetSolicitudCotizacion(t), "DetalleId")
	creado := helpers.CrearRelHistMstDetSolicitudCotizacion(t)
	helpers.ProbarPutFKInexistente(t, rutaRelHistMstDetSolicitudCotizacion, tablaRelHistMstDetSolicitudCotizacion, creado.Id, "DetalleId")
	t.Log("TestRelHistMstDetSolicitudCotizacionDetalleIdInexistente Finalizado Correctamente (OK)")
}
