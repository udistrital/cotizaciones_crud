package helpers

import (
	"fmt"
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
)

// IdInexistente es un Id que nunca se crea en la base de datos de pruebas.
const IdInexistente = 999999

// secuencia genera valores únicos para los datos de prueba.
var secuencia int

func siguiente() int {
	secuencia++
	return secuencia
}

// MismaFecha indica si dos fechas corresponden al mismo instante, tolerando
// la pérdida de precisión al guardarlas en la base de datos.
func MismaFecha(a, b time.Time) bool {
	d := a.Sub(b)
	return d > -time.Second && d < time.Second
}

// Las funciones Nuevo<Modelo> retornan un registro de ejemplo sin guardar
// (creando en la base de datos sus llaves foráneas) y las funciones
// Crear<Modelo> lo guardan en la base de datos. Los Id de otros servicios
// (Parámetros, Terceros, Oikos, Kronos y gestor documental) son valores de
// ejemplo, porque el CRUD solo guarda la referencia.

// NuevoSolicitudCotizacion asigna a cada solicitud una vigencia, un estado y
// una dependencia únicos, para que los filtros de consulta la identifiquen.
func NuevoSolicitudCotizacion(t *testing.T) *models.SolicitudCotizacion {
	t.Helper()
	n := siguiente()
	return &models.SolicitudCotizacion{
		Vigencia:                 int16(2000 + n),
		UnidadEjecutoraId:        1,
		TipoProcesoId:            1,
		EstadoId:                 100 + n,
		Titulo:                   fmt.Sprintf("Solicitud prueba %d", n),
		Objetivo:                 "Adquisición de equipos de prueba",
		DependenciaSolicitanteId: 100 + n,
		DependenciaDestinoId:     1,
		ResponsableTerceroId:     1,
		TipoFormaPagoId:          1,
		FechaApertura:            time.Date(2026, 11, 3, 0, 0, 0, 0, time.UTC),
		FechaCierre:              time.Date(2026, 11, 6, 0, 0, 0, 0, time.UTC),
	}
}

func CrearSolicitudCotizacion(t *testing.T) *models.SolicitudCotizacion {
	t.Helper()
	v := NuevoSolicitudCotizacion(t)
	if _, err := models.AddSolicitudCotizacion(v); err != nil {
		t.Fatal("No se pudo crear SolicitudCotizacion:", err)
	}
	return v
}

func NuevoSolicitudNecesidad(t *testing.T) *models.SolicitudNecesidad {
	t.Helper()
	return &models.SolicitudNecesidad{SolicitudId: CrearSolicitudCotizacion(t), NumeroNecesidad: fmt.Sprintf("N-%d", siguiente()), NumeroDisponibilidad: "CDP-1", ValorCdp: 1500000, RubroCodigo: "1-01", RubroNombre: "Rubro prueba"}
}

func CrearSolicitudNecesidad(t *testing.T) *models.SolicitudNecesidad {
	t.Helper()
	v := NuevoSolicitudNecesidad(t)
	if _, err := models.AddSolicitudNecesidad(v); err != nil {
		t.Fatal("No se pudo crear SolicitudNecesidad:", err)
	}
	return v
}

func NuevoSolicitudItem(t *testing.T) *models.SolicitudItem {
	t.Helper()
	return &models.SolicitudItem{SolicitudId: CrearSolicitudCotizacion(t), TipoBienServicioId: 1, Nombre: "Computador de escritorio", Descripcion: "Procesador de 8 núcleos y 16 GB de RAM", UnidadMedidaId: 1, Cantidad: 25}
}

func CrearSolicitudItem(t *testing.T) *models.SolicitudItem {
	t.Helper()
	v := NuevoSolicitudItem(t)
	if _, err := models.AddSolicitudItem(v); err != nil {
		t.Fatal("No se pudo crear SolicitudItem:", err)
	}
	return v
}

func NuevoCargueMasivo(t *testing.T) *models.CargueMasivo {
	t.Helper()
	return &models.CargueMasivo{SolicitudId: CrearSolicitudCotizacion(t), PlantillaId: 1, DocumentoId: 1, NombreArchivo: "items.xlsx", CargadoPorTerceroId: 1, TotalRegistros: 10, Aceptados: 9, Rechazados: 1, EstadoId: 1}
}

