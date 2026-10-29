package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaSolicitudModificacion = "/v1/solicitud_modificacion"

func TestSolicitudModificacionPost(t *testing.T) {
	nuevo := helpers.NuevoSolicitudModificacion(t)
	w := helpers.Peticion("POST", rutaSolicitudModificacion+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetSolicitudModificacionById(r.IdDato(t)); err != nil || v.Numero != nuevo.Numero || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaSolicitudModificacion+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestSolicitudModificacionPost Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionGetOne(t *testing.T) {
	creado := helpers.CrearSolicitudModificacion(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudModificacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudModificacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudModificacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudModificacionGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionGetAll(t *testing.T) {
	creado := helpers.CrearSolicitudModificacion(t)
	helpers.CrearSolicitudModificacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaSolicitudModificacion, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudModificacion, creado.SolicitudId, creado.Id)
	t.Log("TestSolicitudModificacionGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionActivo(t *testing.T) {
	activo := helpers.CrearSolicitudModificacion(t)
	inactivo := helpers.CrearSolicitudModificacion(t)
	if err := models.DeleteSolicitudModificacion(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaSolicitudModificacion, activo.Id, inactivo.Id)
	t.Log("TestSolicitudModificacionActivo Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionPut(t *testing.T) {
	creado := helpers.CrearSolicitudModificacion(t)
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudModificacion, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Numero = 2
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetSolicitudModificacionById(creado.Id); err != nil || v.Numero != 2 {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaSolicitudModificacion+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudModificacion, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestSolicitudModificacionPut Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionDelete(t *testing.T) {
	creado := helpers.CrearSolicitudModificacion(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudModificacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetSolicitudModificacionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudModificacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaSolicitudModificacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudModificacionDelete Finalizado Correctamente (OK)")
}
