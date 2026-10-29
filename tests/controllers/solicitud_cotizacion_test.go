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
	helpers.CrearSolicitudCotizacion(t) // el total debe superar el limit de las pruebas
	helpers.ProbarGetAllController(t, rutaSolicitudCotizacion, creado.Id)
	helpers.ProbarFiltrosSolicitud(t, rutaSolicitudCotizacion, creado, creado.Id)
	t.Log("TestSolicitudCotizacionGetAll Finalizado Correctamente (OK)")
}

func TestSolicitudCotizacionActivo(t *testing.T) {
	activo := helpers.CrearSolicitudCotizacion(t)
	inactivo := helpers.CrearSolicitudCotizacion(t)
	if err := models.DeleteSolicitudCotizacion(inactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el registro:", err)
	}
	helpers.ProbarActivoController(t, rutaSolicitudCotizacion, activo.Id, inactivo.Id)
	t.Log("TestSolicitudCotizacionActivo Finalizado Correctamente (OK)")
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
	itemInactivo := helpers.NuevoSolicitudItem(t)
	itemInactivo.SolicitudId = solicitud
	cargue := helpers.NuevoCargueMasivo(t)
	cargue.SolicitudId = solicitud
	for _, err := range []error{
		func() error { _, err := models.AddSolicitudItem(item); return err }(),
		func() error { _, err := models.AddSolicitudItem(itemInactivo); return err }(),
		func() error { _, err := models.AddCargueMasivo(cargue); return err }(),
	} {
		if err != nil {
			t.Fatal("No se pudieron crear las entidades asociadas:", err)
		}
	}
	errorCargue := helpers.NuevoCargueMasivoError(t)
	errorCargue.CargueMasivoId = cargue
	if _, err := models.AddCargueMasivoError(errorCargue); err != nil {
		t.Fatal("No se pudo crear el error de cargue:", err)
	}
	if err := models.DeleteSolicitudItem(itemInactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el ítem:", err)
	}
	ruta := fmt.Sprintf("%s/%d/detalle", rutaSolicitudCotizacion, solicitud.Id)

	// ids retorna los Id de la colección indicada del detalle
	ids := func(t *testing.T, d map[string]interface{}, coleccion string) []float64 {
		t.Helper()
		l, ok := d[coleccion].([]interface{})
		if !ok {
			t.Fatalf("Se esperaba la colección %s y se obtuvo %v", coleccion, d[coleccion])
		}
		res := []float64{}
		for _, r := range l {
			res = append(res, r.(map[string]interface{})["Id"].(float64))
		}
		return res
	}

	t.Run("AsociadasActivasPorDefecto", func(t *testing.T) {
		d := helpers.VerificarRespuesta(t, helpers.Peticion("GET", ruta, ""), http.StatusOK).Dato(t)
		if d["Id"] != float64(solicitud.Id) {
			t.Error("Se esperaba la solicitud", solicitud.Id, "y se obtuvo", d["Id"])
		}
		items, _ := d["Items"].([]interface{})
		if len(items) != 1 || items[0].(map[string]interface{})["Id"] != float64(item.Id) || items[0].(map[string]interface{})["SolicitudId"] != nil {
			t.Error("Se esperaba solo el ítem activo", item.Id, "sin la solicitud repetida y se obtuvo", items)
		}
		for _, lista := range []string{"Necesidades", "PagosParciales", "ActividadesEconomicas", "Documentos", "RequisitosDocumentales", "Proveedores"} {
			if l := ids(t, d, lista); len(l) != 0 {
				t.Errorf("Se esperaba una lista vacía en %s y se obtuvo %v", lista, l)
			}
		}
		if l := ids(t, d, "CarguesMasivos"); len(l) != 1 || l[0] != float64(cargue.Id) {
			t.Error("Se esperaba el cargue", cargue.Id, "y se obtuvo", l)
		}
		// Los errores son una colección independiente, no anidada en cada cargue
		if l := ids(t, d, "CarguesMasivosErrores"); len(l) != 1 || l[0] != float64(errorCargue.Id) {
			t.Error("Se esperaba el error de cargue", errorCargue.Id, "y se obtuvo", l)
		}
		for _, c := range d["CarguesMasivos"].([]interface{}) {
			for campo := range c.(map[string]interface{}) {
				if campo == "CarguesMasivosErrores" || campo == "Errores" {
					t.Error("Los errores de cargue no deben estar anidados en el cargue")
				}
			}
		}
	})

	t.Run("AsociadasInactivas", func(t *testing.T) {
		d := helpers.VerificarRespuesta(t, helpers.Peticion("GET", ruta+"?activo=false", ""), http.StatusOK).Dato(t)
		if d["Id"] != float64(solicitud.Id) {
			t.Error("Se esperaba la solicitud", solicitud.Id, "y se obtuvo", d["Id"])
		}
		if l := ids(t, d, "Items"); len(l) != 1 || l[0] != float64(itemInactivo.Id) {
			t.Error("Se esperaba solo el ítem inactivo", itemInactivo.Id, "y se obtuvo", l)
		}
		if l := ids(t, d, "CarguesMasivosErrores"); len(l) != 0 {
			t.Error("Se esperaba una lista vacía de errores inactivos y se obtuvo", l)
		}
	})

	t.Run("Errores", func(t *testing.T) {
		helpers.VerificarError(t, helpers.Peticion("GET", ruta+"?activo=si", ""), http.StatusNotFound, "el parametro activo debe ser true o false")
		w := helpers.Peticion("GET", fmt.Sprintf("%s/%d/detalle", rutaSolicitudCotizacion, helpers.IdInexistente), "")
		helpers.VerificarError(t, w, http.StatusNotFound, "registro no encontrado")
		w = helpers.Peticion("GET", rutaSolicitudCotizacion+"/abc/detalle", "")
		helpers.VerificarError(t, w, http.StatusNotFound, "el parametro id debe ser numerico")
	})
	t.Log("TestSolicitudCotizacionGetDetalle Finalizado Correctamente (OK)")
}
