package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaSolicitudModificacionDetalle = "/v1/solicitud_modificacion_detalle"

func TestSolicitudModificacionDetallePost(t *testing.T) {
	nuevo := helpers.NuevoSolicitudModificacionDetalle(t)
	w := helpers.Peticion("POST", rutaSolicitudModificacionDetalle+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetSolicitudModificacionDetalleById(r.IdDato(t)); err != nil || v.Justificacion != nuevo.Justificacion || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaSolicitudModificacionDetalle+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestSolicitudModificacionDetallePost Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionDetalleGetOne(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionDetalle(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudModificacionDetalle, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudModificacionDetalle, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudModificacionDetalle+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudModificacionDetalleGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionDetalleGetAll(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionDetalle(t)
	helpers.CrearSolicitudModificacionDetalle(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaSolicitudModificacionDetalle, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudModificacionDetalle, creado.SolicitudModificacionId.SolicitudId, creado.Id)
	t.Log("TestSolicitudModificacionDetalleGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionDetalleActivo(t *testing.T) {
	activo := helpers.CrearSolicitudModificacionDetalle(t)
	inactivo := helpers.CrearSolicitudModificacionDetalle(t)
	if err := models.DeleteSolicitudModificacionDetalle(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaSolicitudModificacionDetalle, activo.Id, inactivo.Id)
	t.Log("TestSolicitudModificacionDetalleActivo Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionDetallePut(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionDetalle(t)
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudModificacionDetalle, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Justificacion = "Ampliacion de plazo"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetSolicitudModificacionDetalleById(creado.Id); err != nil || v.Justificacion != "Ampliacion de plazo" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaSolicitudModificacionDetalle+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudModificacionDetalle, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestSolicitudModificacionDetallePut Finalizado Correctamente (OK)")
}

func TestSolicitudModificacionDetalleDelete(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionDetalle(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudModificacionDetalle, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetSolicitudModificacionDetalleById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudModificacionDetalle, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaSolicitudModificacionDetalle+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudModificacionDetalleDelete Finalizado Correctamente (OK)")
}
