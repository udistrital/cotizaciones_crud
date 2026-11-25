// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas unitarias del POST de HistMstCotizacionController (/v1/hist_mst_cotizacion), sin base de datos.
// Cada escenario aplica a todos los campos escalares opcionales de la entidad.

const rutaHistMstCotizacion = "/v1/hist_mst_cotizacion"

// ejemploHistMstCotizacion retorna la entidad con todos sus campos de ejemplo.
func ejemploHistMstCotizacion() *models.HistMstCotizacion {
	return &models.HistMstCotizacion{
		MstCotizacionId:       &models.MstCotizacion{Id: 1},
		FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
		Descripcion:           helpers.Ptr("Prueba"),
		InformacionEntrega:    helpers.Ptr("Prueba"),
		Descuentos:            helpers.Ptr("Prueba"),
		Observaciones:         helpers.Ptr("Prueba"),
		FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
		SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
		RegistroActual:        true,
		UsuarioCreacionId:     1,
		UsuarioModificacionId: 1,
	}
}

func cuerpoHistMstCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, ejemploHistMstCotizacion())
}

// TestPostHistMstCotizacionOpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPostHistMstCotizacionOpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpoHistMstCotizacion(t)
	delete(cuerpo, "Descripcion")
	delete(cuerpo, "InformacionEntrega")
	delete(cuerpo, "Descuentos")
	delete(cuerpo, "Observaciones")
	helpers.ProbarPostOpcionales(t, rutaHistMstCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "Descripcion", Columna: "descripcion", Valor: nil},
		{Campo: "InformacionEntrega", Columna: "informacion_entrega", Valor: nil},
		{Campo: "Descuentos", Columna: "descuentos", Valor: nil},
		{Campo: "Observaciones", Columna: "observaciones", Valor: nil},
	})
}

// TestPostHistMstCotizacionOpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPostHistMstCotizacionOpcionalesNull(t *testing.T) {
	cuerpo := cuerpoHistMstCotizacion(t)
	cuerpo["Descripcion"] = nil
	cuerpo["InformacionEntrega"] = nil
	cuerpo["Descuentos"] = nil
	cuerpo["Observaciones"] = nil
	helpers.ProbarPostOpcionales(t, rutaHistMstCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "Descripcion", Columna: "descripcion", Valor: nil},
		{Campo: "InformacionEntrega", Columna: "informacion_entrega", Valor: nil},
		{Campo: "Descuentos", Columna: "descuentos", Valor: nil},
		{Campo: "Observaciones", Columna: "observaciones", Valor: nil},
	})
}

// TestPostHistMstCotizacionOpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPostHistMstCotizacionOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploHistMstCotizacion()
	helpers.ProbarPostOpcionales(t, rutaHistMstCotizacion, cuerpoHistMstCotizacion(t), []helpers.CampoOpcional{
		{Campo: "Descripcion", Columna: "descripcion", Valor: helpers.Desreferenciar(ejemplo.Descripcion)},
		{Campo: "InformacionEntrega", Columna: "informacion_entrega", Valor: helpers.Desreferenciar(ejemplo.InformacionEntrega)},
		{Campo: "Descuentos", Columna: "descuentos", Valor: helpers.Desreferenciar(ejemplo.Descuentos)},
		{Campo: "Observaciones", Columna: "observaciones", Valor: helpers.Desreferenciar(ejemplo.Observaciones)},
	})
}

// TestGetHistMstCotizacionOpcionalesNull: la base de datos devuelve NULL en cada campo opcional.
func TestGetHistMstCotizacionOpcionalesNull(t *testing.T) {
	helpers.ProbarGetOpcionales(t, rutaHistMstCotizacion, []helpers.CampoOpcional{
		{Campo: "Descripcion", Columna: "descripcion", Valor: nil},
		{Campo: "InformacionEntrega", Columna: "informacion_entrega", Valor: nil},
		{Campo: "Descuentos", Columna: "descuentos", Valor: nil},
		{Campo: "Observaciones", Columna: "observaciones", Valor: nil},
	})
}

// TestGetHistMstCotizacionOpcionalesConValor: la base de datos devuelve el valor de cada campo opcional.
func TestGetHistMstCotizacionOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploHistMstCotizacion()
	helpers.ProbarGetOpcionales(t, rutaHistMstCotizacion, []helpers.CampoOpcional{
		{Campo: "Descripcion", Columna: "descripcion", Valor: helpers.Desreferenciar(ejemplo.Descripcion)},
		{Campo: "InformacionEntrega", Columna: "informacion_entrega", Valor: helpers.Desreferenciar(ejemplo.InformacionEntrega)},
		{Campo: "Descuentos", Columna: "descuentos", Valor: helpers.Desreferenciar(ejemplo.Descuentos)},
		{Campo: "Observaciones", Columna: "observaciones", Valor: helpers.Desreferenciar(ejemplo.Observaciones)},
	})
}

// TestPutHistMstCotizacionOpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPutHistMstCotizacionOpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpoHistMstCotizacion(t)
	delete(cuerpo, "Descripcion")
	delete(cuerpo, "InformacionEntrega")
	delete(cuerpo, "Descuentos")
	delete(cuerpo, "Observaciones")
	helpers.ProbarPutOpcionales(t, rutaHistMstCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "Descripcion", Columna: "descripcion", Valor: nil},
		{Campo: "InformacionEntrega", Columna: "informacion_entrega", Valor: nil},
		{Campo: "Descuentos", Columna: "descuentos", Valor: nil},
		{Campo: "Observaciones", Columna: "observaciones", Valor: nil},
	})
}

// TestPutHistMstCotizacionOpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPutHistMstCotizacionOpcionalesNull(t *testing.T) {
	cuerpo := cuerpoHistMstCotizacion(t)
	cuerpo["Descripcion"] = nil
	cuerpo["InformacionEntrega"] = nil
	cuerpo["Descuentos"] = nil
	cuerpo["Observaciones"] = nil
	helpers.ProbarPutOpcionales(t, rutaHistMstCotizacion, cuerpo, []helpers.CampoOpcional{
		{Campo: "Descripcion", Columna: "descripcion", Valor: nil},
		{Campo: "InformacionEntrega", Columna: "informacion_entrega", Valor: nil},
		{Campo: "Descuentos", Columna: "descuentos", Valor: nil},
		{Campo: "Observaciones", Columna: "observaciones", Valor: nil},
	})
}

// TestPutHistMstCotizacionOpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPutHistMstCotizacionOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploHistMstCotizacion()
	helpers.ProbarPutOpcionales(t, rutaHistMstCotizacion, cuerpoHistMstCotizacion(t), []helpers.CampoOpcional{
		{Campo: "Descripcion", Columna: "descripcion", Valor: helpers.Desreferenciar(ejemplo.Descripcion)},
		{Campo: "InformacionEntrega", Columna: "informacion_entrega", Valor: helpers.Desreferenciar(ejemplo.InformacionEntrega)},
		{Campo: "Descuentos", Columna: "descuentos", Valor: helpers.Desreferenciar(ejemplo.Descuentos)},
		{Campo: "Observaciones", Columna: "observaciones", Valor: helpers.Desreferenciar(ejemplo.Observaciones)},
	})
}
