// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package modelos

import (
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas unitarias de models.AddCargueMasivo para los campos escalares opcionales, sin base de datos.

// TestCargueMasivoEnlacePlantilla verifica el valor de enlace_plantilla en el INSERT para
// los tres casos del campo EnlacePlantilla.
func TestCargueMasivoEnlacePlantilla(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.CargueMasivo{
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
		}
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "enlace_plantilla", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.CargueMasivo{
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
			EnlacePlantilla:              nil,
		}
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "enlace_plantilla", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.CargueMasivo{
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
		esperado := helpers.Desreferenciar(m.EnlacePlantilla)
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "enlace_plantilla", esperado)
	})
}

// TestCargueMasivoEnlacePlantillaGet verifica la lectura de enlace_plantilla: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestCargueMasivoEnlacePlantillaGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"enlace_plantilla": nil})
		v, err := models.GetCargueMasivoById(1)
		if err != nil {
			t.Fatal("No se pudo consultar CargueMasivo:", err)
		}
		helpers.VerificarValorLeido(t, "EnlacePlantilla", nil, helpers.Desreferenciar(v.EnlacePlantilla))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr("Prueba"))
		helpers.DefinirFila(map[string]interface{}{"enlace_plantilla": esperado})
		v, err := models.GetCargueMasivoById(1)
		if err != nil {
			t.Fatal("No se pudo consultar CargueMasivo:", err)
		}
		helpers.VerificarValorLeido(t, "EnlacePlantilla", esperado, helpers.Desreferenciar(v.EnlacePlantilla))
	})
}

// TestCargueMasivoEnlacePlantillaUpdate verifica el valor de enlace_plantilla en el UPDATE
// para los tres casos del campo EnlacePlantilla.
func TestCargueMasivoEnlacePlantillaUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
		}
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "enlace_plantilla", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
			EnlacePlantilla:              nil,
		}
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "enlace_plantilla", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
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
		esperado := helpers.Desreferenciar(m.EnlacePlantilla)
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "enlace_plantilla", esperado)
	})
}

// TestCargueMasivoTotalRegistros verifica el valor de total_registros en el INSERT para
// los tres casos del campo TotalRegistros.
func TestCargueMasivoTotalRegistros(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.CargueMasivo{
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
		}
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "total_registros", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.CargueMasivo{
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
			TotalRegistros:               nil,
		}
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "total_registros", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.CargueMasivo{
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
		esperado := helpers.Desreferenciar(m.TotalRegistros)
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "total_registros", esperado)
	})
}

// TestCargueMasivoTotalRegistrosGet verifica la lectura de total_registros: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestCargueMasivoTotalRegistrosGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"total_registros": nil})
		v, err := models.GetCargueMasivoById(1)
		if err != nil {
			t.Fatal("No se pudo consultar CargueMasivo:", err)
		}
		helpers.VerificarValorLeido(t, "TotalRegistros", nil, helpers.Desreferenciar(v.TotalRegistros))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr(1))
		helpers.DefinirFila(map[string]interface{}{"total_registros": esperado})
		v, err := models.GetCargueMasivoById(1)
		if err != nil {
			t.Fatal("No se pudo consultar CargueMasivo:", err)
		}
		helpers.VerificarValorLeido(t, "TotalRegistros", esperado, helpers.Desreferenciar(v.TotalRegistros))
	})
}

// TestCargueMasivoTotalRegistrosUpdate verifica el valor de total_registros en el UPDATE
// para los tres casos del campo TotalRegistros.
func TestCargueMasivoTotalRegistrosUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
		}
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "total_registros", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
			TotalRegistros:               nil,
		}
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "total_registros", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
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
		esperado := helpers.Desreferenciar(m.TotalRegistros)
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "total_registros", esperado)
	})
}

// TestCargueMasivoAceptados verifica el valor de aceptados en el INSERT para
// los tres casos del campo Aceptados.
func TestCargueMasivoAceptados(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.CargueMasivo{
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
		}
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "aceptados", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.CargueMasivo{
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
			Aceptados:                    nil,
		}
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "aceptados", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.CargueMasivo{
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
		esperado := helpers.Desreferenciar(m.Aceptados)
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "aceptados", esperado)
	})
}

// TestCargueMasivoAceptadosGet verifica la lectura de aceptados: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestCargueMasivoAceptadosGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"aceptados": nil})
		v, err := models.GetCargueMasivoById(1)
		if err != nil {
			t.Fatal("No se pudo consultar CargueMasivo:", err)
		}
		helpers.VerificarValorLeido(t, "Aceptados", nil, helpers.Desreferenciar(v.Aceptados))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr(1))
		helpers.DefinirFila(map[string]interface{}{"aceptados": esperado})
		v, err := models.GetCargueMasivoById(1)
		if err != nil {
			t.Fatal("No se pudo consultar CargueMasivo:", err)
		}
		helpers.VerificarValorLeido(t, "Aceptados", esperado, helpers.Desreferenciar(v.Aceptados))
	})
}

