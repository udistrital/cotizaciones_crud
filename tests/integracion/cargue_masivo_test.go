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

// Pruebas del controlador CargueMasivoController (/v1/cargue_masivo).

const (
	rutaCargueMasivo  = "/v1/cargue_masivo"
	tablaCargueMasivo = "cargue_masivo"
)

// cuerpoPostCargueMasivo retorna un cuerpo válido para POST con valores que el
// servicio debe ignorar: Activo=false y fechas enviadas por el cliente y un Id (lo asigna la base de datos).
func cuerpoPostCargueMasivo(t *testing.T) map[string]interface{} {
	t.Helper()
	cuerpo := helpers.AMapa(t, helpers.NuevoCargueMasivo(t))
	cuerpo["Activo"] = false
	cuerpo["FechaCreacion"] = helpers.FechaCliente
	cuerpo["FechaModificacion"] = helpers.FechaCliente
	cuerpo["Id"] = helpers.IdInexistente
	return cuerpo
}

func TestCargueMasivoPost(t *testing.T) {
	helpers.RequerirBaseDatos(t)

	// El Id lo asigna la base de datos (autoincremental) y se ignora el enviado
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpoPostCargueMasivo(t))
	id2 := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpoPostCargueMasivo(t))
	if id <= 0 || id == helpers.IdInexistente || id2 <= id {
		t.Errorf("Se esperaban Id autoincrementales asignados por la base de datos y se obtuvo %d y %d", id, id2)
	}

	// Activo=true y fechas asignadas por el servicio
	helpers.VerificarCamposAdministrados(t, tablaCargueMasivo, id)

	// La llave foránea HistMstSolicitudCotizacionId es obligatoria
	sinHistMstSolicitudCotizacionId := cuerpoPostCargueMasivo(t)
	sinHistMstSolicitudCotizacionId["HistMstSolicitudCotizacionId"] = nil
	helpers.VerificarError(t, helpers.Peticion(http.MethodPost, rutaCargueMasivo+"/", helpers.JSON(t, sinHistMstSolicitudCotizacionId)), http.StatusBadRequest)

	helpers.ProbarPostInvalido(t, rutaCargueMasivo)
	t.Log("TestCargueMasivoPost Finalizado Correctamente (OK)")
}

func TestCargueMasivoGetOne(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCargueMasivo(t)
	dato := helpers.ProbarGetOne(t, rutaCargueMasivo, creado.Id)

	// La llave foránea HistMstSolicitudCotizacionId viene expandida
	if relacionado, err := models.GetHistMstSolicitudCotizacionById(creado.HistMstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "HistMstSolicitudCotizacionId", relacionado)
	}

	helpers.ProbarGetOneInexistente(t, rutaCargueMasivo)
	t.Log("TestCargueMasivoGetOne Finalizado Correctamente (OK)")
}

func TestCargueMasivoGetAll(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCargueMasivo(t)
	dato := helpers.ProbarGetAll(t, rutaCargueMasivo, tablaCargueMasivo, creado.Id)

	// La llave foránea HistMstSolicitudCotizacionId viene expandida y se puede filtrar por ella
	if relacionado, err := models.GetHistMstSolicitudCotizacionById(creado.HistMstSolicitudCotizacionId.Id); err != nil {
		t.Fatal("No se pudo consultar HistMstSolicitudCotizacion:", err)
	} else {
		helpers.VerificarFKExpandida(t, dato, "HistMstSolicitudCotizacionId", relacionado)
	}
	if l := helpers.ConsultarGetAll(t, rutaCargueMasivo, fmt.Sprintf("query=HistMstSolicitudCotizacionId.Id:%d", creado.HistMstSolicitudCotizacionId.Id)); !helpers.ContieneId(l, creado.Id) {
		t.Error("Se esperaba el registro al filtrar por HistMstSolicitudCotizacionId.Id y se obtuvo", l)
	}
	t.Log("TestCargueMasivoGetAll Finalizado Correctamente (OK)")
}

