package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

const rutaSolicitudCotizacion = "/v1/solicitud_cotizacion"

func TestSolicitudCotizacionPost(t *testing.T) {
	nuevo := helpers.NuevoSolicitudCotizacion(t)
	w := helpers.Peticion("POST", rutaSolicitudCotizacion+"/", helpers.JSON(t, nuevo))
	r := helpers.VerificarRespuesta(t, w, http.StatusCreated)
	if v, err := models.GetSolicitudCotizacionById(r.IdDato(t)); err != nil || v.Titulo != nuevo.Titulo || !v.Activo {
		t.Error("Se esperaba el registro creado y activo y se obtuvo", v, err)
	}
	w = helpers.Peticion("POST", rutaSolicitudCotizacion+"/", `{"Id":"texto"}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "cuerpo de la solicitud invalido")
	t.Log("TestSolicitudCotizacionPost Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionGetOne(t *testing.T) {
	creado := helpers.CrearSolicitudCotizacion(t)
	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudCotizacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el registro", creado.Id, "y se obtuvo", r.Data)
	}
	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d", rutaSolicitudCotizacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudCotizacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudCotizacionGetOne Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionGetAll(t *testing.T) {
	creado := helpers.CrearSolicitudCotizacion(t)
	helpers.ProbarGetAllController(t, rutaSolicitudCotizacion, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudCotizacion, creado, creado.Id)
	t.Log("TestSolicitudCotizacionGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionPut(t *testing.T) {
	creado := helpers.CrearSolicitudCotizacion(t)
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudCotizacion, creado.Id)
	cambios := *creado
	cambios.Id = 0
	cambios.Titulo = "Titulo actualizado"
	helpers.VerificarRespuesta(t, helpers.Peticion("PUT", ruta, helpers.JSON(t, cambios)), http.StatusOK)
	if v, err := models.GetSolicitudCotizacionById(creado.Id); err != nil || v.Titulo != "Titulo actualizado" {
		t.Error("Se esperaba el registro actualizado y se obtuvo", v, err)
	}

	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, `{"Id":"texto"}`), http.StatusBadRequest, "cuerpo de la solicitud invalido")
	helpers.VerificarError(t, helpers.Peticion("PUT", rutaSolicitudCotizacion+"/abc", helpers.JSON(t, cambios)), http.StatusBadRequest, "el parametro id debe ser numerico")
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudCotizacion, helpers.IdInexistente), helpers.JSON(t, cambios))
	helpers.VerificarError(t, w, http.StatusBadRequest, "no existe un registro con el id indicado")
	t.Log("TestSolicitudCotizacionPut Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionDelete(t *testing.T) {
	creado := helpers.CrearSolicitudCotizacion(t)
	w := helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudCotizacion, creado.Id), "")
	if r := helpers.VerificarRespuesta(t, w, http.StatusOK); r.IdDato(t) != creado.Id {
		t.Error("Se esperaba el Id eliminado", creado.Id, "y se obtuvo", r.Data)
	}
	if v, _ := models.GetSolicitudCotizacionById(creado.Id); v == nil || v.Activo {
		t.Error("Se esperaba el registro inactivo (borrado lógico) y se obtuvo", v)
	}
	w = helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", rutaSolicitudCotizacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "no existe un registro con el id indicado")
	w = helpers.Peticion("DELETE", rutaSolicitudCotizacion+"/abc", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudCotizacionDelete Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionGetDetalle(t *testing.T) {
	solicitud := helpers.CrearSolicitudCotizacion(t)
	item := helpers.NuevoSolicitudItem(t)
	item.SolicitudId = solicitud
	if _, err := models.AddSolicitudItem(item); err != nil {
		t.Fatal("No se pudo crear el ítem:", err)
	}

	w := helpers.Peticion("GET", fmt.Sprintf("%s/%d/detalle", rutaSolicitudCotizacion, solicitud.Id), "")
	d := helpers.VerificarRespuesta(t, w, http.StatusOK).Dato(t)
	// Verificar la ruta sin /v1 utilizada por el MID.
	w = helpers.Peticion("GET", fmt.Sprintf("/solicitud_cotizacion/%d/detalle", solicitud.Id), "")
	dCompat := helpers.VerificarRespuesta(t, w, http.StatusOK).Dato(t)
	if dCompat["Id"] != float64(solicitud.Id) {
		t.Errorf("La ruta de compatibilidad debía devolver la solicitud %d y devolvió %v", solicitud.Id, dCompat["Id"])
	}
	if d["Id"] != float64(solicitud.Id) {
		t.Error("Se esperaba la solicitud", solicitud.Id, "y se obtuvo", d["Id"])
	}
	items, ok := d["Items"].([]interface{})
	if !ok || len(items) != 1 {
		t.Fatal("Se esperaba 1 ítem y se obtuvo", d["Items"])
	}
	if i, ok := items[0].(map[string]interface{}); !ok || i["Id"] != float64(item.Id) || i["SolicitudId"] != nil {
		t.Error("Se esperaba el ítem", item.Id, "sin la solicitud repetida y se obtuvo", items[0])
	}
	for _, lista := range []string{"Necesidades", "PagosParciales", "ActividadesEconomicas", "Documentos", "RequisitosDocumentales", "Proveedores", "CarguesMasivos"} {
		if l, ok := d[lista].([]interface{}); !ok || len(l) != 0 {
			t.Errorf("Se esperaba una lista vacía en %s y se obtuvo %v", lista, d[lista])
		}
	}

	w = helpers.Peticion("GET", fmt.Sprintf("%s/%d/detalle", rutaSolicitudCotizacion, helpers.IdInexistente), "")
	helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
	w = helpers.Peticion("GET", rutaSolicitudCotizacion+"/abc/detalle", "")
	helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	t.Log("TestSolicitudCotizacionGetDetalle Finalizado Correctamente (OK)")
}
