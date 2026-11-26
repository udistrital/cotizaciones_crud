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
// RelHistMstDetCotizacionController (/v1/rel_hist_mst_det_cotizacion), sin base de datos.

const rutaValidacionRelHistMstDetCotizacion = "/v1/rel_hist_mst_det_cotizacion"

// cuerpoValidacionRelHistMstDetCotizacion retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionRelHistMstDetCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.RelHistMstDetCotizacion{
		MaestroId:             &models.HistMstCotizacion{Id: 1},
		DetalleId:             &models.HistDetCotizacion{Id: 1},
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestRelHistMstDetCotizacionGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionRelHistMstDetCotizacion, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestRelHistMstDetCotizacionPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionRelHistMstDetCotizacion, helpers.JSON(t, cuerpoValidacionRelHistMstDetCotizacion(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestRelHistMstDetCotizacionDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionRelHistMstDetCotizacion, "")
}