// TestCargueMasivoAceptadosUpdate verifica el valor de aceptados en el UPDATE
// para los tres casos del campo Aceptados.
func TestCargueMasivoAceptadosUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
		}
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "aceptados", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
			Aceptados:                    nil,
		}
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "aceptados", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
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
		esperado := helpers.Desreferenciar(m.Aceptados)
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "aceptados", esperado)
	})
}

// TestCargueMasivoRechazados verifica el valor de rechazados en el INSERT para
// los tres casos del campo Rechazados.
func TestCargueMasivoRechazados(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.CargueMasivo{
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
		}
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "rechazados", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.CargueMasivo{
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
			Rechazados:                   nil,
		}
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "rechazados", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.CargueMasivo{
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
		esperado := helpers.Desreferenciar(m.Rechazados)
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "rechazados", esperado)
	})
}

// TestCargueMasivoRechazadosGet verifica la lectura de rechazados: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestCargueMasivoRechazadosGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"rechazados": nil})
		v, err := models.GetCargueMasivoById(1)
		if err != nil {
			t.Fatal("No se pudo consultar CargueMasivo:", err)
		}
		helpers.VerificarValorLeido(t, "Rechazados", nil, helpers.Desreferenciar(v.Rechazados))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr(1))
		helpers.DefinirFila(map[string]interface{}{"rechazados": esperado})
		v, err := models.GetCargueMasivoById(1)
		if err != nil {
			t.Fatal("No se pudo consultar CargueMasivo:", err)
		}
		helpers.VerificarValorLeido(t, "Rechazados", esperado, helpers.Desreferenciar(v.Rechazados))
	})
}

// TestCargueMasivoRechazadosUpdate verifica el valor de rechazados en el UPDATE
// para los tres casos del campo Rechazados.
func TestCargueMasivoRechazadosUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
		}
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "rechazados", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Respuesta:                    helpers.Ptr("{\"prueba\": true}"),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
			Rechazados:                   nil,
		}
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "rechazados", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
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
		esperado := helpers.Desreferenciar(m.Rechazados)
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "rechazados", esperado)
	})
}

// TestCargueMasivoRespuesta verifica el valor de respuesta en el INSERT para
// los tres casos del campo Respuesta.
func TestCargueMasivoRespuesta(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.CargueMasivo{
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
		}
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "respuesta", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.CargueMasivo{
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
			Respuesta:                    nil,
		}
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "respuesta", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.CargueMasivo{
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
		esperado := helpers.Desreferenciar(m.Respuesta)
		if _, err := models.AddCargueMasivo(m); err != nil {
			t.Fatal("No se pudo agregar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "respuesta", esperado)
	})
}

// TestCargueMasivoRespuestaGet verifica la lectura de respuesta: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestCargueMasivoRespuestaGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"respuesta": nil})
		v, err := models.GetCargueMasivoById(1)
		if err != nil {
			t.Fatal("No se pudo consultar CargueMasivo:", err)
		}
		helpers.VerificarValorLeido(t, "Respuesta", nil, helpers.Desreferenciar(v.Respuesta))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr("{\"prueba\": true}"))
		helpers.DefinirFila(map[string]interface{}{"respuesta": esperado})
		v, err := models.GetCargueMasivoById(1)
		if err != nil {
			t.Fatal("No se pudo consultar CargueMasivo:", err)
		}
		helpers.VerificarValorLeido(t, "Respuesta", esperado, helpers.Desreferenciar(v.Respuesta))
	})
}

// TestCargueMasivoRespuestaUpdate verifica el valor de respuesta en el UPDATE
// para los tres casos del campo Respuesta.
func TestCargueMasivoRespuestaUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
		}
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "respuesta", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
			HistMstSolicitudCotizacionId: &models.HistMstSolicitudCotizacion{Id: 1},
			EnlacePlantilla:              helpers.Ptr("Prueba"),
			CargadoPorTerceroId:          1,
			TotalRegistros:               helpers.Ptr(1),
			Aceptados:                    helpers.Ptr(1),
			Rechazados:                   helpers.Ptr(1),
			UsuarioCreacionId:            1,
			UsuarioModificacionId:        1,
			Respuesta:                    nil,
		}
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "respuesta", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.CargueMasivo{
			Id:                           1,
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
		esperado := helpers.Desreferenciar(m.Respuesta)
		if err := models.UpdateCargueMasivoById(m); err != nil {
			t.Fatal("No se pudo actualizar CargueMasivo:", err)
		}
		helpers.VerificarColumnaEnviada(t, "respuesta", esperado)
	})
}
