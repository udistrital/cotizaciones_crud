// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas del controlador SolicitudCotizacionObservacionProveedorController (/v1/solicitud_cotizacion_observacion_proveedor).

const (
	rutaSolicitudCotizacionObservacionProveedor  = "/v1/solicitud_cotizacion_observacion_proveedor"
	tablaSolicitudCotizacionObservacionProveedor = "solicitud_cotizacion_observacion_proveedor"
)

// cuerpoPostSolicitudCotizacionObservacionProveedor retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostSolicitudCotizacionObservacionProveedor(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoSolicitudCotizacionObservacionProveedor(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestSolicitudCotizacionObservacionProveedorPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaSolicitudCotizacionObservacionProveedor, tablaSolicitudCotizacionObservacionProveedor, cuerpoPostSolicitudCotizacionObservacionProveedor(t))
	id2 := helpers.CrearPorPost(t, rutaSolicitudCotizacionObservacionProveedor, tablaSolicitudCotizacionObservacionProveedor, cuerpoPostSolicitudCotizacionObservacionProveedor(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaSolicitudCotizacionObservacionProveedor, id)

	// La llave foránea RelSolicitudCotizacionProveedorId es obligatoria
	sinRelSolicitudCotizacionProveedorId := cuerpoPostSolicitudCotizacionObservacionProveedor(t)
	sinRelSolicitudCotizacionProveedorId["RelSolicitudCotizacionProveedorId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaSolicitudCotizacionObservacionProveedor+"/", helpers.JSON(t, sinRelSolicitudCotizacionProveedorId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaSolicitudCotizacionObservacionProveedor)
	t.Log("TestSolicitudCotizacionObservacionProveedorPost Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionObservacionProveedorGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionObservacionProveedor(t)
	dato := helpers.ProbarGetOne(t, rutaSolicitudCotizacionObservacionProveedor, creado.Id)

	// La llave foránea RelSolicitudCotizacionProveedorId viene expandida
	if relacionado, err := models.GetRelSolicitudCotizacionProveedorById(creado.RelSolicitudCotizacionProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar RelSolicitudCotizacionProveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "RelSolicitudCotizacionProveedorId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaSolicitudCotizacionObservacionProveedor)
	t.Log("TestSolicitudCotizacionObservacionProveedorGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionObservacionProveedorGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionObservacionProveedor(t)
	dato := helpers.ProbarGetAll(t, rutaSolicitudCotizacionObservacionProveedor, tablaSolicitudCotizacionObservacionProveedor, creado.Id)

	// La llave foránea RelSolicitudCotizacionProveedorId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetRelSolicitudCotizacionProveedorById(creado.RelSolicitudCotizacionProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar RelSolicitudCotizacionProveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "RelSolicitudCotizacionProveedorId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaSolicitudCotizacionObservacionProveedor, fmt.Sprintf("query=RelSolicitudCotizacionProveedorId.Id:%d", creado.RelSolicitudCotizacionProveedorId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por RelSolicitudCotizacionProveedorId.Id y se obtuvo", l)
	}
	t.Log("TestSolicitudCotizacionObservacionProveedorGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionObservacionProveedorPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionObservacionProveedor(t)
	helpers.ProbarPut(t, rutaSolicitudCotizacionObservacionProveedor, tablaSolicitudCotizacionObservacionProveedor, creado.Id, "Texto", "Prueba-2", "texto")
	t.Log("TestSolicitudCotizacionObservacionProveedorPut Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionObservacionProveedorDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionObservacionProveedor(t)
	helpers.ProbarDelete(t, rutaSolicitudCotizacionObservacionProveedor, tablaSolicitudCotizacionObservacionProveedor, creado.Id)
	t.Log("TestSolicitudCotizacionObservacionProveedorDelete Finalizado Correctamente (OK)")
}
