package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddSolicitudPagoParcial(t *testing.T) {
	v := helpers.NuevoSolicitudPagoParcial(t)
	v.Activo = false
	id, err := models.AddSolicitudPagoParcial(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddSolicitudPagoParcial Finalizado Correctamente (OK)")
}

func TestGetSolicitudPagoParcialById(t *testing.T) {
	creado := helpers.CrearSolicitudPagoParcial(t)
	v, err := models.GetSolicitudPagoParcialById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Condicion != creado.Condicion {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetSolicitudPagoParcialById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetSolicitudPagoParcialById Finalizado Correctamente (OK)")
}

func TestGetAllSolicitudPagoParcial(t *testing.T) {
	creado := helpers.CrearSolicitudPagoParcial(t)
	helpers.ProbarGetAll(t, models.GetAllSolicitudPagoParcial, creado.Id)
	t.Log("TestGetAllSolicitudPagoParcial Finalizado Correctamente (OK)")
}

func TestUpdateSolicitudPagoParcialById(t *testing.T) {
	creado := helpers.CrearSolicitudPagoParcial(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Condicion = "Contra entrega total"
	if err := models.UpdateSolicitudPagoParcialById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetSolicitudPagoParcialById(creado.Id)
	if v == nil || v.Condicion != "Contra entrega total" {
		t.Fatal("Se esperaba Condicion actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateSolicitudPagoParcialById(&models.SolicitudPagoParcial{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateSolicitudPagoParcialById Finalizado Correctamente (OK)")
}

func TestDeleteSolicitudPagoParcial(t *testing.T) {
	creado := helpers.CrearSolicitudPagoParcial(t)
	if err := models.DeleteSolicitudPagoParcial(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetSolicitudPagoParcialById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteSolicitudPagoParcial(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteSolicitudPagoParcial Finalizado Correctamente (OK)")
}
