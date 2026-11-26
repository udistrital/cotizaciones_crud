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
// MstCotizacionController (/v1/mst_cotizacion), sin base de datos.

const rutaValidacionMstCotizacion = "/v1/mst_cotizacion"

// cuerpoValidacionMstCotizacion retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionMstCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.MstCotizacion{
		RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
		UsuarioCreacionId:                 1,
		UsuarioModificacionId:             1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestMstCotizacionGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionMstCotizacion, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestMstCotizacionPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionMstCotizacion, helpers.JSON(t, cuerpoValidacionMstCotizacion(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestMstCotizacionDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionMstCotizacion, "")
}
