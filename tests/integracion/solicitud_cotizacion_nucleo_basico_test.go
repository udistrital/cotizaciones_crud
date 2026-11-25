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

// Pruebas del controlador SolicitudCotizacionNucleoBasicoController (/v1/solicitud_cotizacion_nucleo_basico).

const (
	rutaSolicitudCotizacionNucleoBasico  = "/v1/solicitud_cotizacion_nucleo_basico"
	tablaSolicitudCotizacionNucleoBasico = "solicitud_cotizacion_nucleo_basico"
)

// cuerpoPostSolicitudCotizacionNucleoBasico retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostSolicitudCotizacionNucleoBasico(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoSolicitudCotizacionNucleoBasico(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestSolicitudCotizacionNucleoBasicoPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaSolicitudCotizacionNucleoBasico, tablaSolicitudCotizacionNucleoBasico, cuerpoPostSolicitudCotizacionNucleoBasico(t))
	id2 := helpers.CrearPorPost(t, rutaSolicitudCotizacionNucleoBasico, tablaSolicitudCotizacionNucleoBasico, cuerpoPostSolicitudCotizacionNucleoBasico(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaSolicitudCotizacionNucleoBasico, id)

	// La llave foránea MstSolicitudCotizacionId es obligatoria
	sinMstSolicitudCotizacionId := cuerpoPostSolicitudCotizacionNucleoBasico(t)
	sinMstSolicitudCotizacionId["MstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaSolicitudCotizacionNucleoBasico+"/", helpers.JSON(t, sinMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaSolicitudCotizacionNucleoBasico)
	t.Log("TestSolicitudCotizacionNucleoBasicoPost Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionNucleoBasicoGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionNucleoBasico(t)
	dato := helpers.ProbarGetOne(t, rutaSolicitudCotizacionNucleoBasico, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaSolicitudCotizacionNucleoBasico)
	t.Log("TestSolicitudCotizacionNucleoBasicoGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionNucleoBasicoGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionNucleoBasico(t)
	dato := helpers.ProbarGetAll(t, rutaSolicitudCotizacionNucleoBasico, tablaSolicitudCotizacionNucleoBasico, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaSolicitudCotizacionNucleoBasico, fmt.Sprintf("query=MstSolicitudCotizacionId.Id:%d", creado.MstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MstSolicitudCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestSolicitudCotizacionNucleoBasicoGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionNucleoBasicoPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionNucleoBasico(t)
	helpers.ProbarPut(t, rutaSolicitudCotizacionNucleoBasico, tablaSolicitudCotizacionNucleoBasico, creado.Id, "NucleoBasicoId", 2, "nucleo_basico_id")
	t.Log("TestSolicitudCotizacionNucleoBasicoPut Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionNucleoBasicoDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionNucleoBasico(t)
	helpers.ProbarDelete(t, rutaSolicitudCotizacionNucleoBasico, tablaSolicitudCotizacionNucleoBasico, creado.Id)
	t.Log("TestSolicitudCotizacionNucleoBasicoDelete Finalizado Correctamente (OK)")
}
