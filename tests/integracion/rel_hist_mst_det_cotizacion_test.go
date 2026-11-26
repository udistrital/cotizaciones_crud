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

// Pruebas del controlador RelHistMstDetCotizacionController (/v1/rel_hist_mst_det_cotizacion).

const (
	rutaRelHistMstDetCotizacion  = "/v1/rel_hist_mst_det_cotizacion"
	tablaRelHistMstDetCotizacion = "rel_hist_mst_det_cotizacion"
)

// cuerpoPostRelHistMstDetCotizacion retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostRelHistMstDetCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoRelHistMstDetCotizacion(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestRelHistMstDetCotizacionPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaRelHistMstDetCotizacion, tablaRelHistMstDetCotizacion, cuerpoPostRelHistMstDetCotizacion(t))
	id2 := helpers.CrearPorPost(t, rutaRelHistMstDetCotizacion, tablaRelHistMstDetCotizacion, cuerpoPostRelHistMstDetCotizacion(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaRelHistMstDetCotizacion, id)

	// La llave foránea MaestroId es obligatoria
	sinMaestroId := cuerpoPostRelHistMstDetCotizacion(t)
	sinMaestroId["MaestroId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaRelHistMstDetCotizacion+"/", helpers.JSON(t, sinMaestroId)), http.StatusBadRequest)

	// La llave foránea DetalleId es obligatoria
	sinDetalleId := cuerpoPostRelHistMstDetCotizacion(t)
	sinDetalleId["DetalleId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaRelHistMstDetCotizacion+"/", helpers.JSON(t, sinDetalleId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaRelHistMstDetCotizacion)
	t.Log("TestRelHistMstDetCotizacionPost Finalizado Correctamente (OK)")
}

func TestRelHistMstDetCotizacionGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelHistMstDetCotizacion(t)
	dato := helpers.ProbarGetOne(t, rutaRelHistMstDetCotizacion, creado.Id)

	// La llave foránea MaestroId viene expandida
	if relacionado, err := models.GetHistMstCotizacionById(creado.MaestroId.Id); err != nil {
		t.Fatal("No se pudo consultar HistMstCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MaestroId", relacionado)
	}

	// La llave foránea DetalleId viene expandida
	if relacionado, err := models.GetHistDetCotizacionById(creado.DetalleId.Id); err != nil {
		t.Fatal("No se pudo consultar HistDetCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "DetalleId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaRelHistMstDetCotizacion)
	t.Log("TestRelHistMstDetCotizacionGetOne Finalizado Correctamente (OK)")
}

func TestRelHistMstDetCotizacionGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelHistMstDetCotizacion(t)
	dato := helpers.ProbarGetAll(t, rutaRelHistMstDetCotizacion, tablaRelHistMstDetCotizacion, creado.Id)

	// La llave foránea MaestroId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetHistMstCotizacionById(creado.MaestroId.Id); err != nil {
		t.Fatal("No se pudo consultar HistMstCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MaestroId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaRelHistMstDetCotizacion, fmt.Sprintf("query=MaestroId.Id:%d", creado.MaestroId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MaestroId.Id y se obtuvo", l)
	}

	// La llave foránea DetalleId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetHistDetCotizacionById(creado.DetalleId.Id); err != nil {
		t.Fatal("No se pudo consultar HistDetCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "DetalleId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaRelHistMstDetCotizacion, fmt.Sprintf("query=DetalleId.Id:%d", creado.DetalleId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por DetalleId.Id y se obtuvo", l)
	}
	t.Log("TestRelHistMstDetCotizacionGetAll Finalizado Correctamente (OK)")
}

func TestRelHistMstDetCotizacionPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelHistMstDetCotizacion(t)
	helpers.ProbarPut(t, rutaRelHistMstDetCotizacion, tablaRelHistMstDetCotizacion, creado.Id, "UsuarioModificacionId", 2, "usuario_modificacion_id")
	t.Log("TestRelHistMstDetCotizacionPut Finalizado Correctamente (OK)")
}

func TestRelHistMstDetCotizacionDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearRelHistMstDetCotizacion(t)
	helpers.ProbarDelete(t, rutaRelHistMstDetCotizacion, tablaRelHistMstDetCotizacion, creado.Id)
	t.Log("TestRelHistMstDetCotizacionDelete Finalizado Correctamente (OK)")
}

// La llave foránea MaestroId debe referenciar un registro existente: POST y PUT
// responden 400 y no crean ni modifican registros.
func TestRelHistMstDetCotizacionMaestroIdInexistente(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	helpers.ProbarPostFKInexistente(t, rutaRelHistMstDetCotizacion, tablaRelHistMstDetCotizacion, cuerpoPostRelHistMstDetCotizacion(t), "MaestroId")
	creado := helpers.CrearRelHistMstDetCotizacion(t)
	helpers.ProbarPutFKInexistente(t, rutaRelHistMstDetCotizacion, tablaRelHistMstDetCotizacion, creado.Id, "MaestroId")
	t.Log("TestRelHistMstDetCotizacionMaestroIdInexistente Finalizado Correctamente (OK)")
}

// La llave foránea DetalleId debe referenciar un registro existente: POST y PUT
// responden 400 y no crean ni modifican registros.
func TestRelHistMstDetCotizacionDetalleIdInexistente(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	helpers.ProbarPostFKInexistente(t, rutaRelHistMstDetCotizacion, tablaRelHistMstDetCotizacion, cuerpoPostRelHistMstDetCotizacion(t), "DetalleId")
	creado := helpers.CrearRelHistMstDetCotizacion(t)
	helpers.ProbarPutFKInexistente(t, rutaRelHistMstDetCotizacion, tablaRelHistMstDetCotizacion, creado.Id, "DetalleId")
	t.Log("TestRelHistMstDetCotizacionDetalleIdInexistente Finalizado Correctamente (OK)")
}
