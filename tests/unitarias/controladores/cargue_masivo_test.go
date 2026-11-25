// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas unitarias del POST de CargueMasivoController (/v1/cargue_masivo), sin base de datos.
// Cada escenario aplica a todos los campos escalares opcionales de la entidad.

const rutaCargueMasivo = "/v1/cargue_masivo"

// ejemploCargueMasivo retorna la entidad con todos sus campos de ejemplo.
func ejemploCargueMasivo() *models.CargueMasivo {
	return &models.CargueMasivo{
		HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
		EnlacePlantilla:              helpers.Ptr("Prueba"),
		CargadoPorTerceroId:          1,
		TotalRegistros:               helpers.Ptr(1),
		Aceptados:                    helpers.Ptr(1),
		Rechazados:                   helpers.Ptr(1),
		Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
		UsuarioCreacionId:            1,
		UsuarioModificacionId:        1,
	}
}

func cuerpoCargueMasivo(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, ejemploCargueMasivo())
}

// TestPostCargueMasivoOpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPostCargueMasivoOpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpoCargueMasivo(t)
	delete(cuerpo, "EnlacePlantilla")
	delete(cuerpo, "TotalRegistros")
	delete(cuerpo, "Aceptados")
	delete(cuerpo, "Rechazados")
	delete(cuerpo, "Respuesta")
	helpers.ProbarPostOpcionales(t, rutaCargueMasivo, cuerpo, []helpers.CampoOpcional{
		{Campo: "EnlacePlantilla", Columna: "enlace_plantilla", Valor: nil},
		{Campo: "TotalRegistros", Columna: "total_registros", Valor: nil},
		{Campo: "Aceptados", Columna: "aceptados", Valor: nil},
		{Campo: "Rechazados", Columna: "rechazados", Valor: nil},
		{Campo: "Respuesta", Columna: "respuesta", Valor: nil},
	})
}

// TestPostCargueMasivoOpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPostCargueMasivoOpcionalesNull(t *testing.T) {
	cuerpo := cuerpoCargueMasivo(t)
	cuerpo["EnlacePlantilla"] = nil
	cuerpo["TotalRegistros"] = nil
	cuerpo["Aceptados"] = nil
	cuerpo["Rechazados"] = nil
	cuerpo["Respuesta"] = nil
	helpers.ProbarPostOpcionales(t, rutaCargueMasivo, cuerpo, []helpers.CampoOpcional{
		{Campo: "EnlacePlantilla", Columna: "enlace_plantilla", Valor: nil},
		{Campo: "TotalRegistros", Columna: "total_registros", Valor: nil},
		{Campo: "Aceptados", Columna: "aceptados", Valor: nil},
		{Campo: "Rechazados", Columna: "rechazados", Valor: nil},
		{Campo: "Respuesta", Columna: "respuesta", Valor: nil},
	})
}

// TestPostCargueMasivoOpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPostCargueMasivoOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploCargueMasivo()
	helpers.ProbarPostOpcionales(t, rutaCargueMasivo, cuerpoCargueMasivo(t), []helpers.CampoOpcional{
		{Campo: "EnlacePlantilla", Columna: "enlace_plantilla", Valor: helpers.Desreferenciar(ejemplo.EnlacePlantilla)},
		{Campo: "TotalRegistros", Columna: "total_registros", Valor: helpers.Desreferenciar(ejemplo.TotalRegistros)},
		{Campo: "Aceptados", Columna: "aceptados", Valor: helpers.Desreferenciar(ejemplo.Aceptados)},
		{Campo: "Rechazados", Columna: "rechazados", Valor: helpers.Desreferenciar(ejemplo.Rechazados)},
		{Campo: "Respuesta", Columna: "respuesta", Valor: helpers.Desreferenciar(ejemplo.Respuesta)},
	})
}

// TestGetCargueMasivoOpcionalesNull: la base de datos devuelve NULL en cada campo opcional.
func TestGetCargueMasivoOpcionalesNull(t *testing.T) {
	helpers.ProbarGetOpcionales(t, rutaCargueMasivo, []helpers.CampoOpcional{
		{Campo: "EnlacePlantilla", Columna: "enlace_plantilla", Valor: nil},
		{Campo: "TotalRegistros", Columna: "total_registros", Valor: nil},
		{Campo: "Aceptados", Columna: "aceptados", Valor: nil},
		{Campo: "Rechazados", Columna: "rechazados", Valor: nil},
		{Campo: "Respuesta", Columna: "respuesta", Valor: nil},
	})
}

