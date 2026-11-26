// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador de DecisionSolicitanteController (/v1/decision_solicitante) ante errores de
// persistencia simulados, sin base de datos.

// cuerpoPutPersistenciaDecisionSolicitante retorna un cuerpo válido para PUT al registro 1.
func cuerpoPutPersistenciaDecisionSolicitante(t *testing.T) string {
	t.Helper()
	cuerpo := cuerpoValidacionDecisionSolicitante(t)
	cuerpo["Id"] = 1
	return helpers.JSON(t, cuerpo)
}

func TestDecisionSolicitantePostErrorPersistencia(t *testing.T) {
	cuerpo := helpers.JSON(t, cuerpoValidacionDecisionSolicitante(t))
	t.Run("INSERT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionDecisionSolicitante+"/", cuerpo, "INSERT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "POST")
	})
	t.Run("INSERT restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPost, rutaValidacionDecisionSolicitante+"/", cuerpo, "INSERT", helpers.ErrRestriccion, http.StatusBadRequest, "POST")
	})
}

func TestDecisionSolicitanteGetOneErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionDecisionSolicitante+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetOne")
	})
}

func TestDecisionSolicitanteGetAllErrorPersistencia(t *testing.T) {
	t.Run("SELECT base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodGet, rutaValidacionDecisionSolicitante+"/", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "GetAll")
	})
}

func TestDecisionSolicitantePutErrorPersistencia(t *testing.T) {
	cuerpo := cuerpoPutPersistenciaDecisionSolicitante(t)
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionDecisionSolicitante+"/1", cuerpo, "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionDecisionSolicitante+"/1", cuerpo, "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Put")
	})
	t.Run("UPDATE restriccion", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodPut, rutaValidacionDecisionSolicitante+"/1", cuerpo, "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Put")
	})
}

func TestDecisionSolicitanteDeleteErrorPersistencia(t *testing.T) {
	t.Run("lectura base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionDecisionSolicitante+"/1", "", "SELECT", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE base no disponible", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionDecisionSolicitante+"/1", "", "UPDATE", helpers.ErrBaseNoDisponible, http.StatusInternalServerError, "Delete")
	})
	t.Run("UPDATE otro error", func(t *testing.T) {
		helpers.ProbarErrorPersistencia(t, http.MethodDelete, rutaValidacionDecisionSolicitante+"/1", "", "UPDATE", helpers.ErrRestriccion, http.StatusBadRequest, "Delete")
	})
}
