// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador de SolicitudCotizacionEspecificacionTecnicaController (/v1/solicitud_cotizacion_especificacion_tecnica) ante errores de
// persistencia simulados, sin base de datos.

// cuerpoPutPersistenciaSolicitudCotizacionEspecificacionTecnica retorna un cuerpo válido para PUT al registro 1.
func cuerpoPutPersistenciaSolicitudCotizacionEspecificacionTecnica(t *testing.T) string {
	t.Helper()
	cuerpo := cuerpoValidacionSolicitudCotizacionEspecificacionTecnica(t)
	cuerpo["Id"] = 1
	return helpers.JSON(t, cuerpo)
}

func TestSolicitudCotizacionEspecificacionTecnicaPostErrorPersistencia(t *testing.T) {
	cuerpo := helpers.JSON(t, cuerpoValidacionSolicitudCotizacionEspecificacionTecnica(t))
	t.Run("INSERT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionSolicitudCotizacionEspecificacionTecnica+"/", cuerpo, "INSERT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "POST")
	})
	t.Run("INSERT restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionSolicitudCotizacionEspecificacionTecnica+"/", cuerpo, "INSERT", helpers.ErrRestriccion, http.StatusBadRequest, "POST")
	})
}

func TestSolicitudCotizacionEspecificacionTecnicaGetOneErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionSolicitudCotizacionEspecificacionTecnica+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetOne")
	})
}

func TestSolicitudCotizacionEspecificacionTecnicaGetAllErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionSolicitudCotizacionEspecificacionTecnica+"/", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetAll")
	})
}

func TestSolicitudCotizacionEspecificacionTecnicaPutErrorPersistencia(t *testing.T) {
	cuerpo := cuerpoPutPersistenciaSolicitudCotizacionEspecificacionTecnica(t)
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionSolicitudCotizacionEspecificacionTecnica+"/1", cuerpo, "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionSolicitudCotizacionEspecificacionTecnica+"/1", cuerpo, "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionSolicitudCotizacionEspecificacionTecnica+"/1", cuerpo, "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Put")
	})
}

func TestSolicitudCotizacionEspecificacionTecnicaDeleteErrorPersistencia(t *testing.T) {
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionSolicitudCotizacionEspecificacionTecnica+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionSolicitudCotizacionEspecificacionTecnica+"/1", "", "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE otro error", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionSolicitudCotizacionEspecificacionTecnica+"/1", "", "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Delete")
	})
}
