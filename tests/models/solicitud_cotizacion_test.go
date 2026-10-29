package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddSolicitudCotizacion(t *testing.T) {
	v := helpers.NuevoSolicitudCotizacion(t)
	v.Activo = false
	id, err := models.AddSolicitudCotizacion(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddSolicitudCotizacion Finalizado Correctamente (OK)")
}

func TestGetSolicitudCotizacionById(t *testing.T) {
	creado := helpers.CrearSolicitudCotizacion(t)
	v, err := models.GetSolicitudCotizacionById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Titulo != creado.Titulo {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetSolicitudCotizacionById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetSolicitudCotizacionById Finalizado Correctamente (OK)")
}

func TestGetAllSolicitudCotizacion(t *testing.T) {
	creado := helpers.CrearSolicitudCotizacion(t)
	helpers.CrearSolicitudCotizacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAll(t, models.GetAllSolicitudCotizacion, creado.Id)
	t.Log("TestGetAllSolicitudCotizacion Finalizado Correctamente (OK)")
}

func TestUpdateSolicitudCotizacionById(t *testing.T) {
	creado := helpers.CrearSolicitudCotizacion(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Titulo = "Titulo actualizado"
	if err := models.UpdateSolicitudCotizacionById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetSolicitudCotizacionById(creado.Id)
	if v == nil || v.Titulo != "Titulo actualizado" {
		t.Fatal("Se esperaba Titulo actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateSolicitudCotizacionById(&models.SolicitudCotizacion{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateSolicitudCotizacionById Finalizado Correctamente (OK)")
}

func TestDeleteSolicitudCotizacion(t *testing.T) {
	creado := helpers.CrearSolicitudCotizacion(t)
	if err := models.DeleteSolicitudCotizacion(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetSolicitudCotizacionById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteSolicitudCotizacion(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteSolicitudCotizacion Finalizado Correctamente (OK)")
}

func TestAddSolicitudCotizacionNumeroSolicitud(t *testing.T) {
	// Sin número la solicitud se guarda con NULL (el número se asigna al procesarla)
	sinNumero := helpers.CrearSolicitudCotizacion(t)
	if v, err := models.GetSolicitudCotizacionById(sinNumero.Id); err != nil || v.NumeroSolicitud != nil {
		t.Error("Se esperaba NumeroSolicitud nulo y se obtuvo", v, err)
	}
	numero := "SC-2026-0001"
	conNumero := helpers.NuevoSolicitudCotizacion(t)
	conNumero.NumeroSolicitud = &numero
	if _, err := models.AddSolicitudCotizacion(conNumero); err != nil {
		t.Fatal("Se esperaba crear la solicitud con número y se obtuvo", err)
	}
	if v, err := models.GetSolicitudCotizacionById(conNumero.Id); err != nil || v.NumeroSolicitud == nil || *v.NumeroSolicitud != numero {
		t.Error("Se esperaba NumeroSolicitud", numero, "y se obtuvo", v, err)
	}
	t.Log("TestAddSolicitudCotizacionNumeroSolicitud Finalizado Correctamente (OK)")
}
