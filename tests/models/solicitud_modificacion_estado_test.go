package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddSolicitudModificacionEstado(t *testing.T) {
	v := helpers.NuevoSolicitudModificacionEstado(t)
	v.Activo = false
	id, err := models.AddSolicitudModificacionEstado(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddSolicitudModificacionEstado Finalizado Correctamente (OK)")
}

func TestGetSolicitudModificacionEstadoById(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionEstado(t)
	v, err := models.GetSolicitudModificacionEstadoById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Justificacion != creado.Justificacion {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetSolicitudModificacionEstadoById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetSolicitudModificacionEstadoById Finalizado Correctamente (OK)")
}

func TestGetAllSolicitudModificacionEstado(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionEstado(t)
	helpers.CrearSolicitudModificacionEstado(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllSolicitudModificacionEstado, creado.Id)
	t.Log("TestGetAllSolicitudModificacionEstado Finalizado Correctamente (OK)")
}

func TestUpdateSolicitudModificacionEstadoById(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionEstado(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Justificacion = "Se aprueba la modificacion"
	if err := models.UpdateSolicitudModificacionEstadoById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetSolicitudModificacionEstadoById(creado.Id)
	if v == nil || v.Justificacion != "Se aprueba la modificacion" {
		t.Fatal("Se esperaba Justificacion actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateSolicitudModificacionEstadoById(&models.SolicitudModificacionEstado{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateSolicitudModificacionEstadoById Finalizado Correctamente (OK)")
}

func TestDeleteSolicitudModificacionEstado(t *testing.T) {
	creado := helpers.CrearSolicitudModificacionEstado(t)
	if err := models.DeleteSolicitudModificacionEstado(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetSolicitudModificacionEstadoById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteSolicitudModificacionEstado(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteSolicitudModificacionEstado Finalizado Correctamente (OK)")
}
