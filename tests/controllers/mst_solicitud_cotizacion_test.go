// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controllers

import (
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas del controlador MstSolicitudCotizacionController (/v1/mst_solicitud_cotizacion).

const (
	rutaMstSolicitudCotizacion  = "/v1/mst_solicitud_cotizacion"
	tablaMstSolicitudCotizacion = "mst_solicitud_cotizacion"
)

// cuerpoPostMstSolicitudCotizacion retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostMstSolicitudCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoMstSolicitudCotizacion(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestMstSolicitudCotizacionPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaMstSolicitudCotizacion, tablaMstSolicitudCotizacion, cuerpoPostMstSolicitudCotizacion(t))
	id2 := helpers.CrearPorPost(t, rutaMstSolicitudCotizacion, tablaMstSolicitudCotizacion, cuerpoPostMstSolicitudCotizacion(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaMstSolicitudCotizacion, id)

	helpers.ProbarPostInvalido(t, rutaMstSolicitudCotizacion)
	t.Log("TestMstSolicitudCotizacionPost Finalizado Correctamente (OK)")
}

func TestMstSolicitudCotizacionGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearMstSolicitudCotizacion(t)
	helpers.ProbarGetOne(t, rutaMstSolicitudCotizacion, creado.Id)

	helpers.ProbarGetOneInexistente(t, rutaMstSolicitudCotizacion)
	t.Log("TestMstSolicitudCotizacionGetOne Finalizado Correctamente (OK)")
}

func TestMstSolicitudCotizacionGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearMstSolicitudCotizacion(t)
	helpers.ProbarGetAll(t, rutaMstSolicitudCotizacion, tablaMstSolicitudCotizacion, creado.Id)
	t.Log("TestMstSolicitudCotizacionGetAll Finalizado Correctamente (OK)")
}

func TestMstSolicitudCotizacionPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearMstSolicitudCotizacion(t)
	helpers.ProbarPut(t, rutaMstSolicitudCotizacion, tablaMstSolicitudCotizacion, creado.Id, "UsuarioModificacionId", 2, "usuario_modificacion_id")
	t.Log("TestMstSolicitudCotizacionPut Finalizado Correctamente (OK)")
}

func TestMstSolicitudCotizacionDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearMstSolicitudCotizacion(t)
	helpers.ProbarDelete(t, rutaMstSolicitudCotizacion, tablaMstSolicitudCotizacion, creado.Id)
	t.Log("TestMstSolicitudCotizacionDelete Finalizado Correctamente (OK)")
}