func CrearCargueMasivo(t *testing.T) *models.CargueMasivo {
	t.Helper()
	v := NuevoCargueMasivo(t)
	if _, err := models.AddCargueMasivo(v); err != nil {
		t.Fatal("No se pudo crear CargueMasivo:", err)
	}
	return v
}

func NuevoCargueMasivoError(t *testing.T) *models.CargueMasivoError {
	t.Helper()
	return &models.CargueMasivoError{CargueMasivoId: CrearCargueMasivo(t), Hoja: "Items", Fila: 5, Campo: "Cantidad", ValorDetectado: "abc", MotivoId: 1, Detalle: "La cantidad debe ser numérica"}
}

func CrearCargueMasivoError(t *testing.T) *models.CargueMasivoError {
	t.Helper()
	v := NuevoCargueMasivoError(t)
	if _, err := models.AddCargueMasivoError(v); err != nil {
		t.Fatal("No se pudo crear CargueMasivoError:", err)
	}
	return v
}

func NuevoSolicitudPagoParcial(t *testing.T) *models.SolicitudPagoParcial {
	t.Helper()
	return &models.SolicitudPagoParcial{SolicitudId: CrearSolicitudCotizacion(t), Condicion: "Contra entrega del 50 % de los equipos", Porcentaje: 0.5}
}

func CrearSolicitudPagoParcial(t *testing.T) *models.SolicitudPagoParcial {
	t.Helper()
	v := NuevoSolicitudPagoParcial(t)
	if _, err := models.AddSolicitudPagoParcial(v); err != nil {
		t.Fatal("No se pudo crear SolicitudPagoParcial:", err)
	}
	return v
}

func NuevoSolicitudCiiu(t *testing.T) *models.SolicitudCiiu {
	t.Helper()
	return &models.SolicitudCiiu{SolicitudId: CrearSolicitudCotizacion(t), ActividadEconomicaId: 1}
}

func CrearSolicitudCiiu(t *testing.T) *models.SolicitudCiiu {
	t.Helper()
	v := NuevoSolicitudCiiu(t)
	if _, err := models.AddSolicitudCiiu(v); err != nil {
		t.Fatal("No se pudo crear SolicitudCiiu:", err)
	}
	return v
}

func NuevoDocumentoSolicitud(t *testing.T) *models.DocumentoSolicitud {
	t.Helper()
	return &models.DocumentoSolicitud{SolicitudId: CrearSolicitudCotizacion(t), TipoDocumentoId: 1, DocumentoId: 1, NombreArchivo: "estudios_previos.pdf", Extension: "pdf", TamanoBytes: 2048, CargadoPorTerceroId: 1}
}

func CrearDocumentoSolicitud(t *testing.T) *models.DocumentoSolicitud {
	t.Helper()
	v := NuevoDocumentoSolicitud(t)
	if _, err := models.AddDocumentoSolicitud(v); err != nil {
		t.Fatal("No se pudo crear DocumentoSolicitud:", err)
	}
	return v
}

func NuevoSolicitudRequisitoDocumental(t *testing.T) *models.SolicitudRequisitoDocumental {
	t.Helper()
	return &models.SolicitudRequisitoDocumental{SolicitudId: CrearSolicitudCotizacion(t), TipoDocumentoId: 1, ClasificacionId: 1, Instruccion: "Adjuntar RUT actualizado"}
}

func CrearSolicitudRequisitoDocumental(t *testing.T) *models.SolicitudRequisitoDocumental {
	t.Helper()
	v := NuevoSolicitudRequisitoDocumental(t)
	if _, err := models.AddSolicitudRequisitoDocumental(v); err != nil {
		t.Fatal("No se pudo crear SolicitudRequisitoDocumental:", err)
	}
	return v
}

func NuevoSolicitudProveedor(t *testing.T) *models.SolicitudProveedor {
	t.Helper()
	return &models.SolicitudProveedor{SolicitudId: CrearSolicitudCotizacion(t), ProveedorId: 1, FechaConvocatoria: time.Now()}
}

func CrearSolicitudProveedor(t *testing.T) *models.SolicitudProveedor {
	t.Helper()
	v := NuevoSolicitudProveedor(t)
	if _, err := models.AddSolicitudProveedor(v); err != nil {
		t.Fatal("No se pudo crear SolicitudProveedor:", err)
	}
	return v
}

