// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador de MstCotizacionController (/v1/mst_cotizacion) ante errores de
// persistencia simulados, sin base de datos.

// cuerpoPutPersistenciaMstCotizacion retorna un cuerpo válido para PUT al registro 1.
func cuerpoPutPersistenciaMstCotizacion(t *testing.T) string {
	t.Helper()
	cuerpo := cuerpoValidacionMstCotizacion(t)
	cuerpo["Id"] = 1
	return helpers.JSON(t, cuerpo)
}

func TestMstCotizacionPostErrorPersistencia(t *testing.T) {
	cuerpo := helpers.JSON(t, cuerpoValidacionMstCotizacion(t))
	t.Run("INSERT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionMstCotizacion+"/", cuerpo, "INSERT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "POST")
	})
	t.Run("INSERT restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionMstCotizacion+"/", cuerpo, "INSERT", helpers.ErrRestriccion, http.StatusBadRequest, "POST")
	})
}

func TestMstCotizacionGetOneErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionMstCotizacion+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetOne")
	})
}

func TestMstCotizacionGetAllErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionMstCotizacion+"/", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetAll")
	})
}

func TestMstCotizacionPutErrorPersistencia(t *testing.T) {
	cuerpo := cuerpoPutPersistenciaMstCotizacion(t)
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionMstCotizacion+"/1", cuerpo, "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionMstCotizacion+"/1", cuerpo, "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionMstCotizacion+"/1", cuerpo, "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Put")
	})
}

func TestMstCotizacionDeleteErrorPersistencia(t *testing.T) {
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionMstCotizacion+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionMstCotizacion+"/1", "", "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE otro error", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionMstCotizacion+"/1", "", "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Delete")
	})
}
