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
// ValidacionOrdenadorController (/v1/validacion_ordenador), sin base de datos.

const rutaValidacionValidacionOrdenador = "/v1/validacion_ordenador"

// cuerpoValidacionValidacionOrdenador retorna un cuerpo válido para POST y PUT, con todos los campos de ejemplo.
func cuerpoValidacionValidacionOrdenador(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, &models.ValidacionOrdenador{
		MstSolicitudCotizacionId: &models.MstSolicitudCotizacion{Id: 1},
		Observaciones:            "Prueba",
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	})
}

// GetOne con un Id de ruta no numérico responde 404 sin consultar la base.
func TestValidacionOrdenadorGetOneIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodGet, rutaValidacionValidacionOrdenador, "")
}

// PUT con un Id de ruta no numérico responde 404 sin consultar la base.
func TestValidacionOrdenadorPutIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodPut, rutaValidacionValidacionOrdenador, helpers.JSON(t, cuerpoValidacionValidacionOrdenador(t)))
}

// DELETE con un Id de ruta no numérico responde 404 sin consultar la base.
func TestValidacionOrdenadorDeleteIdNoNumerico(t *testing.T) {
	helpers.ProbarIdNoNumerico(t, http.MethodDelete, rutaValidacionValidacionOrdenador, "")
}

// La columna observaciones es NOT NULL: POST sin Observaciones, con null o vacío
// responde 400 sin insertar (el valor cero de Go no reemplaza al campo ausente).
func TestValidacionOrdenadorPostSinObservaciones(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Observaciones") },
		"null":     func(c map[string]interface{}) { c["Observaciones"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Observaciones"] = "" },
		"espacios": func(c map[string]interface{}) { c["Observaciones"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionValidacionOrdenador(t)
			modificar(cuerpo)
			helpers.ProbarPostRechazado(t, rutaValidacionValidacionOrdenador, cuerpo)
		})
	}
}

// La columna observaciones es NOT NULL: PUT sin Observaciones, con null o vacío
// responde 400 sin actualizar.
func TestValidacionOrdenadorPutSinObservaciones(t *testing.T) {
	for nombre, modificar := range map[string]func(map[string]interface{}){
		"omitido":  func(c map[string]interface{}) { delete(c, "Observaciones") },
		"null":     func(c map[string]interface{}) { c["Observaciones"] = nil },
		"vacio":    func(c map[string]interface{}) { c["Observaciones"] = "" },
		"espacios": func(c map[string]interface{}) { c["Observaciones"] = "   " },
	} {
		t.Run(nombre, func(t *testing.T) {
			cuerpo := cuerpoValidacionValidacionOrdenador(t)
			modificar(cuerpo)
			helpers.ProbarPutRechazado(t, rutaValidacionValidacionOrdenador, cuerpo)
		})
	}
}
