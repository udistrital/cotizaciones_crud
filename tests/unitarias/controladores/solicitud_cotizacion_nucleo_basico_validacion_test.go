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
// SolicitudCotizacionNucleoBasicoController (/v1/solicitud_cotizacion_nucleo_basico), sin base de datos.

const rutaValidacionSolicitudCotizacionNucleoBasico = "/v1/solicitud_cotizacion_nucleo_basico"

// cuerpoValidacionSolicitudCotizacionNucleoBasico retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionSolicitudCotizacionNucleoBasico(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.SolicitudCotizacionNucleoBasico{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
		NucleoBasicoId:           1,
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionNucleoBasicoGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionSolicitudCotizacionNucleoBasico, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionNucleoBasicoPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionSolicitudCotizacionNucleoBasico, helpers.JSON(t, cuerpoValidacionSolicitudCotizacionNucleoBasico(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestSolicitudCotizacionNucleoBasicoDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionSolicitudCotizacionNucleoBasico, "")
}

// La columna nucleo_basico_id es NOT NULL: POST sin NucleoBasicoId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestSolicitudCotizacionNucleoBasicoPostSinNucleoBasicoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "NucleoBasicoId") },
		"null":    func(c map[string]interface{}) { c["NucleoBasicoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionSolicitudCotizacionNucleoBasico(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionSolicitudCotizacionNucleoBasico, cuerpo)
		})
	}
}

// La columna nucleo_basico_id es NOT NULL: PUT sin NucleoBasicoId, con null
// responde 400 sin actualizar.
func TestSolicitudCotizacionNucleoBasicoPutSinNucleoBasicoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "NucleoBasicoId") },
		"null":    func(c map[string]interface{}) { c["NucleoBasicoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionSolicitudCotizacionNucleoBasico(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionSolicitudCotizacionNucleoBasico, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestSolicitudCotizacionNucleoBasicoPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionSolicitudCotizacionNucleoBasico(t)
	cuerpo["NucleoBasicoId"] = 0
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionSolicitudCotizacionNucleoBasico+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
