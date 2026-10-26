package helpers

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
)

// FuncionGetAll corresponde a la firma de las funciones GetAll<Modelo> del
// paquete models.
type FuncionGetAll func(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) ([]interface{}, error)

// ProbarGetAll verifica el comportamiento común de una función GetAll<Modelo>:
// filtros, selección de campos, ordenamiento, paginación y validaciones.
// Requiere que exista en la base de datos el registro con Id idExistente.
func ProbarGetAll(t *testing.T, getAll FuncionGetAll, idExistente int) {
	t.Helper()
	id := strconv.Itoa(idExistente)

	t.Run("SinFiltros", func(t *testing.T) {
		l, err := getAll(map[string]string{}, nil, nil, nil, 0, 0)
		if err != nil || len(l) == 0 {
			t.Error("Se esperaban registros y se obtuvo", len(l), err)
		}
	})

	t.Run("FiltroPorId", func(t *testing.T) {
		l, err := getAll(map[string]string{"Id": id}, nil, nil, nil, 0, 10)
		if err != nil || len(l) != 1 {
			t.Error("Se esperaba 1 registro y se obtuvo", len(l), err)
		}
	})

	t.Run("FiltroIsNull", func(t *testing.T) {
		l, err := getAll(map[string]string{"Id__isnull": "false"}, nil, nil, nil, 0, 10)
		if err != nil || len(l) == 0 {
			t.Error("Se esperaban registros y se obtuvo", len(l), err)
		}
	})

	t.Run("SinResultados", func(t *testing.T) {
		l, err := getAll(map[string]string{"Id": strconv.Itoa(IdInexistente)}, nil, nil, nil, 0, 10)
		if err != nil || l != nil {
			t.Error("Se esperaba una lista vacía y se obtuvo", l, err)
		}
	})

	t.Run("SeleccionDeCampos", func(t *testing.T) {
		l, err := getAll(map[string]string{"Id": id}, []string{"Id", "Activo"}, nil, nil, 0, 10)
		if err != nil || len(l) != 1 {
			t.Fatal("Se esperaba 1 registro y se obtuvo", len(l), err)
		}
		m, ok := l[0].(map[string]interface{})
		if !ok || len(m) != 2 || m["Id"] != idExistente {
			t.Error("Se esperaba un mapa con los campos Id y Activo y se obtuvo", l[0])
		}
	})

	t.Run("OrdenPorCampo", func(t *testing.T) {
		l, err := getAll(map[string]string{}, nil, []string{"Id", "Activo"}, []string{"desc", "asc"}, 0, 10)
		if err != nil || len(l) == 0 {
			t.Error("Se esperaban registros ordenados y se obtuvo", len(l), err)
		}
	})

	t.Run("OrdenUnicoParaVariosCampos", func(t *testing.T) {
		for _, orden := range []string{"asc", "desc"} {
			l, err := getAll(map[string]string{}, nil, []string{"Id", "Activo"}, []string{orden}, 0, 10)
			if err != nil || len(l) == 0 {
				t.Error("Se esperaban registros ordenados", orden, "y se obtuvo", len(l), err)
			}
		}
	})

	t.Run("Paginacion", func(t *testing.T) {
		l, err := getAll(map[string]string{}, nil, nil, nil, 0, 1)
		if err != nil || len(l) != 1 {
			t.Error("Se esperaba 1 registro con limit=1 y se obtuvo", len(l), err)
		}
	})

	// El ORM lanza un panic con campos inexistentes; la función debe
	// convertirlo en error (models/consultas.go).
	casosCampoInexistente := []struct {
		nombre string
		query  map[string]string
		fields []string
		sortby []string
		order  []string
	}{
		{"CampoInexistenteEnFields", map[string]string{}, []string{"CampoInexistente"}, nil, nil},
		{"CampoInexistenteEnQuery", map[string]string{"CampoInexistente": "1"}, nil, nil, nil},
		{"CampoInexistenteEnSortby", map[string]string{}, nil, []string{"CampoInexistente"}, []string{"asc"}},
	}
	for _, c := range casosCampoInexistente {
		t.Run(c.nombre, func(t *testing.T) {
			if _, err := getAll(c.query, c.fields, c.sortby, c.order, 0, 10); err == nil {
				t.Error("Se esperaba error al usar un campo inexistente")
			}
		})
	}

	casosOrdenInvalido := []struct {
		nombre string
		sortby []string
		order  []string
	}{
		{"OrdenInvalidoPorCampo", []string{"Id"}, []string{"abc"}},
		{"OrdenUnicoInvalido", []string{"Id", "Activo"}, []string{"abc"}},
		{"TamanosNoCoinciden", []string{"Id", "Activo"}, []string{"asc", "desc", "asc"}},
		{"OrderSinSortby", nil, []string{"asc"}},
	}
	for _, c := range casosOrdenInvalido {
		t.Run(c.nombre, func(t *testing.T) {
			if _, err := getAll(map[string]string{}, nil, c.sortby, c.order, 0, 10); err == nil {
				t.Error("Se esperaba error con sortby", c.sortby, "y order", c.order)
			}
		})
	}
}

