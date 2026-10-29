package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddSolicitudEstado(t *testing.T) {
	v := helpers.NuevoSolicitudEstado(t)
	v.Activo = false
	id, err := models.AddSolicitudEstado(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddSolicitudEstado Finalizado Correctamente (OK)")
}

func TestGetSolicitudEstadoById(t *testing.T) {
	creado := helpers.CrearSolicitudEstado(t)
	v, err := models.GetSolicitudEstadoById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Justificacion != creado.Justificacion {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetSolicitudEstadoById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetSolicitudEstadoById Finalizado Correctamente (OK)")
}

func TestGetAllSolicitudEstado(t *testing.T) {
	creado := helpers.CrearSolicitudEstado(t)
	helpers.CrearSolicitudEstado(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllSolicitudEstado, creado.Id)
	t.Log("TestGetAllSolicitudEstado Finalizado Correctamente (OK)")
}

func TestUpdateSolicitudEstadoById(t *testing.T) {
	creado := helpers.CrearSolicitudEstado(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Justificacion = "Envio a proveedores"
	if err := models.UpdateSolicitudEstadoById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetSolicitudEstadoById(creado.Id)
	if v == nil || v.Justificacion != "Envio a proveedores" {
		t.Fatal("Se esperaba Justificacion actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateSolicitudEstadoById(&models.SolicitudEstado{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateSolicitudEstadoById Finalizado Correctamente (OK)")
}

func TestDeleteSolicitudEstado(t *testing.T) {
	creado := helpers.CrearSolicitudEstado(t)
	if err := models.DeleteSolicitudEstado(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetSolicitudEstadoById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteSolicitudEstado(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteSolicitudEstado Finalizado Correctamente (OK)")
}
