package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddSolicitudNecesidad(t *testing.T) {
	v := helpers.NuevoSolicitudNecesidad(t)
	v.Activo = false
	id, err := models.AddSolicitudNecesidad(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddSolicitudNecesidad Finalizado Correctamente (OK)")
}

func TestGetSolicitudNecesidadById(t *testing.T) {
	creado := helpers.CrearSolicitudNecesidad(t)
	v, err := models.GetSolicitudNecesidadById(creado.Id)
	if err != nil || v.Id != creado.Id || v.RubroNombre != creado.RubroNombre {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetSolicitudNecesidadById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetSolicitudNecesidadById Finalizado Correctamente (OK)")
}

func TestGetAllSolicitudNecesidad(t *testing.T) {
	creado := helpers.CrearSolicitudNecesidad(t)
	helpers.ProbarGetAll(t, models.GetAllSolicitudNecesidad, creado.Id)
	t.Log("TestGetAllSolicitudNecesidad Finalizado Correctamente (OK)")
}

func TestUpdateSolicitudNecesidadById(t *testing.T) {
	creado := helpers.CrearSolicitudNecesidad(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.RubroNombre = "Rubro actualizado"
	if err := models.UpdateSolicitudNecesidadById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetSolicitudNecesidadById(creado.Id)
	if v == nil || v.RubroNombre != "Rubro actualizado" {
		t.Fatal("Se esperaba RubroNombre actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateSolicitudNecesidadById(&models.SolicitudNecesidad{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateSolicitudNecesidadById Finalizado Correctamente (OK)")
}

func TestDeleteSolicitudNecesidad(t *testing.T) {
	creado := helpers.CrearSolicitudNecesidad(t)
	if err := models.DeleteSolicitudNecesidad(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetSolicitudNecesidadById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteSolicitudNecesidad(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteSolicitudNecesidad Finalizado Correctamente (OK)")
}
