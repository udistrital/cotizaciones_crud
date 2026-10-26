package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddSolicitudRequisitoDocumental(t *testing.T) {
	v := helpers.NuevoSolicitudRequisitoDocumental(t)
	v.Activo = false
	id, err := models.AddSolicitudRequisitoDocumental(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddSolicitudRequisitoDocumental Finalizado Correctamente (OK)")
}

func TestGetSolicitudRequisitoDocumentalById(t *testing.T) {
	creado := helpers.CrearSolicitudRequisitoDocumental(t)
	v, err := models.GetSolicitudRequisitoDocumentalById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Instruccion != creado.Instruccion {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetSolicitudRequisitoDocumentalById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetSolicitudRequisitoDocumentalById Finalizado Correctamente (OK)")
}

func TestGetAllSolicitudRequisitoDocumental(t *testing.T) {
	creado := helpers.CrearSolicitudRequisitoDocumental(t)
	helpers.ProbarGetAll(t, models.GetAllSolicitudRequisitoDocumental, creado.Id)
	t.Log("TestGetAllSolicitudRequisitoDocumental Finalizado Correctamente (OK)")
}

func TestUpdateSolicitudRequisitoDocumentalById(t *testing.T) {
	creado := helpers.CrearSolicitudRequisitoDocumental(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Instruccion = "Adjuntar camara de comercio"
	if err := models.UpdateSolicitudRequisitoDocumentalById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetSolicitudRequisitoDocumentalById(creado.Id)
	if v == nil || v.Instruccion != "Adjuntar camara de comercio" {
		t.Fatal("Se esperaba Instruccion actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateSolicitudRequisitoDocumentalById(&models.SolicitudRequisitoDocumental{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateSolicitudRequisitoDocumentalById Finalizado Correctamente (OK)")
}

func TestDeleteSolicitudRequisitoDocumental(t *testing.T) {
	creado := helpers.CrearSolicitudRequisitoDocumental(t)
	if err := models.DeleteSolicitudRequisitoDocumental(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetSolicitudRequisitoDocumentalById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteSolicitudRequisitoDocumental(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteSolicitudRequisitoDocumental Finalizado Correctamente (OK)")
}
