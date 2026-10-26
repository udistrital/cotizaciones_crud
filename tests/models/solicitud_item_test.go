package models

import (
	"strconv"
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestAddSolicitudItem(t *testing.T) {
	v := helpers.NuevoSolicitudItem(t)
	v.Activo = false
	id, err := models.AddSolicitudItem(v)
	if err != nil || id == 0 || v.Id != int(id) {
		t.Fatal("Se esperaba el registro creado y se obtuvo", id, err)
	}
	if !v.Activo || v.FechaCreacion.IsZero() || !v.FechaModificacion.Equal(v.FechaCreacion) {
		t.Error("Se esperaba Activo=true y FechaCreacion y FechaModificacion asignadas y se obtuvo", v.Activo, v.FechaCreacion, v.FechaModificacion)
	}
	t.Log("TestAddSolicitudItem Finalizado Correctamente (OK)")
}

func TestGetSolicitudItemById(t *testing.T) {
	creado := helpers.CrearSolicitudItem(t)
	v, err := models.GetSolicitudItemById(creado.Id)
	if err != nil || v.Id != creado.Id || v.Nombre != creado.Nombre {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", v, err)
	}
	if v, err := models.GetSolicitudItemById(helpers.IdInexistente); err == nil || v != nil {
		t.Error("Se esperaba error al consultar un Id inexistente")
	}
	t.Log("TestGetSolicitudItemById Finalizado Correctamente (OK)")
}

func TestGetAllSolicitudItem(t *testing.T) {
	creado := helpers.CrearSolicitudItem(t)
	helpers.ProbarGetAll(t, models.GetAllSolicitudItem, creado.Id)
	t.Log("TestGetAllSolicitudItem Finalizado Correctamente (OK)")
}

func TestUpdateSolicitudItemById(t *testing.T) {
	creado := helpers.CrearSolicitudItem(t)
	fechaCreacion := creado.FechaCreacion
	cambios := *creado
	cambios.FechaCreacion = time.Time{}
	cambios.Nombre = "Portatil"
	if err := models.UpdateSolicitudItemById(&cambios); err != nil {
		t.Fatal("Se esperaba actualizar el registro y se obtuvo", err)
	}
	v, _ := models.GetSolicitudItemById(creado.Id)
	if v == nil || v.Nombre != "Portatil" {
		t.Fatal("Se esperaba Nombre actualizado y se obtuvo", v)
	}
	if !helpers.MismaFecha(v.FechaCreacion, fechaCreacion) || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba conservar FechaCreacion y actualizar FechaModificacion y se obtuvo", v.FechaCreacion, v.FechaModificacion)
	}
	if err := models.UpdateSolicitudItemById(&models.SolicitudItem{Id: helpers.IdInexistente}); err == nil {
		t.Error("Se esperaba error al actualizar un Id inexistente")
	}
	t.Log("TestUpdateSolicitudItemById Finalizado Correctamente (OK)")
}

func TestDeleteSolicitudItem(t *testing.T) {
	creado := helpers.CrearSolicitudItem(t)
	if err := models.DeleteSolicitudItem(creado.Id); err != nil {
		t.Fatal("Se esperaba eliminar el registro y se obtuvo", err)
	}
	if v, _ := models.GetSolicitudItemById(creado.Id); v == nil || v.Activo || v.FechaModificacion.Before(v.FechaCreacion) {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	if err := models.DeleteSolicitudItem(helpers.IdInexistente); err == nil {
		t.Error("Se esperaba error al eliminar un Id inexistente")
	}
	t.Log("TestDeleteSolicitudItem Finalizado Correctamente (OK)")
}

func TestAddSolicitudItemCargueMasivo(t *testing.T) {
	// Un ítem creado a mano no tiene cargue masivo (llave foránea opcional)
	sinCargue := helpers.CrearSolicitudItem(t)
	if sinCargue.CargueMasivoId != nil {
		t.Fatal("Se esperaba un ítem sin cargue masivo y se obtuvo", sinCargue.CargueMasivoId)
	}
	l, err := models.GetAllSolicitudItem(map[string]string{"Id": strconv.Itoa(sinCargue.Id)}, nil, nil, nil, 0, 10)
	if err != nil || len(l) != 1 {
		t.Error("Se esperaba consultar el ítem sin cargue masivo y se obtuvo", l, err)
	}

	conCargue := helpers.NuevoSolicitudItem(t)
	conCargue.CargueMasivoId = helpers.CrearCargueMasivo(t)
	if _, err := models.AddSolicitudItem(conCargue); err != nil {
		t.Fatal("Se esperaba crear el ítem con cargue masivo y se obtuvo", err)
	}
	l, err = models.GetAllSolicitudItem(map[string]string{"CargueMasivoId.Id": strconv.Itoa(conCargue.CargueMasivoId.Id)}, nil, nil, nil, 0, 10)
	if err != nil || len(l) != 1 {
		t.Error("Se esperaba consultar el ítem por su cargue masivo y se obtuvo", l, err)
	}
	t.Log("TestAddSolicitudItemCargueMasivo Finalizado Correctamente (OK)")
}

func TestAddSolicitudItemSinSolicitud(t *testing.T) {
	v := helpers.NuevoSolicitudItem(t)
	v.SolicitudId = nil
	if _, err := models.AddSolicitudItem(v); err == nil {
		t.Error("Se esperaba error al crear un ítem sin solicitud")
	}
	t.Log("TestAddSolicitudItemSinSolicitud Finalizado Correctamente (OK)")
}
