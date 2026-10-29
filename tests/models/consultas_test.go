package models

import (
	"strconv"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// TestRecuperarConsulta verifica que el panic que lanza el ORM con campos
// inexistentes se convierta en error en todas las funciones GetAll.
func TestRecuperarConsulta(t *testing.T) {
	funciones := map[string]helpers.FuncionGetAll{
		"SolicitudCotizacion":          models.GetAllSolicitudCotizacion,
		"SolicitudNecesidad":           models.GetAllSolicitudNecesidad,
		"SolicitudItem":                models.GetAllSolicitudItem,
		"CargueMasivo":                 models.GetAllCargueMasivo,
		"CargueMasivoError":            models.GetAllCargueMasivoError,
		"SolicitudPagoParcial":         models.GetAllSolicitudPagoParcial,
		"SolicitudCiiu":                models.GetAllSolicitudCiiu,
		"DocumentoSolicitud":           models.GetAllDocumentoSolicitud,
		"SolicitudRequisitoDocumental": models.GetAllSolicitudRequisitoDocumental,
		"SolicitudProveedor":           models.GetAllSolicitudProveedor,
		"SolicitudEstado":              models.GetAllSolicitudEstado,
		"SolicitudModificacion":        models.GetAllSolicitudModificacion,
		"SolicitudModificacionEstado":  models.GetAllSolicitudModificacionEstado,
		"SolicitudModificacionDetalle": models.GetAllSolicitudModificacionDetalle,
		"Observacion":                  models.GetAllObservacion,
		"RespuestaObservacion":         models.GetAllRespuestaObservacion,
	}
	for nombre, getAll := range funciones {
		t.Run(nombre, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("No se esperaba un panic y se obtuvo", r)
				}
			}()
			if _, _, err := getAll(map[string]string{"CampoInexistente": "1"}, nil, nil, nil, 0, 10); err == nil {
				t.Error("Se esperaba error al filtrar por un campo inexistente")
			}
		})
	}
	t.Log("TestRecuperarConsulta Finalizado Correctamente (OK)")
}

// TestFiltroBooleano verifica que los filtros del parámetro query sobre campos
// booleanos, propios o de una relación, se comparen como booleanos.
func TestFiltroBooleano(t *testing.T) {
	respuesta := helpers.CrearRespuestaObservacion(t) // Publica = true
	inactiva := helpers.CrearRespuestaObservacion(t)
	if err := models.DeleteRespuestaObservacion(inactiva.Id); err != nil {
		t.Fatal("No se pudo inactivar la respuesta:", err)
	}
	id := strconv.Itoa(respuesta.Id)
	idInactiva := strconv.Itoa(inactiva.Id)

	casos := []struct {
		nombre   string
		query    map[string]string
		esperado int64
	}{
		{"CampoPropioTrue", map[string]string{"Id": id, "Activo": "true"}, 1},
		{"CampoPropioFalse", map[string]string{"Id": idInactiva, "Activo": "false"}, 1},
		{"CampoPropioEnMinuscula", map[string]string{"Id": idInactiva, "activo": "false"}, 1},
		{"CampoPropioNoCoincide", map[string]string{"Id": id, "Activo": "false"}, 0},
		{"OtroCampoBooleano", map[string]string{"Id": id, "Publica": "true"}, 1},
		{"CampoDeRelacion", map[string]string{"Id": id, "ObservacionId__Activo": "true"}, 1},
		{"CampoDeRelacionNoCoincide", map[string]string{"Id": id, "ObservacionId__Activo": "false"}, 0},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			if _, total, err := models.GetAllRespuestaObservacion(c.query, nil, nil, nil, 0, 10); err != nil || total != c.esperado {
				t.Errorf("Se esperaba un total de %d con %v y se obtuvo %d %v", c.esperado, c.query, total, err)
			}
		})
	}
	t.Log("TestFiltroBooleano Finalizado Correctamente (OK)")
}
