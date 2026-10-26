package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaSolicitudPagoParcial = "/v1/solicitud_pago_parcial"

func TestSolicitudPagoParcialPost(t *testing.T) {
	nuevo := helpers.NuevoSolicitudPagoParcial(t)
	w := helpers.Peticion("POST", rutaSolicitudPagoParcial+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetSolicitudPagoParcialById(r.IdDato(t)); err != nil || v.Condicion != nuevo.Condicion || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaSolicitudPagoParcial+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestSolicitudPagoParcialPost Finalizado Correctamente (OK)")
}

func TestSolicitudPagoParcialGetOne(t *testing.T) {
	creado := helpers.CrearSolicitudPagoParcial(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudPagoParcial, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudPagoParcial, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudPagoParcial+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudPagoParcialGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudPagoParcialGetAll(t *testing.T) {
	creado := helpers.CrearSolicitudPagoParcial(t)
	helpers.ProbarGetAllController(t, rutaSolicitudPagoParcial, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudPagoParcial, creado.SolicitudId, creado.Id)
	t.Log("TestSolicitudPagoParcialGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudPagoParcialPut(t *testing.T) {
	creado := helpers.CrearSolicitudPagoParcial(t)
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudPagoParcial, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Condicion = "Contra entrega total"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetSolicitudPagoParcialById(creado.Id); err != nil || v.Condicion != "Contra entrega total" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaSolicitudPagoParcial+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudPagoParcial, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestSolicitudPagoParcialPut Finalizado Correctamente (OK)")
}

func TestSolicitudPagoParcialDelete(t *testing.T) {
	creado := helpers.CrearSolicitudPagoParcial(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudPagoParcial, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetSolicitudPagoParcialById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudPagoParcial, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaSolicitudPagoParcial+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudPagoParcialDelete Finalizado Correctamente (OK)")
}
