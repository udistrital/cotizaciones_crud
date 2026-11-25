// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package modelos

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas unitarias de models.AddHistMstSolicitudCotizacion para los campos escalares opcionales, sin base de datos.

// TestHistMstSolicitudCotizacionNumeroSolicitud verifica el valor de numero_solicitud en el INSERT para
// los tres casos del campo NumeroSolicitud.
func TestHistMstSolicitudCotizacionNumeroSolicitud(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
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
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_solicitud", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
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
			NumeroSolicitud:          nil,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_solicitud", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
		}
		esperado := helpers.Desreferenciar(m.NumeroSolicitud)
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_solicitud", esperado)
	})
}

// TestHistMstSolicitudCotizacionNumeroSolicitudGet verifica la lectura de numero_solicitud: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstSolicitudCotizacionNumeroSolicitudGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"numero_solicitud": nil})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "NumeroSolicitud", nil, helpers.Desreferenciar(v.NumeroSolicitud))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr("Prueba"))
		helpers.DefinirFila(map[string]interface{}{"numero_solicitud": esperado})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "NumeroSolicitud", esperado, helpers.Desreferenciar(v.NumeroSolicitud))
	})
}

// TestHistMstSolicitudCotizacionNumeroSolicitudUpdate verifica el valor de numero_solicitud en el UPDATE
// para los tres casos del campo NumeroSolicitud.
func TestHistMstSolicitudCotizacionNumeroSolicitudUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
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
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_solicitud", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
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
			NumeroSolicitud:          nil,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_solicitud", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
		}
		esperado := helpers.Desreferenciar(m.NumeroSolicitud)
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_solicitud", esperado)
	})
}

// TestHistMstSolicitudCotizacionVigencia verifica el valor de vigencia en el INSERT para
// los tres casos del campo Vigencia.
func TestHistMstSolicitudCotizacionVigencia(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
			NumeroSolicitud:          helpers.Ptr("Prueba"),
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
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "vigencia", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
			NumeroSolicitud:          helpers.Ptr("Prueba"),
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
			Vigencia:                 nil,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "vigencia", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
		}
		esperado := helpers.Desreferenciar(m.Vigencia)
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "vigencia", esperado)
	})
}

// TestHistMstSolicitudCotizacionVigenciaGet verifica la lectura de vigencia: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstSolicitudCotizacionVigenciaGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"vigencia": nil})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "Vigencia", nil, helpers.Desreferenciar(v.Vigencia))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr(int16(2026)))
		helpers.DefinirFila(map[string]interface{}{"vigencia": esperado})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "Vigencia", esperado, helpers.Desreferenciar(v.Vigencia))
	})
}

// TestHistMstSolicitudCotizacionVigenciaUpdate verifica el valor de vigencia en el UPDATE
// para los tres casos del campo Vigencia.
func TestHistMstSolicitudCotizacionVigenciaUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
			NumeroSolicitud:          helpers.Ptr("Prueba"),
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
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "vigencia", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
			NumeroSolicitud:          helpers.Ptr("Prueba"),
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
			Vigencia:                 nil,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "vigencia", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
		}
		esperado := helpers.Desreferenciar(m.Vigencia)
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "vigencia", esperado)
	})
}

// TestHistMstSolicitudCotizacionConsecutivo verifica el valor de consecutivo en el INSERT para
// los tres casos del campo Consecutivo.
func TestHistMstSolicitudCotizacionConsecutivo(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
			NumeroSolicitud:          helpers.Ptr("Prueba"),
			Vigencia:                 helpers.Ptr(int16(2026)),
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
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "consecutivo", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
			NumeroSolicitud:          helpers.Ptr("Prueba"),
			Vigencia:                 helpers.Ptr(int16(2026)),
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
			Consecutivo:              nil,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "consecutivo", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
		}
		esperado := helpers.Desreferenciar(m.Consecutivo)
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "consecutivo", esperado)
	})
}

// TestHistMstSolicitudCotizacionConsecutivoGet verifica la lectura de consecutivo: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstSolicitudCotizacionConsecutivoGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"consecutivo": nil})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "Consecutivo", nil, helpers.Desreferenciar(v.Consecutivo))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr(1))
		helpers.DefinirFila(map[string]interface{}{"consecutivo": esperado})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "Consecutivo", esperado, helpers.Desreferenciar(v.Consecutivo))
	})
}

// TestHistMstSolicitudCotizacionConsecutivoUpdate verifica el valor de consecutivo en el UPDATE
// para los tres casos del campo Consecutivo.
func TestHistMstSolicitudCotizacionConsecutivoUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
			NumeroSolicitud:          helpers.Ptr("Prueba"),
			Vigencia:                 helpers.Ptr(int16(2026)),
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
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "consecutivo", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
			MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
			NumeroSolicitud:          helpers.Ptr("Prueba"),
			Vigencia:                 helpers.Ptr(int16(2026)),
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
			Consecutivo:              nil,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "consecutivo", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
		}
		esperado := helpers.Desreferenciar(m.Consecutivo)
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "consecutivo", esperado)
	})
}

