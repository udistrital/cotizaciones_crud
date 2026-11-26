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
// SolicitudCotizacionEspecificacionTecnicaController (/v1/solicitud_cotizacion_especificacion_tecnica), sin base de datos.

const rutaValidacionSolicitudCotizacionEspecificacionTecnica = "/v1/solicitud_cotizacion_especificacion_tecnica"

// cuerpoValidacionSolicitudCotizacionEspecificacionTecnica retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionSolicitudCotizacionEspecificacionTecnica(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.SolicitudCotizacionEspecificacionTecnica{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
		Enlace:                   "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionEspecificacionTecnicaGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionSolicitudCotizacionEspecificacionTecnica, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionEspecificacionTecnicaPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionSolicitudCotizacionEspecificacionTecnica, helpers.JSON(t, cuerpoValidacionSolicitudCotizacionEspecificacionTecnica(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionEspecificacionTecnicaDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionSolicitudCotizacionEspecificacionTecnica, "")
}

// La columna enlace es NOT NULL: POST sin Enlace, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestSolicitudCotizacionEspecificacionTecnicaPostSinEnlace(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Enlace") },
		"null":     func(c map[string]interface{}) { c["Enlace"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Enlace"] = "" },
		"espacios": func(c map[string]interface{}) { c["Enlace"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionSolicitudCotizacionEspecificacionTecnica(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionSolicitudCotizacionEspecificacionTecnica, cuerpo)
		})
	}
}

// La columna enlace es NOT NULL: PUT sin Enlace, con null o vacío
// responde 400 sin actualizar.
func TestSolicitudCotizacionEspecificacionTecnicaPutSinEnlace(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Enlace") },
		"null":     func(c map[string]interface{}) { c["Enlace"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Enlace"] = "" },
		"espacios": func(c map[string]interface{}) { c["Enlace"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionSolicitudCotizacionEspecificacionTecnica(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionSolicitudCotizacionEspecificacionTecnica, cuerpo)
		})
	}
}
