package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaSolicitudModificacionEstado = "/v1/solicitud_modificacion_estado"

func TestSolicitudModificacionEstadoPost(t *testing.T) {
	nuevo := helpers.NuevoSolicitudModificacionEstado(t)
	w := helpers.Peticion("POST", rutaSolicitudModificacionEstado+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetSolicitudModificacionEstadoById(r.IdDato(t)); err != nil || v.Justificacion != nuevo.Justificacion || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaSolicitudModificacionEstado+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestSolicitudModificacionEstadoPost Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionEstadoGetOne(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionEstado(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudModificacionEstado, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudModificacionEstado, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudModificacionEstado+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudModificacionEstadoGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionEstadoGetAll(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionEstado(t)
	helpers.CrearSolicitudModificacionEstado(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaSolicitudModificacionEstado, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudModificacionEstado, creado.SolicitudModificacionId.SolicitudId, creado.Id)
	t.Log("TestSolicitudModificacionEstadoGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionEstadoActivo(t *testing.T) {
	activo := helpers.CrearSolicitudModificacionEstado(t)
	inactivo := helpers.CrearSolicitudModificacionEstado(t)
	if err := models.DeleteSolicitudModificacionEstado(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaSolicitudModificacionEstado, activo.Id, inactivo.Id)
	t.Log("TestSolicitudModificacionEstadoActivo Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionEstadoPut(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionEstado(t)
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudModificacionEstado, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Justificacion = "Se aprueba la modificacion"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetSolicitudModificacionEstadoById(creado.Id); err != nil || v.Justificacion != "Se aprueba la modificacion" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaSolicitudModificacionEstado+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudModificacionEstado, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestSolicitudModificacionEstadoPut Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionEstadoDelete(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionEstado(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudModificacionEstado, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetSolicitudModificacionEstadoById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudModificacionEstado, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaSolicitudModificacionEstado+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudModificacionEstadoDelete Finalizado Correctamente (OK)")
}
