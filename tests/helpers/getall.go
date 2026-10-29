package helpers

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
)

// FuncionGetAll corresponde a la firma de las funciones GetAll<Modelo> del
// paquete models, que retornan la lista, el total de registros que cumplen los
// filtros y el error.
type FuncionGetAll func(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) ([]interface{}, int64, error)

// ProbarGetAll verifica el comportamiento común de una función GetAll<Modelo>:
// filtros, selección de campos, ordenamiento, paginación, total y validaciones.
// Requiere al menos dos registros en la tabla, uno de ellos con Id idExistente.
func ProbarGetAll(t *testing.T, getAll FuncionGetAll, idExistente int) {
	t.Helper()
	id := strconv.Itoa(idExistente)

	t.Run("SinFiltros", func(t *testing.T) {
		l, _, err := getAll(map[string]string{}, nil, nil, nil, 0, 0)
		if err != nil || len(l) == 0 {
			t.Error("Se esperaban registros y se obtuvo", len(l), err)
		}
	})

	t.Run("FiltroPorId", func(t *testing.T) {
		l, _, err := getAll(map[string]string{"Id": id}, nil, nil, nil, 0, 10)
		if err != nil || len(l) != 1 {
			t.Error("Se esperaba 1 registro y se obtuvo", len(l), err)
		}
	})

	t.Run("FiltroIsNull", func(t *testing.T) {
		l, _, err := getAll(map[string]string{"Id__isnull": "false"}, nil, nil, nil, 0, 10)
		if err != nil || len(l) == 0 {
			t.Error("Se esperaban registros y se obtuvo", len(l), err)
		}
	})

	t.Run("SinResultados", func(t *testing.T) {
		l, _, err := getAll(map[string]string{"Id": strconv.Itoa(IdInexistente)}, nil, nil, nil, 0, 10)
		if err != nil || l != nil {
			t.Error("Se esperaba una lista vacía y se obtuvo", l, err)
		}
	})

	t.Run("SeleccionDeCampos", func(t *testing.T) {
		l, _, err := getAll(map[string]string{"Id": id}, []string{"Id", "Activo"}, nil, nil, 0, 10)
		if err != nil || len(l) != 1 {
			t.Fatal("Se esperaba 1 registro y se obtuvo", len(l), err)
		}
		m, ok := l[0].(map[string]interface{})
		if !ok || len(m) != 2 || m["Id"] != idExistente {
			t.Error("Se esperaba un mapa con los campos Id y Activo y se obtuvo", l[0])
		}
	})

	t.Run("OrdenPorCampo", func(t *testing.T) {
		l, _, err := getAll(map[string]string{}, nil, []string{"Id", "Activo"}, []string{"desc", "asc"}, 0, 10)
		if err != nil || len(l) == 0 {
			t.Error("Se esperaban registros ordenados y se obtuvo", len(l), err)
		}
	})

	t.Run("OrdenUnicoParaVariosCampos", func(t *testing.T) {
		for _, orden := range []string{"asc", "desc"} {
			l, _, err := getAll(map[string]string{}, nil, []string{"Id", "Activo"}, []string{orden}, 0, 10)
			if err != nil || len(l) == 0 {
				t.Error("Se esperaban registros ordenados", orden, "y se obtuvo", len(l), err)
			}
		}
	})

	t.Run("Paginacion", func(t *testing.T) {
		l, _, err := getAll(map[string]string{}, nil, nil, nil, 0, 1)
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
			if _, _, err := getAll(c.query, c.fields, c.sortby, c.order, 0, 10); err == nil {
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
			if _, _, err := getAll(map[string]string{}, nil, c.sortby, c.order, 0, 10); err == nil {
				t.Error("Se esperaba error con sortby", c.sortby, "y order", c.order)
			}
		})
	}

	// El total cuenta los registros que cumplen los filtros sin aplicar limit
	// ni offset
	t.Run("TotalIgnoraLimit", func(t *testing.T) {
		l, total, err := getAll(map[string]string{}, nil, nil, nil, 0, 1)
		if err != nil || len(l) != 1 || total < 2 {
			t.Error("Se esperaba 1 registro y un total mayor o igual a 2 y se obtuvo", len(l), total, err)
		}
	})

	t.Run("TotalIgnoraOffset", func(t *testing.T) {
		_, total, _ := getAll(map[string]string{}, nil, nil, nil, 0, 1)
		l, totalConOffset, err := getAll(map[string]string{}, nil, nil, nil, total, 10)
		if err != nil || l != nil || totalConOffset != total {
			t.Error("Se esperaba una página vacía con el mismo total", total, "y se obtuvo", len(l), totalConOffset, err)
		}
	})

	t.Run("TotalRespetaFiltros", func(t *testing.T) {
		if _, total, err := getAll(map[string]string{"Id": id}, nil, nil, nil, 0, 10); err != nil || total != 1 {
			t.Error("Se esperaba un total de 1 con el filtro por Id y se obtuvo", total, err)
		}
		if _, total, err := getAll(map[string]string{"Id": strconv.Itoa(IdInexistente)}, nil, nil, nil, 0, 10); err != nil || total != 0 {
			t.Error("Se esperaba un total de 0 sin resultados y se obtuvo", total, err)
		}
	})
}

