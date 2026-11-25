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

// Pruebas del controlador CargueMasivoController (/v1/cargue_masivo).

const (
	rutaCargueMasivo  = "/v1/cargue_masivo"
	tablaCargueMasivo = "cargue_masivo"
)

// cuerpoPostCargueMasivo retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostCargueMasivo(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoCargueMasivo(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestCargueMasivoPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpoPostCargueMasivo(t))
	id2 := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpoPostCargueMasivo(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaCargueMasivo, id)

	// La llave foránea HistMstSolicitudCotizacionId es obligatoria
	sinHistMstSolicitudCotizacionId := cuerpoPostCargueMasivo(t)
	sinHistMstSolicitudCotizacionId["HistMstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaCargueMasivo+"/", helpers.JSON(t, sinHistMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaCargueMasivo)
	t.Log("TestCargueMasivoPost Finalizado Correctamente (OK)")
}

func TestCargueMasivoGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCargueMasivo(t)
	dato := helpers.ProbarGetOne(t, rutaCargueMasivo, creado.Id)

	// La llave foránea HistMstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetHistMstSolicitudCotizacionById(creado.HistMstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "HistMstSolicitudCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaCargueMasivo)
	t.Log("TestCargueMasivoGetOne Finalizado Correctamente (OK)")
}

func TestCargueMasivoGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCargueMasivo(t)
	dato := helpers.ProbarGetAll(t, rutaCargueMasivo, tablaCargueMasivo, creado.Id)

	// La llave foránea HistMstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetHistMstSolicitudCotizacionById(creado.HistMstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "HistMstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaCargueMasivo, fmt.Sprintf("query=HistMstSolicitudCotizacionId.Id:%d", creado.HistMstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por HistMstSolicitudCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestCargueMasivoGetAll Finalizado Correctamente (OK)")
}

func TestCargueMasivoPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCargueMasivo(t)
	helpers.ProbarPut(t, rutaCargueMasivo, tablaCargueMasivo, creado.Id, "EnlacePlantilla", "Prueba-2", "enlace_plantilla")
	t.Log("TestCargueMasivoPut Finalizado Correctamente (OK)")
}

func TestCargueMasivoDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCargueMasivo(t)
	helpers.ProbarDelete(t, rutaCargueMasivo, tablaCargueMasivo, creado.Id)
	t.Log("TestCargueMasivoDelete Finalizado Correctamente (OK)")
}
