// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador de HistDetSolicitudCotizacionController (/v1/hist_det_solicitud_cotizacion) ante errores de
// persistencia simulados, sin base de datos.

// cuerpoPutPersistenciaHistDetSolicitudCotizacion retorna un cuerpo válido para PUT al registro 1.
func cuerpoPutPersistenciaHistDetSolicitudCotizacion(t *testing.T) string {
	t.Helper()
	cuerpo := cuerpoValidacionHistDetSolicitudCotizacion(t)
	cuerpo["Id"] = 1
	return helpers.JSON(t, cuerpo)
}

func TestHistDetSolicitudCotizacionPostErrorPersistencia(t *testing.T) {
	cuerpo := helpers.JSON(t, cuerpoValidacionHistDetSolicitudCotizacion(t))
	t.Run("INSERT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionHistDetSolicitudCotizacion+"/", cuerpo, "INSERT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "POST")
	})
	t.Run("INSERT restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionHistDetSolicitudCotizacion+"/", cuerpo, "INSERT", helpers.ErrRestriccion, http.StatusBadRequest, "POST")
	})
}

func TestHistDetSolicitudCotizacionGetOneErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionHistDetSolicitudCotizacion+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetOne")
	})
}

func TestHistDetSolicitudCotizacionGetAllErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionHistDetSolicitudCotizacion+"/", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetAll")
	})
}

func TestHistDetSolicitudCotizacionPutErrorPersistencia(t *testing.T) {
	cuerpo := cuerpoPutPersistenciaHistDetSolicitudCotizacion(t)
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionHistDetSolicitudCotizacion+"/1", cuerpo, "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionHistDetSolicitudCotizacion+"/1", cuerpo, "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionHistDetSolicitudCotizacion+"/1", cuerpo, "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Put")
	})
}

func TestHistDetSolicitudCotizacionDeleteErrorPersistencia(t *testing.T) {
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionHistDetSolicitudCotizacion+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionHistDetSolicitudCotizacion+"/1", "", "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE otro error", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionHistDetSolicitudCotizacion+"/1", "", "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Delete")
	})
}
