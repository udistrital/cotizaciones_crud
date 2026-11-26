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

// Pruebas del controlador HistMstSolicitudCotizacionController (/v1/hist_mst_solicitud_cotizacion).

const (
	rutaHistMstSolicitudCotizacion  = "/v1/hist_mst_solicitud_cotizacion"
	tablaHistMstSolicitudCotizacion = "hist_mst_solicitud_cotizacion"
)

// cuerpoPostHistMstSolicitudCotizacion retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostHistMstSolicitudCotizacion(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoHistMstSolicitudCotizacion(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestHistMstSolicitudCotizacionPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpoPostHistMstSolicitudCotizacion(t))
	id2 := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpoPostHistMstSolicitudCotizacion(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaHistMstSolicitudCotizacion, id)

	// La llave foránea MstSolicitudCotizacionId es obligatoria
	sinMstSolicitudCotizacionId := cuerpoPostHistMstSolicitudCotizacion(t)
	sinMstSolicitudCotizacionId["MstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaHistMstSolicitudCotizacion+"/", helpers.JSON(t, sinMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaHistMstSolicitudCotizacion)
	t.Log("TestHistMstSolicitudCotizacionPost Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacion(t)
	dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaHistMstSolicitudCotizacion)
	t.Log("TestHistMstSolicitudCotizacionGetOne Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacion(t)
	dato := helpers.ProbarGetAll(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, creado.Id)

	// La llave foránea MstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetMstSolicitudCotizacionById(creado.MstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar MstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "MstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaHistMstSolicitudCotizacion, fmt.Sprintf("query=MstSolicitudCotizacionId.Id:%d", creado.MstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por MstSolicitudCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestHistMstSolicitudCotizacionGetAll Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacion(t)
	helpers.ProbarPut(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, creado.Id, "Titulo", "Prueba-2", "titulo")
	t.Log("TestHistMstSolicitudCotizacionPut Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearHistMstSolicitudCotizacion(t)
	helpers.ProbarDelete(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, creado.Id)
	t.Log("TestHistMstSolicitudCotizacionDelete Finalizado Correctamente (OK)")
}

// La llave foránea MstSolicitudCotizacionId debe referenciar un registro existente: POST y PUT
// responden 400 y no crean ni modifican registros.
func TestHistMstSolicitudCotizacionMstSolicitudCotizacionIdInexistente(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	helpers.ProbarPostFKInexistente(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpoPostHistMstSolicitudCotizacion(t), "MstSolicitudCotizacionId")
	creado := helpers.CrearHistMstSolicitudCotizacion(t)
	helpers.ProbarPutFKInexistente(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, creado.Id, "MstSolicitudCotizacionId")
	t.Log("TestHistMstSolicitudCotizacionMstSolicitudCotizacionIdInexistente Finalizado Correctamente (OK)")
}

// El campo NumeroSolicitud es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstSolicitudCotizacionNumeroSolicitudOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	delete(cuerpo, "NumeroSolicitud")
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["NumeroSolicitud"] != nil {
		t.Error("Se esperaba NumeroSolicitud nulo al omitirlo y se obtuvo", dato["NumeroSolicitud"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["numero_solicitud"]; valor != nil {
		t.Error("Se esperaba numero_solicitud NULL en la base de datos al omitir NumeroSolicitud y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionNumeroSolicitudOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionNumeroSolicitudOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	cuerpo["NumeroSolicitud"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["NumeroSolicitud"] != nil {
		t.Error("Se esperaba NumeroSolicitud nulo al enviar null y se obtuvo", dato["NumeroSolicitud"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["numero_solicitud"]; valor != nil {
		t.Error("Se esperaba numero_solicitud NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionNumeroSolicitudOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo Vigencia es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstSolicitudCotizacionVigenciaOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	delete(cuerpo, "Vigencia")
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["Vigencia"] != nil {
		t.Error("Se esperaba Vigencia nulo al omitirlo y se obtuvo", dato["Vigencia"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["vigencia"]; valor != nil {
		t.Error("Se esperaba vigencia NULL en la base de datos al omitir Vigencia y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionVigenciaOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionVigenciaOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	cuerpo["Vigencia"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["Vigencia"] != nil {
		t.Error("Se esperaba Vigencia nulo al enviar null y se obtuvo", dato["Vigencia"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["vigencia"]; valor != nil {
		t.Error("Se esperaba vigencia NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionVigenciaOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo Consecutivo es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstSolicitudCotizacionConsecutivoOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	delete(cuerpo, "Consecutivo")
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["Consecutivo"] != nil {
		t.Error("Se esperaba Consecutivo nulo al omitirlo y se obtuvo", dato["Consecutivo"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["consecutivo"]; valor != nil {
		t.Error("Se esperaba consecutivo NULL en la base de datos al omitir Consecutivo y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionConsecutivoOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionConsecutivoOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	cuerpo["Consecutivo"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["Consecutivo"] != nil {
		t.Error("Se esperaba Consecutivo nulo al enviar null y se obtuvo", dato["Consecutivo"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["consecutivo"]; valor != nil {
		t.Error("Se esperaba consecutivo NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionConsecutivoOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo PlanAccionId es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstSolicitudCotizacionPlanAccionIdOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	delete(cuerpo, "PlanAccionId")
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["PlanAccionId"] != nil {
		t.Error("Se esperaba PlanAccionId nulo al omitirlo y se obtuvo", dato["PlanAccionId"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["plan_accion_id"]; valor != nil {
		t.Error("Se esperaba plan_accion_id NULL en la base de datos al omitir PlanAccionId y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionPlanAccionIdOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionPlanAccionIdOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	cuerpo["PlanAccionId"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["PlanAccionId"] != nil {
		t.Error("Se esperaba PlanAccionId nulo al enviar null y se obtuvo", dato["PlanAccionId"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["plan_accion_id"]; valor != nil {
		t.Error("Se esperaba plan_accion_id NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionPlanAccionIdOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo MetodoSeleccionId es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstSolicitudCotizacionMetodoSeleccionIdOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	delete(cuerpo, "MetodoSeleccionId")
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["MetodoSeleccionId"] != nil {
		t.Error("Se esperaba MetodoSeleccionId nulo al omitirlo y se obtuvo", dato["MetodoSeleccionId"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["metodo_seleccion_id"]; valor != nil {
		t.Error("Se esperaba metodo_seleccion_id NULL en la base de datos al omitir MetodoSeleccionId y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionMetodoSeleccionIdOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionMetodoSeleccionIdOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	cuerpo["MetodoSeleccionId"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["MetodoSeleccionId"] != nil {
		t.Error("Se esperaba MetodoSeleccionId nulo al enviar null y se obtuvo", dato["MetodoSeleccionId"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["metodo_seleccion_id"]; valor != nil {
		t.Error("Se esperaba metodo_seleccion_id NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionMetodoSeleccionIdOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo NumeroNecesidad es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstSolicitudCotizacionNumeroNecesidadOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	delete(cuerpo, "NumeroNecesidad")
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["NumeroNecesidad"] != nil {
		t.Error("Se esperaba NumeroNecesidad nulo al omitirlo y se obtuvo", dato["NumeroNecesidad"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["numero_necesidad"]; valor != nil {
		t.Error("Se esperaba numero_necesidad NULL en la base de datos al omitir NumeroNecesidad y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionNumeroNecesidadOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionNumeroNecesidadOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	cuerpo["NumeroNecesidad"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["NumeroNecesidad"] != nil {
		t.Error("Se esperaba NumeroNecesidad nulo al enviar null y se obtuvo", dato["NumeroNecesidad"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["numero_necesidad"]; valor != nil {
		t.Error("Se esperaba numero_necesidad NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionNumeroNecesidadOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo TipoContratoId es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstSolicitudCotizacionTipoContratoIdOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	delete(cuerpo, "TipoContratoId")
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["TipoContratoId"] != nil {
		t.Error("Se esperaba TipoContratoId nulo al omitirlo y se obtuvo", dato["TipoContratoId"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["tipo_contrato_id"]; valor != nil {
		t.Error("Se esperaba tipo_contrato_id NULL en la base de datos al omitir TipoContratoId y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionTipoContratoIdOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionTipoContratoIdOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	cuerpo["TipoContratoId"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["TipoContratoId"] != nil {
		t.Error("Se esperaba TipoContratoId nulo al enviar null y se obtuvo", dato["TipoContratoId"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["tipo_contrato_id"]; valor != nil {
		t.Error("Se esperaba tipo_contrato_id NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionTipoContratoIdOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo CriterioSeleccion es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstSolicitudCotizacionCriterioSeleccionOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	delete(cuerpo, "CriterioSeleccion")
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["CriterioSeleccion"] != nil {
		t.Error("Se esperaba CriterioSeleccion nulo al omitirlo y se obtuvo", dato["CriterioSeleccion"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["criterio_seleccion"]; valor != nil {
		t.Error("Se esperaba criterio_seleccion NULL en la base de datos al omitir CriterioSeleccion y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionCriterioSeleccionOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionCriterioSeleccionOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	cuerpo["CriterioSeleccion"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["CriterioSeleccion"] != nil {
		t.Error("Se esperaba CriterioSeleccion nulo al enviar null y se obtuvo", dato["CriterioSeleccion"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["criterio_seleccion"]; valor != nil {
		t.Error("Se esperaba criterio_seleccion NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionCriterioSeleccionOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo PlazoEjecucion es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestHistMstSolicitudCotizacionPlazoEjecucionOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	delete(cuerpo, "PlazoEjecucion")
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["PlazoEjecucion"] != nil {
		t.Error("Se esperaba PlazoEjecucion nulo al omitirlo y se obtuvo", dato["PlazoEjecucion"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["plazo_ejecucion"]; valor != nil {
		t.Error("Se esperaba plazo_ejecucion NULL en la base de datos al omitir PlazoEjecucion y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionPlazoEjecucionOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestHistMstSolicitudCotizacionPlazoEjecucionOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostHistMstSolicitudCotizacion(t)
	cuerpo["PlazoEjecucion"] = nil
	id := helpers.CrearPorPost(t, rutaHistMstSolicitudCotizacion, tablaHistMstSolicitudCotizacion, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaHistMstSolicitudCotizacion, id); dato["PlazoEjecucion"] != nil {
		t.Error("Se esperaba PlazoEjecucion nulo al enviar null y se obtuvo", dato["PlazoEjecucion"])
	}
	if valor := helpers.Fila(t, tablaHistMstSolicitudCotizacion, id)["plazo_ejecucion"]; valor != nil {
		t.Error("Se esperaba plazo_ejecucion NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestHistMstSolicitudCotizacionPlazoEjecucionOpcionalNulo Finalizado Correctamente (OK)")
}
