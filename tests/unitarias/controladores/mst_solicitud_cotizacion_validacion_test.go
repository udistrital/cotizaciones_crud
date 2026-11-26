// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador de la validación de entradas de
// MstSolicitudCotizacionController (/v1/mst_solicitud_cotizacion), sin base de datos.

const rutaValidacionMstSolicitudCotizacion = "/v1/mst_solicitud_cotizacion"

// cuerpoValidacionMstSolicitudCotizacion retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionMstSolicitudCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.MstSolicitudCotizacion{
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestMstSolicitudCotizacionGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionMstSolicitudCotizacion, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestMstSolicitudCotizacionPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionMstSolicitudCotizacion, helpers.JSON(t, cuerpoValidacionMstSolicitudCotizacion(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestMstSolicitudCotizacionDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionMstSolicitudCotizacion, "")
}
