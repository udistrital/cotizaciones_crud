// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador de HistMstSolicitudCotizacionEstadoHistController (/v1/hist_mst_solicitud_cotizacion_estado_hist) ante errores de
// persistencia simulados, sin base de datos.

// cuerpoPutPersistenciaHistMstSolicitudCotizacionEstadoHist retorna un cuerpo válido para PUT al registro 1.
func cuerpoPutPersistenciaHistMstSolicitudCotizacionEstadoHist(t *testing.T) string {
	t.Helper()
	cuerpo := cuerpoValidacionHistMstSolicitudCotizacionEstadoHist(t)
	cuerpo["Id"] = 1
	return helpers.JSON(t, cuerpo)
}

func TestHistMstSolicitudCotizacionEstadoHistPostErrorPersistencia(t *testing.T) {
	cuerpo := helpers.JSON(t, cuerpoValidacionHistMstSolicitudCotizacionEstadoHist(t))
	t.Run("INSERT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionHistMstSolicitudCotizacionEstadoHist+"/", cuerpo, "INSERT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "POST")
	})
	t.Run("INSERT restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionHistMstSolicitudCotizacionEstadoHist+"/", cuerpo, "INSERT", helpers.ErrRestriccion, http.StatusBadRequest, "POST")
	})
}

func TestHistMstSolicitudCotizacionEstadoHistGetOneErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionHistMstSolicitudCotizacionEstadoHist+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetOne")
	})
}

func TestHistMstSolicitudCotizacionEstadoHistGetAllErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionHistMstSolicitudCotizacionEstadoHist+"/", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetAll")
	})
}

func TestHistMstSolicitudCotizacionEstadoHistPutErrorPersistencia(t *testing.T) {
	cuerpo := cuerpoPutPersistenciaHistMstSolicitudCotizacionEstadoHist(t)
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionHistMstSolicitudCotizacionEstadoHist+"/1", cuerpo, "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionHistMstSolicitudCotizacionEstadoHist+"/1", cuerpo, "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionHistMstSolicitudCotizacionEstadoHist+"/1", cuerpo, "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Put")
	})
}

func TestHistMstSolicitudCotizacionEstadoHistDeleteErrorPersistencia(t *testing.T) {
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionHistMstSolicitudCotizacionEstadoHist+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionHistMstSolicitudCotizacionEstadoHist+"/1", "", "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE otro error", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionHistMstSolicitudCotizacionEstadoHist+"/1", "", "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Delete")
	})
}
