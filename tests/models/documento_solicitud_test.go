package models

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddDocumentoSolicitud(t *testing.T) {
	v := helpers.NuevoDocumentoSolicitud(t)
	v.Activo = false
	id, err := models.AddDocumentoSolicitud(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddDocumentoSolicitud Finalizado Correctamente (OK)")
}

func TestGetDocumentoSolicitudById(t *testing.T) {
	creado := helpers.CrearDocumentoSolicitud(t)
	v, err := models.GetDocumentoSolicitudById(creado.Id)
	if err != nil || v.Id != creado.Id || v.NombreArchivo != creado.NombreArchivo {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetDocumentoSolicitudById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetDocumentoSolicitudById Finalizado Correctamente (OK)")
}

func TestGetAllDocumentoSolicitud(t *testing.T) {
	creado := helpers.CrearDocumentoSolicitud(t)
	helpers.ProbarGetAll(t, models.GetAllDocumentoSolicitud, creado.Id)
	t.Log("TestGetAllDocumentoSolicitud Finalizado Correctamente (OK)")
}

func TestUpdateDocumentoSolicitudById(t *testing.T) {
	creado := helpers.CrearDocumentoSolicitud(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.NombreArchivo = "anexo_tecnico.pdf"
	if err := models.UpdateDocumentoSolicitudById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetDocumentoSolicitudById(creado.Id)
	if v == nil || v.NombreArchivo != "anexo_tecnico.pdf" {
		t.Fatal("Se esperaba NombreArchivo actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateDocumentoSolicitudById(&models.DocumentoSolicitud{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateDocumentoSolicitudById Finalizado Correctamente (OK)")
}

func TestDeleteDocumentoSolicitud(t *testing.T) {
	creado := helpers.CrearDocumentoSolicitud(t)
	if err := models.DeleteDocumentoSolicitud(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetDocumentoSolicitudById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteDocumentoSolicitud(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteDocumentoSolicitud Finalizado Correctamente (OK)")
}
