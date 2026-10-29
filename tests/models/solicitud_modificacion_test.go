package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddSolicitudModificacion(t *testing.T) {
	v := helpers.NuevoSolicitudModificacion(t)
	v.Activo = false
	id, err := models.AddSolicitudModificacion(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddSolicitudModificacion Finalizado Correctamente (OK)")
}

func TestGetSolicitudModificacionById(t *testing.T) {
	creado := helpers.CrearSolicitudModificacion(t)
	v, err := models.GetSolicitudModificacionById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Numero != creado.Numero {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetSolicitudModificacionById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetSolicitudModificacionById Finalizado Correctamente (OK)")
}

func TestGetAllSolicitudModificacion(t *testing.T) {
	creado := helpers.CrearSolicitudModificacion(t)
	helpers.CrearSolicitudModificacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllSolicitudModificacion, creado.Id)
	t.Log("TestGetAllSolicitudModificacion Finalizado Correctamente (OK)")
}

func TestUpdateSolicitudModificacionById(t *testing.T) {
	creado := helpers.CrearSolicitudModificacion(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Numero = 2
	if err := models.UpdateSolicitudModificacionById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetSolicitudModificacionById(creado.Id)
	if v == nil || v.Numero != 2 {
		t.Fatal("Se esperaba Numero actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateSolicitudModificacionById(&models.SolicitudModificacion{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateSolicitudModificacionById Finalizado Correctamente (OK)")
}

func TestDeleteSolicitudModificacion(t *testing.T) {
	creado := helpers.CrearSolicitudModificacion(t)
	if err := models.DeleteSolicitudModificacion(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetSolicitudModificacionById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteSolicitudModificacion(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteSolicitudModificacion Finalizado Correctamente (OK)")
}
