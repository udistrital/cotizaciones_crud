// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas de la capa de controlador de la validación de entradas de
// CargueMasivoController (/v1/cargue_masivo), sin base de datos.

const rutaValidacionCargueMasivo = "/v1/cargue_masivo"

// cuerpoValidacionCargueMasivo retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionCargueMasivo(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.CargueMasivo{
		HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
		EnlacePlantilla:              helpers.Ptr("Prueba"),
		CargadoPorTerceroId:          1,
		TotalRegistros:               helpers.Ptr(1),
		Aceptados:                    helpers.Ptr(1),
		Rechazados:                   helpers.Ptr(1),
		Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
		UsuarioCreacionId:            1,
		UsuarioModificacionId:        1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestCargueMasivoGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionCargueMasivo, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestCargueMasivoPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionCargueMasivo, helpers.JSON(t, cuerpoValidacionCargueMasivo(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestCargueMasivoDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionCargueMasivo, "")
}

// La columna cargado_por_tercero_id es NOT NULL: POST sin CargadoPorTerceroId, con null
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestCargueMasivoPostSinCargadoPorTerceroId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "CargadoPorTerceroId") },
		"null":    func(c map[string]interface{}) { c["CargadoPorTerceroId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionCargueMasivo(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionCargueMasivo, cuerpo)
		})
	}
}

// La columna cargado_por_tercero_id es NOT NULL: PUT sin CargadoPorTerceroId, con null
// responde 400 sin actualizar.
func TestCargueMasivoPutSinCargadoPorTerceroId(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido": func(c map[string]interface{}) { delete(c, "CargadoPorTerceroId") },
		"null":    func(c map[string]interface{}) { c["CargadoPorTerceroId"] = nil },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionCargueMasivo(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionCargueMasivo, cuerpo)
		})
	}
}

// Los campos obligatorios numéricos o booleanos enviados explícitamente con 0 o false
// son válidos: la validación solo rechaza el campo ausente o null.
func TestCargueMasivoPostObligatoriosConCeroOFalse(t *testing.T) {
	cuerpo := cuerpoValidacionCargueMasivo(t)
	cuerpo["CargadoPorTerceroId"] = 0
	helpers.VerificarExito(t, helpers.Peticion(http.MethodPost, rutaValidacionCargueMasivo+"/", helpers.JSON(t, cuerpo)), http.StatusCreated)
}
