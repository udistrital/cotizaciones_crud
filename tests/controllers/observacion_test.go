package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaObservacion = "/v1/observacion"

func TestObservacionPost(t *testing.T) {
	nuevo := helpers.NuevoObservacion(t)
	w := helpers.Peticion("POST", rutaObservacion+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetObservacionById(r.IdDato(t)); err != nil || v.Texto != nuevo.Texto || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaObservacion+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestObservacionPost Finalizado Correctamente (OK)")
}

func TestObservacionGetOne(t *testing.T) {
	creado := helpers.CrearObservacion(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaObservacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaObservacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaObservacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestObservacionGetOne Finalizado Correctamente (OK)")
}

func TestObservacionGetAll(t *testing.T) {
	creado := helpers.CrearObservacion(t)
	helpers.CrearObservacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaObservacion, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaObservacion, creado.SolicitudId, creado.Id)
	t.Log("TestObservacionGetAll Finalizado Correctamente (OK)")
}

func TestObservacionActivo(t *testing.T) {
	activo := helpers.CrearObservacion(t)
	inactivo := helpers.CrearObservacion(t)
	if err := models.DeleteObservacion(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaObservacion, activo.Id, inactivo.Id)
	t.Log("TestObservacionActivo Finalizado Correctamente (OK)")
}

func TestObservacionPut(t *testing.T) {
	creado := helpers.CrearObservacion(t)
	ruta := fmt.Sprintf("%s/%d", rutaObservacion, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Texto = "Texto actualizado"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetObservacionById(creado.Id); err != nil || v.Texto != "Texto actualizado" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaObservacion+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaObservacion, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestObservacionPut Finalizado Correctamente (OK)")
}

func TestObservacionDelete(t *testing.T) {
	creado := helpers.CrearObservacion(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaObservacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetObservacionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaObservacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaObservacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestObservacionDelete Finalizado Correctamente (OK)")
}
