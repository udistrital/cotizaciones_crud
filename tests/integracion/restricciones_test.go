package integracion

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// crear envía el registro al servicio POST de la ruta y retorna su Id.
func crear(t *testing.T, ruta string, registro interface{}) int {
	t.Helper()
	w := helpers.Peticion("POST", ruta+"/", helpers.JSON(t, registro))
	return helpers.VerificarRespuesta(t, w, http.StatusCreated).IdDato(t)
}

// conRelacionInexistente asigna a la relación campo del registro un Id que no
// existe en la base de datos.
func conRelacionInexistente(registro interface{}, campo string) interface{} {
	f := reflect.ValueOf(registro).Elem().FieldByName(campo)
	relacion := reflect.New(f.Type().Elem())
	relacion.Elem().FieldByName("Id").SetInt(helpers.IdInexistente)
	f.Set(relacion)
	return registro
}

// TestIndicesUnicos verifica que cada índice único del script rechace un
// registro duplicado con el mensaje que indica las columnas y los valores.
func TestIndicesUnicos(t *testing.T) {
	casos := []struct {
		nombre string
		ruta   string
		// registro retorna el registro a duplicar y las columnas del índice
		registro func(t *testing.T) (interface{}, string)
	}{
		{"SolicitudCotizacion", "/v1/solicitud_cotizacion", func(t *testing.T) (interface{}, string) {
			v := helpers.NuevoSolicitudCotizacion(t)
			numero := fmt.Sprintf("SC-%d", v.Vigencia)
			v.NumeroSolicitud = &numero
			return v, "numero_solicitud = " + numero
		}},
		{"SolicitudNecesidad", "/v1/solicitud_necesidad", func(t *testing.T) (interface{}, string) {
			v := helpers.NuevoSolicitudNecesidad(t)
			return v, fmt.Sprintf("solicitud_id = %d, numero_necesidad = %s", v.SolicitudId.Id, v.NumeroNecesidad)
		}},
		{"SolicitudCiiu", "/v1/solicitud_ciiu", func(t *testing.T) (interface{}, string) {
			v := helpers.NuevoSolicitudCiiu(t)
			return v, fmt.Sprintf("solicitud_id = %d, actividad_economica_id = %d", v.SolicitudId.Id, v.ActividadEconomicaId)
		}},
		{"SolicitudRequisitoDocumental", "/v1/solicitud_requisito_documental", func(t *testing.T) (interface{}, string) {
			v := helpers.NuevoSolicitudRequisitoDocumental(t)
			return v, fmt.Sprintf("solicitud_id = %d, tipo_documento_id = %d", v.SolicitudId.Id, v.TipoDocumentoId)
		}},
		{"SolicitudProveedor", "/v1/solicitud_proveedor", func(t *testing.T) (interface{}, string) {
			v := helpers.NuevoSolicitudProveedor(t)
			return v, fmt.Sprintf("solicitud_id = %d, proveedor_id = %d", v.SolicitudId.Id, v.ProveedorId)
		}},
		{"SolicitudModificacion", "/v1/solicitud_modificacion", func(t *testing.T) (interface{}, string) {
			v := helpers.NuevoSolicitudModificacion(t)
			return v, fmt.Sprintf("solicitud_id = %d, numero = %d", v.SolicitudId.Id, v.Numero)
		}},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			registro, columnas := c.registro(t)
			id := crear(t, c.ruta, registro)
			mensaje := "error al crear el registro: ya existe un registro activo o inactivo con " + columnas

			t.Run("Duplicado", func(t *testing.T) {
				w := helpers.Peticion("POST", c.ruta+"/", helpers.JSON(t, registro))
				helpers.VerificarError(t, w, http.StatusBadRequest, mensaje)
			})

			// El borrado es lógico, por lo que el registro inactivo sigue ocupando el índice
			t.Run("DuplicadoDeRegistroInactivo", func(t *testing.T) {
				helpers.VerificarRespuesta(t, helpers.Peticion("DELETE", fmt.Sprintf("%s/%d", c.ruta, id), ""), http.StatusOK)
				w := helpers.Peticion("POST", c.ruta+"/", helpers.JSON(t, registro))
				helpers.VerificarError(t, w, http.StatusBadRequest, mensaje)
			})
		})
	}
	t.Log("TestIndicesUnicos Finalizado Correctamente (OK)")
}

