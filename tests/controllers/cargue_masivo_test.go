package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaCargueMasivo = "/v1/cargue_masivo"

func TestCargueMasivoPost(t *testing.T) {
	nuevo := helpers.NuevoCargueMasivo(t)
	w := helpers.Peticion("POST", rutaCargueMasivo+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetCargueMasivoById(r.IdDato(t)); err != nil || v.NombreArchivo != nuevo.NombreArchivo || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaCargueMasivo+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestCargueMasivoPost Finalizado Correctamente (OK)")
}

func TestCargueMasivoGetOne(t *testing.T) {
	creado := helpers.CrearCargueMasivo(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaCargueMasivo, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaCargueMasivo, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaCargueMasivo+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestCargueMasivoGetOne Finalizado Correctamente (OK)")
}

func TestCargueMasivoGetAll(t *testing.T) {
	creado := helpers.CrearCargueMasivo(t)
	helpers.ProbarGetAllController(t, rutaCargueMasivo, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaCargueMasivo, creado.SolicitudId, creado.Id)
	t.Log("TestCargueMasivoGetAll Finalizado Correctamente (OK)")
}

func TestCargueMasivoPut(t *testing.T) {
	creado := helpers.CrearCargueMasivo(t)
	ruta := fmt.Sprintf("%s/%d", rutaCargueMasivo, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.NombreArchivo = "items_v2.xlsx"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetCargueMasivoById(creado.Id); err != nil || v.NombreArchivo != "items_v2.xlsx" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaCargueMasivo+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaCargueMasivo, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestCargueMasivoPut Finalizado Correctamente (OK)")
}

func TestCargueMasivoDelete(t *testing.T) {
	creado := helpers.CrearCargueMasivo(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaCargueMasivo, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetCargueMasivoById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaCargueMasivo, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaCargueMasivo+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestCargueMasivoDelete Finalizado Correctamente (OK)")
}