// ProbarGetAllController verifica el comportamiento común del servicio GetAll
// de un controlador, cuya ruta base es ruta (p. ej. /v1/solicitud_item).
// Requiere al menos dos registros activos, uno de ellos con Id idExistente.
func ProbarGetAllController(t *testing.T, ruta string, idExistente int) {
	t.Helper()
	id := strconv.Itoa(idExistente)

	t.Run("ConParametros", func(t *testing.T) {
		w := Peticion("GET", ruta+"/?query=Id:"+id+"&fields=Id,Activo&sortby=Id&order=desc&limit=5&offset=0", "")
		r := VerificarRespuesta(t, w, http.StatusOK)
		l := r.Lista(t)
		if len(l) != 1 {
			t.Fatal("Se esperaba 1 registro y se obtuvo", l)
		}
		if m, ok := l[0].(map[string]interface{}); !ok || m["Id"] != float64(idExistente) {
			t.Error("Se esperaba el registro", id, "y se obtuvo", l[0])
		}
		if r.Total == nil || *r.Total != 1 {
			t.Error("Se esperaba Total 1 con el filtro por Id y se obtuvo", r.Total)
		}
	})

	t.Run("SinResultados", func(t *testing.T) {
		w := Peticion("GET", ruta+"/?query=Id:"+strconv.Itoa(IdInexistente), "")
		r := VerificarRespuesta(t, w, http.StatusOK)
		l := r.Lista(t)
		if m, ok := l[0].(map[string]interface{}); len(l) != 1 || !ok || len(m) != 0 {
			t.Error("Se esperaba una lista con un objeto vacío y se obtuvo", l)
		}
		if r.Total == nil || *r.Total != 0 {
			t.Error("Se esperaba Total 0 sin resultados y se obtuvo", r.Total)
		}
	})

	// El total cuenta los registros que cumplen los filtros sin aplicar limit
	// ni offset
	t.Run("TotalIgnoraLimit", func(t *testing.T) {
		r := VerificarRespuesta(t, Peticion("GET", ruta+"/?limit=1", ""), http.StatusOK)
		if l := r.Lista(t); len(l) != 1 || r.Total == nil || *r.Total < 2 {
			t.Error("Se esperaba 1 registro y un Total mayor o igual a 2 y se obtuvo", l, r.Total)
		}
	})

	t.Run("TotalIgnoraOffset", func(t *testing.T) {
		r := VerificarRespuesta(t, Peticion("GET", ruta+"/?offset="+strconv.Itoa(IdInexistente), ""), http.StatusOK)
		if r.Total == nil || *r.Total < 2 {
			t.Error("Se esperaba un Total mayor o igual a 2 en una página vacía y se obtuvo", r.Total)
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

// ProbarActivoController verifica el parámetro activo en las consultas de lista
// y por id: por defecto solo se retornan registros activos y con activo=false
// solo los inactivos. idActivo e idInactivo son registros de la ruta.
func ProbarActivoController(t *testing.T, ruta string, idActivo, idInactivo int) {
	t.Helper()
	activo, inactivo := strconv.Itoa(idActivo), strconv.Itoa(idInactivo)

	// total retorna el Total de la lista consultada con los parámetros indicados
	total := func(t *testing.T, params string) int64 {
		t.Helper()
		r := VerificarRespuesta(t, Peticion("GET", ruta+"/?"+params, ""), http.StatusOK)
		if r.Total == nil {
			t.Fatal("La respuesta no contiene Total")
		}
		return *r.Total
	}

	casosLista := []struct {
		nombre   string
		params   string
		esperado int64
	}{
		{"ListaActivoPorDefecto", "query=Id:" + activo, 1},
		{"ListaInactivoExcluidoPorDefecto", "query=Id:" + inactivo, 0},
		{"ListaActivoTrue", "query=Id:" + activo + "&activo=true", 1},
		{"ListaInactivoConActivoFalse", "query=Id:" + inactivo + "&activo=false", 1},
		{"ListaActivoExcluidoConActivoFalse", "query=Id:" + activo + "&activo=false", 0},
		{"ListaActivoDentroDeQuery", "query=Id:" + inactivo + ",Activo:false", 1},
		{"ListaParametroActivoPrevaleceSobreQuery", "query=Id:" + inactivo + ",Activo:false&activo=true", 0},
	}
	for _, c := range casosLista {
		t.Run(c.nombre, func(t *testing.T) {
			if n := total(t, c.params); n != c.esperado {
				t.Errorf("Se esperaba Total %d con %q y se obtuvo %d", c.esperado, c.params, n)
			}
		})
	}

	t.Run("ListaActivoInvalido", func(t *testing.T) {
		VerificarError(t, Peticion("GET", ruta+"/?activo=si", ""), http.StatusBadRequest, "el parametro activo debe ser true o false")
	})

	casosPorId := []struct {
		nombre string
		ruta   string
		estado int
	}{
		{"PorIdActivoPorDefecto", ruta + "/" + activo, http.StatusOK},
		{"PorIdInactivoPorDefecto", ruta + "/" + inactivo, http.StatusNotFound},
		{"PorIdInactivoConActivoFalse", ruta + "/" + inactivo + "?activo=false", http.StatusOK},
		{"PorIdActivoConActivoFalse", ruta + "/" + activo + "?activo=false", http.StatusNotFound},
	}
	for _, c := range casosPorId {
		t.Run(c.nombre, func(t *testing.T) {
			w := Peticion("GET", c.ruta, "")
			if c.estado == http.StatusOK {
				VerificarRespuesta(t, w, http.StatusOK)
			} else {
				VerificarError(t, w, c.estado, "registro no encontrado")
			}
		})
	}

	t.Run("PorIdActivoInvalido", func(t *testing.T) {
		VerificarError(t, Peticion("GET", ruta+"/"+activo+"?activo=si", ""), http.StatusNotFound, "el parametro activo debe ser true o false")
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
