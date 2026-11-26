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
// HistMstSolicitudCotizacionController (/v1/hist_mst_solicitud_cotizacion), sin base de datos.

const rutaValidacionHistMstSolicitudCotizacion = "/v1/hist_mst_solicitud_cotizacion"

// cuerpoValidacionHistMstSolicitudCotizacion retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionHistMstSolicitudCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.HistMstSolicitudCotizacion{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
		NumeroSolicitud:          helpers.Ptr("Prueba"),
		Vigencia:                 helpers.Ptr(int16(2026)),
		Consecutivo:              helpers.Ptr(1),
		FechaProcesamiento:       time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
		EstadoId:                 1,
		UnidadEjecutoraId:        1,
		TipoNecesidadId:          1,
		Titulo:                   "Prueba",
		FechaApertura:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
		FechaCierre:              time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
		Objetivo:                 "Prueba",
		Requisitos:               "Prueba",
		Observaciones:            "Prueba",
		OrdenadorGastoId:         1,
		DependenciaSolicitanteId: 1,
		JustificacionSeleccion:   "Prueba",
		MedioPagoId:              1,
		PlanAccionId:             helpers.Ptr(1),
		ResponsableTerceroId:     1,
		MetodoSeleccionId:        helpers.Ptr(1),
		NumeroNecesidad:          helpers.Ptr(1.5),
		DependenciaDestinoId:     1,
		TipoContratoId:           helpers.Ptr(1),
		CriterioSeleccion:        helpers.Ptr("Prueba"),
		PlazoEjecucion:           helpers.Ptr("Prueba"),
		RegistroActual:           true,
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistMstSolicitudCotizacionGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionHistMstSolicitudCotizacion, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistMstSolicitudCotizacionPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionHistMstSolicitudCotizacion, helpers.JSON(t, cuerpoValidacionHistMstSolicitudCotizacion(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestHistMstSolicitudCotizacionDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionHistMstSolicitudCotizacion, "")
}

