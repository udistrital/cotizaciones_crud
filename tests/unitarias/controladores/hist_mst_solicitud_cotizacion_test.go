// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas unitarias del POST de HistMstSolicitudCotizacionController (/v1/hist_mst_solicitud_cotizacion), sin base de datos.
// Cada escenario aplica a todos los campos escalares opcionales de la entidad.

const rutaHistMstSolicitudCotizacion = "/v1/hist_mst_solicitud_cotizacion"

// ejemploHistMstSolicitudCotizacion retorna la entidad con todos sus campos de ejemplo.
func ejemploHistMstSolicitudCotizacion() *models.HistMstSolicitudCotizacion {
	return &models.HistMstSolicitudCotizacion{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
		NumeroSolicitud:          helpers.Ptr("Prueba"),
		Vigencia:                 helpers.Ptr(int16(2026)),
		Consecutivo:              helpers.Ptr(1),
		FechaProcesamiento:       time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
		EstadoId:                 1,
		UnidadEjecutoraId:        1,
		TipoNecesidadId:          1,
		Titulo:                   "Prueba",
		FechaApertura:            models.NuevaFecha(2026, time.January, 15),
		FechaCierre:              models.NuevaFecha(2026, time.January, 15),
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
	}
}

func cuerpoHistMstSolicitudCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, ejemploHistMstSolicitudCotizacion())
}

// TestPostHistMstSolicitudCotizacionOpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPostHistMstSolicitudCotizacionOpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpoHistMstSolicitudCotizacion(t)
	delete(cuerpo, "NumeroSolicitud")
	delete(cuerpo, "Vigencia")
	delete(cuerpo, "Consecutivo")
	delete(cuerpo, "PlanAccionId")
	delete(cuerpo, "MetodoSeleccionId")
	delete(cuerpo, "NumeroNecesidad")
	delete(cuerpo, "TipoContratoId")
	delete(cuerpo, "CriterioSeleccion")
	delete(cuerpo, "PlazoEjecucion")
	helpers.ProbarPostOpcionales(t, rutaHistMstSolicitudCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "NumeroSolicitud", Columna: "numero_solicitud", Valor: nil},
		{Campo: "Vigencia", Columna: "vigencia", Valor: nil},
		{Campo: "Consecutivo", Columna: "consecutivo", Valor: nil},
		{Campo: "PlanAccionId", Columna: "plan_accion_id", Valor: nil},
		{Campo: "MetodoSeleccionId", Columna: "metodo_seleccion_id", Valor: nil},
		{Campo: "NumeroNecesidad", Columna: "numero_necesidad", Valor: nil},
		{Campo: "TipoContratoId", Columna: "tipo_contrato_id", Valor: nil},
		{Campo: "CriterioSeleccion", Columna: "criterio_seleccion", Valor: nil},
		{Campo: "PlazoEjecucion", Columna: "plazo_ejecucion", Valor: nil},
	})
}

// TestPostHistMstSolicitudCotizacionOpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPostHistMstSolicitudCotizacionOpcionalesNull(t *testing.T) {
	cuerpo := cuerpoHistMstSolicitudCotizacion(t)
	cuerpo["NumeroSolicitud"] = nil
	cuerpo["Vigencia"] = nil
	cuerpo["Consecutivo"] = nil
	cuerpo["PlanAccionId"] = nil
	cuerpo["MetodoSeleccionId"] = nil
	cuerpo["NumeroNecesidad"] = nil
	cuerpo["TipoContratoId"] = nil
	cuerpo["CriterioSeleccion"] = nil
	cuerpo["PlazoEjecucion"] = nil
	helpers.ProbarPostOpcionales(t, rutaHistMstSolicitudCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "NumeroSolicitud", Columna: "numero_solicitud", Valor: nil},
		{Campo: "Vigencia", Columna: "vigencia", Valor: nil},
		{Campo: "Consecutivo", Columna: "consecutivo", Valor: nil},
		{Campo: "PlanAccionId", Columna: "plan_accion_id", Valor: nil},
		{Campo: "MetodoSeleccionId", Columna: "metodo_seleccion_id", Valor: nil},
		{Campo: "NumeroNecesidad", Columna: "numero_necesidad", Valor: nil},
		{Campo: "TipoContratoId", Columna: "tipo_contrato_id", Valor: nil},
		{Campo: "CriterioSeleccion", Columna: "criterio_seleccion", Valor: nil},
		{Campo: "PlazoEjecucion", Columna: "plazo_ejecucion", Valor: nil},
	})
}

// TestPostHistMstSolicitudCotizacionOpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPostHistMstSolicitudCotizacionOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploHistMstSolicitudCotizacion()
	helpers.ProbarPostOpcionales(t, rutaHistMstSolicitudCotizacion, cuerpoHistMstSolicitudCotizacion(t), []helpers.CampoOpcional{
		{Campo: "NumeroSolicitud", Columna: "numero_solicitud", Valor: helpers.Desreferenciar(ejemplo.NumeroSolicitud)},
		{Campo: "Vigencia", Columna: "vigencia", Valor: helpers.Desreferenciar(ejemplo.Vigencia)},
		{Campo: "Consecutivo", Columna: "consecutivo", Valor: helpers.Desreferenciar(ejemplo.Consecutivo)},
		{Campo: "PlanAccionId", Columna: "plan_accion_id", Valor: helpers.Desreferenciar(ejemplo.PlanAccionId)},
		{Campo: "MetodoSeleccionId", Columna: "metodo_seleccion_id", Valor: helpers.Desreferenciar(ejemplo.MetodoSeleccionId)},
		{Campo: "NumeroNecesidad", Columna: "numero_necesidad", Valor: helpers.Desreferenciar(ejemplo.NumeroNecesidad)},
		{Campo: "TipoContratoId", Columna: "tipo_contrato_id", Valor: helpers.Desreferenciar(ejemplo.TipoContratoId)},
		{Campo: "CriterioSeleccion", Columna: "criterio_seleccion", Valor: helpers.Desreferenciar(ejemplo.CriterioSeleccion)},
		{Campo: "PlazoEjecucion", Columna: "plazo_ejecucion", Valor: helpers.Desreferenciar(ejemplo.PlazoEjecucion)},
	})
}

// TestGetHistMstSolicitudCotizacionOpcionalesNull: la base de datos devuelve NULL en cada campo opcional.
func TestGetHistMstSolicitudCotizacionOpcionalesNull(t *testing.T) {
	helpers.ProbarGetOpcionales(t, rutaHistMstSolicitudCotizacion, []helpers.CampoOpcional{
		{Campo: "NumeroSolicitud", Columna: "numero_solicitud", Valor: nil},
		{Campo: "Vigencia", Columna: "vigencia", Valor: nil},
		{Campo: "Consecutivo", Columna: "consecutivo", Valor: nil},
		{Campo: "PlanAccionId", Columna: "plan_accion_id", Valor: nil},
		{Campo: "MetodoSeleccionId", Columna: "metodo_seleccion_id", Valor: nil},
		{Campo: "NumeroNecesidad", Columna: "numero_necesidad", Valor: nil},
		{Campo: "TipoContratoId", Columna: "tipo_contrato_id", Valor: nil},
		{Campo: "CriterioSeleccion", Columna: "criterio_seleccion", Valor: nil},
		{Campo: "PlazoEjecucion", Columna: "plazo_ejecucion", Valor: nil},
	})
}

// TestGetHistMstSolicitudCotizacionOpcionalesConValor: la base de datos devuelve el valor de cada campo opcional.
func TestGetHistMstSolicitudCotizacionOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploHistMstSolicitudCotizacion()
	helpers.ProbarGetOpcionales(t, rutaHistMstSolicitudCotizacion, []helpers.CampoOpcional{
		{Campo: "NumeroSolicitud", Columna: "numero_solicitud", Valor: helpers.Desreferenciar(ejemplo.NumeroSolicitud)},
		{Campo: "Vigencia", Columna: "vigencia", Valor: helpers.Desreferenciar(ejemplo.Vigencia)},
		{Campo: "Consecutivo", Columna: "consecutivo", Valor: helpers.Desreferenciar(ejemplo.Consecutivo)},
		{Campo: "PlanAccionId", Columna: "plan_accion_id", Valor: helpers.Desreferenciar(ejemplo.PlanAccionId)},
		{Campo: "MetodoSeleccionId", Columna: "metodo_seleccion_id", Valor: helpers.Desreferenciar(ejemplo.MetodoSeleccionId)},
		{Campo: "NumeroNecesidad", Columna: "numero_necesidad", Valor: helpers.Desreferenciar(ejemplo.NumeroNecesidad)},
		{Campo: "TipoContratoId", Columna: "tipo_contrato_id", Valor: helpers.Desreferenciar(ejemplo.TipoContratoId)},
		{Campo: "CriterioSeleccion", Columna: "criterio_seleccion", Valor: helpers.Desreferenciar(ejemplo.CriterioSeleccion)},
		{Campo: "PlazoEjecucion", Columna: "plazo_ejecucion", Valor: helpers.Desreferenciar(ejemplo.PlazoEjecucion)},
	})
}

// TestPutHistMstSolicitudCotizacionOpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPutHistMstSolicitudCotizacionOpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpoHistMstSolicitudCotizacion(t)
	delete(cuerpo, "NumeroSolicitud")
	delete(cuerpo, "Vigencia")
	delete(cuerpo, "Consecutivo")
	delete(cuerpo, "PlanAccionId")
	delete(cuerpo, "MetodoSeleccionId")
	delete(cuerpo, "NumeroNecesidad")
	delete(cuerpo, "TipoContratoId")
	delete(cuerpo, "CriterioSeleccion")
	delete(cuerpo, "PlazoEjecucion")
	helpers.ProbarPutOpcionales(t, rutaHistMstSolicitudCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "NumeroSolicitud", Columna: "numero_solicitud", Valor: nil},
		{Campo: "Vigencia", Columna: "vigencia", Valor: nil},
		{Campo: "Consecutivo", Columna: "consecutivo", Valor: nil},
		{Campo: "PlanAccionId", Columna: "plan_accion_id", Valor: nil},
		{Campo: "MetodoSeleccionId", Columna: "metodo_seleccion_id", Valor: nil},
		{Campo: "NumeroNecesidad", Columna: "numero_necesidad", Valor: nil},
		{Campo: "TipoContratoId", Columna: "tipo_contrato_id", Valor: nil},
		{Campo: "CriterioSeleccion", Columna: "criterio_seleccion", Valor: nil},
		{Campo: "PlazoEjecucion", Columna: "plazo_ejecucion", Valor: nil},
	})
}

// TestPutHistMstSolicitudCotizacionOpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPutHistMstSolicitudCotizacionOpcionalesNull(t *testing.T) {
	cuerpo := cuerpoHistMstSolicitudCotizacion(t)
	cuerpo["NumeroSolicitud"] = nil
	cuerpo["Vigencia"] = nil
	cuerpo["Consecutivo"] = nil
	cuerpo["PlanAccionId"] = nil
	cuerpo["MetodoSeleccionId"] = nil
	cuerpo["NumeroNecesidad"] = nil
	cuerpo["TipoContratoId"] = nil
	cuerpo["CriterioSeleccion"] = nil
	cuerpo["PlazoEjecucion"] = nil
	helpers.ProbarPutOpcionales(t, rutaHistMstSolicitudCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "NumeroSolicitud", Columna: "numero_solicitud", Valor: nil},
		{Campo: "Vigencia", Columna: "vigencia", Valor: nil},
		{Campo: "Consecutivo", Columna: "consecutivo", Valor: nil},
		{Campo: "PlanAccionId", Columna: "plan_accion_id", Valor: nil},
		{Campo: "MetodoSeleccionId", Columna: "metodo_seleccion_id", Valor: nil},
		{Campo: "NumeroNecesidad", Columna: "numero_necesidad", Valor: nil},
		{Campo: "TipoContratoId", Columna: "tipo_contrato_id", Valor: nil},
		{Campo: "CriterioSeleccion", Columna: "criterio_seleccion", Valor: nil},
		{Campo: "PlazoEjecucion", Columna: "plazo_ejecucion", Valor: nil},
	})
}

// TestPutHistMstSolicitudCotizacionOpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPutHistMstSolicitudCotizacionOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploHistMstSolicitudCotizacion()
	helpers.ProbarPutOpcionales(t, rutaHistMstSolicitudCotizacion, cuerpoHistMstSolicitudCotizacion(t), []helpers.CampoOpcional{
		{Campo: "NumeroSolicitud", Columna: "numero_solicitud", Valor: helpers.Desreferenciar(ejemplo.NumeroSolicitud)},
		{Campo: "Vigencia", Columna: "vigencia", Valor: helpers.Desreferenciar(ejemplo.Vigencia)},
		{Campo: "Consecutivo", Columna: "consecutivo", Valor: helpers.Desreferenciar(ejemplo.Consecutivo)},
		{Campo: "PlanAccionId", Columna: "plan_accion_id", Valor: helpers.Desreferenciar(ejemplo.PlanAccionId)},
		{Campo: "MetodoSeleccionId", Columna: "metodo_seleccion_id", Valor: helpers.Desreferenciar(ejemplo.MetodoSeleccionId)},
		{Campo: "NumeroNecesidad", Columna: "numero_necesidad", Valor: helpers.Desreferenciar(ejemplo.NumeroNecesidad)},
		{Campo: "TipoContratoId", Columna: "tipo_contrato_id", Valor: helpers.Desreferenciar(ejemplo.TipoContratoId)},
		{Campo: "CriterioSeleccion", Columna: "criterio_seleccion", Valor: helpers.Desreferenciar(ejemplo.CriterioSeleccion)},
		{Campo: "PlazoEjecucion", Columna: "plazo_ejecucion", Valor: helpers.Desreferenciar(ejemplo.PlazoEjecucion)},
	})
}
