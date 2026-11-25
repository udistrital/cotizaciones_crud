// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controllers

import (
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas del controlador NotificacionController (/v1/notificacion).

const (
	rutaNotificacion  = "/v1/notificacion"
	tablaNotificacion = "notificacion"
)

// cuerpoPostNotificacion retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostNotificacion(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoNotificacion(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestNotificacionPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaNotificacion, tablaNotificacion, cuerpoPostNotificacion(t))
	id2 := helpers.CrearPorPost(t, rutaNotificacion, tablaNotificacion, cuerpoPostNotificacion(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaNotificacion, id)

	helpers.ProbarPostInvalido(t, rutaNotificacion)
	t.Log("TestNotificacionPost Finalizado Correctamente (OK)")
}

func TestNotificacionGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearNotificacion(t)
	helpers.ProbarGetOne(t, rutaNotificacion, creado.Id)

	helpers.ProbarGetOneInexistente(t, rutaNotificacion)
	t.Log("TestNotificacionGetOne Finalizado Correctamente (OK)")
}

func TestNotificacionGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearNotificacion(t)
	helpers.ProbarGetAll(t, rutaNotificacion, tablaNotificacion, creado.Id)
	t.Log("TestNotificacionGetAll Finalizado Correctamente (OK)")
}

func TestNotificacionPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearNotificacion(t)
	helpers.ProbarPut(t, rutaNotificacion, tablaNotificacion, creado.Id, "ProveedorNotificadoId", 2, "proveedor_notificado_id")
	t.Log("TestNotificacionPut Finalizado Correctamente (OK)")
}

func TestNotificacionDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearNotificacion(t)
	helpers.ProbarDelete(t, rutaNotificacion, tablaNotificacion, creado.Id)
	t.Log("TestNotificacionDelete Finalizado Correctamente (OK)")
}