func NuevoSolicitudEstado(t *testing.T) *models.SolicitudEstado {
	t.Helper()
	return &models.SolicitudEstado{SolicitudId: CrearSolicitudCotizacion(t), EstadoId: 1, ResponsableTerceroId: 1, Justificacion: "Creación de la solicitud"}
}

func CrearSolicitudEstado(t *testing.T) *models.SolicitudEstado {
	t.Helper()
	v := NuevoSolicitudEstado(t)
	if _, err := models.AddSolicitudEstado(v); err != nil {
		t.Fatal("No se pudo crear SolicitudEstado:", err)
	}
	return v
}

func NuevoSolicitudModificacion(t *testing.T) *models.SolicitudModificacion {
	t.Helper()
	return &models.SolicitudModificacion{SolicitudId: CrearSolicitudCotizacion(t), Numero: 1}
}

func CrearSolicitudModificacion(t *testing.T) *models.SolicitudModificacion {
	t.Helper()
	v := NuevoSolicitudModificacion(t)
	if _, err := models.AddSolicitudModificacion(v); err != nil {
		t.Fatal("No se pudo crear SolicitudModificacion:", err)
	}
	return v
}

func NuevoSolicitudModificacionEstado(t *testing.T) *models.SolicitudModificacionEstado {
	t.Helper()
	return &models.SolicitudModificacionEstado{SolicitudModificacionId: CrearSolicitudModificacion(t), EstadoId: 1, ResponsableTerceroId: 1, Justificacion: "Ajuste de cantidades"}
}

func CrearSolicitudModificacionEstado(t *testing.T) *models.SolicitudModificacionEstado {
	t.Helper()
	v := NuevoSolicitudModificacionEstado(t)
	if _, err := models.AddSolicitudModificacionEstado(v); err != nil {
		t.Fatal("No se pudo crear SolicitudModificacionEstado:", err)
	}
	return v
}

func NuevoSolicitudModificacionDetalle(t *testing.T) *models.SolicitudModificacionDetalle {
	t.Helper()
	anterior, nuevo := `{"Cantidad": 20}`, `{"Cantidad": 25}`
	return &models.SolicitudModificacionDetalle{SolicitudModificacionId: CrearSolicitudModificacion(t), TipoCambioId: 1, ValorAnterior: &anterior, ValorNuevo: &nuevo, Justificacion: "Cambio de cantidad"}
}

func CrearSolicitudModificacionDetalle(t *testing.T) *models.SolicitudModificacionDetalle {
	t.Helper()
	v := NuevoSolicitudModificacionDetalle(t)
	if _, err := models.AddSolicitudModificacionDetalle(v); err != nil {
		t.Fatal("No se pudo crear SolicitudModificacionDetalle:", err)
	}
	return v
}

// NuevoObservacion crea un proveedor convocado de la misma solicitud, que es
// quien escribe la observación.
func NuevoObservacion(t *testing.T) *models.Observacion {
	t.Helper()
	proveedor := CrearSolicitudProveedor(t)
	return &models.Observacion{SolicitudId: proveedor.SolicitudId, ProveedorConvocadoId: proveedor, Texto: "¿Se aceptan equipos con procesador de 10 núcleos?", EstadoId: 1}
}

func CrearObservacion(t *testing.T) *models.Observacion {
	t.Helper()
	v := NuevoObservacion(t)
	if _, err := models.AddObservacion(v); err != nil {
		t.Fatal("No se pudo crear Observacion:", err)
	}
	return v
}

func NuevoRespuestaObservacion(t *testing.T) *models.RespuestaObservacion {
	t.Helper()
	return &models.RespuestaObservacion{ObservacionId: CrearObservacion(t), AutorTerceroId: 1, Texto: "Sí, se aceptan procesadores iguales o superiores.", Publica: true}
}

func CrearRespuestaObservacion(t *testing.T) *models.RespuestaObservacion {
	t.Helper()
	v := NuevoRespuestaObservacion(t)
	if _, err := models.AddRespuestaObservacion(v); err != nil {
		t.Fatal("No se pudo crear RespuestaObservacion:", err)
	}
	return v
}
