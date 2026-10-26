package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaSolicitudNecesidad = "/v1/solicitud_necesidad"

func TestSolicitudNecesidadPost(t *testing.T) {
	nuevo := helpers.NuevoSolicitudNecesidad(t)
	w := helpers.Peticion("POST", rutaSolicitudNecesidad+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetSolicitudNecesidadById(r.IdDato(t)); err != nil || v.RubroNombre != nuevo.RubroNombre || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaSolicitudNecesidad+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestSolicitudNecesidadPost Finalizado Correctamente (OK)")
}

func TestSolicitudNecesidadGetOne(t *testing.T) {
	creado := helpers.CrearSolicitudNecesidad(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudNecesidad, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudNecesidad, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudNecesidad+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudNecesidadGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudNecesidadGetAll(t *testing.T) {
	creado := helpers.CrearSolicitudNecesidad(t)
	helpers.ProbarGetAllController(t, rutaSolicitudNecesidad, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudNecesidad, creado.SolicitudId, creado.Id)
	t.Log("TestSolicitudNecesidadGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudNecesidadPut(t *testing.T) {
	creado := helpers.CrearSolicitudNecesidad(t)
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudNecesidad, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.RubroNombre = "Rubro actualizado"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetSolicitudNecesidadById(creado.Id); err != nil || v.RubroNombre != "Rubro actualizado" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaSolicitudNecesidad+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudNecesidad, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestSolicitudNecesidadPut Finalizado Correctamente (OK)")
}

func TestSolicitudNecesidadDelete(t *testing.T) {
	creado := helpers.CrearSolicitudNecesidad(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudNecesidad, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetSolicitudNecesidadById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudNecesidad, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaSolicitudNecesidad+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudNecesidadDelete Finalizado Correctamente (OK)")
}
