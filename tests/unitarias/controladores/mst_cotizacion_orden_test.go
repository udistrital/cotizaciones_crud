// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador del GetAll de MstCotizacionController (/v1/mst_cotizacion), sin base de datos.

const rutaOrdenMstCotizacion = "/v1/mst_cotizacion"

func TestMstCotizacionGetAllOrdenSortbySinOrderEsAsc(t *testing.T) {
	helpers.DefinirFilas(map[string]interface{}{"id": 1}, map[string]interface{}{"id": 2}, map[string]interface{}{"id": 3})
	lista, total := helpers.ConsultarOrdenado(t, rutaOrdenMstCotizacion, "sortby=Id", "ORDER BY T0.\"id\" ASC")
	helpers.VerificarIgual(t, "Registros devueltos", len(lista), 3)
	helpers.VerificarIgual(t, "Total", total, 3)
}

func TestMstCotizacionGetAllOrdenOrderDesc(t *testing.T) {
	helpers.DefinirFilas(map[string]interface{}{"id": 3}, map[string]interface{}{"id": 2}, map[string]interface{}{"id": 1})
	helpers.ConsultarOrdenado(t, rutaOrdenMstCotizacion, "sortby=Id&order=desc", "ORDER BY T0.\"id\" DESC")
}

func TestMstCotizacionGetAllOrdenCampoInexistente(t *testing.T) {
	helpers.ProbarSortInvalido(t, rutaOrdenMstCotizacion, "sortby=campo_falso&order=asc")
	helpers.ProbarSortInvalido(t, rutaOrdenMstCotizacion, "sortby=campo_falso")
}

func TestMstCotizacionGetAllOrdenOrderInvalido(t *testing.T) {
	helpers.ProbarSortInvalido(t, rutaOrdenMstCotizacion, "sortby=Id&order=ascending")
}

func TestMstCotizacionGetAllPaginacionYTotal(t *testing.T) {
	helpers.DefinirFilas(map[string]interface{}{"id": 1}, map[string]interface{}{"id": 2}, map[string]interface{}{"id": 3})
	_, total := helpers.ConsultarOrdenado(t, rutaOrdenMstCotizacion, "sortby=Id&order=asc&limit=1&offset=1", "LIMIT 1 OFFSET 1")
	helpers.VerificarIgual(t, "Total con limit=1 y offset=1", total, 3)
}
