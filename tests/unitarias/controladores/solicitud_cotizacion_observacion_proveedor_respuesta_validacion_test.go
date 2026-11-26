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
// SolicitudCotizacionObservacionProveedorRespuestaController (/v1/solicitud_cotizacion_observacion_proveedor_respuesta), sin base de datos.

const rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta = "/v1/solicitud_cotizacion_observacion_proveedor_respuesta"

// cuerpoValidacionSolicitudCotizacionObservacionProveedorRespuesta retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionSolicitudCotizacionObservacionProveedorRespuesta(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.SolicitudCotizacionObservacionProveedorRespuesta{
		ObservacionId:         &models.SolicitudCotizacionObservacionProveedor{Id: 1},
		Texto:                 "Prueba",
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionObservacionProveedorRespuestaGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionObservacionProveedorRespuestaPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta, helpers.JSON(t, cuerpoValidacionSolicitudCotizacionObservacionProveedorRespuesta(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionObservacionProveedorRespuestaDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta, "")
}

// La columna texto es NOT NULL: POST sin Texto, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestSolicitudCotizacionObservacionProveedorRespuestaPostSinTexto(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Texto") },
		"null":     func(c map[string]interface{}) { c["Texto"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Texto"] = "" },
		"espacios": func(c map[string]interface{}) { c["Texto"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionSolicitudCotizacionObservacionProveedorRespuesta(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta, cuerpo)
		})
	}
}

// La columna texto es NOT NULL: PUT sin Texto, con null o vacío
// responde 400 sin actualizar.
func TestSolicitudCotizacionObservacionProveedorRespuestaPutSinTexto(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Texto") },
		"null":     func(c map[string]interface{}) { c["Texto"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Texto"] = "" },
		"espacios": func(c map[string]interface{}) { c["Texto"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionSolicitudCotizacionObservacionProveedorRespuesta(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta, cuerpo)
		})
	}
}
