package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddRespuestaObservacion(t *testing.T) {
	v := helpers.NuevoRespuestaObservacion(t)
	v.Activo = false
	id, err := models.AddRespuestaObservacion(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddRespuestaObservacion Finalizado Correctamente (OK)")
}

func TestGetRespuestaObservacionById(t *testing.T) {
	creado := helpers.CrearRespuestaObservacion(t)
	v, err := models.GetRespuestaObservacionById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Texto != creado.Texto {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetRespuestaObservacionById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetRespuestaObservacionById Finalizado Correctamente (OK)")
}

func TestGetAllRespuestaObservacion(t *testing.T) {
	creado := helpers.CrearRespuestaObservacion(t)
	helpers.CrearRespuestaObservacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllRespuestaObservacion, creado.Id)
	t.Log("TestGetAllRespuestaObservacion Finalizado Correctamente (OK)")
}

func TestUpdateRespuestaObservacionById(t *testing.T) {
	creado := helpers.CrearRespuestaObservacion(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Texto = "Respuesta actualizada"
	if err := models.UpdateRespuestaObservacionById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetRespuestaObservacionById(creado.Id)
	if v == nil || v.Texto != "Respuesta actualizada" {
		t.Fatal("Se esperaba Texto actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateRespuestaObservacionById(&models.RespuestaObservacion{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateRespuestaObservacionById Finalizado Correctamente (OK)")
}

func TestDeleteRespuestaObservacion(t *testing.T) {
	creado := helpers.CrearRespuestaObservacion(t)
	if err := models.DeleteRespuestaObservacion(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetRespuestaObservacionById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteRespuestaObservacion(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteRespuestaObservacion Finalizado Correctamente (OK)")
}
