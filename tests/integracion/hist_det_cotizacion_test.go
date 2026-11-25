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

// Pruebas del controlador HistDetCotizacionController (/v1/hist_det_cotizacion).

const (
	rutaHistDetCotizacion  = "/v1/hist_det_cotizacion"
	tablaHistDetCotizacion = "hist_det_cotizacion"
)

// cuerpoPostHistDetCotizacion retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostHistDetCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoHistDetCotizacion(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestHistDetCotizacionPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaHistDetCotizacion, tablaHistDetCotizacion, cuerpoPostHistDetCotizacion(t))
	id2 := helpers.CrearPorPost(t, rutaHistDetCotizacion, tablaHistDetCotizacion, cuerpoPostHistDetCotizacion(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaHistDetCotizacion, id)

	// La llave foránea MstCotizacionId es obligatoria
	sinMstCotizacionId := cuerpoPostHistDetCotizacion(t)
	sinMstCotizacionId["MstCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaHistDetCotizacion+"/", helpers.JSON(t, sinMstCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaHistDetCotizacion)
	t.Log("TestHistDetCotizacionPost Finalizado Correctamente (OK)")
}

func TestHistDetCotizacionGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistDetCotizacion(t)
	dato := helpers.ProbarGetOne(t, rutaHistDetCotizacion, creado.Id)

	// La llave foránea MstCotizacionId viene expandida
	if relacionado, err := models.GetMstCotizacionById(creado.MstCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaHistDetCotizacion)
	t.Log("TestHistDetCotizacionGetOne Finalizado Correctamente (OK)")
}

func TestHistDetCotizacionGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistDetCotizacion(t)
	dato := helpers.ProbarGetAll(t, rutaHistDetCotizacion, tablaHistDetCotizacion, creado.Id)

	// La llave foránea MstCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetMstCotizacionById(creado.MstCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaHistDetCotizacion, fmt.Sprintf("query=MstCotizacionId.Id:%d", creado.MstCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MstCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestHistDetCotizacionGetAll Finalizado Correctamente (OK)")
}

func TestHistDetCotizacionPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistDetCotizacion(t)
	helpers.ProbarPut(t, rutaHistDetCotizacion, tablaHistDetCotizacion, creado.Id, "FichaTecnica", "Prueba-2", "ficha_tecnica")
	t.Log("TestHistDetCotizacionPut Finalizado Correctamente (OK)")
}

func TestHistDetCotizacionDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistDetCotizacion(t)
	helpers.ProbarDelete(t, rutaHistDetCotizacion, tablaHistDetCotizacion, creado.Id)
	t.Log("TestHistDetCotizacionDelete Finalizado Correctamente (OK)")
}

// El campo ValorUnitario es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistDetCotizacionValorUnitarioOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistDetCotizacion(t)
	delete(cuerpo, "ValorUnitario")
	id := helpers.CrearPorPost(t, rutaHistDetCotizacion, tablaHistDetCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistDetCotizacion, id); dato["ValorUnitario"] != nil {
		t.Error("Se esperaba ValorUnitario nulo al omitirlo y se obtuvo", dato["ValorUnitario"])
	}
	if valor := helpers.Fila(t, tablaHistDetCotizacion, id)["valor_unitario"]; valor != nil {
		t.Error("Se esperaba valor_unitario NULL en la base de datos al omitir ValorUnitario y se obtuvo", valor)
	}
	t.Log("TestHistDetCotizacionValorUnitarioOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistDetCotizacionValorUnitarioOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistDetCotizacion(t)
	cuerpo["ValorUnitario"] = nil
	id := helpers.CrearPorPost(t, rutaHistDetCotizacion, tablaHistDetCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistDetCotizacion, id); dato["ValorUnitario"] != nil {
		t.Error("Se esperaba ValorUnitario nulo al enviar null y se obtuvo", dato["ValorUnitario"])
	}
	if valor := helpers.Fila(t, tablaHistDetCotizacion, id)["valor_unitario"]; valor != nil {
		t.Error("Se esperaba valor_unitario NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistDetCotizacionValorUnitarioOpcionalNulo Finalizado Correctamente (OK)")
}