// TestHistMstSolicitudCotizacionPlanAccionId verifica el valor de plan_accion_id en el INSERT para
// los tres casos del campo PlanAccionId.
func TestHistMstSolicitudCotizacionPlanAccionId(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plan_accion_id", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
			PlanAccionId:             nil,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plan_accion_id", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
		}
		esperado := helpers.Desreferenciar(m.PlanAccionId)
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plan_accion_id", esperado)
	})
}

// TestHistMstSolicitudCotizacionPlanAccionIdGet verifica la lectura de plan_accion_id: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstSolicitudCotizacionPlanAccionIdGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"plan_accion_id": nil})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "PlanAccionId", nil, helpers.Desreferenciar(v.PlanAccionId))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr(1))
		helpers.DefinirFila(map[string]interface{}{"plan_accion_id": esperado})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "PlanAccionId", esperado, helpers.Desreferenciar(v.PlanAccionId))
	})
}

// TestHistMstSolicitudCotizacionPlanAccionIdUpdate verifica el valor de plan_accion_id en el UPDATE
// para los tres casos del campo PlanAccionId.
func TestHistMstSolicitudCotizacionPlanAccionIdUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plan_accion_id", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
			PlanAccionId:             nil,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plan_accion_id", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
		}
		esperado := helpers.Desreferenciar(m.PlanAccionId)
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plan_accion_id", esperado)
	})
}

// TestHistMstSolicitudCotizacionMetodoSeleccionId verifica el valor de metodo_seleccion_id en el INSERT para
// los tres casos del campo MetodoSeleccionId.
func TestHistMstSolicitudCotizacionMetodoSeleccionId(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
			NumeroNecesidad:          helpers.Ptr(1.5),
			DependenciaDestinoId:     1,
			TipoContratoId:           helpers.Ptr(1),
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "metodo_seleccion_id", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
			NumeroNecesidad:          helpers.Ptr(1.5),
			DependenciaDestinoId:     1,
			TipoContratoId:           helpers.Ptr(1),
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			MetodoSeleccionId:        nil,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "metodo_seleccion_id", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
		}
		esperado := helpers.Desreferenciar(m.MetodoSeleccionId)
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "metodo_seleccion_id", esperado)
	})
}

// TestHistMstSolicitudCotizacionMetodoSeleccionIdGet verifica la lectura de metodo_seleccion_id: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstSolicitudCotizacionMetodoSeleccionIdGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"metodo_seleccion_id": nil})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "MetodoSeleccionId", nil, helpers.Desreferenciar(v.MetodoSeleccionId))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr(1))
		helpers.DefinirFila(map[string]interface{}{"metodo_seleccion_id": esperado})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "MetodoSeleccionId", esperado, helpers.Desreferenciar(v.MetodoSeleccionId))
	})
}

// TestHistMstSolicitudCotizacionMetodoSeleccionIdUpdate verifica el valor de metodo_seleccion_id en el UPDATE
// para los tres casos del campo MetodoSeleccionId.
func TestHistMstSolicitudCotizacionMetodoSeleccionIdUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
			NumeroNecesidad:          helpers.Ptr(1.5),
			DependenciaDestinoId:     1,
			TipoContratoId:           helpers.Ptr(1),
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "metodo_seleccion_id", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
			NumeroNecesidad:          helpers.Ptr(1.5),
			DependenciaDestinoId:     1,
			TipoContratoId:           helpers.Ptr(1),
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			MetodoSeleccionId:        nil,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "metodo_seleccion_id", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
		}
		esperado := helpers.Desreferenciar(m.MetodoSeleccionId)
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "metodo_seleccion_id", esperado)
	})
}

// TestHistMstSolicitudCotizacionNumeroNecesidad verifica el valor de numero_necesidad en el INSERT para
// los tres casos del campo NumeroNecesidad.
func TestHistMstSolicitudCotizacionNumeroNecesidad(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
			DependenciaDestinoId:     1,
			TipoContratoId:           helpers.Ptr(1),
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_necesidad", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
			DependenciaDestinoId:     1,
			TipoContratoId:           helpers.Ptr(1),
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			NumeroNecesidad:          nil,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_necesidad", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
		}
		esperado := helpers.Desreferenciar(m.NumeroNecesidad)
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_necesidad", esperado)
	})
}

// TestHistMstSolicitudCotizacionNumeroNecesidadGet verifica la lectura de numero_necesidad: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstSolicitudCotizacionNumeroNecesidadGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"numero_necesidad": nil})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "NumeroNecesidad", nil, helpers.Desreferenciar(v.NumeroNecesidad))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr(1.5))
		helpers.DefinirFila(map[string]interface{}{"numero_necesidad": esperado})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "NumeroNecesidad", esperado, helpers.Desreferenciar(v.NumeroNecesidad))
	})
}

// TestHistMstSolicitudCotizacionNumeroNecesidadUpdate verifica el valor de numero_necesidad en el UPDATE
// para los tres casos del campo NumeroNecesidad.
func TestHistMstSolicitudCotizacionNumeroNecesidadUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
			DependenciaDestinoId:     1,
			TipoContratoId:           helpers.Ptr(1),
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_necesidad", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
			DependenciaDestinoId:     1,
			TipoContratoId:           helpers.Ptr(1),
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			NumeroNecesidad:          nil,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_necesidad", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
		}
		esperado := helpers.Desreferenciar(m.NumeroNecesidad)
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "numero_necesidad", esperado)
	})
}