func TestCargueMasivoPut(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCargueMasivo(t)
	helpers.ProbarPut(t, rutaCargueMasivo, tablaCargueMasivo, creado.Id, "CargadoPorTerceroId", 2, "cargado_por_tercero_id")
	t.Log("TestCargueMasivoPut Finalizado Correctamente (OK)")
}

func TestCargueMasivoDelete(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	creado := helpers.CrearCargueMasivo(t)
	helpers.ProbarDelete(t, rutaCargueMasivo, tablaCargueMasivo, creado.Id)
	t.Log("TestCargueMasivoDelete Finalizado Correctamente (OK)")
}

// La llave foránea HistMstSolicitudCotizacionId debe referenciar un registro existente: POST y PUT
// responden 400 y no crean ni modifican registros.
func TestCargueMasivoHistMstSolicitudCotizacionIdInexistente(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	helpers.ProbarPostFKInexistente(t, rutaCargueMasivo, tablaCargueMasivo, cuerpoPostCargueMasivo(t), "HistMstSolicitudCotizacionId")
	creado := helpers.CrearCargueMasivo(t)
	helpers.ProbarPutFKInexistente(t, rutaCargueMasivo, tablaCargueMasivo, creado.Id, "HistMstSolicitudCotizacionId")
	t.Log("TestCargueMasivoHistMstSolicitudCotizacionIdInexistente Finalizado Correctamente (OK)")
}

