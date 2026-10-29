package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaDocumentoSolicitud = "/v1/documento_solicitud"

func TestDocumentoSolicitudPost(t *testing.T) {
	nuevo := helpers.NuevoDocumentoSolicitud(t)
	w := helpers.Peticion("POST", rutaDocumentoSolicitud+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetDocumentoSolicitudById(r.IdDato(t)); err != nil || v.NombreArchivo != nuevo.NombreArchivo || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaDocumentoSolicitud+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestDocumentoSolicitudPost Finalizado Correctamente (OK)")
}

func TestDocumentoSolicitudGetOne(t *testing.T) {
	creado := helpers.CrearDocumentoSolicitud(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaDocumentoSolicitud, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaDocumentoSolicitud, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaDocumentoSolicitud+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestDocumentoSolicitudGetOne Finalizado Correctamente (OK)")
}

func TestDocumentoSolicitudGetAll(t *testing.T) {
	creado := helpers.CrearDocumentoSolicitud(t)
	helpers.CrearDocumentoSolicitud(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaDocumentoSolicitud, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaDocumentoSolicitud, creado.SolicitudId, creado.Id)
	t.Log("TestDocumentoSolicitudGetAll Finalizado Correctamente (OK)")
}

func TestDocumentoSolicitudActivo(t *testing.T) {
	activo := helpers.CrearDocumentoSolicitud(t)
	inactivo := helpers.CrearDocumentoSolicitud(t)
	if err := models.DeleteDocumentoSolicitud(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaDocumentoSolicitud, activo.Id, inactivo.Id)
	t.Log("TestDocumentoSolicitudActivo Finalizado Correctamente (OK)")
}

func TestDocumentoSolicitudPut(t *testing.T) {
	creado := helpers.CrearDocumentoSolicitud(t)
	ruta := fmt.Sprintf("%s/%d", rutaDocumentoSolicitud, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.NombreArchivo = "anexo_tecnico.pdf"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetDocumentoSolicitudById(creado.Id); err != nil || v.NombreArchivo != "anexo_tecnico.pdf" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaDocumentoSolicitud+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaDocumentoSolicitud, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestDocumentoSolicitudPut Finalizado Correctamente (OK)")
}

func TestDocumentoSolicitudDelete(t *testing.T) {
	creado := helpers.CrearDocumentoSolicitud(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaDocumentoSolicitud, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetDocumentoSolicitudById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaDocumentoSolicitud, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaDocumentoSolicitud+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestDocumentoSolicitudDelete Finalizado Correctamente (OK)")
}