// TestHistMstSolicitudCotizacionTipoContratoId verifica el valor de tipo_contrato_id en el INSERT para
// los tres casos del campo TipoContratoId.
func TestHistMstSolicitudCotizacionTipoContratoId(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "tipo_contrato_id", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			TipoContratoId:           nil,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "tipo_contrato_id", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
		}
		esperado := helpers.Desreferenciar(m.TipoContratoId)
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "tipo_contrato_id", esperado)
	})
}

// TestHistMstSolicitudCotizacionTipoContratoIdGet verifica la lectura de tipo_contrato_id: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstSolicitudCotizacionTipoContratoIdGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"tipo_contrato_id": nil})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "TipoContratoId", nil, helpers.Desreferenciar(v.TipoContratoId))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr(1))
		helpers.DefinirFila(map[string]interface{}{"tipo_contrato_id": esperado})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "TipoContratoId", esperado, helpers.Desreferenciar(v.TipoContratoId))
	})
}

// TestHistMstSolicitudCotizacionTipoContratoIdUpdate verifica el valor de tipo_contrato_id en el UPDATE
// para los tres casos del campo TipoContratoId.
func TestHistMstSolicitudCotizacionTipoContratoIdUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "tipo_contrato_id", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
			CriterioSeleccion:        helpers.Ptr("Prueba"),
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			TipoContratoId:           nil,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "tipo_contrato_id", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
		}
		esperado := helpers.Desreferenciar(m.TipoContratoId)
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "tipo_contrato_id", esperado)
	})
}

// TestHistMstSolicitudCotizacionCriterioSeleccion verifica el valor de criterio_seleccion en el INSERT para
// los tres casos del campo CriterioSeleccion.
func TestHistMstSolicitudCotizacionCriterioSeleccion(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "criterio_seleccion", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			CriterioSeleccion:        nil,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "criterio_seleccion", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
		}
		esperado := helpers.Desreferenciar(m.CriterioSeleccion)
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "criterio_seleccion", esperado)
	})
}

// TestHistMstSolicitudCotizacionCriterioSeleccionGet verifica la lectura de criterio_seleccion: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstSolicitudCotizacionCriterioSeleccionGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"criterio_seleccion": nil})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "CriterioSeleccion", nil, helpers.Desreferenciar(v.CriterioSeleccion))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr("Prueba"))
		helpers.DefinirFila(map[string]interface{}{"criterio_seleccion": esperado})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "CriterioSeleccion", esperado, helpers.Desreferenciar(v.CriterioSeleccion))
	})
}

// TestHistMstSolicitudCotizacionCriterioSeleccionUpdate verifica el valor de criterio_seleccion en el UPDATE
// para los tres casos del campo CriterioSeleccion.
func TestHistMstSolicitudCotizacionCriterioSeleccionUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "criterio_seleccion", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
			PlazoEjecucion:           helpers.Ptr("Prueba"),
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			CriterioSeleccion:        nil,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "criterio_seleccion", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
		}
		esperado := helpers.Desreferenciar(m.CriterioSeleccion)
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "criterio_seleccion", esperado)
	})
}

// TestHistMstSolicitudCotizacionPlazoEjecucion verifica el valor de plazo_ejecucion en el INSERT para
// los tres casos del campo PlazoEjecucion.
func TestHistMstSolicitudCotizacionPlazoEjecucion(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plazo_ejecucion", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			PlazoEjecucion:           nil,
		}
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plazo_ejecucion", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstSolicitudCotizacion{
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
		}
		esperado := helpers.Desreferenciar(m.PlazoEjecucion)
		if _, err := models.AddHistMstSolicitudCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plazo_ejecucion", esperado)
	})
}

// TestHistMstSolicitudCotizacionPlazoEjecucionGet verifica la lectura de plazo_ejecucion: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstSolicitudCotizacionPlazoEjecucionGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"plazo_ejecucion": nil})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "PlazoEjecucion", nil, helpers.Desreferenciar(v.PlazoEjecucion))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr("Prueba"))
		helpers.DefinirFila(map[string]interface{}{"plazo_ejecucion": esperado})
		v, err := models.GetHistMstSolicitudCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "PlazoEjecucion", esperado, helpers.Desreferenciar(v.PlazoEjecucion))
	})
}

// TestHistMstSolicitudCotizacionPlazoEjecucionUpdate verifica el valor de plazo_ejecucion en el UPDATE
// para los tres casos del campo PlazoEjecucion.
func TestHistMstSolicitudCotizacionPlazoEjecucionUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plazo_ejecucion", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			PlazoEjecucion:           nil,
		}
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plazo_ejecucion", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstSolicitudCotizacion{
			Id:                       1,
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
		}
		esperado := helpers.Desreferenciar(m.PlazoEjecucion)
		if err := models.UpdateHistMstSolicitudCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstSolicitudCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "plazo_ejecucion", esperado)
	})
}
