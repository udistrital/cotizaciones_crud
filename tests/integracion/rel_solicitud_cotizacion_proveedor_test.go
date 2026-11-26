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

// Pruebas del controlador RelSolicitudCotizacionProveedorController (/v1/rel_solicitud_cotizacion_proveedor).

const (
	rutaRelSolicitudCotizacionProveedor  = "/v1/rel_solicitud_cotizacion_proveedor"
	tablaRelSolicitudCotizacionProveedor = "rel_solicitud_cotizacion_proveedor"
)

// cuerpoPostRelSolicitudCotizacionProveedor retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostRelSolicitudCotizacionProveedor(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoRelSolicitudCotizacionProveedor(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestRelSolicitudCotizacionProveedorPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, cuerpoPostRelSolicitudCotizacionProveedor(t))
	id2 := helpers.CrearPorPost(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, cuerpoPostRelSolicitudCotizacionProveedor(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaRelSolicitudCotizacionProveedor, id)

	// La llave foránea MstSolicitudCotizacionId es obligatoria
	sinMstSolicitudCotizacionId := cuerpoPostRelSolicitudCotizacionProveedor(t)
	sinMstSolicitudCotizacionId["MstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaRelSolicitudCotizacionProveedor+"/", helpers.JSON(t, sinMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaRelSolicitudCotizacionProveedor)
	t.Log("TestRelSolicitudCotizacionProveedorPost Finalizado Correctamente (OK)")
}

func TestRelSolicitudCotizacionProveedorGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelSolicitudCotizacionProveedor(t)
	dato := helpers.ProbarGetOne(t, rutaRelSolicitudCotizacionProveedor, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}

	// La llave foránea NotificacionId viene expandida
	if relacionado, err := models.GetNotificacionById(creado.NotificacionId.Id); err != nil {
		t.Fatal("No se pudo consultar Notificacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "NotificacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaRelSolicitudCotizacionProveedor)
	t.Log("TestRelSolicitudCotizacionProveedorGetOne Finalizado Correctamente (OK)")
}

func TestRelSolicitudCotizacionProveedorGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelSolicitudCotizacionProveedor(t)
	dato := helpers.ProbarGetAll(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaRelSolicitudCotizacionProveedor, fmt.Sprintf("query=MstSolicitudCotizacionId.Id:%d", creado.MstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MstSolicitudCotizacionId.Id y se obtuvo", l)
	}

	// La llave foránea NotificacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetNotificacionById(creado.NotificacionId.Id); err != nil {
		t.Fatal("No se pudo consultar Notificacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "NotificacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaRelSolicitudCotizacionProveedor, fmt.Sprintf("query=NotificacionId.Id:%d", creado.NotificacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por NotificacionId.Id y se obtuvo", l)
	}
	t.Log("TestRelSolicitudCotizacionProveedorGetAll Finalizado Correctamente (OK)")
}

func TestRelSolicitudCotizacionProveedorPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelSolicitudCotizacionProveedor(t)
	helpers.ProbarPut(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, creado.Id, "ProveedorId", 2, "proveedor_id")
	t.Log("TestRelSolicitudCotizacionProveedorPut Finalizado Correctamente (OK)")
}

func TestRelSolicitudCotizacionProveedorDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelSolicitudCotizacionProveedor(t)
	helpers.ProbarDelete(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, creado.Id)
	t.Log("TestRelSolicitudCotizacionProveedorDelete Finalizado Correctamente (OK)")
}

// La llave foránea MstSolicitudCotizacionId debe referenciar un registro existente: POST y PUT
// responden 400 y no crean ni modifican registros.
func TestRelSolicitudCotizacionProveedorMstSolicitudCotizacionIdInexistente(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	helpers.ProbarPostFKInexistente(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, cuerpoPostRelSolicitudCotizacionProveedor(t), "MstSolicitudCotizacionId")
	creado := helpers.CrearRelSolicitudCotizacionProveedor(t)
	helpers.ProbarPutFKInexistente(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, creado.Id, "MstSolicitudCotizacionId")
	t.Log("TestRelSolicitudCotizacionProveedorMstSolicitudCotizacionIdInexistente Finalizado Correctamente (OK)")
}

// La llave foránea NotificacionId debe referenciar un registro existente: POST y PUT
// responden 400 y no crean ni modifican registros.
func TestRelSolicitudCotizacionProveedorNotificacionIdInexistente(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	helpers.ProbarPostFKInexistente(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, cuerpoPostRelSolicitudCotizacionProveedor(t), "NotificacionId")
	creado := helpers.CrearRelSolicitudCotizacionProveedor(t)
	helpers.ProbarPutFKInexistente(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, creado.Id, "NotificacionId")
	t.Log("TestRelSolicitudCotizacionProveedorNotificacionIdInexistente Finalizado Correctamente (OK)")
}

// La llave foránea NotificacionId es opcional (null): el registro sin ella se
// puede crear, consultar, actualizar y borrar.
func TestRelSolicitudCotizacionProveedorNotificacionIdOpcional(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostRelSolicitudCotizacionProveedor(t)
	cuerpo["NotificacionId"] = nil
	id := helpers.CrearPorPost(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, cuerpo)

	if dato := helpers.ProbarGetOne(t, rutaRelSolicitudCotizacionProveedor, id); dato["NotificacionId"] != nil {
		t.Error("Se esperaba NotificacionId nulo y se obtuvo", dato["NotificacionId"])
	}
	if l := helpers.ConsultarGetAll(t, rutaRelSolicitudCotizacionProveedor, fmt.Sprintf("query=Id:%d", id)); !helpers.ContieneId(l, id) {
		t.Error("Se esperaba el registro sin NotificacionId en GetAll y se obtuvo", l)
	}
	helpers.ProbarPut(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, id, "ProveedorId", 2, "proveedor_id")
	helpers.ProbarDelete(t, rutaRelSolicitudCotizacionProveedor, tablaRelSolicitudCotizacionProveedor, id)
	t.Log("TestRelSolicitudCotizacionProveedorNotificacionIdOpcional Finalizado Correctamente (OK)")
}
