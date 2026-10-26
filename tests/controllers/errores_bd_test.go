package controllers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// Los mensajes de duplicados, llaves foráneas inexistentes, longitud y rango
// dependen de los códigos de error de PostgreSQL (23505, 23503, 22001 y 22003),
// que SQLite no genera; se prueban en las pruebas de integración.

func TestErroresBDRelacionObligatoria(t *testing.T) {
	// Sin SolicitudId el ORM rechaza el registro antes de llegar a la base de datos
	w := helpers.Peticion("POST", rutaSolicitudItem+"/", `{"Nombre":"Item","Descripcion":"x","TipoBienServicioId":1,"Cantidad":1}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "error al crear el registro: el campo SolicitudId es obligatorio y debe incluir un Id valido")

	creado := helpers.CrearSolicitudItem(t)
	w = helpers.Peticion("PUT", fmt.Sprintf("%s/%d", rutaSolicitudItem, creado.Id), `{"Nombre":"Item","Descripcion":"x","TipoBienServicioId":1,"Cantidad":1}`)
	helpers.VerificarError(t, w, http.StatusBadRequest, "error al actualizar el registro: el campo SolicitudId es obligatorio y debe incluir un Id valido")
	t.Log("TestErroresBDRelacionObligatoria Finalizado Correctamente (OK)")
}

func TestErroresBDRegistroInexistente(t *testing.T) {
	// PUT responde 400 y DELETE 404, según el lineamiento de control de errores del CRUD
	ruta := fmt.Sprintf("%s/%d", rutaSolicitudCotizacion, helpers.IdInexistente)
	cuerpo := helpers.JSON(t, helpers.NuevoSolicitudCotizacion(t))
	helpers.VerificarError(t, helpers.Peticion("PUT", ruta, cuerpo), http.StatusBadRequest, "error al actualizar el registro: no existe un registro con el id indicado")
	helpers.VerificarError(t, helpers.Peticion("DELETE", ruta, ""), http.StatusNotFound, "error al eliminar el registro: no existe un registro con el id indicado")
	t.Log("TestErroresBDRegistroInexistente Finalizado Correctamente (OK)")
}
