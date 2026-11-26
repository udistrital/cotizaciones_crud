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
// HistDetCotizacionController (/v1/hist_det_cotizacion), sin base de datos.

const rutaValidacionHistDetCotizacion = "/v1/hist_det_cotizacion"

// cuerpoValidacionHistDetCotizacion retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionHistDetCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.HistDetCotizacion{
		DetCotizacionId:          1,
		MstCotizacionId:          &models.MstCotizacion{Id: 1},
		DetSolicitudCotizacionId: 1,
		ValorUnitario:            helpers.Ptr(1.5),
		Iva:                      1.5,
		FichaTecnica:             "Prueba",
		RegistroActual:           true,
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistDetCotizacionGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionHistDetCotizacion, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistDetCotizacionPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionHistDetCotizacion, helpers.JSON(t, cuerpoValidacionHistDetCotizacion(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistDetCotizacionDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionHistDetCotizacion, "")
}

// La columna det_cotizacion_id es NOT NULL: POST sin DetCotizacionId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetCotizacionPostSinDetCotizacionId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "DetCotizacionId") },
		"null":    func(c map[string]interface{}) { c["DetCotizacionId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetCotizacion, cuerpo)
		})
	}
}

// La columna det_cotizacion_id es NOT NULL: PUT sin DetCotizacionId, con null
// responde 400 sin actualizar.
func TestHistDetCotizacionPutSinDetCotizacionId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "DetCotizacionId") },
		"null":    func(c map[string]interface{}) { c["DetCotizacionId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetCotizacion, cuerpo)
		})
	}
}

// La columna det_solicitud_cotizacion_id es NOT NULL: POST sin DetSolicitudCotizacionId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetCotizacionPostSinDetSolicitudCotizacionId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "DetSolicitudCotizacionId") },
		"null":    func(c map[string]interface{}) { c["DetSolicitudCotizacionId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetCotizacion, cuerpo)
		})
	}
}

// La columna det_solicitud_cotizacion_id es NOT NULL: PUT sin DetSolicitudCotizacionId, con null
// responde 400 sin actualizar.
func TestHistDetCotizacionPutSinDetSolicitudCotizacionId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "DetSolicitudCotizacionId") },
		"null":    func(c map[string]interface{}) { c["DetSolicitudCotizacionId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetCotizacion, cuerpo)
		})
	}
}

// La columna iva es NOT NULL: POST sin Iva, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetCotizacionPostSinIva(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "Iva") },
		"null":    func(c map[string]interface{}) { c["Iva"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetCotizacion, cuerpo)
		})
	}
}

// La columna iva es NOT NULL: PUT sin Iva, con null
// responde 400 sin actualizar.
func TestHistDetCotizacionPutSinIva(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "Iva") },
		"null":    func(c map[string]interface{}) { c["Iva"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetCotizacion, cuerpo)
		})
	}
}

// La columna ficha_tecnica es NOT NULL: POST sin FichaTecnica, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetCotizacionPostSinFichaTecnica(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "FichaTecnica") },
		"null":     func(c map[string]interface{}) { c["FichaTecnica"] = nil },
		"vacio":    func(c map[string]interface{}) { c["FichaTecnica"] = "" },
		"espacios": func(c map[string]interface{}) { c["FichaTecnica"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetCotizacion, cuerpo)
		})
	}
}

// La columna ficha_tecnica es NOT NULL: PUT sin FichaTecnica, con null o vacío
// responde 400 sin actualizar.
func TestHistDetCotizacionPutSinFichaTecnica(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "FichaTecnica") },
		"null":     func(c map[string]interface{}) { c["FichaTecnica"] = nil },
		"vacio":    func(c map[string]interface{}) { c["FichaTecnica"] = "" },
		"espacios": func(c map[string]interface{}) { c["FichaTecnica"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetCotizacion, cuerpo)
		})
	}
}

// La columna registro_actual es NOT NULL: POST sin RegistroActual, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetCotizacionPostSinRegistroActual(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "RegistroActual") },
		"null":    func(c map[string]interface{}) { c["RegistroActual"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetCotizacion, cuerpo)
		})
	}
}

// La columna registro_actual es NOT NULL: PUT sin RegistroActual, con null
// responde 400 sin actualizar.
func TestHistDetCotizacionPutSinRegistroActual(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "RegistroActual") },
		"null":    func(c map[string]interface{}) { c["RegistroActual"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetCotizacion, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestHistDetCotizacionPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionHistDetCotizacion(t)
	cuerpo["DetCotizacionId"] = 0
	cuerpo["DetSolicitudCotizacionId"] = 0
	cuerpo["Iva"] = 0
	cuerpo["RegistroActual"] = false
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionHistDetCotizacion+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
