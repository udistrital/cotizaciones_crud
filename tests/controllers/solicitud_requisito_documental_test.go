package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaSolicitudRequisitoDocumental = "/v1/solicitud_requisito_documental"

func TestSolicitudRequisitoDocumentalPost(t *testing.T) {
	nuevo := helpers.NuevoSolicitudRequisitoDocumental(t)
	w := helpers.Peticion("POST", rutaSolicitudRequisitoDocumental+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetSolicitudRequisitoDocumentalById(r.IdDato(t)); err != nil || v.Instruccion != nuevo.Instruccion || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaSolicitudRequisitoDocumental+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestSolicitudRequisitoDocumentalPost Finalizado Correctamente (OK)")
}

func TestSolicitudRequisitoDocumentalGetOne(t *testing.T) {
	creado := helpers.CrearSolicitudRequisitoDocumental(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudRequisitoDocumental, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudRequisitoDocumental, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudRequisitoDocumental+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudRequisitoDocumentalGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudRequisitoDocumentalGetAll(t *testing.T) {
	creado := helpers.CrearSolicitudRequisitoDocumental(t)
	helpers.CrearSolicitudRequisitoDocumental(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaSolicitudRequisitoDocumental, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudRequisitoDocumental, creado.SolicitudId, creado.Id)
	t.Log("TestSolicitudRequisitoDocumentalGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudRequisitoDocumentalActivo(t *testing.T) {
	activo := helpers.CrearSolicitudRequisitoDocumental(t)
	inactivo := helpers.CrearSolicitudRequisitoDocumental(t)
	if err := models.DeleteSolicitudRequisitoDocumental(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaSolicitudRequisitoDocumental, activo.Id, inactivo.Id)
	t.Log("TestSolicitudRequisitoDocumentalActivo Finalizado Correctamente (OK)")
}

func TestSolicitudRequisitoDocumentalPut(t *testing.T) {
	creado := helpers.CrearSolicitudRequisitoDocumental(t)
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudRequisitoDocumental, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Instruccion = "Adjuntar camara de comercio"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetSolicitudRequisitoDocumentalById(creado.Id); err != nil || v.Instruccion != "Adjuntar camara de comercio" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaSolicitudRequisitoDocumental+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudRequisitoDocumental, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestSolicitudRequisitoDocumentalPut Finalizado Correctamente (OK)")
}

func TestSolicitudRequisitoDocumentalDelete(t *testing.T) {
	creado := helpers.CrearSolicitudRequisitoDocumental(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudRequisitoDocumental, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetSolicitudRequisitoDocumentalById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudRequisitoDocumental, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaSolicitudRequisitoDocumental+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudRequisitoDocumentalDelete Finalizado Correctamente (OK)")
}
