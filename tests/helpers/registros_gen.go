// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package helpers

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
)

// Las funciones Nuevo<Modelo> retornan un registro de ejemplo válido sin
// guardar (creando en la base de datos sus llaves foráneas). Las funciones
// Crear<Modelo> lo guardan directamente con el modelo y programan su
// eliminación al terminar la prueba.

func NuevoCargueMasivo(t *testing.T) *models.CargueMasivo {
	t.Helper()
	return &models.CargueMasivo{
		HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: CrearHistMstSolicitudCotizacion(t).Id},
		EnlacePlantilla:              Ptr("Prueba"),
		CargadoPorTerceroId:          1,
		TotalRegistros:               Ptr(1),
		Aceptados:                    Ptr(1),
		Rechazados:                   Ptr(1),
		Respuesta:                    Ptr("{\"prueba\": true}"),
		UsuarioCreacionId:            1,
		UsuarioModificacionId:        1,
	}
}

func CrearCargueMasivo(t *testing.T) *models.CargueMasivo {
	t.Helper()
	v := NuevoCargueMasivo(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddCargueMasivo(v); err != nil {
		t.Fatal("No se pudo crear CargueMasivo:", err)
	}
	EliminarAlFinal(t, "cargue_masivo", v.Id)
	return v
}

func NuevoDecisionSolicitante(t *testing.T) *models.DecisionSolicitante {
	t.Helper()
	return &models.DecisionSolicitante{
		RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: CrearRelSolicitudCotizacionProveedor(t).Id},
		MensajeAProveedor:                 "Prueba",
		ResultadoId:                       1,
		UsuarioCreacionId:                 1,
		UsuarioModificacionId:             1,
	}
}

func CrearDecisionSolicitante(t *testing.T) *models.DecisionSolicitante {
	t.Helper()
	v := NuevoDecisionSolicitante(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddDecisionSolicitante(v); err != nil {
		t.Fatal("No se pudo crear DecisionSolicitante:", err)
	}
	EliminarAlFinal(t, "decision_solicitante", v.Id)
	return v
}

func NuevoFormaPago(t *testing.T) *models.FormaPago {
	t.Helper()
	return &models.FormaPago{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: CrearMstSolicitudCotizacion(t).Id},
		TipoFormaPagoId:          1,
		Condicion:                "Prueba",
		PorcentajePago:           1.5,
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	}
}

