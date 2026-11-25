// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package integracion

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas del controlador SolicitudCotizacionObservacionProveedorRespuestaController (/v1/solicitud_cotizacion_observacion_proveedor_respuesta).

const (
	rutaSolicitudCotizacionObservacionProveedorRespuesta  = "/v1/solicitud_cotizacion_observacion_proveedor_respuesta"
	tablaSolicitudCotizacionObservacionProveedorRespuesta = "solicitud_cotizacion_observacion_proveedor_respuesta"
)

// cuerpoPostSolicitudCotizacionObservacionProveedorRespuesta retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostSolicitudCotizacionObservacionProveedorRespuesta(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoSolicitudCotizacionObservacionProveedorRespuesta(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestSolicitudCotizacionObservacionProveedorRespuestaPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaSolicitudCotizacionObservacionProveedorRespuesta, tablaSolicitudCotizacionObservacionProveedorRespuesta, cuerpoPostSolicitudCotizacionObservacionProveedorRespuesta(t))
	id2 := helpers.CrearPorPost(t, rutaSolicitudCotizacionObservacionProveedorRespuesta, tablaSolicitudCotizacionObservacionProveedorRespuesta, cuerpoPostSolicitudCotizacionObservacionProveedorRespuesta(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaSolicitudCotizacionObservacionProveedorRespuesta, id)

	// La llave foránea ObservacionId es obligatoria
	sinObservacionId := cuerpoPostSolicitudCotizacionObservacionProveedorRespuesta(t)
	sinObservacionId["ObservacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaSolicitudCotizacionObservacionProveedorRespuesta+"/", helpers.JSON(t, sinObservacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaSolicitudCotizacionObservacionProveedorRespuesta)
	t.Log("TestSolicitudCotizacionObservacionProveedorRespuestaPost Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionObservacionProveedorRespuestaGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionObservacionProveedorRespuesta(t)
	dato := helpers.ProbarGetOne(t, rutaSolicitudCotizacionObservacionProveedorRespuesta, creado.Id)

	// La llave foránea ObservacionId viene expandida
	if relacionado, err := models.GetSolicitudCotizacionObservacionProveedorById(creado.ObservacionId.Id); err != nil {
		t.Fatal("No se pudo consultar SolicitudCotizacionObservacionProveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ObservacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaSolicitudCotizacionObservacionProveedorRespuesta)
	t.Log("TestSolicitudCotizacionObservacionProveedorRespuestaGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionObservacionProveedorRespuestaGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionObservacionProveedorRespuesta(t)
	dato := helpers.ProbarGetAll(t, rutaSolicitudCotizacionObservacionProveedorRespuesta, tablaSolicitudCotizacionObservacionProveedorRespuesta, creado.Id)

	// La llave foránea ObservacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetSolicitudCotizacionObservacionProveedorById(creado.ObservacionId.Id); err != nil {
		t.Fatal("No se pudo consultar SolicitudCotizacionObservacionProveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "ObservacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaSolicitudCotizacionObservacionProveedorRespuesta, fmt.Sprintf("query=ObservacionId.Id:%d", creado.ObservacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por ObservacionId.Id y se obtuvo", l)
	}
	t.Log("TestSolicitudCotizacionObservacionProveedorRespuestaGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionObservacionProveedorRespuestaPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionObservacionProveedorRespuesta(t)
	helpers.ProbarPut(t, rutaSolicitudCotizacionObservacionProveedorRespuesta, tablaSolicitudCotizacionObservacionProveedorRespuesta, creado.Id, "Texto", "Prueba-2", "texto")
	t.Log("TestSolicitudCotizacionObservacionProveedorRespuestaPut Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionObservacionProveedorRespuestaDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionObservacionProveedorRespuesta(t)
	helpers.ProbarDelete(t, rutaSolicitudCotizacionObservacionProveedorRespuesta, tablaSolicitudCotizacionObservacionProveedorRespuesta, creado.Id)
	t.Log("TestSolicitudCotizacionObservacionProveedorRespuestaDelete Finalizado Correctamente (OK)")
}
