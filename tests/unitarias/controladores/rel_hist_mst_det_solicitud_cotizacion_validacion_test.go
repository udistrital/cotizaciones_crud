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
// RelHistMstDetSolicitudCotizacionController (/v1/rel_hist_mst_det_solicitud_cotizacion), sin base de datos.

const rutaValidacionRelHistMstDetSolicitudCotizacion = "/v1/rel_hist_mst_det_solicitud_cotizacion"

// cuerpoValidacionRelHistMstDetSolicitudCotizacion retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionRelHistMstDetSolicitudCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.RelHistMstDetSolicitudCotizacion{
		MaestroId:             &models.HistMstSolicitudCotizacion{Id: 1},
		DetalleId:             &models.HistDetSolicitudCotizacion{Id: 1},
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestRelHistMstDetSolicitudCotizacionGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionRelHistMstDetSolicitudCotizacion, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestRelHistMstDetSolicitudCotizacionPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionRelHistMstDetSolicitudCotizacion, helpers.JSON(t, cuerpoValidacionRelHistMstDetSolicitudCotizacion(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestRelHistMstDetSolicitudCotizacionDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionRelHistMstDetSolicitudCotizacion, "")
}