func CrearFormaPago(t *testing.T) *models.FormaPago {
	t.Helper()
	v := NuevoFormaPago(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddFormaPago(v); err != nil {
		t.Fatal("No se pudo crear FormaPago:", err)
	}
	EliminarAlFinal(t, "forma_pago", v.Id)
	return v
}

func NuevoHistDetCotizacion(t *testing.T) *models.HistDetCotizacion {
	t.Helper()
	return &models.HistDetCotizacion{
		DetCotizacionId:          1,
		MstCotizacionId:          &models.MstCotizacion{Id: CrearMstCotizacion(t).Id},
		DetSolicitudCotizacionId: 1,
		ValorUnitario:            Ptr(1.5),
		Iva:                      1.5,
		FichaTecnica:             "Prueba",
		RegistroActual:           true,
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	}
}

func CrearHistDetCotizacion(t *testing.T) *models.HistDetCotizacion {
	t.Helper()
	v := NuevoHistDetCotizacion(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddHistDetCotizacion(v); err != nil {
		t.Fatal("No se pudo crear HistDetCotizacion:", err)
	}
	EliminarAlFinal(t, "hist_det_cotizacion", v.Id)
	return v
}

func NuevoHistDetSolicitudCotizacion(t *testing.T) *models.HistDetSolicitudCotizacion {
	t.Helper()
	return &models.HistDetSolicitudCotizacion{
		DetSolicitudCotizacionId: 1,
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: CrearMstSolicitudCotizacion(t).Id},
		Nombre:                   "Prueba",
		Descripcion:              "Prueba",
		TipoBienServicioId:       1,
		UnidadMedidaId:           1,
		Cantidad:                 1.5,
		RegistroActual:           true,
		TipoRegistro:             "Prueba",
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	}
}

func CrearHistDetSolicitudCotizacion(t *testing.T) *models.HistDetSolicitudCotizacion {
	t.Helper()
	v := NuevoHistDetSolicitudCotizacion(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddHistDetSolicitudCotizacion(v); err != nil {
		t.Fatal("No se pudo crear HistDetSolicitudCotizacion:", err)
	}
	EliminarAlFinal(t, "hist_det_solicitud_cotizacion", v.Id)
	return v
}

func NuevoHistMstCotizacion(t *testing.T) *models.HistMstCotizacion {
	t.Helper()
	return &models.HistMstCotizacion{
		MstCotizacionId:       &models.MstCotizacion{Id: CrearMstCotizacion(t).Id},
		FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
		Descripcion:           Ptr("Prueba"),
		InformacionEntrega:    Ptr("Prueba"),
		Descuentos:            Ptr("Prueba"),
		Observaciones:         Ptr("Prueba"),
		FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
		SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
		RegistroActual:        true,
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	}
}

func CrearHistMstCotizacion(t *testing.T) *models.HistMstCotizacion {
	t.Helper()
	v := NuevoHistMstCotizacion(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddHistMstCotizacion(v); err != nil {
		t.Fatal("No se pudo crear HistMstCotizacion:", err)
	}
	EliminarAlFinal(t, "hist_mst_cotizacion", v.Id)
	return v
}

func NuevoHistMstSolicitudCotizacion(t *testing.T) *models.HistMstSolicitudCotizacion {
	t.Helper()
	return &models.HistMstSolicitudCotizacion{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: CrearMstSolicitudCotizacion(t).Id},
		NumeroSolicitud:          Ptr("Prueba"),
		Vigencia:                 Ptr(int16(2026)),
		Consecutivo:              Ptr(1),
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
		PlanAccionId:             Ptr(1),
		ResponsableTerceroId:     1,
		MetodoSeleccionId:        Ptr(1),
		NumeroNecesidad:          Ptr(1.5),
		DependenciaDestinoId:     1,
		TipoContratoId:           Ptr(1),
		CriterioSeleccion:        Ptr("Prueba"),
		PlazoEjecucion:           Ptr("Prueba"),
		RegistroActual:           true,
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	}
}

func CrearHistMstSolicitudCotizacion(t *testing.T) *models.HistMstSolicitudCotizacion {
	t.Helper()
	v := NuevoHistMstSolicitudCotizacion(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddHistMstSolicitudCotizacion(v); err != nil {
		t.Fatal("No se pudo crear HistMstSolicitudCotizacion:", err)
	}
	EliminarAlFinal(t, "hist_mst_solicitud_cotizacion", v.Id)
	return v
}

func NuevoHistMstSolicitudCotizacionEstadoHist(t *testing.T) *models.HistMstSolicitudCotizacionEstadoHist {
	t.Helper()
	return &models.HistMstSolicitudCotizacionEstadoHist{
		ResponsableTerceroId:         1,
		Justificacion:                "Prueba",
		EstadoHistId:                 1,
		HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: CrearHistMstSolicitudCotizacion(t).Id},
		UsuarioCreacionId:            1,
		UsuarioModificacionId:        1,
	}
}

func CrearHistMstSolicitudCotizacionEstadoHist(t *testing.T) *models.HistMstSolicitudCotizacionEstadoHist {
	t.Helper()
	v := NuevoHistMstSolicitudCotizacionEstadoHist(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddHistMstSolicitudCotizacionEstadoHist(v); err != nil {
		t.Fatal("No se pudo crear HistMstSolicitudCotizacionEstadoHist:", err)
	}
	EliminarAlFinal(t, "hist_mst_solicitud_cotizacion_estado_hist", v.Id)
	return v
}

func NuevoMstCotizacion(t *testing.T) *models.MstCotizacion {
	t.Helper()
	return &models.MstCotizacion{
		RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: CrearRelSolicitudCotizacionProveedor(t).Id},
		UsuarioCreacionId:                 1,
		UsuarioModificacionId:             1,
	}
}

func CrearMstCotizacion(t *testing.T) *models.MstCotizacion {
	t.Helper()
	v := NuevoMstCotizacion(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddMstCotizacion(v); err != nil {
		t.Fatal("No se pudo crear MstCotizacion:", err)
	}
	EliminarAlFinal(t, "mst_cotizacion", v.Id)
	return v
}

func NuevoMstSolicitudCotizacion(t *testing.T) *models.MstSolicitudCotizacion {
	t.Helper()
	return &models.MstSolicitudCotizacion{
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	}
}

func CrearMstSolicitudCotizacion(t *testing.T) *models.MstSolicitudCotizacion {
	t.Helper()
	v := NuevoMstSolicitudCotizacion(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddMstSolicitudCotizacion(v); err != nil {
		t.Fatal("No se pudo crear MstSolicitudCotizacion:", err)
	}
	EliminarAlFinal(t, "mst_solicitud_cotizacion", v.Id)
	return v
}

func NuevoNotificacion(t *testing.T) *models.Notificacion {
	t.Helper()
	return &models.Notificacion{
		ProveedorNotificadoId: 1,
		TerceroNotificaId:     1,
		EstadoEnvioId:         1,
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	}
}

func CrearNotificacion(t *testing.T) *models.Notificacion {
	t.Helper()
	v := NuevoNotificacion(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddNotificacion(v); err != nil {
		t.Fatal("No se pudo crear Notificacion:", err)
	}
	EliminarAlFinal(t, "notificacion", v.Id)
	return v
}

func NuevoRelHistMstDetCotizacion(t *testing.T) *models.RelHistMstDetCotizacion {
	t.Helper()
	return &models.RelHistMstDetCotizacion{
		MaestroId:             &models.HistMstCotizacion{Id: CrearHistMstCotizacion(t).Id},
		DetalleId:             &models.HistDetCotizacion{Id: CrearHistDetCotizacion(t).Id},
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	}
}

func CrearRelHistMstDetCotizacion(t *testing.T) *models.RelHistMstDetCotizacion {
	t.Helper()
	v := NuevoRelHistMstDetCotizacion(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddRelHistMstDetCotizacion(v); err != nil {
		t.Fatal("No se pudo crear RelHistMstDetCotizacion:", err)
	}
	EliminarAlFinal(t, "rel_hist_mst_det_cotizacion", v.Id)
	return v
}

func NuevoRelHistMstDetSolicitudCotizacion(t *testing.T) *models.RelHistMstDetSolicitudCotizacion {
	t.Helper()
	return &models.RelHistMstDetSolicitudCotizacion{
		MaestroId:             &models.HistMstSolicitudCotizacion{Id: CrearHistMstSolicitudCotizacion(t).Id},
		DetalleId:             &models.HistDetSolicitudCotizacion{Id: CrearHistDetSolicitudCotizacion(t).Id},
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	}
}

func CrearRelHistMstDetSolicitudCotizacion(t *testing.T) *models.RelHistMstDetSolicitudCotizacion {
	t.Helper()
	v := NuevoRelHistMstDetSolicitudCotizacion(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddRelHistMstDetSolicitudCotizacion(v); err != nil {
		t.Fatal("No se pudo crear RelHistMstDetSolicitudCotizacion:", err)
	}
	EliminarAlFinal(t, "rel_hist_mst_det_solicitud_cotizacion", v.Id)
	return v
}

func NuevoRelSolicitudCotizacionProveedor(t *testing.T) *models.RelSolicitudCotizacionProveedor {
	t.Helper()
	return &models.RelSolicitudCotizacionProveedor{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: CrearMstSolicitudCotizacion(t).Id},
		ProveedorId:              1,
		NotificacionId:           &models.Notificacion{Id: CrearNotificacion(t).Id},
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	}
}

func CrearRelSolicitudCotizacionProveedor(t *testing.T) *models.RelSolicitudCotizacionProveedor {
	t.Helper()
	v := NuevoRelSolicitudCotizacionProveedor(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddRelSolicitudCotizacionProveedor(v); err != nil {
		t.Fatal("No se pudo crear RelSolicitudCotizacionProveedor:", err)
	}
	EliminarAlFinal(t, "rel_solicitud_cotizacion_proveedor", v.Id)
	return v
}

func NuevoSolicitudCotizacionActividadEconomica(t *testing.T) *models.SolicitudCotizacionActividadEconomica {
	t.Helper()
	return &models.SolicitudCotizacionActividadEconomica{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: CrearMstSolicitudCotizacion(t).Id},
		ActividadEconomicaId:     1,
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	}
}

func CrearSolicitudCotizacionActividadEconomica(t *testing.T) *models.SolicitudCotizacionActividadEconomica {
	t.Helper()
	v := NuevoSolicitudCotizacionActividadEconomica(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddSolicitudCotizacionActividadEconomica(v); err != nil {
		t.Fatal("No se pudo crear SolicitudCotizacionActividadEconomica:", err)
	}
	EliminarAlFinal(t, "solicitud_cotizacion_actividad_economica", v.Id)
	return v
}

func NuevoSolicitudCotizacionEspecificacionTecnica(t *testing.T) *models.SolicitudCotizacionEspecificacionTecnica {
	t.Helper()
	return &models.SolicitudCotizacionEspecificacionTecnica{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: CrearMstSolicitudCotizacion(t).Id},
		Enlace:                   "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	}
}

func CrearSolicitudCotizacionEspecificacionTecnica(t *testing.T) *models.SolicitudCotizacionEspecificacionTecnica {
	t.Helper()
	v := NuevoSolicitudCotizacionEspecificacionTecnica(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddSolicitudCotizacionEspecificacionTecnica(v); err != nil {
		t.Fatal("No se pudo crear SolicitudCotizacionEspecificacionTecnica:", err)
	}
	EliminarAlFinal(t, "solicitud_cotizacion_especificacion_tecnica", v.Id)
	return v
}

func NuevoSolicitudCotizacionNucleoBasico(t *testing.T) *models.SolicitudCotizacionNucleoBasico {
	t.Helper()
	return &models.SolicitudCotizacionNucleoBasico{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: CrearMstSolicitudCotizacion(t).Id},
		NucleoBasicoId:           1,
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	}
}

func CrearSolicitudCotizacionNucleoBasico(t *testing.T) *models.SolicitudCotizacionNucleoBasico {
	t.Helper()
	v := NuevoSolicitudCotizacionNucleoBasico(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddSolicitudCotizacionNucleoBasico(v); err != nil {
		t.Fatal("No se pudo crear SolicitudCotizacionNucleoBasico:", err)
	}
	EliminarAlFinal(t, "solicitud_cotizacion_nucleo_basico", v.Id)
	return v
}

func NuevoSolicitudCotizacionObservacionProveedor(t *testing.T) *models.SolicitudCotizacionObservacionProveedor {
	t.Helper()
	return &models.SolicitudCotizacionObservacionProveedor{
		Texto:                             Ptr("Prueba"),
		RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: CrearRelSolicitudCotizacionProveedor(t).Id},
		VistoSolicitante:                  true,
		FechaVistoSolicitante:             Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
		UsuarioCreacionId:                 1,
		UsuarioModificacionId:             1,
	}
}

func CrearSolicitudCotizacionObservacionProveedor(t *testing.T) *models.SolicitudCotizacionObservacionProveedor {
	t.Helper()
	v := NuevoSolicitudCotizacionObservacionProveedor(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddSolicitudCotizacionObservacionProveedor(v); err != nil {
		t.Fatal("No se pudo crear SolicitudCotizacionObservacionProveedor:", err)
	}
	EliminarAlFinal(t, "solicitud_cotizacion_observacion_proveedor", v.Id)
	return v
}

func NuevoSolicitudCotizacionObservacionProveedorRespuesta(t *testing.T) *models.SolicitudCotizacionObservacionProveedorRespuesta {
	t.Helper()
	return &models.SolicitudCotizacionObservacionProveedorRespuesta{
		ObservacionId:         &models.SolicitudCotizacionObservacionProveedor{Id: CrearSolicitudCotizacionObservacionProveedor(t).Id},
		Texto:                 "Prueba",
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	}
}

func CrearSolicitudCotizacionObservacionProveedorRespuesta(t *testing.T) *models.SolicitudCotizacionObservacionProveedorRespuesta {
	t.Helper()
	v := NuevoSolicitudCotizacionObservacionProveedorRespuesta(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddSolicitudCotizacionObservacionProveedorRespuesta(v); err != nil {
		t.Fatal("No se pudo crear SolicitudCotizacionObservacionProveedorRespuesta:", err)
	}
	EliminarAlFinal(t, "solicitud_cotizacion_observacion_proveedor_respuesta", v.Id)
	return v
}

func NuevoValidacionOrdenador(t *testing.T) *models.ValidacionOrdenador {
	t.Helper()
	return &models.ValidacionOrdenador{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: CrearMstSolicitudCotizacion(t).Id},
		Observaciones:            "Prueba",
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	}
}

func CrearValidacionOrdenador(t *testing.T) *models.ValidacionOrdenador {
	t.Helper()
	v := NuevoValidacionOrdenador(t)
	v.Activo = true
	v.FechaCreacion = time.Now()
	v.FechaModificacion = time.Now()
	if _, err := models.AddValidacionOrdenador(v); err != nil {
		t.Fatal("No se pudo crear ValidacionOrdenador:", err)
	}
	EliminarAlFinal(t, "validacion_ordenador", v.Id)
	return v
}
