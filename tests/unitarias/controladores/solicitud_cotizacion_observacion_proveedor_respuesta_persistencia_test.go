// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador de SolicitudCotizacionObservacionProveedorRespuestaController (/v1/solicitud_cotizacion_observacion_proveedor_respuesta) ante errores de
// persistencia simulados, sin base de datos.

// cuerpoPutPersistenciaSolicitudCotizacionObservacionProveedorRespuesta retorna un cuerpo válido para PUT al registro 1.
func cuerpoPutPersistenciaSolicitudCotizacionObservacionProveedorRespuesta(t *testing.T) string {
	t.Helper()
	cuerpo := cuerpoValidacionSolicitudCotizacionObservacionProveedorRespuesta(t)
	cuerpo["Id"] = 1
	return helpers.JSON(t, cuerpo)
}

func TestSolicitudCotizacionObservacionProveedorRespuestaPostErrorPersistencia(t *testing.T) {
	cuerpo := helpers.JSON(t, cuerpoValidacionSolicitudCotizacionObservacionProveedorRespuesta(t))
	t.Run("INSERT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta+"/", cuerpo, "INSERT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "POST")
	})
	t.Run("INSERT restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta+"/", cuerpo, "INSERT", helpers.ErrRestriccion, http.StatusBadRequest, "POST")
	})
}

func TestSolicitudCotizacionObservacionProveedorRespuestaGetOneErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetOne")
	})
}

func TestSolicitudCotizacionObservacionProveedorRespuestaGetAllErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta+"/", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetAll")
	})
}

func TestSolicitudCotizacionObservacionProveedorRespuestaPutErrorPersistencia(t *testing.T) {
	cuerpo := cuerpoPutPersistenciaSolicitudCotizacionObservacionProveedorRespuesta(t)
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta+"/1", cuerpo, "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta+"/1", cuerpo, "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta+"/1", cuerpo, "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Put")
	})
}

func TestSolicitudCotizacionObservacionProveedorRespuestaDeleteErrorPersistencia(t *testing.T) {
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta+"/1", "", "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE otro error", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionSolicitudCotizacionObservacionProveedorRespuesta+"/1", "", "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Delete")
	})
}
