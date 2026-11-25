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

// Pruebas del controlador SolicitudCotizacionActividadEconomicaController (/v1/solicitud_cotizacion_actividad_economica).

const (
	rutaSolicitudCotizacionActividadEconomica  = "/v1/solicitud_cotizacion_actividad_economica"
	tablaSolicitudCotizacionActividadEconomica = "solicitud_cotizacion_actividad_economica"
)

// cuerpoPostSolicitudCotizacionActividadEconomica retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostSolicitudCotizacionActividadEconomica(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoSolicitudCotizacionActividadEconomica(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestSolicitudCotizacionActividadEconomicaPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaSolicitudCotizacionActividadEconomica, tablaSolicitudCotizacionActividadEconomica, cuerpoPostSolicitudCotizacionActividadEconomica(t))
	id2 := helpers.CrearPorPost(t, rutaSolicitudCotizacionActividadEconomica, tablaSolicitudCotizacionActividadEconomica, cuerpoPostSolicitudCotizacionActividadEconomica(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaSolicitudCotizacionActividadEconomica, id)

	// La llave foránea MstSolicitudCotizacionId es obligatoria
	sinMstSolicitudCotizacionId := cuerpoPostSolicitudCotizacionActividadEconomica(t)
	sinMstSolicitudCotizacionId["MstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaSolicitudCotizacionActividadEconomica+"/", helpers.JSON(t, sinMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaSolicitudCotizacionActividadEconomica)
	t.Log("TestSolicitudCotizacionActividadEconomicaPost Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionActividadEconomicaGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionActividadEconomica(t)
	dato := helpers.ProbarGetOne(t, rutaSolicitudCotizacionActividadEconomica, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaSolicitudCotizacionActividadEconomica)
	t.Log("TestSolicitudCotizacionActividadEconomicaGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionActividadEconomicaGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionActividadEconomica(t)
	dato := helpers.ProbarGetAll(t, rutaSolicitudCotizacionActividadEconomica, tablaSolicitudCotizacionActividadEconomica, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaSolicitudCotizacionActividadEconomica, fmt.Sprintf("query=MstSolicitudCotizacionId.Id:%d", creado.MstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MstSolicitudCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestSolicitudCotizacionActividadEconomicaGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionActividadEconomicaPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionActividadEconomica(t)
	helpers.ProbarPut(t, rutaSolicitudCotizacionActividadEconomica, tablaSolicitudCotizacionActividadEconomica, creado.Id, "ActividadEconomicaId", 2, "actividad_economica_id")
	t.Log("TestSolicitudCotizacionActividadEconomicaPut Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionActividadEconomicaDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionActividadEconomica(t)
	helpers.ProbarDelete(t, rutaSolicitudCotizacionActividadEconomica, tablaSolicitudCotizacionActividadEconomica, creado.Id)
	t.Log("TestSolicitudCotizacionActividadEconomicaDelete Finalizado Correctamente (OK)")
}
