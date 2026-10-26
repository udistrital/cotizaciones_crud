package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddCargueMasivoError(t *testing.T) {
	v := helpers.NuevoCargueMasivoError(t)
	v.Activo = false
	id, err := models.AddCargueMasivoError(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddCargueMasivoError Finalizado Correctamente (OK)")
}

func TestGetCargueMasivoErrorById(t *testing.T) {
	creado := helpers.CrearCargueMasivoError(t)
	v, err := models.GetCargueMasivoErrorById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Detalle != creado.Detalle {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetCargueMasivoErrorById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetCargueMasivoErrorById Finalizado Correctamente (OK)")
}

func TestGetAllCargueMasivoError(t *testing.T) {
	creado := helpers.CrearCargueMasivoError(t)
	helpers.ProbarGetAll(t, models.GetAllCargueMasivoError, creado.Id)
	t.Log("TestGetAllCargueMasivoError Finalizado Correctamente (OK)")
}

func TestUpdateCargueMasivoErrorById(t *testing.T) {
	creado := helpers.CrearCargueMasivoError(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Detalle = "Detalle actualizado"
	if err := models.UpdateCargueMasivoErrorById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetCargueMasivoErrorById(creado.Id)
	if v == nil || v.Detalle != "Detalle actualizado" {
		t.Fatal("Se esperaba Detalle actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateCargueMasivoErrorById(&models.CargueMasivoError{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateCargueMasivoErrorById Finalizado Correctamente (OK)")
}

func TestDeleteCargueMasivoError(t *testing.T) {
	creado := helpers.CrearCargueMasivoError(t)
	if err := models.DeleteCargueMasivoError(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetCargueMasivoErrorById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteCargueMasivoError(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteCargueMasivoError Finalizado Correctamente (OK)")
}
