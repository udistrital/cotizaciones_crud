package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaSolicitudProveedor = "/v1/solicitud_proveedor"

func TestSolicitudProveedorPost(t *testing.T) {
	nuevo := helpers.NuevoSolicitudProveedor(t)
	w := helpers.Peticion("POST", rutaSolicitudProveedor+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetSolicitudProveedorById(r.IdDato(t)); err != nil || v.ProveedorId != nuevo.ProveedorId || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaSolicitudProveedor+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestSolicitudProveedorPost Finalizado Correctamente (OK)")
}

func TestSolicitudProveedorGetOne(t *testing.T) {
	creado := helpers.CrearSolicitudProveedor(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudProveedor, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudProveedor, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudProveedor+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudProveedorGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudProveedorGetAll(t *testing.T) {
	creado := helpers.CrearSolicitudProveedor(t)
	helpers.ProbarGetAllController(t, rutaSolicitudProveedor, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudProveedor, creado.SolicitudId, creado.Id)
	t.Log("TestSolicitudProveedorGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudProveedorPut(t *testing.T) {
	creado := helpers.CrearSolicitudProveedor(t)
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudProveedor, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.ProveedorId = 2
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetSolicitudProveedorById(creado.Id); err != nil || v.ProveedorId != 2 {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaSolicitudProveedor+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudProveedor, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestSolicitudProveedorPut Finalizado Correctamente (OK)")
}

func TestSolicitudProveedorDelete(t *testing.T) {
	creado := helpers.CrearSolicitudProveedor(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudProveedor, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetSolicitudProveedorById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudProveedor, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaSolicitudProveedor+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudProveedorDelete Finalizado Correctamente (OK)")
}
