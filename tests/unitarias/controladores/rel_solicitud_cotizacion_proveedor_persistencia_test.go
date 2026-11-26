// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador de RelSolicitudCotizacionProveedorController (/v1/rel_solicitud_cotizacion_proveedor) ante errores de
// persistencia simulados, sin base de datos.

// cuerpoPutPersistenciaRelSolicitudCotizacionProveedor retorna un cuerpo válido para PUT al registro 1.
func cuerpoPutPersistenciaRelSolicitudCotizacionProveedor(t *testing.T) string {
	t.Helper()
	cuerpo := cuerpoValidacionRelSolicitudCotizacionProveedor(t)
	cuerpo["Id"] = 1
	return helpers.JSON(t, cuerpo)
}

func TestRelSolicitudCotizacionProveedorPostErrorPersistencia(t *testing.T) {
	cuerpo := helpers.JSON(t, cuerpoValidacionRelSolicitudCotizacionProveedor(t))
	t.Run("INSERT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionRelSolicitudCotizacionProveedor+"/", cuerpo, "INSERT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "POST")
	})
	t.Run("INSERT restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionRelSolicitudCotizacionProveedor+"/", cuerpo, "INSERT", helpers.ErrRestriccion, http.StatusBadRequest, "POST")
	})
}

func TestRelSolicitudCotizacionProveedorGetOneErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionRelSolicitudCotizacionProveedor+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetOne")
	})
}

func TestRelSolicitudCotizacionProveedorGetAllErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionRelSolicitudCotizacionProveedor+"/", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetAll")
	})
}

func TestRelSolicitudCotizacionProveedorPutErrorPersistencia(t *testing.T) {
	cuerpo := cuerpoPutPersistenciaRelSolicitudCotizacionProveedor(t)
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionRelSolicitudCotizacionProveedor+"/1", cuerpo, "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionRelSolicitudCotizacionProveedor+"/1", cuerpo, "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionRelSolicitudCotizacionProveedor+"/1", cuerpo, "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Put")
	})
}

func TestRelSolicitudCotizacionProveedorDeleteErrorPersistencia(t *testing.T) {
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionRelSolicitudCotizacionProveedor+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionRelSolicitudCotizacionProveedor+"/1", "", "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE otro error", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionRelSolicitudCotizacionProveedor+"/1", "", "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Delete")
	})
}
