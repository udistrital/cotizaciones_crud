package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddSolicitudProveedor(t *testing.T) {
	v := helpers.NuevoSolicitudProveedor(t)
	v.Activo = false
	id, err := models.AddSolicitudProveedor(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddSolicitudProveedor Finalizado Correctamente (OK)")
}

func TestGetSolicitudProveedorById(t *testing.T) {
	creado := helpers.CrearSolicitudProveedor(t)
	v, err := models.GetSolicitudProveedorById(creado.Id)
	if err != nil || v.Id != creado.Id || v.ProveedorId != creado.ProveedorId {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetSolicitudProveedorById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetSolicitudProveedorById Finalizado Correctamente (OK)")
}

func TestGetAllSolicitudProveedor(t *testing.T) {
	creado := helpers.CrearSolicitudProveedor(t)
	helpers.CrearSolicitudProveedor(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllSolicitudProveedor, creado.Id)
	t.Log("TestGetAllSolicitudProveedor Finalizado Correctamente (OK)")
}

func TestUpdateSolicitudProveedorById(t *testing.T) {
	creado := helpers.CrearSolicitudProveedor(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.ProveedorId = 2
	if err := models.UpdateSolicitudProveedorById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetSolicitudProveedorById(creado.Id)
	if v == nil || v.ProveedorId != 2 {
		t.Fatal("Se esperaba ProveedorId actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateSolicitudProveedorById(&models.SolicitudProveedor{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateSolicitudProveedorById Finalizado Correctamente (OK)")
}

func TestDeleteSolicitudProveedor(t *testing.T) {
	creado := helpers.CrearSolicitudProveedor(t)
	if err := models.DeleteSolicitudProveedor(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetSolicitudProveedorById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteSolicitudProveedor(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteSolicitudProveedor Finalizado Correctamente (OK)")
}