// TestGetCargueMasivoOpcionalesConValor: la base de datos devuelve el valor de cada campo opcional.
func TestGetCargueMasivoOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploCargueMasivo()
	helpers.ProbarGetOpcionales(t, rutaCargueMasivo, []helpers.CampoOpcional{
		{Campo: "EnlacePlantilla", Columna: "enlace_plantilla", Valor: helpers.Desreferenciar(ejemplo.EnlacePlantilla)},
		{Campo: "TotalRegistros", Columna: "total_registros", Valor: helpers.Desreferenciar(ejemplo.TotalRegistros)},
		{Campo: "Aceptados", Columna: "aceptados", Valor: helpers.Desreferenciar(ejemplo.Aceptados)},
		{Campo: "Rechazados", Columna: "rechazados", Valor: helpers.Desreferenciar(ejemplo.Rechazados)},
		{Campo: "Respuesta", Columna: "respuesta", Valor: helpers.Desreferenciar(ejemplo.Respuesta)},
	})
}

// TestPutCargueMasivoOpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPutCargueMasivoOpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpoCargueMasivo(t)
	delete(cuerpo, "EnlacePlantilla")
	delete(cuerpo, "TotalRegistros")
	delete(cuerpo, "Aceptados")
	delete(cuerpo, "Rechazados")
	delete(cuerpo, "Respuesta")
	helpers.ProbarPutOpcionales(t, rutaCargueMasivo, cuerpo, []helpers.CampoOpcional{
		{Campo: "EnlacePlantilla", Columna: "enlace_plantilla", Valor: nil},
		{Campo: "TotalRegistros", Columna: "total_registros", Valor: nil},
		{Campo: "Aceptados", Columna: "aceptados", Valor: nil},
		{Campo: "Rechazados", Columna: "rechazados", Valor: nil},
		{Campo: "Respuesta", Columna: "respuesta", Valor: nil},
	})
}

// TestPutCargueMasivoOpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPutCargueMasivoOpcionalesNull(t *testing.T) {
	cuerpo := cuerpoCargueMasivo(t)
	cuerpo["EnlacePlantilla"] = nil
	cuerpo["TotalRegistros"] = nil
	cuerpo["Aceptados"] = nil
	cuerpo["Rechazados"] = nil
	cuerpo["Respuesta"] = nil
	helpers.ProbarPutOpcionales(t, rutaCargueMasivo, cuerpo, []helpers.CampoOpcional{
		{Campo: "EnlacePlantilla", Columna: "enlace_plantilla", Valor: nil},
		{Campo: "TotalRegistros", Columna: "total_registros", Valor: nil},
		{Campo: "Aceptados", Columna: "aceptados", Valor: nil},
		{Campo: "Rechazados", Columna: "rechazados", Valor: nil},
		{Campo: "Respuesta", Columna: "respuesta", Valor: nil},
	})
}

// TestPutCargueMasivoOpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPutCargueMasivoOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploCargueMasivo()
	helpers.ProbarPutOpcionales(t, rutaCargueMasivo, cuerpoCargueMasivo(t), []helpers.CampoOpcional{
		{Campo: "EnlacePlantilla", Columna: "enlace_plantilla", Valor: helpers.Desreferenciar(ejemplo.EnlacePlantilla)},
		{Campo: "TotalRegistros", Columna: "total_registros", Valor: helpers.Desreferenciar(ejemplo.TotalRegistros)},
		{Campo: "Aceptados", Columna: "aceptados", Valor: helpers.Desreferenciar(ejemplo.Aceptados)},
		{Campo: "Rechazados", Columna: "rechazados", Valor: helpers.Desreferenciar(ejemplo.Rechazados)},
		{Campo: "Respuesta", Columna: "respuesta", Valor: helpers.Desreferenciar(ejemplo.Respuesta)},
	})
}
