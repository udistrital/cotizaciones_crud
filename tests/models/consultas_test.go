package models

import (
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
)

// TestRecuperarConsulta verifica que el panic que lanza el ORM con campos
// inexistentes se convierta en error en todas las funciones GetAll.
func TestRecuperarConsulta(t *testing.T) {
	funciones := map[string]func(map[string]string, []string, []string, []string, int64, int64) ([]interface{}, error){
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
	}
	for nombre, getAll := range funciones {
		t.Run(nombre, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("No se esperaba un panic y se obtuvo", r)
				}
			}()
			if _, err := getAll(map[string]string{"CampoInexistente": "1"}, nil, nil, nil, 0, 10); err == nil {
				t.Error("Se esperaba error al filtrar por un campo inexistente")
			}
		})
	}
	t.Log("TestRecuperarConsulta Finalizado Correctamente (OK)")
}
