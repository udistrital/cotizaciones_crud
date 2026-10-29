package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaRespuestaObservacion = "/v1/respuesta_observacion"

func TestRespuestaObservacionPost(t *testing.T) {
	nuevo := helpers.NuevoRespuestaObservacion(t)
	w := helpers.Peticion("POST", rutaRespuestaObservacion+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetRespuestaObservacionById(r.IdDato(t)); err != nil || v.Texto != nuevo.Texto || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaRespuestaObservacion+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestRespuestaObservacionPost Finalizado Correctamente (OK)")
}

func TestRespuestaObservacionGetOne(t *testing.T) {
	creado := helpers.CrearRespuestaObservacion(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaRespuestaObservacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaRespuestaObservacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaRespuestaObservacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestRespuestaObservacionGetOne Finalizado Correctamente (OK)")
}

func TestRespuestaObservacionGetAll(t *testing.T) {
	creado := helpers.CrearRespuestaObservacion(t)
	helpers.CrearRespuestaObservacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaRespuestaObservacion, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaRespuestaObservacion, creado.ObservacionId.SolicitudId, creado.Id)
	t.Log("TestRespuestaObservacionGetAll Finalizado Correctamente (OK)")
}

func TestRespuestaObservacionActivo(t *testing.T) {
	activo := helpers.CrearRespuestaObservacion(t)
	inactivo := helpers.CrearRespuestaObservacion(t)
	if err := models.DeleteRespuestaObservacion(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaRespuestaObservacion, activo.Id, inactivo.Id)
	t.Log("TestRespuestaObservacionActivo Finalizado Correctamente (OK)")
}

func TestRespuestaObservacionPut(t *testing.T) {
	creado := helpers.CrearRespuestaObservacion(t)
	ruta := fmt.Sprintf("%s/%d", rutaRespuestaObservacion, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Texto = "Respuesta actualizada"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetRespuestaObservacionById(creado.Id); err != nil || v.Texto != "Respuesta actualizada" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaRespuestaObservacion+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaRespuestaObservacion, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestRespuestaObservacionPut Finalizado Correctamente (OK)")
}

func TestRespuestaObservacionDelete(t *testing.T) {
	creado := helpers.CrearRespuestaObservacion(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaRespuestaObservacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetRespuestaObservacionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaRespuestaObservacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaRespuestaObservacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestRespuestaObservacionDelete Finalizado Correctamente (OK)")
}
