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
// HistDetSolicitudCotizacionController (/v1/hist_det_solicitud_cotizacion), sin base de datos.

const rutaValidacionHistDetSolicitudCotizacion = "/v1/hist_det_solicitud_cotizacion"

// cuerpoValidacionHistDetSolicitudCotizacion retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionHistDetSolicitudCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.HistDetSolicitudCotizacion{
		DetSolicitudCotizacionId: 1,
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
		Nombre:                   "Prueba",
		Descripcion:              "Prueba",
		TipoBienServicioId:       1,
		UnidadMedidaId:           1,
		Cantidad:                 1.5,
		RegistroActual:           true,
		TipoRegistro:             "Prueba",
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistDetSolicitudCotizacionGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionHistDetSolicitudCotizacion, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistDetSolicitudCotizacionPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionHistDetSolicitudCotizacion, helpers.JSON(t, cuerpoValidacionHistDetSolicitudCotizacion(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistDetSolicitudCotizacionDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionHistDetSolicitudCotizacion, "")
}

// La columna det_solicitud_cotizacion_id es NOT NULL: POST sin DetSolicitudCotizacionId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetSolicitudCotizacionPostSinDetSolicitudCotizacionId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "DetSolicitudCotizacionId") },
		"null":    func(c map[string]interface{}) { c["DetSolicitudCotizacionId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna det_solicitud_cotizacion_id es NOT NULL: PUT sin DetSolicitudCotizacionId, con null
// responde 400 sin actualizar.
func TestHistDetSolicitudCotizacionPutSinDetSolicitudCotizacionId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "DetSolicitudCotizacionId") },
		"null":    func(c map[string]interface{}) { c["DetSolicitudCotizacionId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna nombre es NOT NULL: POST sin Nombre, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetSolicitudCotizacionPostSinNombre(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Nombre") },
		"null":     func(c map[string]interface{}) { c["Nombre"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Nombre"] = "" },
		"espacios": func(c map[string]interface{}) { c["Nombre"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna nombre es NOT NULL: PUT sin Nombre, con null o vacío
// responde 400 sin actualizar.
func TestHistDetSolicitudCotizacionPutSinNombre(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Nombre") },
		"null":     func(c map[string]interface{}) { c["Nombre"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Nombre"] = "" },
		"espacios": func(c map[string]interface{}) { c["Nombre"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna descripcion es NOT NULL: POST sin Descripcion, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetSolicitudCotizacionPostSinDescripcion(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Descripcion") },
		"null":     func(c map[string]interface{}) { c["Descripcion"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Descripcion"] = "" },
		"espacios": func(c map[string]interface{}) { c["Descripcion"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna descripcion es NOT NULL: PUT sin Descripcion, con null o vacío
// responde 400 sin actualizar.
func TestHistDetSolicitudCotizacionPutSinDescripcion(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Descripcion") },
		"null":     func(c map[string]interface{}) { c["Descripcion"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Descripcion"] = "" },
		"espacios": func(c map[string]interface{}) { c["Descripcion"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna tipo_bien_servicio_id es NOT NULL: POST sin TipoBienServicioId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetSolicitudCotizacionPostSinTipoBienServicioId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "TipoBienServicioId") },
		"null":    func(c map[string]interface{}) { c["TipoBienServicioId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna tipo_bien_servicio_id es NOT NULL: PUT sin TipoBienServicioId, con null
// responde 400 sin actualizar.
func TestHistDetSolicitudCotizacionPutSinTipoBienServicioId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "TipoBienServicioId") },
		"null":    func(c map[string]interface{}) { c["TipoBienServicioId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna unidad_medida_id es NOT NULL: POST sin UnidadMedidaId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetSolicitudCotizacionPostSinUnidadMedidaId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "UnidadMedidaId") },
		"null":    func(c map[string]interface{}) { c["UnidadMedidaId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna unidad_medida_id es NOT NULL: PUT sin UnidadMedidaId, con null
// responde 400 sin actualizar.
func TestHistDetSolicitudCotizacionPutSinUnidadMedidaId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "UnidadMedidaId") },
		"null":    func(c map[string]interface{}) { c["UnidadMedidaId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna cantidad es NOT NULL: POST sin Cantidad, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetSolicitudCotizacionPostSinCantidad(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "Cantidad") },
		"null":    func(c map[string]interface{}) { c["Cantidad"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna cantidad es NOT NULL: PUT sin Cantidad, con null
// responde 400 sin actualizar.
func TestHistDetSolicitudCotizacionPutSinCantidad(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "Cantidad") },
		"null":    func(c map[string]interface{}) { c["Cantidad"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna registro_actual es NOT NULL: POST sin RegistroActual, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetSolicitudCotizacionPostSinRegistroActual(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "RegistroActual") },
		"null":    func(c map[string]interface{}) { c["RegistroActual"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna registro_actual es NOT NULL: PUT sin RegistroActual, con null
// responde 400 sin actualizar.
func TestHistDetSolicitudCotizacionPutSinRegistroActual(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "RegistroActual") },
		"null":    func(c map[string]interface{}) { c["RegistroActual"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna tipo_registro es NOT NULL: POST sin TipoRegistro, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistDetSolicitudCotizacionPostSinTipoRegistro(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "TipoRegistro") },
		"null":     func(c map[string]interface{}) { c["TipoRegistro"] = nil },
		"vacio":    func(c map[string]interface{}) { c["TipoRegistro"] = "" },
		"espacios": func(c map[string]interface{}) { c["TipoRegistro"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna tipo_registro es NOT NULL: PUT sin TipoRegistro, con null o vacío
// responde 400 sin actualizar.
func TestHistDetSolicitudCotizacionPutSinTipoRegistro(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "TipoRegistro") },
		"null":     func(c map[string]interface{}) { c["TipoRegistro"] = nil },
		"vacio":    func(c map[string]interface{}) { c["TipoRegistro"] = "" },
		"espacios": func(c map[string]interface{}) { c["TipoRegistro"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistDetSolicitudCotizacion, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestHistDetSolicitudCotizacionPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
	cuerpo["DetSolicitudCotizacionId"] = 0
	cuerpo["TipoBienServicioId"] = 0
	cuerpo["UnidadMedidaId"] = 0
	cuerpo["Cantidad"] = 0
	cuerpo["RegistroActual"] = false
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionHistDetSolicitudCotizacion+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
