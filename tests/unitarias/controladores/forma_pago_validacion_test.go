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
// FormaPagoController (/v1/forma_pago), sin base de datos.

const rutaValidacionFormaPago = "/v1/forma_pago"

// cuerpoValidacionFormaPago retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionFormaPago(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.FormaPago{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
		TipoFormaPagoId:          1,
		Condicion:                "Prueba",
		PorcentajePago:           1.5,
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestFormaPagoGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionFormaPago, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestFormaPagoPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionFormaPago, helpers.JSON(t, cuerpoValidacionFormaPago(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestFormaPagoDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionFormaPago, "")
}

// La columna tipo_forma_pago_id es NOT NULL: POST sin TipoFormaPagoId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestFormaPagoPostSinTipoFormaPagoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "TipoFormaPagoId") },
		"null":    func(c map[string]interface{}) { c["TipoFormaPagoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionFormaPago(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionFormaPago, cuerpo)
		})
	}
}

// La columna tipo_forma_pago_id es NOT NULL: PUT sin TipoFormaPagoId, con null
// responde 400 sin actualizar.
func TestFormaPagoPutSinTipoFormaPagoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "TipoFormaPagoId") },
		"null":    func(c map[string]interface{}) { c["TipoFormaPagoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionFormaPago(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionFormaPago, cuerpo)
		})
	}
}

// La columna condicion es NOT NULL: POST sin Condicion, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestFormaPagoPostSinCondicion(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Condicion") },
		"null":     func(c map[string]interface{}) { c["Condicion"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Condicion"] = "" },
		"espacios": func(c map[string]interface{}) { c["Condicion"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionFormaPago(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionFormaPago, cuerpo)
		})
	}
}

// La columna condicion es NOT NULL: PUT sin Condicion, con null o vacío
// responde 400 sin actualizar.
func TestFormaPagoPutSinCondicion(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Condicion") },
		"null":     func(c map[string]interface{}) { c["Condicion"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Condicion"] = "" },
		"espacios": func(c map[string]interface{}) { c["Condicion"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionFormaPago(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionFormaPago, cuerpo)
		})
	}
}

// La columna porcentaje_pago es NOT NULL: POST sin PorcentajePago, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestFormaPagoPostSinPorcentajePago(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "PorcentajePago") },
		"null":    func(c map[string]interface{}) { c["PorcentajePago"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionFormaPago(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionFormaPago, cuerpo)
		})
	}
}

// La columna porcentaje_pago es NOT NULL: PUT sin PorcentajePago, con null
// responde 400 sin actualizar.
func TestFormaPagoPutSinPorcentajePago(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "PorcentajePago") },
		"null":    func(c map[string]interface{}) { c["PorcentajePago"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionFormaPago(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionFormaPago, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestFormaPagoPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionFormaPago(t)
	cuerpo["TipoFormaPagoId"] = 0
	cuerpo["PorcentajePago"] = 0
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionFormaPago+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
