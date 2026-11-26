// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador del formato de las columnas date de
// HistMstSolicitudCotizacionController (/v1/hist_mst_solicitud_cotizacion), sin base de datos.

// La base devuelve las columnas date como las entrega lib/pq (medianoche marcada como
// UTC): el JSON trae solo el día guardado.
func TestHistMstSolicitudCotizacionGetFechasSoloDia(t *testing.T) {
	helpers.ProbarGetFechas(t, rutaValidacionHistMstSolicitudCotizacion, map[string]interface{}{
		"fecha_apertura": helpers.FechaGuardada,
		"fecha_cierre":   helpers.FechaGuardada,
	}, map[string]string{
		"FechaApertura": helpers.FechaJSON,
		"FechaCierre":   helpers.FechaJSON,
	})
}

// FechaApertura es una columna date: POST acepta solo el día o una fecha y hora RFC 3339 (de la
// que toma el día escrito), guarda ese día sin correrlo por la zona horaria, lo devuelve
// solo con el día y rechaza otros formatos.
func TestHistMstSolicitudCotizacionPostFechaAperturaSoloDia(t *testing.T) {
	for enviado, guardado := range map[string]string{
		"2026-01-15":                "2026-01-15",
		"2026-01-15T00:00:00Z":      "2026-01-15",
		"2026-01-15T23:30:00-05:00": "2026-01-15",
	} {
		t.Run(enviado, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			cuerpo["FechaApertura"] = enviado
			helpers.LimpiarSentencia()
			r := helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionHistMstSolicitudCotizacion+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
			helpers.VerificarColumnaEnviada(t, "fecha_apertura", guardado)
			helpers.VerificarIgual(t, "Campo FechaApertura en la respuesta", r.Dato(t)["FechaApertura"], guardado)
		})
	}
	for _, invalido := range []string{"15/01/2026", "2026-13-45", "texto"} {
		t.Run("invalido "+invalido, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			cuerpo["FechaApertura"] = invalido
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}

// FechaCierre es una columna date: POST acepta solo el día o una fecha y hora RFC 3339 (de la
// que toma el día escrito), guarda ese día sin correrlo por la zona horaria, lo devuelve
// solo con el día y rechaza otros formatos.
func TestHistMstSolicitudCotizacionPostFechaCierreSoloDia(t *testing.T) {
	for enviado, guardado := range map[string]string{
		"2026-01-15":                "2026-01-15",
		"2026-01-15T00:00:00Z":      "2026-01-15",
		"2026-01-15T23:30:00-05:00": "2026-01-15",
	} {
		t.Run(enviado, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			cuerpo["FechaCierre"] = enviado
			helpers.LimpiarSentencia()
			r := helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionHistMstSolicitudCotizacion+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
			helpers.VerificarColumnaEnviada(t, "fecha_cierre", guardado)
			helpers.VerificarIgual(t, "Campo FechaCierre en la respuesta", r.Dato(t)["FechaCierre"], guardado)
		})
	}
	for _, invalido := range []string{"15/01/2026", "2026-13-45", "texto"} {
		t.Run("invalido "+invalido, func(t *testing.T) {
			cuerpo := cuerpoValidacionHistMstSolicitudCotizacion(t)
			cuerpo["FechaCierre"] = invalido
			helpers.ProbarPostRechazado(t, rutaValidacionHistMstSolicitudCotizacion, cuerpo)
		})
	}
}
