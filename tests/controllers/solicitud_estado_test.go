package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaSolicitudEstado = "/v1/solicitud_estado"

func TestSolicitudEstadoPost(t *testing.T) {
	nuevo := helpers.NuevoSolicitudEstado(t)
	w := helpers.Peticion("POST", rutaSolicitudEstado+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetSolicitudEstadoById(r.IdDato(t)); err != nil || v.Justificacion != nuevo.Justificacion || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaSolicitudEstado+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestSolicitudEstadoPost Finalizado Correctamente (OK)")
}

func TestSolicitudEstadoGetOne(t *testing.T) {
	creado := helpers.CrearSolicitudEstado(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudEstado, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudEstado, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudEstado+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudEstadoGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudEstadoGetAll(t *testing.T) {
	creado := helpers.CrearSolicitudEstado(t)
	helpers.CrearSolicitudEstado(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaSolicitudEstado, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudEstado, creado.SolicitudId, creado.Id)
	t.Log("TestSolicitudEstadoGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudEstadoActivo(t *testing.T) {
	activo := helpers.CrearSolicitudEstado(t)
	inactivo := helpers.CrearSolicitudEstado(t)
	if err := models.DeleteSolicitudEstado(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaSolicitudEstado, activo.Id, inactivo.Id)
	t.Log("TestSolicitudEstadoActivo Finalizado Correctamente (OK)")
}

func TestSolicitudEstadoPut(t *testing.T) {
	creado := helpers.CrearSolicitudEstado(t)
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudEstado, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Justificacion = "Envio a proveedores"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetSolicitudEstadoById(creado.Id); err != nil || v.Justificacion != "Envio a proveedores" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaSolicitudEstado+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudEstado, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestSolicitudEstadoPut Finalizado Correctamente (OK)")
}

func TestSolicitudEstadoDelete(t *testing.T) {
	creado := helpers.CrearSolicitudEstado(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudEstado, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetSolicitudEstadoById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudEstado, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaSolicitudEstado+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudEstadoDelete Finalizado Correctamente (OK)")
}