// TestIndiceUnicoAlActualizar verifica que el índice único también se valide al
// actualizar un registro con los valores de otro.
func TestIndiceUnicoAlActualizar(t *testing.T) {
	ruta := "/v1/solicitud_necesidad"
	primera := helpers.NuevoSolicitudNecesidad(t)
	crear(t, ruta, primera)
	segunda := *primera
	segunda.NumeroNecesidad = primera.NumeroNecesidad + "-2"
	idSegunda := crear(t, ruta, &segunda)

	segunda.NumeroNecesidad = primera.NumeroNecesidad
	w := helpers.Peticion("PUT", fmt.Sprintf("%s/%d", ruta, idSegunda), helpers.JSON(t, &segunda))
	helpers.VerificarError(t, w, http.StatusBadRequest, fmt.Sprintf(
		"error al actualizar el registro: ya existe un registro activo o inactivo con solicitud_id = %d, numero_necesidad = %s",
		primera.SolicitudId.Id, primera.NumeroNecesidad))
	t.Log("TestIndiceUnicoAlActualizar Finalizado Correctamente (OK)")
}

// TestLlavesForaneas verifica que cada llave foránea del script rechace una
// relación con un registro inexistente, indicando la columna y el valor.
func TestLlavesForaneas(t *testing.T) {
	casos := []struct {
		nombre   string
		ruta     string
		campo    string
		columna  string
		registro func(t *testing.T) interface{}
	}{
		{"SolicitudNecesidad", "/v1/solicitud_necesidad", "SolicitudId", "solicitud_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudNecesidad(t) }},
		{"SolicitudItem", "/v1/solicitud_item", "SolicitudId", "solicitud_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudItem(t) }},
		{"SolicitudItemCargueMasivo", "/v1/solicitud_item", "CargueMasivoId", "cargue_masivo_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudItem(t) }},
		{"CargueMasivo", "/v1/cargue_masivo", "SolicitudId", "solicitud_id", func(t *testing.T) interface{} { return helpers.NuevoCargueMasivo(t) }},
		{"CargueMasivoError", "/v1/cargue_masivo_error", "CargueMasivoId", "cargue_masivo_id", func(t *testing.T) interface{} { return helpers.NuevoCargueMasivoError(t) }},
		{"SolicitudPagoParcial", "/v1/solicitud_pago_parcial", "SolicitudId", "solicitud_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudPagoParcial(t) }},
		{"SolicitudCiiu", "/v1/solicitud_ciiu", "SolicitudId", "solicitud_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudCiiu(t) }},
		{"DocumentoSolicitud", "/v1/documento_solicitud", "SolicitudId", "solicitud_id", func(t *testing.T) interface{} { return helpers.NuevoDocumentoSolicitud(t) }},
		{"SolicitudRequisitoDocumental", "/v1/solicitud_requisito_documental", "SolicitudId", "solicitud_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudRequisitoDocumental(t) }},
		{"SolicitudProveedor", "/v1/solicitud_proveedor", "SolicitudId", "solicitud_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudProveedor(t) }},
		{"SolicitudEstado", "/v1/solicitud_estado", "SolicitudId", "solicitud_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudEstado(t) }},
		{"SolicitudModificacion", "/v1/solicitud_modificacion", "SolicitudId", "solicitud_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudModificacion(t) }},
		{"SolicitudModificacionEstado", "/v1/solicitud_modificacion_estado", "SolicitudModificacionId", "solicitud_modificacion_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudModificacionEstado(t) }},
		{"SolicitudModificacionDetalle", "/v1/solicitud_modificacion_detalle", "SolicitudModificacionId", "solicitud_modificacion_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudModificacionDetalle(t) }},
		{"SolicitudModificacionDetalleItem", "/v1/solicitud_modificacion_detalle", "SolicitudItemId", "solicitud_item_id", func(t *testing.T) interface{} { return helpers.NuevoSolicitudModificacionDetalle(t) }},
		{"Observacion", "/v1/observacion", "SolicitudId", "solicitud_id", func(t *testing.T) interface{} { return helpers.NuevoObservacion(t) }},
		{"ObservacionProveedorConvocado", "/v1/observacion", "ProveedorConvocadoId", "proveedor_convocado_id", func(t *testing.T) interface{} { return helpers.NuevoObservacion(t) }},
		{"RespuestaObservacion", "/v1/respuesta_observacion", "ObservacionId", "observacion_id", func(t *testing.T) interface{} { return helpers.NuevoRespuestaObservacion(t) }},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			registro := conRelacionInexistente(c.registro(t), c.campo)
			w := helpers.Peticion("POST", c.ruta+"/", helpers.JSON(t, registro))
			helpers.VerificarError(t, w, http.StatusBadRequest, fmt.Sprintf(
				"error al crear el registro: no existe el registro referenciado por %s = %d", c.columna, helpers.IdInexistente))
		})
	}

	t.Run("AlActualizar", func(t *testing.T) {
		item := helpers.CrearSolicitudItem(t)
		conRelacionInexistente(item, "SolicitudId")
		w := helpers.Peticion("PUT", fmt.Sprintf("/v1/solicitud_item/%d", item.Id), helpers.JSON(t, item))
		helpers.VerificarError(t, w, http.StatusBadRequest, fmt.Sprintf(
			"error al actualizar el registro: no existe el registro referenciado por solicitud_id = %d", helpers.IdInexistente))
	})
	t.Log("TestLlavesForaneas Finalizado Correctamente (OK)")
}

// TestLongitudMaxima verifica que un texto más largo que su columna varchar se
// rechace indicando la longitud permitida.
func TestLongitudMaxima(t *testing.T) {
	casos := []struct {
		nombre    string
		modificar func(v interface{})
		tipo      string
	}{
		{"PlazoEjecucion", func(v interface{}) {
			reflect.ValueOf(v).Elem().FieldByName("PlazoEjecucion").SetString(strings.Repeat("a", 251))
		}, "character varying(250)"},
		{"NumeroSolicitud", func(v interface{}) {
			numero := strings.Repeat("1", 21)
			reflect.ValueOf(v).Elem().FieldByName("NumeroSolicitud").Set(reflect.ValueOf(&numero))
		}, "character varying(20)"},
	}
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			v := helpers.NuevoSolicitudCotizacion(t)
			c.modificar(v)
			w := helpers.Peticion("POST", "/v1/solicitud_cotizacion/", helpers.JSON(t, v))
			helpers.VerificarError(t, w, http.StatusBadRequest, "error al crear el registro: un texto supera la longitud maxima permitida ("+c.tipo+")")
		})
	}
	t.Log("TestLongitudMaxima Finalizado Correctamente (OK)")
}

// TestRangoNumerico verifica que un valor fuera de la precisión de una columna
// numeric se rechace. porcentaje es numeric(5,4), por lo que su máximo es 9.9999.
func TestRangoNumerico(t *testing.T) {
	v := helpers.NuevoSolicitudPagoParcial(t)
	v.Porcentaje = 10
	w := helpers.Peticion("POST", "/v1/solicitud_pago_parcial/", helpers.JSON(t, v))
	helpers.VerificarError(t, w, http.StatusBadRequest, "error al crear el registro: un valor numerico esta fuera del rango permitido")
	t.Log("TestRangoNumerico Finalizado Correctamente (OK)")
}

// TestJSONInvalido verifica que un texto que no es JSON en una columna jsonb se rechace.
func TestJSONInvalido(t *testing.T) {
	v := helpers.NuevoSolicitudModificacionDetalle(t)
	invalido := "no es un JSON"
	v.ValorAnterior = &invalido
	w := helpers.Peticion("POST", "/v1/solicitud_modificacion_detalle/", helpers.JSON(t, v))
	helpers.VerificarError(t, w, http.StatusBadRequest, "error al crear el registro: un valor no tiene el formato esperado, por ejemplo un JSON invalido")
	t.Log("TestJSONInvalido Finalizado Correctamente (OK)")
}
