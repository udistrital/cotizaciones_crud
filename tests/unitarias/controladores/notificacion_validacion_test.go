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
// NotificacionController (/v1/notificacion), sin base de datos.

const rutaValidacionNotificacion = "/v1/notificacion"

// cuerpoValidacionNotificacion retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionNotificacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.Notificacion{
		ProveedorNotificadoId: 1,
		TerceroNotificaId:     1,
		EstadoEnvioId:         1,
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestNotificacionGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionNotificacion, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestNotificacionPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionNotificacion, helpers.JSON(t, cuerpoValidacionNotificacion(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestNotificacionDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionNotificacion, "")
}

// La columna proveedor_notificado_id es NOT NULL: POST sin ProveedorNotificadoId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestNotificacionPostSinProveedorNotificadoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ProveedorNotificadoId") },
		"null":    func(c map[string]interface{}) { c["ProveedorNotificadoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionNotificacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionNotificacion, cuerpo)
		})
	}
}

// La columna proveedor_notificado_id es NOT NULL: PUT sin ProveedorNotificadoId, con null
// responde 400 sin actualizar.
func TestNotificacionPutSinProveedorNotificadoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ProveedorNotificadoId") },
		"null":    func(c map[string]interface{}) { c["ProveedorNotificadoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionNotificacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionNotificacion, cuerpo)
		})
	}
}

// La columna tercero_notifica_id es NOT NULL: POST sin TerceroNotificaId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestNotificacionPostSinTerceroNotificaId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "TerceroNotificaId") },
		"null":    func(c map[string]interface{}) { c["TerceroNotificaId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionNotificacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionNotificacion, cuerpo)
		})
	}
}

// La columna tercero_notifica_id es NOT NULL: PUT sin TerceroNotificaId, con null
// responde 400 sin actualizar.
func TestNotificacionPutSinTerceroNotificaId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "TerceroNotificaId") },
		"null":    func(c map[string]interface{}) { c["TerceroNotificaId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionNotificacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionNotificacion, cuerpo)
		})
	}
}

// La columna estado_envio_id es NOT NULL: POST sin EstadoEnvioId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestNotificacionPostSinEstadoEnvioId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "EstadoEnvioId") },
		"null":    func(c map[string]interface{}) { c["EstadoEnvioId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionNotificacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionNotificacion, cuerpo)
		})
	}
}

// La columna estado_envio_id es NOT NULL: PUT sin EstadoEnvioId, con null
// responde 400 sin actualizar.
func TestNotificacionPutSinEstadoEnvioId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "EstadoEnvioId") },
		"null":    func(c map[string]interface{}) { c["EstadoEnvioId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionNotificacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionNotificacion, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestNotificacionPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionNotificacion(t)
	cuerpo["ProveedorNotificadoId"] = 0
	cuerpo["TerceroNotificaId"] = 0
	cuerpo["EstadoEnvioId"] = 0
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionNotificacion+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
