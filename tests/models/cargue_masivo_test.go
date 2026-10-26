package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddCargueMasivo(t *testing.T) {
	v := helpers.NuevoCargueMasivo(t)
	v.Activo = false
	id, err := models.AddCargueMasivo(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddCargueMasivo Finalizado Correctamente (OK)")
}

func TestGetCargueMasivoById(t *testing.T) {
	creado := helpers.CrearCargueMasivo(t)
	v, err := models.GetCargueMasivoById(creado.Id)
	if err != nil || v.Id != creado.Id || v.NombreArchivo != creado.NombreArchivo {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetCargueMasivoById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetCargueMasivoById Finalizado Correctamente (OK)")
}

func TestGetAllCargueMasivo(t *testing.T) {
	creado := helpers.CrearCargueMasivo(t)
	helpers.ProbarGetAll(t, models.GetAllCargueMasivo, creado.Id)
	t.Log("TestGetAllCargueMasivo Finalizado Correctamente (OK)")
}

func TestUpdateCargueMasivoById(t *testing.T) {
	creado := helpers.CrearCargueMasivo(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.NombreArchivo = "items_v2.xlsx"
	if err := models.UpdateCargueMasivoById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetCargueMasivoById(creado.Id)
	if v == nil || v.NombreArchivo != "items_v2.xlsx" {
		t.Fatal("Se esperaba NombreArchivo actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateCargueMasivoById(&models.CargueMasivo{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateCargueMasivoById Finalizado Correctamente (OK)")
}

func TestDeleteCargueMasivo(t *testing.T) {
	creado := helpers.CrearCargueMasivo(t)
	if err := models.DeleteCargueMasivo(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetCargueMasivoById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteCargueMasivo(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteCargueMasivo Finalizado Correctamente (OK)")
}
