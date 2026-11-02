package controllers

import "github.com/udistrital/cotizaciones_crud/models"

// Tipos que documentan en Swagger la respuesta de GetAll de cada entidad: los campos
// comunes de APIResponseLista y Data como lista del modelo correspondiente. Solo se
// usan en las anotaciones @Success; la respuesta la construye writeSuccessLista.

// CargueMasivoLista es la respuesta de GetAll de CargueMasivo
type CargueMasivoLista struct {
	APIResponseLista
	Data []models.CargueMasivo `json:"Data"`
}

// CargueMasivoErrorLista es la respuesta de GetAll de CargueMasivoError
type CargueMasivoErrorLista struct {
	APIResponseLista
	Data []models.CargueMasivoError `json:"Data"`
}

// DocumentoSolicitudLista es la respuesta de GetAll de DocumentoSolicitud
type DocumentoSolicitudLista struct {
	APIResponseLista
	Data []models.DocumentoSolicitud `json:"Data"`
}

// ObservacionLista es la respuesta de GetAll de Observacion
type ObservacionLista struct {
	APIResponseLista
	Data []models.Observacion `json:"Data"`
}

// RespuestaObservacionLista es la respuesta de GetAll de RespuestaObservacion
type RespuestaObservacionLista struct {
	APIResponseLista
	Data []models.RespuestaObservacion `json:"Data"`
}

// SolicitudCiiuLista es la respuesta de GetAll de SolicitudCiiu
type SolicitudCiiuLista struct {
	APIResponseLista
	Data []models.SolicitudCiiu `json:"Data"`
}

// SolicitudCotizacionLista es la respuesta de GetAll de SolicitudCotizacion
type SolicitudCotizacionLista struct {
	APIResponseLista
	Data []models.SolicitudCotizacion `json:"Data"`
}

// SolicitudEstadoLista es la respuesta de GetAll de SolicitudEstado
type SolicitudEstadoLista struct {
	APIResponseLista
	Data []models.SolicitudEstado `json:"Data"`
}

// SolicitudItemLista es la respuesta de GetAll de SolicitudItem
type SolicitudItemLista struct {
	APIResponseLista
	Data []models.SolicitudItem `json:"Data"`
}

// SolicitudModificacionLista es la respuesta de GetAll de SolicitudModificacion
type SolicitudModificacionLista struct {
	APIResponseLista
	Data []models.SolicitudModificacion `json:"Data"`
}

// SolicitudModificacionDetalleLista es la respuesta de GetAll de SolicitudModificacionDetalle
type SolicitudModificacionDetalleLista struct {
	APIResponseLista
	Data []models.SolicitudModificacionDetalle `json:"Data"`
}

// SolicitudModificacionEstadoLista es la respuesta de GetAll de SolicitudModificacionEstado
type SolicitudModificacionEstadoLista struct {
	APIResponseLista
	Data []models.SolicitudModificacionEstado `json:"Data"`
}

// SolicitudNecesidadLista es la respuesta de GetAll de SolicitudNecesidad
type SolicitudNecesidadLista struct {
	APIResponseLista
	Data []models.SolicitudNecesidad `json:"Data"`
}

// SolicitudPagoParcialLista es la respuesta de GetAll de SolicitudPagoParcial
type SolicitudPagoParcialLista struct {
	APIResponseLista
	Data []models.SolicitudPagoParcial `json:"Data"`
}

// SolicitudProveedorLista es la respuesta de GetAll de SolicitudProveedor
type SolicitudProveedorLista struct {
	APIResponseLista
	Data []models.SolicitudProveedor `json:"Data"`
}

// SolicitudRequisitoDocumentalLista es la respuesta de GetAll de SolicitudRequisitoDocumental
type SolicitudRequisitoDocumentalLista struct {
	APIResponseLista
	Data []models.SolicitudRequisitoDocumental `json:"Data"`
}
