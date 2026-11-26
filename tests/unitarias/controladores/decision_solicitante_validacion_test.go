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
// DecisionSolicitanteController (/v1/decision_solicitante), sin base de datos.

const rutaValidacionDecisionSolicitante = "/v1/decision_solicitante"

// cuerpoValidacionDecisionSolicitante retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionDecisionSolicitante(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.DecisionSolicitante{
		RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
		MensajeAProveedor:                 "Prueba",
		ResultadoId:                       1,
		UsuarioCreacionId:                 1,
		UsuarioModificacionId:             1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestDecisionSolicitanteGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionDecisionSolicitante, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestDecisionSolicitantePutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionDecisionSolicitante, helpers.JSON(t, cuerpoValidacionDecisionSolicitante(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestDecisionSolicitanteDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionDecisionSolicitante, "")
}

// La columna mensaje_a_proveedor es NOT NULL: POST sin MensajeAProveedor, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestDecisionSolicitantePostSinMensajeAProveedor(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "MensajeAProveedor") },
		"null":     func(c map[string]interface{}) { c["MensajeAProveedor"] = nil },
		"vacio":    func(c map[string]interface{}) { c["MensajeAProveedor"] = "" },
		"espacios": func(c map[string]interface{}) { c["MensajeAProveedor"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionDecisionSolicitante(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionDecisionSolicitante, cuerpo)
		})
	}
}

// La columna mensaje_a_proveedor es NOT NULL: PUT sin MensajeAProveedor, con null o vacío
// responde 400 sin actualizar.
func TestDecisionSolicitantePutSinMensajeAProveedor(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "MensajeAProveedor") },
		"null":     func(c map[string]interface{}) { c["MensajeAProveedor"] = nil },
		"vacio":    func(c map[string]interface{}) { c["MensajeAProveedor"] = "" },
		"espacios": func(c map[string]interface{}) { c["MensajeAProveedor"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionDecisionSolicitante(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionDecisionSolicitante, cuerpo)
		})
	}
}

// La columna resultado_id es NOT NULL: POST sin ResultadoId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestDecisionSolicitantePostSinResultadoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ResultadoId") },
		"null":    func(c map[string]interface{}) { c["ResultadoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionDecisionSolicitante(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionDecisionSolicitante, cuerpo)
		})
	}
}

// La columna resultado_id es NOT NULL: PUT sin ResultadoId, con null
// responde 400 sin actualizar.
func TestDecisionSolicitantePutSinResultadoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ResultadoId") },
		"null":    func(c map[string]interface{}) { c["ResultadoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionDecisionSolicitante(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionDecisionSolicitante, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestDecisionSolicitantePostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionDecisionSolicitante(t)
	cuerpo["ResultadoId"] = 0
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionDecisionSolicitante+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
