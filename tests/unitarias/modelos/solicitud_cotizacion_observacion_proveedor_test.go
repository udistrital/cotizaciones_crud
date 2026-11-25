// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package modelos

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas unitarias de models.AddSolicitudCotizacionObservacionProveedor para los campos escalares opcionales, sin base de datos.

// TestSolicitudCotizacionObservacionProveedorTexto verifica el valor de texto en el INSERT para
// los tres casos del campo Texto.
func TestSolicitudCotizacionObservacionProveedorTexto(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.SolicitudCotizacionObservacionProveedor{
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			FechaVistoSolicitante:             helpers.Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
		}
		if _, err := models.AddSolicitudCotizacionObservacionProveedor(m); err != nil {
			t.Fatal("No se pudo agregar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "texto", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.SolicitudCotizacionObservacionProveedor{
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			FechaVistoSolicitante:             helpers.Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
			Texto:                             nil,
		}
		if _, err := models.AddSolicitudCotizacionObservacionProveedor(m); err != nil {
			t.Fatal("No se pudo agregar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "texto", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.SolicitudCotizacionObservacionProveedor{
			Texto:                             helpers.Ptr("Prueba"),
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			FechaVistoSolicitante:             helpers.Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
		}
		esperado := helpers.Desreferenciar(m.Texto)
		if _, err := models.AddSolicitudCotizacionObservacionProveedor(m); err != nil {
			t.Fatal("No se pudo agregar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "texto", esperado)
	})
}

// TestSolicitudCotizacionObservacionProveedorTextoGet verifica la lectura de texto: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestSolicitudCotizacionObservacionProveedorTextoGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"texto": nil})
		v, err := models.GetSolicitudCotizacionObservacionProveedorById(1)
		if err != nil {
			t.Fatal("No se pudo consultar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarValorLeido(t, "Texto", nil, helpers.Desreferenciar(v.Texto))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr("Prueba"))
		helpers.DefinirFila(map[string]interface{}{"texto": esperado})
		v, err := models.GetSolicitudCotizacionObservacionProveedorById(1)
		if err != nil {
			t.Fatal("No se pudo consultar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarValorLeido(t, "Texto", esperado, helpers.Desreferenciar(v.Texto))
	})
}

// TestSolicitudCotizacionObservacionProveedorTextoUpdate verifica el valor de texto en el UPDATE
// para los tres casos del campo Texto.
func TestSolicitudCotizacionObservacionProveedorTextoUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.SolicitudCotizacionObservacionProveedor{
			Id:                                1,
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			FechaVistoSolicitante:             helpers.Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
		}
		if err := models.UpdateSolicitudCotizacionObservacionProveedorById(m); err != nil {
			t.Fatal("No se pudo actualizar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "texto", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.SolicitudCotizacionObservacionProveedor{
			Id:                                1,
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			FechaVistoSolicitante:             helpers.Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
			Texto:                             nil,
		}
		if err := models.UpdateSolicitudCotizacionObservacionProveedorById(m); err != nil {
			t.Fatal("No se pudo actualizar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "texto", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.SolicitudCotizacionObservacionProveedor{
			Id:                                1,
			Texto:                             helpers.Ptr("Prueba"),
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			FechaVistoSolicitante:             helpers.Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
		}
		esperado := helpers.Desreferenciar(m.Texto)
		if err := models.UpdateSolicitudCotizacionObservacionProveedorById(m); err != nil {
			t.Fatal("No se pudo actualizar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "texto", esperado)
	})
}

// TestSolicitudCotizacionObservacionProveedorFechaVistoSolicitante verifica el valor de fecha_visto_solicitante en el INSERT para
// los tres casos del campo FechaVistoSolicitante.
func TestSolicitudCotizacionObservacionProveedorFechaVistoSolicitante(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		m := &models.SolicitudCotizacionObservacionProveedor{
			Texto:                             helpers.Ptr("Prueba"),
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
		}
		if _, err := models.AddSolicitudCotizacionObservacionProveedor(m); err != nil {
			t.Fatal("No se pudo agregar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "fecha_visto_solicitante", nil)
	})

	t.Run("null", func(t *testing.T) {
		m := &models.SolicitudCotizacionObservacionProveedor{
			Texto:                             helpers.Ptr("Prueba"),
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
			FechaVistoSolicitante:             nil,
		}
		if _, err := models.AddSolicitudCotizacionObservacionProveedor(m); err != nil {
			t.Fatal("No se pudo agregar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "fecha_visto_solicitante", nil)
	})

	t.Run("valor", func(t *testing.T) {
		m := &models.SolicitudCotizacionObservacionProveedor{
			Texto:                             helpers.Ptr("Prueba"),
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			FechaVistoSolicitante:             helpers.Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
		}
		esperado := helpers.Desreferenciar(m.FechaVistoSolicitante)
		if _, err := models.AddSolicitudCotizacionObservacionProveedor(m); err != nil {
			t.Fatal("No se pudo agregar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "fecha_visto_solicitante", esperado)
	})
}

// TestSolicitudCotizacionObservacionProveedorFechaVistoSolicitanteGet verifica la lectura de fecha_visto_solicitante: NULL en la base
// de datos llega como nil y un valor llega igual.
func TestSolicitudCotizacionObservacionProveedorFechaVistoSolicitanteGet(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(map[string]interface{}{"fecha_visto_solicitante": nil})
		v, err := models.GetSolicitudCotizacionObservacionProveedorById(1)
		if err != nil {
			t.Fatal("No se pudo consultar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarValorLeido(t, "FechaVistoSolicitante", nil, helpers.Desreferenciar(v.FechaVistoSolicitante))
	})

	t.Run("valor", func(t *testing.T) {
		esperado := helpers.Desreferenciar(helpers.Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)))
		helpers.DefinirFila(map[string]interface{}{"fecha_visto_solicitante": esperado})
		v, err := models.GetSolicitudCotizacionObservacionProveedorById(1)
		if err != nil {
			t.Fatal("No se pudo consultar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarValorLeido(t, "FechaVistoSolicitante", esperado, helpers.Desreferenciar(v.FechaVistoSolicitante))
	})
}

// TestSolicitudCotizacionObservacionProveedorFechaVistoSolicitanteUpdate verifica el valor de fecha_visto_solicitante en el UPDATE
// para los tres casos del campo FechaVistoSolicitante.
func TestSolicitudCotizacionObservacionProveedorFechaVistoSolicitanteUpdate(t *testing.T) {
	t.Run("omitido", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.SolicitudCotizacionObservacionProveedor{
			Id:                                1,
			Texto:                             helpers.Ptr("Prueba"),
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
		}
		if err := models.UpdateSolicitudCotizacionObservacionProveedorById(m); err != nil {
			t.Fatal("No se pudo actualizar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "fecha_visto_solicitante", nil)
	})

	t.Run("null", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.SolicitudCotizacionObservacionProveedor{
			Id:                                1,
			Texto:                             helpers.Ptr("Prueba"),
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
			FechaVistoSolicitante:             nil,
		}
		if err := models.UpdateSolicitudCotizacionObservacionProveedorById(m); err != nil {
			t.Fatal("No se pudo actualizar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "fecha_visto_solicitante", nil)
	})

	t.Run("valor", func(t *testing.T) {
		helpers.DefinirFila(nil)
		m := &models.SolicitudCotizacionObservacionProveedor{
			Id:                                1,
			Texto:                             helpers.Ptr("Prueba"),
			RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
			VistoSolicitante:                  true,
			FechaVistoSolicitante:             helpers.Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
			UsuarioCreacionId:                 1,
			UsuarioModificacionId:             1,
		}
		esperado := helpers.Desreferenciar(m.FechaVistoSolicitante)
		if err := models.UpdateSolicitudCotizacionObservacionProveedorById(m); err != nil {
			t.Fatal("No se pudo actualizar SolicitudCotizacionObservacionProveedor:", err)
		}
		helpers.VerificarColumnaEnviada(t, "fecha_visto_solicitante", esperado)
	})
}
