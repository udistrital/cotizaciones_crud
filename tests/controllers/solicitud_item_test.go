package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaSolicitudItem = "/v1/solicitud_item"

func TestSolicitudItemPost(t *testing.T) {
	nuevo := helpers.NuevoSolicitudItem(t)
	w := helpers.Peticion("POST", rutaSolicitudItem+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetSolicitudItemById(r.IdDato(t)); err != nil || v.Nombre != nuevo.Nombre || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaSolicitudItem+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestSolicitudItemPost Finalizado Correctamente (OK)")
}

func TestSolicitudItemGetOne(t *testing.T) {
	creado := helpers.CrearSolicitudItem(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudItem, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudItem, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudItem+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudItemGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudItemGetAll(t *testing.T) {
	creado := helpers.CrearSolicitudItem(t)
	helpers.ProbarGetAllController(t, rutaSolicitudItem, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudItem, creado.SolicitudId, creado.Id)
	t.Log("TestSolicitudItemGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudItemPut(t *testing.T) {
	creado := helpers.CrearSolicitudItem(t)
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudItem, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Nombre = "Portatil"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetSolicitudItemById(creado.Id); err != nil || v.Nombre != "Portatil" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaSolicitudItem+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudItem, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestSolicitudItemPut Finalizado Correctamente (OK)")
}

func TestSolicitudItemDelete(t *testing.T) {
	creado := helpers.CrearSolicitudItem(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudItem, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetSolicitudItemById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudItem, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaSolicitudItem+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudItemDelete Finalizado Correctamente (OK)")
}
