// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package helpers

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/udistrital/cotizaciones_crud/models"
)

// VerificarSQL comprueba que la última sentencia SQL que recibió el controlador falso
// contenga el fragmento indicado, por ejemplo el ORDER BY o el LIMIT esperados.
func VerificarSQL(t *testing.T, fragmento string) {
	t.Helper()
	if sql := UltimaSentencia().SQL; !strings.Contains(sql, fragmento) {
		t.Errorf("Se esperaba %q en la sentencia SQL y se obtuvo: %s", fragmento, sql)
	}
}

// VerificarSinConsulta comprueba que la petición no ejecutó ninguna sentencia SQL.
// Una validación de parámetros debe rechazar la petición antes de consultar la base.
func VerificarSinConsulta(t *testing.T) {
	t.Helper()
	if sql := UltimaSentencia().SQL; sql != "" {
		t.Errorf("Se esperaba que la validación rechazara la petición sin consultar la base y se ejecutó: %s", sql)
	}
}

// ConsultarOrdenado ejecuta GetAll con los parámetros indicados, verifica la respuesta
// 200 y que la sentencia SQL contenga el fragmento. Retorna la lista y el Total.
func ConsultarOrdenado(t *testing.T, ruta, parametros, fragmento string) ([]map[string]interface{}, int) {
	t.Helper()
	LimpiarSentencia()
	lista, total := ConsultarPagina(t, ruta, parametros)
	VerificarSQL(t, fragmento)
	return lista, total
}

// ProbarSortInvalido verifica que GetAll rechace la petición con 400 institucional
// (Data nulo) y sin consultar la base de datos. Sirve para campos de sortby
// inexistentes y para order distinto de asc o desc.
func ProbarSortInvalido(t *testing.T, ruta, parametros string) {
	t.Helper()
	LimpiarSentencia()
	r := VerificarError(t, Peticion(http.MethodGet, ruta+"/?"+parametros, ""), http.StatusBadRequest)
	if r.Data != nil {
		t.Error("Se esperaba Data nulo y se obtuvo", r.Data)
	}
	VerificarSinConsulta(t)
}

// LlamarSinPanico ejecuta f y reporta la prueba como fallida si provoca un pánico. Retorna
// el error que devuelve f.
func LlamarSinPanico(t *testing.T, f func() error) error {
	t.Helper()
	var err error
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Se esperaba un error y la función provocó un pánico: %v", r)
			}
		}()
		err = f()
	}()
	return err
}

// VerificarErrorSortInvalido comprueba que err envuelva models.ErrSortInvalido y, si campo
// no está vacío, que el mensaje lo nombre.
func VerificarErrorSortInvalido(t *testing.T, err error, campo string) {
	t.Helper()
	if err == nil {
		t.Fatal("Se esperaba un error de ordenamiento y se obtuvo nil")
	}
	if !errors.Is(err, models.ErrSortInvalido) {
		t.Errorf("Se esperaba un error que envuelva ErrSortInvalido y se obtuvo: %v", err)
	}
	if campo != "" && !strings.Contains(err.Error(), campo) {
		t.Errorf("Se esperaba que el error nombre %q y se obtuvo: %v", campo, err)
	}
}
