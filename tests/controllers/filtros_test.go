package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Cada controlador prueba los filtros en su Test<Modelo>GetAll. Aquí se prueban
// las combinaciones con el parámetro query y el camino de dos niveles
// (cargue_masivo_error -> cargue_masivo -> solicitud_cotizacion).

func TestFiltrosConQuery(t *testing.T) {
	creado := helpers.CrearSolicitudItem(t)
	solicitud := creado.SolicitudId

	w := helpers.Peticion("GET", fmt.Sprintf("%s/?vigencia=%d&query=Nombre:%s", rutaSolicitudItem, solicitud.Vigencia, "Computador%20de%20escritorio"), "")
	if l := helpers.VerificarRespuesta(t, w, http.StatusOK).Lista(t); len(l) != 1 {
		t.Error("Se esperaba 1 registro al combinar vigencia y query y se obtuvo", l)
	}

	w = helpers.Peticion("GET", fmt.Sprintf("%s/?vigencia=%d&query=Nombre:Otro", rutaSolicitudItem, solicitud.Vigencia), "")
	l := helpers.VerificarRespuesta(t, w, http.StatusOK).Lista(t)
	if m, ok := l[0].(map[string]interface{}); len(l) != 1 || !ok || len(m) != 0 {
		t.Error("Se esperaba una lista con un objeto vacío y se obtuvo", l)
	}
	t.Log("TestFiltrosConQuery Finalizado Correctamente (OK)")
}

func TestFiltrosDosNiveles(t *testing.T) {
	creado := helpers.CrearCargueMasivoError(t)
	solicitud := creado.CargueMasivoId.SolicitudId

	w := helpers.Peticion("GET", fmt.Sprintf("%s/?dependencia=%d", rutaCargueMasivoError, solicitud.DependenciaSolicitanteId), "")
	l := helpers.VerificarRespuesta(t, w, http.StatusOK).Lista(t)
	if m, ok := l[0].(map[string]interface{}); len(l) != 1 || !ok || m["Id"] != float64(creado.Id) {
		t.Error("Se esperaba el error de cargue", creado.Id, "y se obtuvo", l)
	}
	t.Log("TestFiltrosDosNiveles Finalizado Correctamente (OK)")
}