// ProbarGetAllController verifica el comportamiento común del servicio GetAll
// de un controlador, cuya ruta base es ruta (p. ej. /v1/solicitud_item).
func ProbarGetAllController(t *testing.T, ruta string, idExistente int) {
	t.Helper()
	id := strconv.Itoa(idExistente)

	t.Run("ConParametros", func(t *testing.T) {
		w := Peticion("GET", ruta+"/?query=Id:"+id+"&fields=Id,Activo&sortby=Id&order=desc&limit=5&offset=0", "")
		l := VerificarRespuesta(t, w, http.StatusOK).Lista(t)
		if len(l) != 1 {
			t.Fatal("Se esperaba 1 registro y se obtuvo", l)
		}
		if m, ok := l[0].(map[string]interface{}); !ok || m["Id"] != float64(idExistente) {
			t.Error("Se esperaba el registro", id, "y se obtuvo", l[0])
		}
	})

	t.Run("SinResultados", func(t *testing.T) {
		w := Peticion("GET", ruta+"/?query=Id:"+strconv.Itoa(IdInexistente), "")
		l := VerificarRespuesta(t, w, http.StatusOK).Lista(t)
		if m, ok := l[0].(map[string]interface{}); len(l) != 1 || !ok || len(m) != 0 {
			t.Error("Se esperaba una lista con un objeto vacío y se obtuvo", l)
		}
	})

	t.Run("QueryInvalido", func(t *testing.T) {
		VerificarError(t, Peticion("GET", ruta+"/?query=Id", ""), http.StatusBadRequest, "query invalida")
	})

	t.Run("LimitInvalido", func(t *testing.T) {
		VerificarError(t, Peticion("GET", ruta+"/?limit=abc", ""), http.StatusBadRequest, "limit")
	})

	t.Run("OffsetInvalido", func(t *testing.T) {
		VerificarError(t, Peticion("GET", ruta+"/?offset=-1", ""), http.StatusBadRequest, "offset")
	})

	t.Run("OrdenInvalido", func(t *testing.T) {
		VerificarError(t, Peticion("GET", ruta+"/?sortby=Id&order=abc", ""), http.StatusNotFound, "error al consultar")
	})

	t.Run("CampoInexistente", func(t *testing.T) {
		VerificarError(t, Peticion("GET", ruta+"/?query=CampoInexistente:1", ""), http.StatusNotFound, "error al consultar")
	})
}

// ProbarFiltrosSolicitud verifica los filtros vigencia, estado y dependencia
// del servicio GetAll. solicitud es la solicitud a la que pertenece el
// registro idEsperado; cada solicitud de prueba tiene valores únicos en esos
// campos, por lo que el filtro debe retornar solo ese registro.
func ProbarFiltrosSolicitud(t *testing.T, ruta string, solicitud *models.SolicitudCotizacion, idEsperado int) {
	t.Helper()

	t.Run("FiltrosVigenciaEstadoDependencia", func(t *testing.T) {
		w := Peticion("GET", fmt.Sprintf("%s/?vigencia=%d&estado=%d&dependencia=%d",
			ruta, solicitud.Vigencia, solicitud.EstadoId, solicitud.DependenciaSolicitanteId), "")
		l := VerificarRespuesta(t, w, http.StatusOK).Lista(t)
		if len(l) != 1 {
			t.Fatal("Se esperaba 1 registro y se obtuvo", l)
		}
		if m, ok := l[0].(map[string]interface{}); !ok || m["Id"] != float64(idEsperado) {
			t.Error("Se esperaba el registro", idEsperado, "y se obtuvo", l[0])
		}
	})

	t.Run("FiltrosSinCoincidencia", func(t *testing.T) {
		w := Peticion("GET", fmt.Sprintf("%s/?vigencia=%d&estado=%d", ruta, solicitud.Vigencia, IdInexistente), "")
		l := VerificarRespuesta(t, w, http.StatusOK).Lista(t)
		if m, ok := l[0].(map[string]interface{}); len(l) != 1 || !ok || len(m) != 0 {
			t.Error("Se esperaba una lista con un objeto vacío y se obtuvo", l)
		}
	})

	for _, parametro := range []string{"vigencia", "estado", "dependencia"} {
		t.Run("FiltroNoNumerico_"+parametro, func(t *testing.T) {
			w := Peticion("GET", ruta+"/?"+parametro+"=abc", "")
			VerificarError(t, w, http.StatusBadRequest, "el parametro "+parametro+" debe ser numerico")
		})
	}
}
