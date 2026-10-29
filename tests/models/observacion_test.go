package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddObservacion(t *testing.T) {
	v := helpers.NuevoObservacion(t)
	v.Activo = false
	id, err := models.AddObservacion(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddObservacion Finalizado Correctamente (OK)")
}

func TestGetObservacionById(t *testing.T) {
	creado := helpers.CrearObservacion(t)
	v, err := models.GetObservacionById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Texto != creado.Texto {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetObservacionById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetObservacionById Finalizado Correctamente (OK)")
}

func TestGetAllObservacion(t *testing.T) {
	creado := helpers.CrearObservacion(t)
	helpers.CrearObservacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllObservacion, creado.Id)
	t.Log("TestGetAllObservacion Finalizado Correctamente (OK)")
}

func TestUpdateObservacionById(t *testing.T) {
	creado := helpers.CrearObservacion(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Texto = "Texto actualizado"
	if err := models.UpdateObservacionById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetObservacionById(creado.Id)
	if v == nil || v.Texto != "Texto actualizado" {
		t.Fatal("Se esperaba Texto actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateObservacionById(&models.Observacion{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateObservacionById Finalizado Correctamente (OK)")
}

func TestDeleteObservacion(t *testing.T) {
	creado := helpers.CrearObservacion(t)
	if err := models.DeleteObservacion(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetObservacionById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteObservacion(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteObservacion Finalizado Correctamente (OK)")
}
