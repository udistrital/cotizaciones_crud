package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaCargueMasivoError = "/v1/cargue_masivo_error"

func TestCargueMasivoErrorPost(t *testing.T) {
	nuevo := helpers.NuevoCargueMasivoError(t)
	w := helpers.Peticion("POST", rutaCargueMasivoError+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetCargueMasivoErrorById(r.IdDato(t)); err != nil || v.Detalle != nuevo.Detalle || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaCargueMasivoError+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestCargueMasivoErrorPost Finalizado Correctamente (OK)")
}

func TestCargueMasivoErrorGetOne(t *testing.T) {
	creado := helpers.CrearCargueMasivoError(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaCargueMasivoError, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaCargueMasivoError, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaCargueMasivoError+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestCargueMasivoErrorGetOne Finalizado Correctamente (OK)")
}

func TestCargueMasivoErrorGetAll(t *testing.T) {
	creado := helpers.CrearCargueMasivoError(t)
	helpers.ProbarGetAllController(t, rutaCargueMasivoError, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaCargueMasivoError, creado.CargueMasivoId.SolicitudId, creado.Id)
	t.Log("TestCargueMasivoErrorGetAll Finalizado Correctamente (OK)")
}

func TestCargueMasivoErrorPut(t *testing.T) {
	creado := helpers.CrearCargueMasivoError(t)
	ruta := fmt.Sprintf("%s/%d", rutaCargueMasivoError, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Detalle = "Detalle actualizado"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetCargueMasivoErrorById(creado.Id); err != nil || v.Detalle != "Detalle actualizado" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaCargueMasivoError+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaCargueMasivoError, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestCargueMasivoErrorPut Finalizado Correctamente (OK)")
}

func TestCargueMasivoErrorDelete(t *testing.T) {
	creado := helpers.CrearCargueMasivoError(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaCargueMasivoError, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetCargueMasivoErrorById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaCargueMasivoError, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaCargueMasivoError+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestCargueMasivoErrorDelete Finalizado Correctamente (OK)")
}
