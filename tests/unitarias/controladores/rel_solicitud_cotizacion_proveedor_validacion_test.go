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
// RelSolicitudCotizacionProveedorController (/v1/rel_solicitud_cotizacion_proveedor), sin base de datos.

const rutaValidacionRelSolicitudCotizacionProveedor = "/v1/rel_solicitud_cotizacion_proveedor"

// cuerpoValidacionRelSolicitudCotizacionProveedor retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionRelSolicitudCotizacionProveedor(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.RelSolicitudCotizacionProveedor{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
		ProveedorId:              1,
		NotificacionId:           &models.Notificacion{Id: 1},
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestRelSolicitudCotizacionProveedorGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionRelSolicitudCotizacionProveedor, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestRelSolicitudCotizacionProveedorPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionRelSolicitudCotizacionProveedor, helpers.JSON(t, cuerpoValidacionRelSolicitudCotizacionProveedor(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestRelSolicitudCotizacionProveedorDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionRelSolicitudCotizacionProveedor, "")
}

// La columna proveedor_id es NOT NULL: POST sin ProveedorId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestRelSolicitudCotizacionProveedorPostSinProveedorId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ProveedorId") },
		"null":    func(c map[string]interface{}) { c["ProveedorId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionRelSolicitudCotizacionProveedor(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionRelSolicitudCotizacionProveedor, cuerpo)
		})
	}
}

// La columna proveedor_id es NOT NULL: PUT sin ProveedorId, con null
// responde 400 sin actualizar.
func TestRelSolicitudCotizacionProveedorPutSinProveedorId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ProveedorId") },
		"null":    func(c map[string]interface{}) { c["ProveedorId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionRelSolicitudCotizacionProveedor(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionRelSolicitudCotizacionProveedor, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestRelSolicitudCotizacionProveedorPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionRelSolicitudCotizacionProveedor(t)
	cuerpo["ProveedorId"] = 0
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionRelSolicitudCotizacionProveedor+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
