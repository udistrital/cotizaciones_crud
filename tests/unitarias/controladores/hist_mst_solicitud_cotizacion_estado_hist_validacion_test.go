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
// HistMstSolicitudCotizacionEstadoHistController (/v1/hist_mst_solicitud_cotizacion_estado_hist), sin base de datos.

const rutaValidacionHistMstSolicitudCotizacionEstadoHist = "/v1/hist_mst_solicitud_cotizacion_estado_hist"

// cuerpoValidacionHistMstSolicitudCotizacionEstadoHist retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionHistMstSolicitudCotizacionEstadoHist(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.HistMstSolicitudCotizacionEstadoHist{
		ResponsableTerceroId:         1,
		Justificacion:                "Prueba",
		EstadoHistId:                 1,
		HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
		UsuarioCreacionId:            1,
		UsuarioModificacionId:        1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistMstSolicitudCotizacionEstadoHistGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionHistMstSolicitudCotizacionEstadoHist, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistMstSolicitudCotizacionEstadoHistPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionHistMstSolicitudCotizacionEstadoHist, helpers.JSON(t, cuerpoValidacionHistMstSolicitudCotizacionEstadoHist(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistMstSolicitudCotizacionEstadoHistDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionHistMstSolicitudCotizacionEstadoHist, "")
}

// La columna responsable_tercero_id es NOT NULL: POST sin ResponsableTerceroId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionEstadoHistPostSinResponsableTerceroId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ResponsableTerceroId") },
		"null":    func(c map[string]interface{}) { c["ResponsableTerceroId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacionEstadoHist(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacionEstadoHist, cuerpo)
		})
	}
}

// La columna responsable_tercero_id es NOT NULL: PUT sin ResponsableTerceroId, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionEstadoHistPutSinResponsableTerceroId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ResponsableTerceroId") },
		"null":    func(c map[string]interface{}) { c["ResponsableTerceroId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacionEstadoHist(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacionEstadoHist, cuerpo)
		})
	}
}

// La columna justificacion es NOT NULL: POST sin Justificacion, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionEstadoHistPostSinJustificacion(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Justificacion") },
		"null":     func(c map[string]interface{}) { c["Justificacion"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Justificacion"] = "" },
		"espacios": func(c map[string]interface{}) { c["Justificacion"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacionEstadoHist(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacionEstadoHist, cuerpo)
		})
	}
}

// La columna justificacion es NOT NULL: PUT sin Justificacion, con null o vacío
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionEstadoHistPutSinJustificacion(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Justificacion") },
		"null":     func(c map[string]interface{}) { c["Justificacion"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Justificacion"] = "" },
		"espacios": func(c map[string]interface{}) { c["Justificacion"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacionEstadoHist(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacionEstadoHist, cuerpo)
		})
	}
}

// La columna estado_hist_id es NOT NULL: POST sin EstadoHistId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionEstadoHistPostSinEstadoHistId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "EstadoHistId") },
		"null":    func(c map[string]interface{}) { c["EstadoHistId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacionEstadoHist(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacionEstadoHist, cuerpo)
		})
	}
}

// La columna estado_hist_id es NOT NULL: PUT sin EstadoHistId, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionEstadoHistPutSinEstadoHistId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "EstadoHistId") },
		"null":    func(c map[string]interface{}) { c["EstadoHistId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacionEstadoHist(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacionEstadoHist, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestHistMstSolicitudCotizacionEstadoHistPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionHistMstSolicitudCotizacionEstadoHist(t)
	cuerpo["ResponsableTerceroId"] = 0
	cuerpo["EstadoHistId"] = 0
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionHistMstSolicitudCotizacionEstadoHist+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
