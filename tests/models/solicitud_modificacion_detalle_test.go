package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddSolicitudModificacionDetalle(t *testing.T) {
	v := helpers.NuevoSolicitudModificacionDetalle(t)
	v.Activo = false
	id, err := models.AddSolicitudModificacionDetalle(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddSolicitudModificacionDetalle Finalizado Correctamente (OK)")
}

func TestGetSolicitudModificacionDetalleById(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionDetalle(t)
	v, err := models.GetSolicitudModificacionDetalleById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Justificacion != creado.Justificacion {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetSolicitudModificacionDetalleById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetSolicitudModificacionDetalleById Finalizado Correctamente (OK)")
}

func TestGetAllSolicitudModificacionDetalle(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionDetalle(t)
	helpers.CrearSolicitudModificacionDetalle(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllSolicitudModificacionDetalle, creado.Id)
	t.Log("TestGetAllSolicitudModificacionDetalle Finalizado Correctamente (OK)")
}

func TestUpdateSolicitudModificacionDetalleById(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionDetalle(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Justificacion = "Ampliacion de plazo"
	if err := models.UpdateSolicitudModificacionDetalleById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetSolicitudModificacionDetalleById(creado.Id)
	if v == nil || v.Justificacion != "Ampliacion de plazo" {
		t.Fatal("Se esperaba Justificacion actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateSolicitudModificacionDetalleById(&models.SolicitudModificacionDetalle{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateSolicitudModificacionDetalleById Finalizado Correctamente (OK)")
}

func TestDeleteSolicitudModificacionDetalle(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionDetalle(t)
	if err := models.DeleteSolicitudModificacionDetalle(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetSolicitudModificacionDetalleById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteSolicitudModificacionDetalle(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteSolicitudModificacionDetalle Finalizado Correctamente (OK)")
}
