// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package integracion

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
	helpers.ProbarPut(t, rutaHistMstCotizacion, tablaHistMstCotizacion, creado.Id, "UsuarioModificacionId", 2, "usuario_modificacion_id")
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
	helpers.ProbarPut(t, rutaHistMstCotizacion, tablaHistMstCotizacion, id, "UsuarioModificacionId", 2, "usuario_modificacion_id")
	helpers.ProbarDelete(t, rutaHistMstCotizacion, tablaHistMstCotizacion, id)
	t.Log("TestHistMstCotizacionMstCotizacionIdOpcional Finalizado Correctamente (OK)")
}

// El campo Descripcion es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstCotizacionDescripcionOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstCotizacion(t)
	delete(cuerpo, "Descripcion")
	id := helpers.CrearPorPost(t, rutaHistMstCotizacion, tablaHistMstCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstCotizacion, id); dato["Descripcion"] != nil {
		t.Error("Se esperaba Descripcion nulo al omitirlo y se obtuvo", dato["Descripcion"])
	}
	if valor := helpers.Fila(t, tablaHistMstCotizacion, id)["descripcion"]; valor != nil {
		t.Error("Se esperaba descripcion NULL en la base de datos al omitir Descripcion y se obtuvo", valor)
	}
	t.Log("TestHistMstCotizacionDescripcionOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstCotizacionDescripcionOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstCotizacion(t)
	cuerpo["Descripcion"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstCotizacion, tablaHistMstCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstCotizacion, id); dato["Descripcion"] != nil {
		t.Error("Se esperaba Descripcion nulo al enviar null y se obtuvo", dato["Descripcion"])
	}
	if valor := helpers.Fila(t, tablaHistMstCotizacion, id)["descripcion"]; valor != nil {
		t.Error("Se esperaba descripcion NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstCotizacionDescripcionOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo InformacionEntrega es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstCotizacionInformacionEntregaOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstCotizacion(t)
	delete(cuerpo, "InformacionEntrega")
	id := helpers.CrearPorPost(t, rutaHistMstCotizacion, tablaHistMstCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstCotizacion, id); dato["InformacionEntrega"] != nil {
		t.Error("Se esperaba InformacionEntrega nulo al omitirlo y se obtuvo", dato["InformacionEntrega"])
	}
	if valor := helpers.Fila(t, tablaHistMstCotizacion, id)["informacion_entrega"]; valor != nil {
		t.Error("Se esperaba informacion_entrega NULL en la base de datos al omitir InformacionEntrega y se obtuvo", valor)
	}
	t.Log("TestHistMstCotizacionInformacionEntregaOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstCotizacionInformacionEntregaOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstCotizacion(t)
	cuerpo["InformacionEntrega"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstCotizacion, tablaHistMstCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstCotizacion, id); dato["InformacionEntrega"] != nil {
		t.Error("Se esperaba InformacionEntrega nulo al enviar null y se obtuvo", dato["InformacionEntrega"])
	}
	if valor := helpers.Fila(t, tablaHistMstCotizacion, id)["informacion_entrega"]; valor != nil {
		t.Error("Se esperaba informacion_entrega NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstCotizacionInformacionEntregaOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo Descuentos es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstCotizacionDescuentosOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstCotizacion(t)
	delete(cuerpo, "Descuentos")
	id := helpers.CrearPorPost(t, rutaHistMstCotizacion, tablaHistMstCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstCotizacion, id); dato["Descuentos"] != nil {
		t.Error("Se esperaba Descuentos nulo al omitirlo y se obtuvo", dato["Descuentos"])
	}
	if valor := helpers.Fila(t, tablaHistMstCotizacion, id)["descuentos"]; valor != nil {
		t.Error("Se esperaba descuentos NULL en la base de datos al omitir Descuentos y se obtuvo", valor)
	}
	t.Log("TestHistMstCotizacionDescuentosOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstCotizacionDescuentosOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstCotizacion(t)
	cuerpo["Descuentos"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstCotizacion, tablaHistMstCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstCotizacion, id); dato["Descuentos"] != nil {
		t.Error("Se esperaba Descuentos nulo al enviar null y se obtuvo", dato["Descuentos"])
	}
	if valor := helpers.Fila(t, tablaHistMstCotizacion, id)["descuentos"]; valor != nil {
		t.Error("Se esperaba descuentos NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstCotizacionDescuentosOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo Observaciones es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstCotizacionObservacionesOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstCotizacion(t)
	delete(cuerpo, "Observaciones")
	id := helpers.CrearPorPost(t, rutaHistMstCotizacion, tablaHistMstCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstCotizacion, id); dato["Observaciones"] != nil {
		t.Error("Se esperaba Observaciones nulo al omitirlo y se obtuvo", dato["Observaciones"])
	}
	if valor := helpers.Fila(t, tablaHistMstCotizacion, id)["observaciones"]; valor != nil {
		t.Error("Se esperaba observaciones NULL en la base de datos al omitir Observaciones y se obtuvo", valor)
	}
	t.Log("TestHistMstCotizacionObservacionesOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstCotizacionObservacionesOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstCotizacion(t)
	cuerpo["Observaciones"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstCotizacion, tablaHistMstCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstCotizacion, id); dato["Observaciones"] != nil {
		t.Error("Se esperaba Observaciones nulo al enviar null y se obtuvo", dato["Observaciones"])
	}
	if valor := helpers.Fila(t, tablaHistMstCotizacion, id)["observaciones"]; valor != nil {
		t.Error("Se esperaba observaciones NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstCotizacionObservacionesOpcionalNulo Finalizado Correctamente (OK)")
}
