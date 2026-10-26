package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaSolicitudCiiu = "/v1/solicitud_ciiu"

func TestSolicitudCiiuPost(t *testing.T) {
	nuevo := helpers.NuevoSolicitudCiiu(t)
	w := helpers.Peticion("POST", rutaSolicitudCiiu+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetSolicitudCiiuById(r.IdDato(t)); err != nil || v.ActividadEconomicaId != nuevo.ActividadEconomicaId || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaSolicitudCiiu+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestSolicitudCiiuPost Finalizado Correctamente (OK)")
}

func TestSolicitudCiiuGetOne(t *testing.T) {
	creado := helpers.CrearSolicitudCiiu(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudCiiu, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudCiiu, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudCiiu+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudCiiuGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudCiiuGetAll(t *testing.T) {
	creado := helpers.CrearSolicitudCiiu(t)
	helpers.ProbarGetAllController(t, rutaSolicitudCiiu, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudCiiu, creado.SolicitudId, creado.Id)
	t.Log("TestSolicitudCiiuGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudCiiuPut(t *testing.T) {
	creado := helpers.CrearSolicitudCiiu(t)
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudCiiu, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.ActividadEconomicaId = 2
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetSolicitudCiiuById(creado.Id); err != nil || v.ActividadEconomicaId != 2 {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaSolicitudCiiu+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudCiiu, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestSolicitudCiiuPut Finalizado Correctamente (OK)")
}

func TestSolicitudCiiuDelete(t *testing.T) {
	creado := helpers.CrearSolicitudCiiu(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudCiiu, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetSolicitudCiiuById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudCiiu, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaSolicitudCiiu+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudCiiuDelete Finalizado Correctamente (OK)")
}
