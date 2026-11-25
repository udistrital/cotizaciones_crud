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

// Pruebas del controlador HistDetSolicitudCotizacionController (/v1/hist_det_solicitud_cotizacion).

const (
	rutaHistDetSolicitudCotizacion  = "/v1/hist_det_solicitud_cotizacion"
	tablaHistDetSolicitudCotizacion = "hist_det_solicitud_cotizacion"
)

// cuerpoPostHistDetSolicitudCotizacion retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostHistDetSolicitudCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoHistDetSolicitudCotizacion(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestHistDetSolicitudCotizacionPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaHistDetSolicitudCotizacion, tablaHistDetSolicitudCotizacion, cuerpoPostHistDetSolicitudCotizacion(t))
	id2 := helpers.CrearPorPost(t, rutaHistDetSolicitudCotizacion, tablaHistDetSolicitudCotizacion, cuerpoPostHistDetSolicitudCotizacion(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaHistDetSolicitudCotizacion, id)

	// La llave foránea MstSolicitudCotizacionId es obligatoria
	sinMstSolicitudCotizacionId := cuerpoPostHistDetSolicitudCotizacion(t)
	sinMstSolicitudCotizacionId["MstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaHistDetSolicitudCotizacion+"/", helpers.JSON(t, sinMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaHistDetSolicitudCotizacion)
	t.Log("TestHistDetSolicitudCotizacionPost Finalizado Correctamente (OK)")
}

func TestHistDetSolicitudCotizacionGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistDetSolicitudCotizacion(t)
	dato := helpers.ProbarGetOne(t, rutaHistDetSolicitudCotizacion, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaHistDetSolicitudCotizacion)
	t.Log("TestHistDetSolicitudCotizacionGetOne Finalizado Correctamente (OK)")
}

func TestHistDetSolicitudCotizacionGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistDetSolicitudCotizacion(t)
	dato := helpers.ProbarGetAll(t, rutaHistDetSolicitudCotizacion, tablaHistDetSolicitudCotizacion, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaHistDetSolicitudCotizacion, fmt.Sprintf("query=MstSolicitudCotizacionId.Id:%d", creado.MstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MstSolicitudCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestHistDetSolicitudCotizacionGetAll Finalizado Correctamente (OK)")
}

func TestHistDetSolicitudCotizacionPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistDetSolicitudCotizacion(t)
	helpers.ProbarPut(t, rutaHistDetSolicitudCotizacion, tablaHistDetSolicitudCotizacion, creado.Id, "Nombre", "Prueba-2", "nombre")
	t.Log("TestHistDetSolicitudCotizacionPut Finalizado Correctamente (OK)")
}

func TestHistDetSolicitudCotizacionDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistDetSolicitudCotizacion(t)
	helpers.ProbarDelete(t, rutaHistDetSolicitudCotizacion, tablaHistDetSolicitudCotizacion, creado.Id)
	t.Log("TestHistDetSolicitudCotizacionDelete Finalizado Correctamente (OK)")
}
