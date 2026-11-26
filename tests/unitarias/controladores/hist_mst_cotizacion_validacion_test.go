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
// HistMstCotizacionController (/v1/hist_mst_cotizacion), sin base de datos.

const rutaValidacionHistMstCotizacion = "/v1/hist_mst_cotizacion"

// cuerpoValidacionHistMstCotizacion retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionHistMstCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.HistMstCotizacion{
		MstCotizacionId:       &models.MstCotizacion{Id: 1},
		FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
		Descripcion:           helpers.Ptr("Prueba"),
		InformacionEntrega:    helpers.Ptr("Prueba"),
		Descuentos:            helpers.Ptr("Prueba"),
		Observaciones:         helpers.Ptr("Prueba"),
		FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
		SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
		RegistroActual:        true,
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistMstCotizacionGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionHistMstCotizacion, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistMstCotizacionPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionHistMstCotizacion, helpers.JSON(t, cuerpoValidacionHistMstCotizacion(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistMstCotizacionDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionHistMstCotizacion, "")
}

// La columna fecha_envio es NOT NULL: POST sin FechaEnvio, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstCotizacionPostSinFechaEnvio(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "FechaEnvio") },
		"null":    func(c map[string]interface{}) { c["FechaEnvio"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstCotizacion, cuerpo)
		})
	}
}

// La columna fecha_envio es NOT NULL: PUT sin FechaEnvio, con null
// responde 400 sin actualizar.
func TestHistMstCotizacionPutSinFechaEnvio(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "FechaEnvio") },
		"null":    func(c map[string]interface{}) { c["FechaEnvio"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstCotizacion, cuerpo)
		})
	}
}

// La columna fecha_vencimiento es NOT NULL: POST sin FechaVencimiento, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstCotizacionPostSinFechaVencimiento(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "FechaVencimiento") },
		"null":    func(c map[string]interface{}) { c["FechaVencimiento"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstCotizacion, cuerpo)
		})
	}
}

// La columna fecha_vencimiento es NOT NULL: PUT sin FechaVencimiento, con null
// responde 400 sin actualizar.
func TestHistMstCotizacionPutSinFechaVencimiento(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "FechaVencimiento") },
		"null":    func(c map[string]interface{}) { c["FechaVencimiento"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstCotizacion, cuerpo)
		})
	}
}

// La columna soporte_cotizacion es NOT NULL: POST sin SoporteCotizacion, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstCotizacionPostSinSoporteCotizacion(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "SoporteCotizacion") },
		"null":     func(c map[string]interface{}) { c["SoporteCotizacion"] = nil },
		"vacio":    func(c map[string]interface{}) { c["SoporteCotizacion"] = "" },
		"espacios": func(c map[string]interface{}) { c["SoporteCotizacion"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstCotizacion, cuerpo)
		})
	}
}

// La columna soporte_cotizacion es NOT NULL: PUT sin SoporteCotizacion, con null o vacío
// responde 400 sin actualizar.
func TestHistMstCotizacionPutSinSoporteCotizacion(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "SoporteCotizacion") },
		"null":     func(c map[string]interface{}) { c["SoporteCotizacion"] = nil },
		"vacio":    func(c map[string]interface{}) { c["SoporteCotizacion"] = "" },
		"espacios": func(c map[string]interface{}) { c["SoporteCotizacion"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstCotizacion, cuerpo)
		})
	}
}

// La columna registro_actual es NOT NULL: POST sin RegistroActual, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstCotizacionPostSinRegistroActual(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "RegistroActual") },
		"null":    func(c map[string]interface{}) { c["RegistroActual"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstCotizacion, cuerpo)
		})
	}
}

// La columna registro_actual es NOT NULL: PUT sin RegistroActual, con null
// responde 400 sin actualizar.
func TestHistMstCotizacionPutSinRegistroActual(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "RegistroActual") },
		"null":    func(c map[string]interface{}) { c["RegistroActual"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstCotizacion, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestHistMstCotizacionPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionHistMstCotizacion(t)
	cuerpo["RegistroActual"] = false
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionHistMstCotizacion+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
