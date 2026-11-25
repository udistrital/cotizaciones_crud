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

// Pruebas del controlador SolicitudCotizacionEspecificacionTecnicaController (/v1/solicitud_cotizacion_especificacion_tecnica).

const (
	rutaSolicitudCotizacionEspecificacionTecnica  = "/v1/solicitud_cotizacion_especificacion_tecnica"
	tablaSolicitudCotizacionEspecificacionTecnica = "solicitud_cotizacion_especificacion_tecnica"
)

// cuerpoPostSolicitudCotizacionEspecificacionTecnica retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostSolicitudCotizacionEspecificacionTecnica(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoSolicitudCotizacionEspecificacionTecnica(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestSolicitudCotizacionEspecificacionTecnicaPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaSolicitudCotizacionEspecificacionTecnica, tablaSolicitudCotizacionEspecificacionTecnica, cuerpoPostSolicitudCotizacionEspecificacionTecnica(t))
	id2 := helpers.CrearPorPost(t, rutaSolicitudCotizacionEspecificacionTecnica, tablaSolicitudCotizacionEspecificacionTecnica, cuerpoPostSolicitudCotizacionEspecificacionTecnica(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaSolicitudCotizacionEspecificacionTecnica, id)

	// La llave foránea MstSolicitudCotizacionId es obligatoria
	sinMstSolicitudCotizacionId := cuerpoPostSolicitudCotizacionEspecificacionTecnica(t)
	sinMstSolicitudCotizacionId["MstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaSolicitudCotizacionEspecificacionTecnica+"/", helpers.JSON(t, sinMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaSolicitudCotizacionEspecificacionTecnica)
	t.Log("TestSolicitudCotizacionEspecificacionTecnicaPost Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionEspecificacionTecnicaGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionEspecificacionTecnica(t)
	dato := helpers.ProbarGetOne(t, rutaSolicitudCotizacionEspecificacionTecnica, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaSolicitudCotizacionEspecificacionTecnica)
	t.Log("TestSolicitudCotizacionEspecificacionTecnicaGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionEspecificacionTecnicaGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionEspecificacionTecnica(t)
	dato := helpers.ProbarGetAll(t, rutaSolicitudCotizacionEspecificacionTecnica, tablaSolicitudCotizacionEspecificacionTecnica, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaSolicitudCotizacionEspecificacionTecnica, fmt.Sprintf("query=MstSolicitudCotizacionId.Id:%d", creado.MstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MstSolicitudCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestSolicitudCotizacionEspecificacionTecnicaGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionEspecificacionTecnicaPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionEspecificacionTecnica(t)
	helpers.ProbarPut(t, rutaSolicitudCotizacionEspecificacionTecnica, tablaSolicitudCotizacionEspecificacionTecnica, creado.Id, "UsuarioModificacionId", 2, "usuario_modificacion_id")
	t.Log("TestSolicitudCotizacionEspecificacionTecnicaPut Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionEspecificacionTecnicaDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearSolicitudCotizacionEspecificacionTecnica(t)
	helpers.ProbarDelete(t, rutaSolicitudCotizacionEspecificacionTecnica, tablaSolicitudCotizacionEspecificacionTecnica, creado.Id)
	t.Log("TestSolicitudCotizacionEspecificacionTecnicaDelete Finalizado Correctamente (OK)")
}
