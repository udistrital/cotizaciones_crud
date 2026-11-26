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

// Pruebas del controlador ValidacionOrdenadorController (/v1/validacion_ordenador).

const (
	rutaValidacionOrdenador  = "/v1/validacion_ordenador"
	tablaValidacionOrdenador = "validacion_ordenador"
)

// cuerpoPostValidacionOrdenador retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostValidacionOrdenador(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoValidacionOrdenador(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestValidacionOrdenadorPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaValidacionOrdenador, tablaValidacionOrdenador, cuerpoPostValidacionOrdenador(t))
	id2 := helpers.CrearPorPost(t, rutaValidacionOrdenador, tablaValidacionOrdenador, cuerpoPostValidacionOrdenador(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaValidacionOrdenador, id)

	// La llave foránea MstSolicitudCotizacionId es obligatoria
	sinMstSolicitudCotizacionId := cuerpoPostValidacionOrdenador(t)
	sinMstSolicitudCotizacionId["MstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaValidacionOrdenador+"/", helpers.JSON(t, sinMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaValidacionOrdenador)
	t.Log("TestValidacionOrdenadorPost Finalizado Correctamente (OK)")
}

func TestValidacionOrdenadorGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearValidacionOrdenador(t)
	dato := helpers.ProbarGetOne(t, rutaValidacionOrdenador, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaValidacionOrdenador)
	t.Log("TestValidacionOrdenadorGetOne Finalizado Correctamente (OK)")
}

func TestValidacionOrdenadorGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearValidacionOrdenador(t)
	dato := helpers.ProbarGetAll(t, rutaValidacionOrdenador, tablaValidacionOrdenador, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaValidacionOrdenador, fmt.Sprintf("query=MstSolicitudCotizacionId.Id:%d", creado.MstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MstSolicitudCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestValidacionOrdenadorGetAll Finalizado Correctamente (OK)")
}

func TestValidacionOrdenadorPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearValidacionOrdenador(t)
	helpers.ProbarPut(t, rutaValidacionOrdenador, tablaValidacionOrdenador, creado.Id, "Observaciones", "Prueba-2", "observaciones")
	t.Log("TestValidacionOrdenadorPut Finalizado Correctamente (OK)")
}

func TestValidacionOrdenadorDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearValidacionOrdenador(t)
	helpers.ProbarDelete(t, rutaValidacionOrdenador, tablaValidacionOrdenador, creado.Id)
	t.Log("TestValidacionOrdenadorDelete Finalizado Correctamente (OK)")
}

// La llave foránea MstSolicitudCotizacionId debe referenciar un registro existente: POST y PUT
// responden 400 y no crean ni modifican registros.
func TestValidacionOrdenadorMstSolicitudCotizacionIdInexistente(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	helpers.ProbarPostFKInexistente(t, rutaValidacionOrdenador, tablaValidacionOrdenador, cuerpoPostValidacionOrdenador(t), "MstSolicitudCotizacionId")
	creado := helpers.CrearValidacionOrdenador(t)
	helpers.ProbarPutFKInexistente(t, rutaValidacionOrdenador, tablaValidacionOrdenador, creado.Id, "MstSolicitudCotizacionId")
	t.Log("TestValidacionOrdenadorMstSolicitudCotizacionIdInexistente Finalizado Correctamente (OK)")
}