// El campo EnlacePlantilla es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestCargueMasivoEnlacePlantillaOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostCargueMasivo(t)
	delete(cuerpo, "EnlacePlantilla")
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaCargueMasivo, id); dato["EnlacePlantilla"] != nil {
		t.Error("Se esperaba EnlacePlantilla nulo al omitirlo y se obtuvo", dato["EnlacePlantilla"])
	}
	if valor := helpers.Fila(t, tablaCargueMasivo, id)["enlace_plantilla"]; valor != nil {
		t.Error("Se esperaba enlace_plantilla NULL en la base de datos al omitir EnlacePlantilla y se obtuvo", valor)
	}
	t.Log("TestCargueMasivoEnlacePlantillaOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestCargueMasivoEnlacePlantillaOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostCargueMasivo(t)
	cuerpo["EnlacePlantilla"] = nil
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaCargueMasivo, id); dato["EnlacePlantilla"] != nil {
		t.Error("Se esperaba EnlacePlantilla nulo al enviar null y se obtuvo", dato["EnlacePlantilla"])
	}
	if valor := helpers.Fila(t, tablaCargueMasivo, id)["enlace_plantilla"]; valor != nil {
		t.Error("Se esperaba enlace_plantilla NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestCargueMasivoEnlacePlantillaOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo TotalRegistros es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestCargueMasivoTotalRegistrosOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostCargueMasivo(t)
	delete(cuerpo, "TotalRegistros")
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaCargueMasivo, id); dato["TotalRegistros"] != nil {
		t.Error("Se esperaba TotalRegistros nulo al omitirlo y se obtuvo", dato["TotalRegistros"])
	}
	if valor := helpers.Fila(t, tablaCargueMasivo, id)["total_registros"]; valor != nil {
		t.Error("Se esperaba total_registros NULL en la base de datos al omitir TotalRegistros y se obtuvo", valor)
	}
	t.Log("TestCargueMasivoTotalRegistrosOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestCargueMasivoTotalRegistrosOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostCargueMasivo(t)
	cuerpo["TotalRegistros"] = nil
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaCargueMasivo, id); dato["TotalRegistros"] != nil {
		t.Error("Se esperaba TotalRegistros nulo al enviar null y se obtuvo", dato["TotalRegistros"])
	}
	if valor := helpers.Fila(t, tablaCargueMasivo, id)["total_registros"]; valor != nil {
		t.Error("Se esperaba total_registros NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestCargueMasivoTotalRegistrosOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo Aceptados es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestCargueMasivoAceptadosOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostCargueMasivo(t)
	delete(cuerpo, "Aceptados")
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaCargueMasivo, id); dato["Aceptados"] != nil {
		t.Error("Se esperaba Aceptados nulo al omitirlo y se obtuvo", dato["Aceptados"])
	}
	if valor := helpers.Fila(t, tablaCargueMasivo, id)["aceptados"]; valor != nil {
		t.Error("Se esperaba aceptados NULL en la base de datos al omitir Aceptados y se obtuvo", valor)
	}
	t.Log("TestCargueMasivoAceptadosOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestCargueMasivoAceptadosOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostCargueMasivo(t)
	cuerpo["Aceptados"] = nil
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaCargueMasivo, id); dato["Aceptados"] != nil {
		t.Error("Se esperaba Aceptados nulo al enviar null y se obtuvo", dato["Aceptados"])
	}
	if valor := helpers.Fila(t, tablaCargueMasivo, id)["aceptados"]; valor != nil {
		t.Error("Se esperaba aceptados NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestCargueMasivoAceptadosOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo Rechazados es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestCargueMasivoRechazadosOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostCargueMasivo(t)
	delete(cuerpo, "Rechazados")
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaCargueMasivo, id); dato["Rechazados"] != nil {
		t.Error("Se esperaba Rechazados nulo al omitirlo y se obtuvo", dato["Rechazados"])
	}
	if valor := helpers.Fila(t, tablaCargueMasivo, id)["rechazados"]; valor != nil {
		t.Error("Se esperaba rechazados NULL en la base de datos al omitir Rechazados y se obtuvo", valor)
	}
	t.Log("TestCargueMasivoRechazadosOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestCargueMasivoRechazadosOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostCargueMasivo(t)
	cuerpo["Rechazados"] = nil
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaCargueMasivo, id); dato["Rechazados"] != nil {
		t.Error("Se esperaba Rechazados nulo al enviar null y se obtuvo", dato["Rechazados"])
	}
	if valor := helpers.Fila(t, tablaCargueMasivo, id)["rechazados"]; valor != nil {
		t.Error("Se esperaba rechazados NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestCargueMasivoRechazadosOpcionalNulo Finalizado Correctamente (OK)")
}

// El campo Respuesta es opcional: si se omite del payload, o se envía null, el
// servicio lo devuelve como null y la base de datos lo guarda como NULL.
func TestCargueMasivoRespuestaOpcionalOmitido(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostCargueMasivo(t)
	delete(cuerpo, "Respuesta")
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaCargueMasivo, id); dato["Respuesta"] != nil {
		t.Error("Se esperaba Respuesta nulo al omitirlo y se obtuvo", dato["Respuesta"])
	}
	if valor := helpers.Fila(t, tablaCargueMasivo, id)["respuesta"]; valor != nil {
		t.Error("Se esperaba respuesta NULL en la base de datos al omitir Respuesta y se obtuvo", valor)
	}
	t.Log("TestCargueMasivoRespuestaOpcionalOmitido Finalizado Correctamente (OK)")
}

func TestCargueMasivoRespuestaOpcionalNulo(t *testing.T) {
	helpers.RequerirBaseDatos(t)
	cuerpo := cuerpoPostCargueMasivo(t)
	cuerpo["Respuesta"] = nil
	id := helpers.CrearPorPost(t, rutaCargueMasivo, tablaCargueMasivo, cuerpo)
	if dato := helpers.ProbarGetOne(t, rutaCargueMasivo, id); dato["Respuesta"] != nil {
		t.Error("Se esperaba Respuesta nulo al enviar null y se obtuvo", dato["Respuesta"])
	}
	if valor := helpers.Fila(t, tablaCargueMasivo, id)["respuesta"]; valor != nil {
		t.Error("Se esperaba respuesta NULL en la base de datos al enviar null y se obtuvo", valor)
	}
	t.Log("TestCargueMasivoRespuestaOpcionalNulo Finalizado Correctamente (OK)")
}
