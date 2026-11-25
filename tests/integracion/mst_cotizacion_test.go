// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package integracion

import (
	"fmt"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Pruebas del controlador MstCotizacionController (/v1/mst_cotizacion).

const (
	rutaMstCotizacion  = "/v1/mst_cotizacion"
	tablaMstCotizacion = "mst_cotizacion"
)

// cuerpoPostMstCotizacion retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostMstCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoMstCotizacion(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestMstCotizacionPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaMstCotizacion, tablaMstCotizacion, cuerpoPostMstCotizacion(t))
	id2 := helpers.CrearPorPost(t, rutaMstCotizacion, tablaMstCotizacion, cuerpoPostMstCotizacion(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaMstCotizacion, id)

	helpers.ProbarPostInvalido(t, rutaMstCotizacion)
	t.Log("TestMstCotizacionPost Finalizado Correctamente (OK)")
}

func TestMstCotizacionGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearMstCotizacion(t)
	dato := helpers.ProbarGetOne(t, rutaMstCotizacion, creado.Id)

	// La llave foránea RelSolicitudCotizacionProveedorId viene expandida
	if relacionado, err := models.GetRelSolicitudCotizacionProveedorById(creado.RelSolicitudCotizacionProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar RelSolicitudCotizacionProveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "RelSolicitudCotizacionProveedorId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaMstCotizacion)
	t.Log("TestMstCotizacionGetOne Finalizado Correctamente (OK)")
}

func TestMstCotizacionGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearMstCotizacion(t)
	dato := helpers.ProbarGetAll(t, rutaMstCotizacion, tablaMstCotizacion, creado.Id)

	// La llave foránea RelSolicitudCotizacionProveedorId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetRelSolicitudCotizacionProveedorById(creado.RelSolicitudCotizacionProveedorId.Id); err != nil {
		t.Fatal("No se pudo consultar RelSolicitudCotizacionProveedor:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "RelSolicitudCotizacionProveedorId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaMstCotizacion, fmt.Sprintf("query=RelSolicitudCotizacionProveedorId.Id:%d", creado.RelSolicitudCotizacionProveedorId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por RelSolicitudCotizacionProveedorId.Id y se obtuvo", l)
	}
	t.Log("TestMstCotizacionGetAll Finalizado Correctamente (OK)")
}

func TestMstCotizacionPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearMstCotizacion(t)
	helpers.ProbarPut(t, rutaMstCotizacion, tablaMstCotizacion, creado.Id, "UsuarioModificacionId", 2, "usuario_modificacion_id")
	t.Log("TestMstCotizacionPut Finalizado Correctamente (OK)")
}

func TestMstCotizacionDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearMstCotizacion(t)
	helpers.ProbarDelete(t, rutaMstCotizacion, tablaMstCotizacion, creado.Id)
	t.Log("TestMstCotizacionDelete Finalizado Correctamente (OK)")
}

// La llave foránea RelSolicitudCotizacionProveedorId es opcional (null): el registro sin ella se
// puede crear, consultar, actualizar y borrar.
func TestMstCotizacionRelSolicitudCotizacionProveedorIdOpcional(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostMstCotizacion(t)
	cuerpo["RelSolicitudCotizacionProveedorId"] = nil
	id := helpers.CrearPorPost(t, rutaMstCotizacion, tablaMstCotizacion, cuerpo)

	if dato := helpers.ProbarGetOne(t, rutaMstCotizacion, id); dato["RelSolicitudCotizacionProveedorId"] != nil {
		t.Error("Se esperaba RelSolicitudCotizacionProveedorId nulo y se obtuvo", dato["RelSolicitudCotizacionProveedorId"])
	}
	if l := helpers.ConsultarGetAll(t, rutaMstCotizacion, fmt.Sprintf("query=Id:%d", id)); !helpers.ContieneId(l, id) {
		t.Error("Se esperaba el registro sin RelSolicitudCotizacionProveedorId en GetAll y se obtuvo", l)
	}
	helpers.ProbarPut(t, rutaMstCotizacion, tablaMstCotizacion, id, "UsuarioModificacionId", 2, "usuario_modificacion_id")
	helpers.ProbarDelete(t, rutaMstCotizacion, tablaMstCotizacion, id)
	t.Log("TestMstCotizacionRelSolicitudCotizacionProveedorIdOpcional Finalizado Correctamente (OK)")
}
