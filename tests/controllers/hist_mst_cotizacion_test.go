// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controllers

import (
	"fmt"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas del controlador HistMstCotizacionController (/v1/hist_mst_cotizacion).

const (
	rutaHistMstCotizacion  = "/v1/hist_mst_cotizacion"
	tablaHistMstCotizacion = "hist_mst_cotizacion"
)

// cuerpoPostHistMstCotizacion retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostHistMstCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoHistMstCotizacion(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestHistMstCotizacionPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaHistMstCotizacion, tablaHistMstCotizacion, cuerpoPostHistMstCotizacion(t))
	id2 := helpers.CrearPorPost(t, rutaHistMstCotizacion, tablaHistMstCotizacion, cuerpoPostHistMstCotizacion(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaHistMstCotizacion, id)

	helpers.ProbarPostInvalido(t, rutaHistMstCotizacion)
	t.Log("TestHistMstCotizacionPost Finalizado Correctamente (OK)")
}

func TestHistMstCotizacionGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstCotizacion(t)
	dato := helpers.ProbarGetOne(t, rutaHistMstCotizacion, creado.Id)

	// La llave foránea MstCotizacionId viene expandida
	if relacionado, err := models.GetMstCotizacionById(creado.MstCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaHistMstCotizacion)
	t.Log("TestHistMstCotizacionGetOne Finalizado Correctamente (OK)")
}

func TestHistMstCotizacionGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstCotizacion(t)
	dato := helpers.ProbarGetAll(t, rutaHistMstCotizacion, tablaHistMstCotizacion, creado.Id)

	// La llave foránea MstCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetMstCotizacionById(creado.MstCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaHistMstCotizacion, fmt.Sprintf("query=MstCotizacionId.Id:%d", creado.MstCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MstCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestHistMstCotizacionGetAll Finalizado Correctamente (OK)")
}

func TestHistMstCotizacionPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstCotizacion(t)
	helpers.ProbarPut(t, rutaHistMstCotizacion, tablaHistMstCotizacion, creado.Id, "Descripcion", "Prueba-2", "descripcion")
	t.Log("TestHistMstCotizacionPut Finalizado Correctamente (OK)")
}

func TestHistMstCotizacionDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstCotizacion(t)
	helpers.ProbarDelete(t, rutaHistMstCotizacion, tablaHistMstCotizacion, creado.Id)
	t.Log("TestHistMstCotizacionDelete Finalizado Correctamente (OK)")
}

// La llave foránea MstCotizacionId es opcional (null): el registro sin ella se
// puede crear, consultar, actualizar y borrar.
func TestHistMstCotizacionMstCotizacionIdOpcional(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstCotizacion(t)
	cuerpo["MstCotizacionId"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstCotizacion, tablaHistMstCotizacion, cuerpo)

	if dato := helpers.ProbarGetOne(t, rutaHistMstCotizacion, id); dato["MstCotizacionId"] != nil {
		t.Error("Se esperaba MstCotizacionId nulo y se obtuvo", dato["MstCotizacionId"])
	}
	if l := helpers.ConsultarGetAll(t, rutaHistMstCotizacion, fmt.Sprintf("query=Id:%d", id)); !helpers.ContieneId(l, id) {
		t.Error("Se esperaba el registro sin MstCotizacionId en GetAll y se obtuvo", l)
	}
	helpers.ProbarPut(t, rutaHistMstCotizacion, tablaHistMstCotizacion, id, "Descripcion", "Prueba-2", "descripcion")
	helpers.ProbarDelete(t, rutaHistMstCotizacion, tablaHistMstCotizacion, id)
	t.Log("TestHistMstCotizacionMstCotizacionIdOpcional Finalizado Correctamente (OK)")
}
