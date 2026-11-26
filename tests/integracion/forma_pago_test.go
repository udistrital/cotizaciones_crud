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

// Pruebas del controlador FormaPagoController (/v1/forma_pago).

const (
	rutaFormaPago  = "/v1/forma_pago"
	tablaFormaPago = "forma_pago"
)

// cuerpoPostFormaPago retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostFormaPago(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoFormaPago(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestFormaPagoPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaFormaPago, tablaFormaPago, cuerpoPostFormaPago(t))
	id2 := helpers.CrearPorPost(t, rutaFormaPago, tablaFormaPago, cuerpoPostFormaPago(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaFormaPago, id)

	// La llave foránea MstSolicitudCotizacionId es obligatoria
	sinMstSolicitudCotizacionId := cuerpoPostFormaPago(t)
	sinMstSolicitudCotizacionId["MstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaFormaPago+"/", helpers.JSON(t, sinMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaFormaPago)
	t.Log("TestFormaPagoPost Finalizado Correctamente (OK)")
}

func TestFormaPagoGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearFormaPago(t)
	dato := helpers.ProbarGetOne(t, rutaFormaPago, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaFormaPago)
	t.Log("TestFormaPagoGetOne Finalizado Correctamente (OK)")
}

func TestFormaPagoGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearFormaPago(t)
	dato := helpers.ProbarGetAll(t, rutaFormaPago, tablaFormaPago, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaFormaPago, fmt.Sprintf("query=MstSolicitudCotizacionId.Id:%d", creado.MstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MstSolicitudCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestFormaPagoGetAll Finalizado Correctamente (OK)")
}

func TestFormaPagoPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearFormaPago(t)
	helpers.ProbarPut(t, rutaFormaPago, tablaFormaPago, creado.Id, "Condicion", "Prueba-2", "condicion")
	t.Log("TestFormaPagoPut Finalizado Correctamente (OK)")
}

func TestFormaPagoDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearFormaPago(t)
	helpers.ProbarDelete(t, rutaFormaPago, tablaFormaPago, creado.Id)
	t.Log("TestFormaPagoDelete Finalizado Correctamente (OK)")
}

// La llave foránea MstSolicitudCotizacionId debe referenciar un registro existente: POST y PUT
// responden 400 y no crean ni modifican registros.
func TestFormaPagoMstSolicitudCotizacionIdInexistente(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	helpers.ProbarPostFKInexistente(t, rutaFormaPago, tablaFormaPago, cuerpoPostFormaPago(t), "MstSolicitudCotizacionId")
	creado := helpers.CrearFormaPago(t)
	helpers.ProbarPutFKInexistente(t, rutaFormaPago, tablaFormaPago, creado.Id, "MstSolicitudCotizacionId")
	t.Log("TestFormaPagoMstSolicitudCotizacionIdInexistente Finalizado Correctamente (OK)")
}
