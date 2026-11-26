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
// SolicitudCotizacionActividadEconomicaController (/v1/solicitud_cotizacion_actividad_economica), sin base de datos.

const rutaValidacionSolicitudCotizacionActividadEconomica = "/v1/solicitud_cotizacion_actividad_economica"

// cuerpoValidacionSolicitudCotizacionActividadEconomica retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionSolicitudCotizacionActividadEconomica(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.SolicitudCotizacionActividadEconomica{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
		ActividadEconomicaId:     1,
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionActividadEconomicaGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionSolicitudCotizacionActividadEconomica, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionActividadEconomicaPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionSolicitudCotizacionActividadEconomica, helpers.JSON(t, cuerpoValidacionSolicitudCotizacionActividadEconomica(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionActividadEconomicaDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionSolicitudCotizacionActividadEconomica, "")
}

// La columna actividad_economica_id es NOT NULL: POST sin ActividadEconomicaId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestSolicitudCotizacionActividadEconomicaPostSinActividadEconomicaId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ActividadEconomicaId") },
		"null":    func(c map[string]interface{}) { c["ActividadEconomicaId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionSolicitudCotizacionActividadEconomica(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionSolicitudCotizacionActividadEconomica, cuerpo)
		})
	}
}

// La columna actividad_economica_id es NOT NULL: PUT sin ActividadEconomicaId, con null
// responde 400 sin actualizar.
func TestSolicitudCotizacionActividadEconomicaPutSinActividadEconomicaId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ActividadEconomicaId") },
		"null":    func(c map[string]interface{}) { c["ActividadEconomicaId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionSolicitudCotizacionActividadEconomica(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionSolicitudCotizacionActividadEconomica, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestSolicitudCotizacionActividadEconomicaPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionSolicitudCotizacionActividadEconomica(t)
	cuerpo["ActividadEconomicaId"] = 0
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionSolicitudCotizacionActividadEconomica+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
