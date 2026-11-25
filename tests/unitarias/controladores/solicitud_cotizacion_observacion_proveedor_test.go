// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"testing"
	"time"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas unitarias del POST de SolicitudCotizacionObservacionProveedorController (/v1/solicitud_cotizacion_observacion_proveedor), sin base de datos.
// Cada escenario aplica a todos los campos escalares opcionales de la entidad.

const rutaSolicitudCotizacionObservacionProveedor = "/v1/solicitud_cotizacion_observacion_proveedor"

// ejemploSolicitudCotizacionObservacionProveedor retorna la entidad con todos sus campos de ejemplo.
func ejemploSolicitudCotizacionObservacionProveedor() *models.SolicitudCotizacionObservacionProveedor {
	return &models.SolicitudCotizacionObservacionProveedor{
		Texto:                             helpers.Ptr("Prueba"),
		RelSolicitudCotizacionProveedorId: &models.RelSolicitudCotizacionProveedor{Id: 1},
		VistoSolicitante:                  true,
		FechaVistoSolicitante:             helpers.Ptr(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
		UsuarioCreacionId:                 1,
		UsuarioModificacionId:             1,
	}
}

func cuerpoSolicitudCotizacionObservacionProveedor(t *testing.T) map[string]interface{} {
	t.Helper()
	return helpers.AMapa(t, ejemploSolicitudCotizacionObservacionProveedor())
}

// TestPostSolicitudCotizacionObservacionProveedorOpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPostSolicitudCotizacionObservacionProveedorOpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpoSolicitudCotizacionObservacionProveedor(t)
	delete(cuerpo, "Texto")
	delete(cuerpo, "FechaVistoSolicitante")
	helpers.ProbarPostOpcionales(t, rutaSolicitudCotizacionObservacionProveedor, cuerpo, []helpers.CampoOpcional{
		{Campo: "Texto", Columna: "texto", Valor: nil},
		{Campo: "FechaVistoSolicitante", Columna: "fecha_visto_solicitante", Valor: nil},
	})
}

// TestPostSolicitudCotizacionObservacionProveedorOpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPostSolicitudCotizacionObservacionProveedorOpcionalesNull(t *testing.T) {
	cuerpo := cuerpoSolicitudCotizacionObservacionProveedor(t)
	cuerpo["Texto"] = nil
	cuerpo["FechaVistoSolicitante"] = nil
	helpers.ProbarPostOpcionales(t, rutaSolicitudCotizacionObservacionProveedor, cuerpo, []helpers.CampoOpcional{
		{Campo: "Texto", Columna: "texto", Valor: nil},
		{Campo: "FechaVistoSolicitante", Columna: "fecha_visto_solicitante", Valor: nil},
	})
}

// TestPostSolicitudCotizacionObservacionProveedorOpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPostSolicitudCotizacionObservacionProveedorOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploSolicitudCotizacionObservacionProveedor()
	helpers.ProbarPostOpcionales(t, rutaSolicitudCotizacionObservacionProveedor, cuerpoSolicitudCotizacionObservacionProveedor(t), []helpers.CampoOpcional{
		{Campo: "Texto", Columna: "texto", Valor: helpers.Desreferenciar(ejemplo.Texto)},
		{Campo: "FechaVistoSolicitante", Columna: "fecha_visto_solicitante", Valor: helpers.Desreferenciar(ejemplo.FechaVistoSolicitante)},
	})
}

// TestGetSolicitudCotizacionObservacionProveedorOpcionalesNull: la base de datos devuelve NULL en cada campo opcional.
func TestGetSolicitudCotizacionObservacionProveedorOpcionalesNull(t *testing.T) {
	helpers.ProbarGetOpcionales(t, rutaSolicitudCotizacionObservacionProveedor, []helpers.CampoOpcional{
		{Campo: "Texto", Columna: "texto", Valor: nil},
		{Campo: "FechaVistoSolicitante", Columna: "fecha_visto_solicitante", Valor: nil},
	})
}

// TestGetSolicitudCotizacionObservacionProveedorOpcionalesConValor: la base de datos devuelve el valor de cada campo opcional.
func TestGetSolicitudCotizacionObservacionProveedorOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploSolicitudCotizacionObservacionProveedor()
	helpers.ProbarGetOpcionales(t, rutaSolicitudCotizacionObservacionProveedor, []helpers.CampoOpcional{
		{Campo: "Texto", Columna: "texto", Valor: helpers.Desreferenciar(ejemplo.Texto)},
		{Campo: "FechaVistoSolicitante", Columna: "fecha_visto_solicitante", Valor: helpers.Desreferenciar(ejemplo.FechaVistoSolicitante)},
	})
}

// TestPutSolicitudCotizacionObservacionProveedorOpcionalesOmitidos: el cuerpo no incluye los campos opcionales.
func TestPutSolicitudCotizacionObservacionProveedorOpcionalesOmitidos(t *testing.T) {
	cuerpo := cuerpoSolicitudCotizacionObservacionProveedor(t)
	delete(cuerpo, "Texto")
	delete(cuerpo, "FechaVistoSolicitante")
	helpers.ProbarPutOpcionales(t, rutaSolicitudCotizacionObservacionProveedor, cuerpo, []helpers.CampoOpcional{
		{Campo: "Texto", Columna: "texto", Valor: nil},
		{Campo: "FechaVistoSolicitante", Columna: "fecha_visto_solicitante", Valor: nil},
	})
}

// TestPutSolicitudCotizacionObservacionProveedorOpcionalesNull: el cuerpo incluye cada campo opcional con null.
func TestPutSolicitudCotizacionObservacionProveedorOpcionalesNull(t *testing.T) {
	cuerpo := cuerpoSolicitudCotizacionObservacionProveedor(t)
	cuerpo["Texto"] = nil
	cuerpo["FechaVistoSolicitante"] = nil
	helpers.ProbarPutOpcionales(t, rutaSolicitudCotizacionObservacionProveedor, cuerpo, []helpers.CampoOpcional{
		{Campo: "Texto", Columna: "texto", Valor: nil},
		{Campo: "FechaVistoSolicitante", Columna: "fecha_visto_solicitante", Valor: nil},
	})
}

// TestPutSolicitudCotizacionObservacionProveedorOpcionalesConValor: el cuerpo incluye cada campo opcional con su valor.
func TestPutSolicitudCotizacionObservacionProveedorOpcionalesConValor(t *testing.T) {
	ejemplo := ejemploSolicitudCotizacionObservacionProveedor()
	helpers.ProbarPutOpcionales(t, rutaSolicitudCotizacionObservacionProveedor, cuerpoSolicitudCotizacionObservacionProveedor(t), []helpers.CampoOpcional{
		{Campo: "Texto", Columna: "texto", Valor: helpers.Desreferenciar(ejemplo.Texto)},
		{Campo: "FechaVistoSolicitante", Columna: "fecha_visto_solicitante", Valor: helpers.Desreferenciar(ejemplo.FechaVistoSolicitante)},
	})
}