// La columna fecha_procesamiento es NOT NULL: POST sin FechaProcesamiento, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinFechaProcesamiento(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "FechaProcesamiento") },
		"null":    func(c map[string]interface{}) { c["FechaProcesamiento"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna fecha_procesamiento es NOT NULL: PUT sin FechaProcesamiento, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinFechaProcesamiento(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "FechaProcesamiento") },
		"null":    func(c map[string]interface{}) { c["FechaProcesamiento"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna estado_id es NOT NULL: POST sin EstadoId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinEstadoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "EstadoId") },
		"null":    func(c map[string]interface{}) { c["EstadoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna estado_id es NOT NULL: PUT sin EstadoId, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinEstadoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "EstadoId") },
		"null":    func(c map[string]interface{}) { c["EstadoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna unidad_ejecutora_id es NOT NULL: POST sin UnidadEjecutoraId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinUnidadEjecutoraId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "UnidadEjecutoraId") },
		"null":    func(c map[string]interface{}) { c["UnidadEjecutoraId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna unidad_ejecutora_id es NOT NULL: PUT sin UnidadEjecutoraId, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinUnidadEjecutoraId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "UnidadEjecutoraId") },
		"null":    func(c map[string]interface{}) { c["UnidadEjecutoraId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna tipo_necesidad_id es NOT NULL: POST sin TipoNecesidadId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinTipoNecesidadId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "TipoNecesidadId") },
		"null":    func(c map[string]interface{}) { c["TipoNecesidadId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna tipo_necesidad_id es NOT NULL: PUT sin TipoNecesidadId, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinTipoNecesidadId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "TipoNecesidadId") },
		"null":    func(c map[string]interface{}) { c["TipoNecesidadId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna titulo es NOT NULL: POST sin Titulo, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinTitulo(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Titulo") },
		"null":     func(c map[string]interface{}) { c["Titulo"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Titulo"] = "" },
		"espacios": func(c map[string]interface{}) { c["Titulo"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna titulo es NOT NULL: PUT sin Titulo, con null o vacío
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinTitulo(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Titulo") },
		"null":     func(c map[string]interface{}) { c["Titulo"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Titulo"] = "" },
		"espacios": func(c map[string]interface{}) { c["Titulo"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna fecha_apertura es NOT NULL: POST sin FechaApertura, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinFechaApertura(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "FechaApertura") },
		"null":    func(c map[string]interface{}) { c["FechaApertura"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna fecha_apertura es NOT NULL: PUT sin FechaApertura, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinFechaApertura(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "FechaApertura") },
		"null":    func(c map[string]interface{}) { c["FechaApertura"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna fecha_cierre es NOT NULL: POST sin FechaCierre, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinFechaCierre(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "FechaCierre") },
		"null":    func(c map[string]interface{}) { c["FechaCierre"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna fecha_cierre es NOT NULL: PUT sin FechaCierre, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinFechaCierre(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "FechaCierre") },
		"null":    func(c map[string]interface{}) { c["FechaCierre"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna objetivo es NOT NULL: POST sin Objetivo, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinObjetivo(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Objetivo") },
		"null":     func(c map[string]interface{}) { c["Objetivo"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Objetivo"] = "" },
		"espacios": func(c map[string]interface{}) { c["Objetivo"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna objetivo es NOT NULL: PUT sin Objetivo, con null o vacío
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinObjetivo(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Objetivo") },
		"null":     func(c map[string]interface{}) { c["Objetivo"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Objetivo"] = "" },
		"espacios": func(c map[string]interface{}) { c["Objetivo"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna requisitos es NOT NULL: POST sin Requisitos, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinRequisitos(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Requisitos") },
		"null":     func(c map[string]interface{}) { c["Requisitos"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Requisitos"] = "" },
		"espacios": func(c map[string]interface{}) { c["Requisitos"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna requisitos es NOT NULL: PUT sin Requisitos, con null o vacío
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinRequisitos(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Requisitos") },
		"null":     func(c map[string]interface{}) { c["Requisitos"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Requisitos"] = "" },
		"espacios": func(c map[string]interface{}) { c["Requisitos"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna observaciones es NOT NULL: POST sin Observaciones, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinObservaciones(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Observaciones") },
		"null":     func(c map[string]interface{}) { c["Observaciones"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Observaciones"] = "" },
		"espacios": func(c map[string]interface{}) { c["Observaciones"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna observaciones es NOT NULL: PUT sin Observaciones, con null o vacío
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinObservaciones(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Observaciones") },
		"null":     func(c map[string]interface{}) { c["Observaciones"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Observaciones"] = "" },
		"espacios": func(c map[string]interface{}) { c["Observaciones"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna ordenador_gasto_id es NOT NULL: POST sin OrdenadorGastoId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinOrdenadorGastoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "OrdenadorGastoId") },
		"null":    func(c map[string]interface{}) { c["OrdenadorGastoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna ordenador_gasto_id es NOT NULL: PUT sin OrdenadorGastoId, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinOrdenadorGastoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "OrdenadorGastoId") },
		"null":    func(c map[string]interface{}) { c["OrdenadorGastoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna dependencia_solicitante_id es NOT NULL: POST sin DependenciaSolicitanteId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinDependenciaSolicitanteId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "DependenciaSolicitanteId") },
		"null":    func(c map[string]interface{}) { c["DependenciaSolicitanteId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna dependencia_solicitante_id es NOT NULL: PUT sin DependenciaSolicitanteId, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinDependenciaSolicitanteId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "DependenciaSolicitanteId") },
		"null":    func(c map[string]interface{}) { c["DependenciaSolicitanteId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna justificacion_seleccion es NOT NULL: POST sin JustificacionSeleccion, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinJustificacionSeleccion(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "JustificacionSeleccion") },
		"null":     func(c map[string]interface{}) { c["JustificacionSeleccion"] = nil },
		"vacio":    func(c map[string]interface{}) { c["JustificacionSeleccion"] = "" },
		"espacios": func(c map[string]interface{}) { c["JustificacionSeleccion"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna justificacion_seleccion es NOT NULL: PUT sin JustificacionSeleccion, con null o vacío
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinJustificacionSeleccion(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "JustificacionSeleccion") },
		"null":     func(c map[string]interface{}) { c["JustificacionSeleccion"] = nil },
		"vacio":    func(c map[string]interface{}) { c["JustificacionSeleccion"] = "" },
		"espacios": func(c map[string]interface{}) { c["JustificacionSeleccion"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna medio_pago_id es NOT NULL: POST sin MedioPagoId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinMedioPagoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "MedioPagoId") },
		"null":    func(c map[string]interface{}) { c["MedioPagoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna medio_pago_id es NOT NULL: PUT sin MedioPagoId, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinMedioPagoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "MedioPagoId") },
		"null":    func(c map[string]interface{}) { c["MedioPagoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna responsable_tercero_id es NOT NULL: POST sin ResponsableTerceroId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinResponsableTerceroId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ResponsableTerceroId") },
		"null":    func(c map[string]interface{}) { c["ResponsableTerceroId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna responsable_tercero_id es NOT NULL: PUT sin ResponsableTerceroId, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinResponsableTerceroId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "ResponsableTerceroId") },
		"null":    func(c map[string]interface{}) { c["ResponsableTerceroId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna dependencia_destino_id es NOT NULL: POST sin DependenciaDestinoId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinDependenciaDestinoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "DependenciaDestinoId") },
		"null":    func(c map[string]interface{}) { c["DependenciaDestinoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna dependencia_destino_id es NOT NULL: PUT sin DependenciaDestinoId, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinDependenciaDestinoId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "DependenciaDestinoId") },
		"null":    func(c map[string]interface{}) { c["DependenciaDestinoId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna registro_actual es NOT NULL: POST sin RegistroActual, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestHistMstSolicitudCotizacionPostSinRegistroActual(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "RegistroActual") },
		"null":    func(c map[string]interface{}) { c["RegistroActual"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// La columna registro_actual es NOT NULL: PUT sin RegistroActual, con null
// responde 400 sin actualizar.
func TestHistMstSolicitudCotizacionPutSinRegistroActual(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "RegistroActual") },
		"null":    func(c map[string]interface{}) { c["RegistroActual"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestHistMstSolicitudCotizacionPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
	cuerpo["EstadoId"] = 0
	cuerpo["UnidadEjecutoraId"] = 0
	cuerpo["TipoNecesidadId"] = 0
	cuerpo["OrdenadorGastoId"] = 0
	cuerpo["DependenciaSolicitanteId"] = 0
	cuerpo["MedioPagoId"] = 0
	cuerpo["ResponsableTerceroId"] = 0
	cuerpo["DependenciaDestinoId"] = 0
	cuerpo["RegistroActual"] = false
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionHistMstSolicitudCotizacion+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
