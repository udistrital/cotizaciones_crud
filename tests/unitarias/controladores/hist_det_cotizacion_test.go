// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas unitarias del POST de HistDetCotizacionController (/v1/hist_det_cotizacion), sin base de datos.
// Cada escenario aplica a todos los campos escalares opcionales de la entidad.

const rutaHistDetCotizacion = "/v1/hist_det_cotizacion"

// ejemploHistDetCotizacion retorna la entidad con todos sus campos de ejemplo.
func ejemploHistDetCotizacion() *models.HistDetCotizacion {
	return &models.HistDetCotizacion{
		DetCotizacionId:          1,
		MstCotizacionId:          &models.MstCotizacion{Id: 1},
		DetSolicitudCotizacionId: 1,
		ValorUnitario:            helpers.Ptr(1.5),
		Iva:                      1.5,
		FichaTecnica:             "Prueba",
		RegistroActual:           true,
		UsuarioCreacionId:        1,
		UsuarioModificacionId:    1,
	}
}

func cuerpoHistDetCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, ejemploHistDetCotizacion())
}

// TestPostHistDetCotizacionOpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPostHistDetCotizacionOpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpoHistDetCotizacion(t)
	delete(cuerpo, "ValorUnitario")
	helpers.ProbarPostOpcionales(t, rutaHistDetCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "ValorUnitario", Columna: "valor_unitario", Valor: nil},
	})
}

// TestPostHistDetCotizacionOpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPostHistDetCotizacionOpcionalesNull(t *testing.T) {
	cuerpo := cuerpoHistDetCotizacion(t)
	cuerpo["ValorUnitario"] = nil
	helpers.ProbarPostOpcionales(t, rutaHistDetCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "ValorUnitario", Columna: "valor_unitario", Valor: nil},
	})
}

// TestPostHistDetCotizacionOpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPostHistDetCotizacionOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploHistDetCotizacion()
	helpers.ProbarPostOpcionales(t, rutaHistDetCotizacion, cuerpoHistDetCotizacion(t), []helpers.CampoOpcional{
		{Campo: "ValorUnitario", Columna: "valor_unitario", Valor: helpers.Desreferenciar(ejemplo.ValorUnitario)},
	})
}

// TestGetHistDetCotizacionOpcionalesNull: la base de datos devuelve NULL en cada campo opcional.
func TestGetHistDetCotizacionOpcionalesNull(t *testing.T) {
	helpers.ProbarGetOpcionales(t, rutaHistDetCotizacion, []helpers.CampoOpcional{
		{Campo: "ValorUnitario", Columna: "valor_unitario", Valor: nil},
	})
}

// TestGetHistDetCotizacionOpcionalesConValor: la base de datos devuelve el valor de cada campo opcional.
func TestGetHistDetCotizacionOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploHistDetCotizacion()
	helpers.ProbarGetOpcionales(t, rutaHistDetCotizacion, []helpers.CampoOpcional{
		{Campo: "ValorUnitario", Columna: "valor_unitario", Valor: helpers.Desreferenciar(ejemplo.ValorUnitario)},
	})
}

// TestPutHistDetCotizacionOpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPutHistDetCotizacionOpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpoHistDetCotizacion(t)
	delete(cuerpo, "ValorUnitario")
	helpers.ProbarPutOpcionales(t, rutaHistDetCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "ValorUnitario", Columna: "valor_unitario", Valor: nil},
	})
}

// TestPutHistDetCotizacionOpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPutHistDetCotizacionOpcionalesNull(t *testing.T) {
	cuerpo := cuerpoHistDetCotizacion(t)
	cuerpo["ValorUnitario"] = nil
	helpers.ProbarPutOpcionales(t, rutaHistDetCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "ValorUnitario", Columna: "valor_unitario", Valor: nil},
	})
}

// TestPutHistDetCotizacionOpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPutHistDetCotizacionOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploHistDetCotizacion()
	helpers.ProbarPutOpcionales(t, rutaHistDetCotizacion, cuerpoHistDetCotizacion(t), []helpers.CampoOpcional{
		{Campo: "ValorUnitario", Columna: "valor_unitario", Valor: helpers.Desreferenciar(ejemplo.ValorUnitario)},
	})
}
