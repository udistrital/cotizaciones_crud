// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package modelos

import (
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas unitarias de models.AddHistDetCotizacion para los campos escalares opcionales, sin base de datos.

// TestHistDetCotizacionValorUnitario verifica el valor de valor_unitario en el INSERT para
// los tres casos del campo ValorUnitario.
func TestHistDetCotizacionValorUnitario(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistDetCotizacion{
			DetCotizacionId:          1,
			MstCotizacionId:          &models.MstCotizacion{Id: 1},
			DetSolicitudCotizacionId: 1,
			Iva:                      1.5,
			FichaTecnica:             "Prueba",
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if _, err := models.AddHistDetCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistDetCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "valor_unitario", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistDetCotizacion{
			DetCotizacionId:          1,
			MstCotizacionId:          &models.MstCotizacion{Id: 1},
			DetSolicitudCotizacionId: 1,
			Iva:                      1.5,
			FichaTecnica:             "Prueba",
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			ValorUnitario:            nil,
		}
		if _, err := models.AddHistDetCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistDetCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "valor_unitario", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistDetCotizacion{
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
		esperado := helpers.Desreferenciar(m.ValorUnitario)
		if _, err := models.AddHistDetCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistDetCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "valor_unitario", esperado)
	})
}

// TestHistDetCotizacionValorUnitarioGet verifica la lectura de valor_unitario: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistDetCotizacionValorUnitarioGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"valor_unitario": nil})
		v, err := models.GetHistDetCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistDetCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "ValorUnitario", nil, helpers.Desreferenciar(v.ValorUnitario))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr(1.5))
		helpers.DefinirFila(map[string]interface{}{"valor_unitario": esperado})
		v, err := models.GetHistDetCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistDetCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "ValorUnitario", esperado, helpers.Desreferenciar(v.ValorUnitario))
	})
}

// TestHistDetCotizacionValorUnitarioUpdate verifica el valor de valor_unitario en el UPDATE
// para los tres casos del campo ValorUnitario.
func TestHistDetCotizacionValorUnitarioUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistDetCotizacion{
			Id:                       1,
			DetCotizacionId:          1,
			MstCotizacionId:          &models.MstCotizacion{Id: 1},
			DetSolicitudCotizacionId: 1,
			Iva:                      1.5,
			FichaTecnica:             "Prueba",
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
		}
		if err := models.UpdateHistDetCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistDetCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "valor_unitario", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistDetCotizacion{
			Id:                       1,
			DetCotizacionId:          1,
			MstCotizacionId:          &models.MstCotizacion{Id: 1},
			DetSolicitudCotizacionId: 1,
			Iva:                      1.5,
			FichaTecnica:             "Prueba",
			RegistroActual:           true,
			UsuarioCreacionId:        1,
			UsuarioModificacionId:    1,
			ValorUnitario:            nil,
		}
		if err := models.UpdateHistDetCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistDetCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "valor_unitario", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistDetCotizacion{
			Id:                       1,
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
		esperado := helpers.Desreferenciar(m.ValorUnitario)
		if err := models.UpdateHistDetCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistDetCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "valor_unitario", esperado)
	})
}
