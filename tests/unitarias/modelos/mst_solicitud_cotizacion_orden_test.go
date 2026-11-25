// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package modelos

import (
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de modelo del ordenamiento de GetAllMstSolicitudCotizacion, sin base de datos.

func TestMstSolicitudCotizacionGetAllOrdenSinOrderEsAsc(t *testing.T) {
	helpers.LimpiarSentencia()
	helpers.DefinirFilas(map[string]interface{}{"id": 1}, map[string]interface{}{"id": 2})
	if _, _, err := models.GetAllMstSolicitudCotizacion(nil, nil, []string{"Id"}, nil, 0, 10); err != nil {
		t.Fatal("Se esperaba sortby sin order sin error y se obtuvo:", err)
	}
	helpers.VerificarSQL(t, "ORDER BY T0.\"id\" ASC")
}

func TestMstSolicitudCotizacionGetAllOrdenDesc(t *testing.T) {
	helpers.LimpiarSentencia()
	if _, _, err := models.GetAllMstSolicitudCotizacion(nil, nil, []string{"Id"}, []string{"desc"}, 0, 10); err != nil {
		t.Fatal("Se esperaba order desc sin error y se obtuvo:", err)
	}
	helpers.VerificarSQL(t, "ORDER BY T0.\"id\" DESC")
}

func TestMstSolicitudCotizacionGetAllOrdenCampoInexistente(t *testing.T) {
	helpers.LimpiarSentencia()
	err := helpers.LlamarSinPanico(t, func() error {
		_, _, err := models.GetAllMstSolicitudCotizacion(nil, nil, []string{"campo_falso"}, []string{"asc"}, 0, 10)
		return err
	})
	helpers.VerificarErrorSortInvalido(t, err, "campo_falso")
	helpers.VerificarSinConsulta(t)
}

func TestMstSolicitudCotizacionGetAllOrdenOrderInvalido(t *testing.T) {
	helpers.LimpiarSentencia()
	_, _, err := models.GetAllMstSolicitudCotizacion(nil, nil, []string{"Id"}, []string{"ascending"}, 0, 10)
	helpers.VerificarErrorSortInvalido(t, err, "")
	helpers.VerificarSinConsulta(t)
}

func TestMstSolicitudCotizacionGetAllPaginacionYTotal(t *testing.T) {
	helpers.LimpiarSentencia()
	helpers.DefinirFilas(map[string]interface{}{"id": 1}, map[string]interface{}{"id": 2}, map[string]interface{}{"id": 3})
	_, total, err := models.GetAllMstSolicitudCotizacion(nil, nil, []string{"Id"}, []string{"asc"}, 1, 1)
	if err != nil {
		t.Fatal("Se esperaba limit y offset sin error y se obtuvo:", err)
	}
	helpers.VerificarIgual(t, "Total con limit=1 y offset=1", total, int64(3))
	helpers.VerificarSQL(t, "LIMIT 1 OFFSET 1")
}
