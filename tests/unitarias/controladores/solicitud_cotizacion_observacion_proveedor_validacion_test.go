// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"net/http"
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador de la validación de entradas de
// SolicitudCotizacionObservacionProveedorController (/v1/solicitud_cotizacion_observacion_proveedor), sin base de datos.

const rutaValidacionSolicitudCotizacionObservacionProveedor = "/v1/solicitud_cotizacion_observacion_proveedor"

// cuerpoValidacionSolicitudCotizacionObservacionProveedor retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionSolicitudCotizacionObservacionProveedor(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.SolicitudCotizacionObservacionProveedor{
		Texto:                             helpers.Ptr("Prueba"),
		RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
		VistoSolicitante:                  true,
		FechaVistoSolicitante:             helpers.Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
		UsuarioCreacionId:                 1,
		UsuarioModificacionId:             1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionObservacionProveedorGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionSolicitudCotizacionObservacionProveedor, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionObservacionProveedorPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionSolicitudCotizacionObservacionProveedor, helpers.JSON(t, cuerpoValidacionSolicitudCotizacionObservacionProveedor(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionObservacionProveedorDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionSolicitudCotizacionObservacionProveedor, "")
}

// La columna visto_solicitante es NOT NULL: POST sin VistoSolicitante, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestSolicitudCotizacionObservacionProveedorPostSinVistoSolicitante(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "VistoSolicitante") },
		"null":    func(c map[string]interface{}) { c["VistoSolicitante"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionSolicitudCotizacionObservacionProveedor(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionSolicitudCotizacionObservacionProveedor, cuerpo)
		})
	}
}

// La columna visto_solicitante es NOT NULL: PUT sin VistoSolicitante, con null
// responde 400 sin actualizar.
func TestSolicitudCotizacionObservacionProveedorPutSinVistoSolicitante(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "VistoSolicitante") },
		"null":    func(c map[string]interface{}) { c["VistoSolicitante"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionSolicitudCotizacionObservacionProveedor(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionSolicitudCotizacionObservacionProveedor, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestSolicitudCotizacionObservacionProveedorPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionSolicitudCotizacionObservacionProveedor(t)
	cuerpo["VistoSolicitante"] = false
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionSolicitudCotizacionObservacionProveedor+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
