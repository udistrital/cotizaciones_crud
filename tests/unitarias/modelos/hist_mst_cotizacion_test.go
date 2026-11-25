// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package modelos

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas unitarias de models.AddHistMstCotizacion para los campos escalares opcionales, sin base de datos.

// TestHistMstCotizacionDescripcion verifica el valor de descripcion en el INSERT para
// los tres casos del campo Descripcion.
func TestHistMstCotizacionDescripcion(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstCotizacion{
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
		}
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descripcion", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstCotizacion{
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
			Descripcion:           nil,
		}
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descripcion", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstCotizacion{
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
		esperado := helpers.Desreferenciar(m.Descripcion)
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descripcion", esperado)
	})
}

// TestHistMstCotizacionDescripcionGet verifica la lectura de descripcion: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstCotizacionDescripcionGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"descripcion": nil})
		v, err := models.GetHistMstCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "Descripcion", nil, helpers.Desreferenciar(v.Descripcion))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr("Prueba"))
		helpers.DefinirFila(map[string]interface{}{"descripcion": esperado})
		v, err := models.GetHistMstCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "Descripcion", esperado, helpers.Desreferenciar(v.Descripcion))
	})
}

// TestHistMstCotizacionDescripcionUpdate verifica el valor de descripcion en el UPDATE
// para los tres casos del campo Descripcion.
func TestHistMstCotizacionDescripcionUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
		}
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descripcion", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
			Descripcion:           nil,
		}
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descripcion", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
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
		esperado := helpers.Desreferenciar(m.Descripcion)
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descripcion", esperado)
	})
}

// TestHistMstCotizacionInformacionEntrega verifica el valor de informacion_entrega en el INSERT para
// los tres casos del campo InformacionEntrega.
func TestHistMstCotizacionInformacionEntrega(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstCotizacion{
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
		}
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "informacion_entrega", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstCotizacion{
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
			InformacionEntrega:    nil,
		}
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "informacion_entrega", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstCotizacion{
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
		esperado := helpers.Desreferenciar(m.InformacionEntrega)
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "informacion_entrega", esperado)
	})
}

// TestHistMstCotizacionInformacionEntregaGet verifica la lectura de informacion_entrega: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstCotizacionInformacionEntregaGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"informacion_entrega": nil})
		v, err := models.GetHistMstCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "InformacionEntrega", nil, helpers.Desreferenciar(v.InformacionEntrega))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr("Prueba"))
		helpers.DefinirFila(map[string]interface{}{"informacion_entrega": esperado})
		v, err := models.GetHistMstCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "InformacionEntrega", esperado, helpers.Desreferenciar(v.InformacionEntrega))
	})
}

// TestHistMstCotizacionInformacionEntregaUpdate verifica el valor de informacion_entrega en el UPDATE
// para los tres casos del campo InformacionEntrega.
func TestHistMstCotizacionInformacionEntregaUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
		}
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "informacion_entrega", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
			InformacionEntrega:    nil,
		}
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "informacion_entrega", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
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
		esperado := helpers.Desreferenciar(m.InformacionEntrega)
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "informacion_entrega", esperado)
	})
}

// TestHistMstCotizacionDescuentos verifica el valor de descuentos en el INSERT para
// los tres casos del campo Descuentos.
func TestHistMstCotizacionDescuentos(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstCotizacion{
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
		}
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descuentos", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstCotizacion{
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
			Descuentos:            nil,
		}
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descuentos", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstCotizacion{
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
		esperado := helpers.Desreferenciar(m.Descuentos)
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descuentos", esperado)
	})
}

// TestHistMstCotizacionDescuentosGet verifica la lectura de descuentos: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstCotizacionDescuentosGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"descuentos": nil})
		v, err := models.GetHistMstCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "Descuentos", nil, helpers.Desreferenciar(v.Descuentos))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr("Prueba"))
		helpers.DefinirFila(map[string]interface{}{"descuentos": esperado})
		v, err := models.GetHistMstCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "Descuentos", esperado, helpers.Desreferenciar(v.Descuentos))
	})
}

// TestHistMstCotizacionDescuentosUpdate verifica el valor de descuentos en el UPDATE
// para los tres casos del campo Descuentos.
func TestHistMstCotizacionDescuentosUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
		}
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descuentos", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Observaciones:         helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
			Descuentos:            nil,
		}
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descuentos", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
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
		esperado := helpers.Desreferenciar(m.Descuentos)
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "descuentos", esperado)
	})
}

// TestHistMstCotizacionObservaciones verifica el valor de observaciones en el INSERT para
// los tres casos del campo Observaciones.
func TestHistMstCotizacionObservaciones(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.HistMstCotizacion{
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
		}
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "observaciones", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.HistMstCotizacion{
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
			Observaciones:         nil,
		}
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "observaciones", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.HistMstCotizacion{
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
		esperado := helpers.Desreferenciar(m.Observaciones)
		if _, err := models.AddHistMstCotizacion(m); err != nil {
			t.Fatal("No se pudo agregar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "observaciones", esperado)
	})
}

// TestHistMstCotizacionObservacionesGet verifica la lectura de observaciones: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestHistMstCotizacionObservacionesGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"observaciones": nil})
		v, err := models.GetHistMstCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "Observaciones", nil, helpers.Desreferenciar(v.Observaciones))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr("Prueba"))
		helpers.DefinirFila(map[string]interface{}{"observaciones": esperado})
		v, err := models.GetHistMstCotizacionById(1)
		if err != nil {
			t.Fatal("No se pudo consultar HistMstCotizacion:", err)
		}
		helpers.VerificarValorLeido(t, "Observaciones", esperado, helpers.Desreferenciar(v.Observaciones))
	})
}

// TestHistMstCotizacionObservacionesUpdate verifica el valor de observaciones en el UPDATE
// para los tres casos del campo Observaciones.
func TestHistMstCotizacionObservacionesUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
		}
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "observaciones", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
			MstCotizacionId:       &models.MstCotizacion{Id: 1},
			FechaEnvio:            time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion:           helpers.Ptr("Prueba"),
			InformacionEntrega:    helpers.Ptr("Prueba"),
			Descuentos:            helpers.Ptr("Prueba"),
			FechaVencimiento:      time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			SoporteCotizacion:     "3f2b6c1e-8a4d-4f5e-9b7a-1c2d3e4f5a6b",
			RegistroActual:        true,
			UsuarioCreacionId:     1,
			UsuarioModificacionId: 1,
			Observaciones:         nil,
		}
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "observaciones", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.HistMstCotizacion{
			Id:                    1,
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
		esperado := helpers.Desreferenciar(m.Observaciones)
		if err := models.UpdateHistMstCotizacionById(m); err != nil {
			t.Fatal("No se pudo actualizar HistMstCotizacion:", err)
		}
		helpers.VerificarColumnaEnviada(t, "observaciones", esperado)
	})
}
