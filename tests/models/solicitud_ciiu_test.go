package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddSolicitudCiiu(t *testing.T) {
	v := helpers.NuevoSolicitudCiiu(t)
	v.Activo = false
	id, err := models.AddSolicitudCiiu(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddSolicitudCiiu Finalizado Correctamente (OK)")
}

func TestGetSolicitudCiiuById(t *testing.T) {
	creado := helpers.CrearSolicitudCiiu(t)
	v, err := models.GetSolicitudCiiuById(creado.Id)
	if err != nil || v.Id != creado.Id || v.ActividadEconomicaId != creado.ActividadEconomicaId {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetSolicitudCiiuById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetSolicitudCiiuById Finalizado Correctamente (OK)")
}

func TestGetAllSolicitudCiiu(t *testing.T) {
	creado := helpers.CrearSolicitudCiiu(t)
	helpers.ProbarGetAll(t, models.GetAllSolicitudCiiu, creado.Id)
	t.Log("TestGetAllSolicitudCiiu Finalizado Correctamente (OK)")
}

func TestUpdateSolicitudCiiuById(t *testing.T) {
	creado := helpers.CrearSolicitudCiiu(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.ActividadEconomicaId = 2
	if err := models.UpdateSolicitudCiiuById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetSolicitudCiiuById(creado.Id)
	if v == nil || v.ActividadEconomicaId != 2 {
		t.Fatal("Se esperaba ActividadEconomicaId actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateSolicitudCiiuById(&models.SolicitudCiiu{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateSolicitudCiiuById Finalizado Correctamente (OK)")
}

func TestDeleteSolicitudCiiu(t *testing.T) {
	creado := helpers.CrearSolicitudCiiu(t)
	if err := models.DeleteSolicitudCiiu(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetSolicitudCiiuById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteSolicitudCiiu(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteSolicitudCiiu Finalizado Correctamente (OK)")
}
