// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package helpers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

// FechaCliente es una fecha que el cliente envía en los campos administrados
// por el servicio (FechaCreacion/FechaModificacion) y que debe ignorarse.
const FechaCliente = "2000-01-01T00:00:00Z"

// CrearPorPost crea un registro mediante el servicio POST, verifica la
// respuesta 201 y programa su eliminación al terminar la prueba.
func CrearPorPost(t *testing.T, ruta, tabla string, cuerpo interface{}) int {
	t.Helper()
	w := Peticion(http.MethodPost, ruta+"/", JSON(t, cuerpo))
	id := VerificarExito(t, w, http.StatusCreated).IdDato(t)
	EliminarAlFinal(t, tabla, id)
	return id
}

// VerificarCamposAdministrados comprueba que, al crear, el servicio asigna
// Activo=true y las fechas de creación y modificación, ignorando los valores
// enviados por el cliente.
func VerificarCamposAdministrados(t *testing.T, tabla string, id int) {
	t.Helper()
	fila := Fila(t, tabla, id)
	if fila["activo"] != true {
		t.Error("Se esperaba activo=true al crear (aunque el cliente envíe false) y se obtuvo", fila["activo"])
	}
	for _, columna := range []string{"fecha_creacion", "fecha_modificacion"} {
		valor, _ := fila[columna].(string)
		if valor == "" || strings.HasPrefix(valor, "2000-") {
			t.Errorf("Se esperaba %s asignada por el servicio y se obtuvo %v", columna, fila[columna])
		}
	}
}

// ProbarPostInvalido comprueba que POST responda 400 ante un JSON mal formado
// o un cuerpo vacío.
func ProbarPostInvalido(t *testing.T, ruta string) {
	t.Helper()
	VerificarError(t, Peticion(http.MethodPost, ruta+"/", `{"Id":`), http.StatusBadRequest)
	VerificarError(t, Peticion(http.MethodPost, ruta+"/", ""), http.StatusBadRequest)
}

// ProbarGetOne consulta el registro por Id, verifica la respuesta 200 y
// retorna el objeto obtenido.
func ProbarGetOne(t *testing.T, ruta string, id int) map[string]interface{} {
	t.Helper()
	r := VerificarExito(t, Peticion(http.MethodGet, ruta+"/"+strconv.Itoa(id), ""), http.StatusOK)
	VerificarIgual(t, "Id del registro consultado", r.IdDato(t), id)
	return r.Dato(t)
}

// ProbarGetOneInexistente comprueba que GetOne responda 404 ante un Id que no
// existe o que no es numérico.
func ProbarGetOneInexistente(t *testing.T, ruta string) {
	t.Helper()
	VerificarError(t, Peticion(http.MethodGet, ruta+"/"+strconv.Itoa(IdInexistente), ""), http.StatusNotFound)
	VerificarError(t, Peticion(http.MethodGet, ruta+"/abc", ""), http.StatusNotFound)
}

// ConsultarGetAll ejecuta GetAll con los parámetros indicados y retorna la
// lista obtenida.
func ConsultarGetAll(t *testing.T, ruta, parametros string) []map[string]interface{} {
	t.Helper()
	return VerificarExito(t, Peticion(http.MethodGet, ruta+"/?"+parametros, ""), http.StatusOK).Lista(t)
}

// ContieneId indica si la lista contiene un objeto con el Id indicado.
func ContieneId(lista []map[string]interface{}, id int) bool {
	for _, m := range lista {
		if m["Id"] == float64(id) {
			return true
		}
	}
	return false
}

// ConsultarPagina ejecuta GetAll con los parámetros indicados y retorna la
// lista obtenida y el campo Total.
func ConsultarPagina(t *testing.T, ruta, parametros string) ([]map[string]interface{}, int) {
	t.Helper()
	r := VerificarExito(t, Peticion(http.MethodGet, ruta+"/?"+parametros, ""), http.StatusOK)
	return r.Lista(t), r.TotalLista(t)
}

// ProbarGetAll verifica el servicio GetAll: filtro por Id, lista vacía sin
// resultados, Total (registros que cumplen el filtro, no solo los de la
// página), limit/offset, orden asc/desc, fields y HTTP 400 ante filtros
// inválidos. Retorna el registro obtenido con el filtro por Id.
func ProbarGetAll(t *testing.T, ruta, tabla string, id int) map[string]interface{} {
	t.Helper()
	lista, total := ConsultarPagina(t, ruta, fmt.Sprintf("query=Id:%d", id))
	if len(lista) != 1 || !ContieneId(lista, id) {
		t.Fatalf("Se esperaba solo el registro %d y se obtuvo %v", id, lista)
	}
	VerificarIgual(t, "Total con filtro por Id", total, 1)

	l, total := ConsultarPagina(t, ruta, fmt.Sprintf("query=Id:%d", IdInexistente))
	if len(l) != 0 {
		t.Error("Se esperaba una lista vacía (Data: []) y se obtuvo", l)
	}
	VerificarIgual(t, "Total sin resultados", total, 0)

	// Paginación: Total es la cantidad de registros de la tabla aunque la
	// página traiga uno solo, y offset salta los primeros registros
	registros := Contar(t, tabla)
	asc, total := ConsultarPagina(t, ruta, "sortby=Id&order=asc&limit=1")
	if len(asc) != 1 {
		t.Fatal("Se esperaba 1 registro con limit=1 y se obtuvo", len(asc))
	}
	VerificarIgual(t, "Total con limit=1", total, registros)
	desc, _ := ConsultarPagina(t, ruta, "sortby=Id&order=desc&limit=1")
	if len(desc) != 1 {
		t.Fatal("Se esperaba 1 registro con limit=1 y se obtuvo", len(desc))
	}
	primero, ultimo := asc[0]["Id"].(float64), desc[0]["Id"].(float64)
	if registros > 1 && primero >= ultimo {
		t.Errorf("Se esperaba el menor Id con order=asc y el mayor con order=desc y se obtuvo %v y %v", primero, ultimo)
	}
	siguiente, total := ConsultarPagina(t, ruta, "sortby=Id&order=asc&limit=1&offset=1")
	VerificarIgual(t, "Total con offset=1", total, registros)
	if registros > 1 && (len(siguiente) != 1 || siguiente[0]["Id"].(float64) <= primero) {
		t.Error("Se esperaba el segundo registro con offset=1 y se obtuvo", siguiente)
	} else if registros == 1 && len(siguiente) != 0 {
		t.Error("Se esperaba una lista vacía con offset=1 y se obtuvo", siguiente)
	}

	l = ConsultarGetAll(t, ruta, fmt.Sprintf("query=Id:%d&fields=Id", id))
	if len(l) != 1 || len(l[0]) != 1 || l[0]["Id"] != float64(id) {
		t.Error("Se esperaba solo el campo Id con fields=Id y se obtuvo", l)
	}

	// Sin order, sortby ordena ascendente igual que order=asc
	sinOrder, _ := ConsultarPagina(t, ruta, "sortby=Id&limit=1")
	if len(sinOrder) != 1 {
		t.Fatal("Se esperaba 1 registro con sortby=Id sin order y se obtuvo", len(sinOrder))
	}
	VerificarIgual(t, "sortby sin order ordena como asc", sinOrder[0]["Id"], asc[0]["Id"])

	// Un campo de sortby inexistente o un order distinto de asc y desc son solicitudes inválidas
	ProbarSortInvalido(t, ruta, "sortby=campo_falso&order=asc")
	ProbarSortInvalido(t, ruta, "sortby=Id&order=ascending")

	// Filtros inválidos: HTTP 400 con la estructura institucional
	VerificarFiltroInvalido(t, ruta+"/?query=Id")
	VerificarFiltroInvalido(t, fmt.Sprintf("%s/?query=Id:%d,Activo", ruta, id))
	return lista[0]
}

// VerificarFKExpandida comprueba que la llave foránea campo venga expandida
// con los mismos datos del registro relacionado. Las llaves foráneas del
// registro relacionado solo se comparan por Id: el ORM expande hasta dos
// niveles (RelatedSel), por lo que en cadenas más largas el último nivel
// llega solo con su Id.
func VerificarFKExpandida(t *testing.T, dato map[string]interface{}, campo string, relacionado interface{}) {
	t.Helper()
	obtenido, ok := dato[campo].(map[string]interface{})
	if !ok {
		t.Errorf("Se esperaba la llave foránea %s expandida como objeto y se obtuvo %v", campo, dato[campo])
		return
	}
	VerificarIgual(t, "Llave foránea "+campo+" expandida", soloIdAnidados(obtenido), soloIdAnidados(AMapa(t, relacionado)))
}

// soloIdAnidados retorna una copia del objeto en la que cada objeto anidado
// (llave foránea) se reduce a su Id.
func soloIdAnidados(m map[string]interface{}) map[string]interface{} {
	copia := map[string]interface{}{}
	for k, v := range m {
		if anidado, ok := v.(map[string]interface{}); ok {
			copia[k] = anidado["Id"]
		} else {
			copia[k] = v
		}
	}
	return copia
}

// ProbarPut verifica el servicio PUT: actualiza el campo indicado, conserva
// fecha_creacion y usuario_creacion_id (aunque el cliente envíe otro valor)
// tras dos actualizaciones, actualiza fecha_modificacion, responde 400 ante
// Id distinto o JSON inválido y 404 ante un Id inexistente o no numérico.
func ProbarPut(t *testing.T, ruta, tabla string, id int, campo string, valor interface{}, columna string) {
	t.Helper()
	rutaId := ruta + "/" + strconv.Itoa(id)
	antes := Fila(t, tabla, id)
	cuerpo := ProbarGetOne(t, ruta, id)
	delete(cuerpo, "FechaCreacion")
	delete(cuerpo, "FechaModificacion")
	cuerpo[campo] = valor
	if campo != "UsuarioCreacionId" {
		cuerpo["UsuarioCreacionId"] = IdInexistente
	}

	for i := 0; i < 2; i++ {
		VerificarExito(t, Peticion(http.MethodPut, rutaId, JSON(t, cuerpo)), http.StatusOK)
	}
	despues := Fila(t, tabla, id)
	VerificarIgual(t, "Columna "+columna+" actualizada", fmt.Sprint(despues[columna]), fmt.Sprint(valor))
	VerificarIgual(t, "fecha_creacion se conserva tras el PUT", despues["fecha_creacion"], antes["fecha_creacion"])
	VerificarIgual(t, "usuario_creacion_id se conserva tras el PUT", despues["usuario_creacion_id"], antes["usuario_creacion_id"])
	if despues["fecha_modificacion"] == nil || despues["fecha_modificacion"] == antes["fecha_modificacion"] {
		t.Error("Se esperaba fecha_modificacion actualizada y se obtuvo", despues["fecha_modificacion"])
	}

	cuerpo["Id"] = id + 1
	VerificarError(t, Peticion(http.MethodPut, rutaId, JSON(t, cuerpo)), http.StatusBadRequest)
	delete(cuerpo, "Id")
	VerificarError(t, Peticion(http.MethodPut, ruta+"/"+strconv.Itoa(IdInexistente), JSON(t, cuerpo)), http.StatusNotFound)
	ProbarIdNoNumerico(t, http.MethodPut, ruta, JSON(t, cuerpo))
	VerificarError(t, Peticion(http.MethodPut, rutaId, `{"Id":"texto"}`), http.StatusBadRequest)
}

// ProbarDelete verifica el borrado lógico: responde 200, deja activo=false,
// actualiza fecha_modificacion, no modifica las demás columnas, el registro
// sigue consultable y un Id inexistente o no numérico responde 404.
func ProbarDelete(t *testing.T, ruta, tabla string, id int) {
	t.Helper()
	antes := Fila(t, tabla, id)
	VerificarExito(t, Peticion(http.MethodDelete, ruta+"/"+strconv.Itoa(id), ""), http.StatusOK)
	despues := Fila(t, tabla, id)

	if despues["activo"] != false {
		t.Error("Se esperaba activo=false (borrado lógico) y se obtuvo", despues["activo"])
	}
	if despues["fecha_modificacion"] == nil || despues["fecha_modificacion"] == antes["fecha_modificacion"] {
		t.Error("Se esperaba fecha_modificacion actualizada y se obtuvo", despues["fecha_modificacion"])
	}
	VerificarIgual(t, "Las demás columnas no cambian con el borrado lógico",
		SinColumnas(despues, "activo", "fecha_modificacion"), SinColumnas(antes, "activo", "fecha_modificacion"))

	if dato := ProbarGetOne(t, ruta, id); dato["Activo"] != false {
		t.Error("Se esperaba el registro consultable con Activo=false y se obtuvo", dato["Activo"])
	}
	VerificarError(t, Peticion(http.MethodDelete, ruta+"/"+strconv.Itoa(IdInexistente), ""), http.StatusNotFound)
	ProbarIdNoNumerico(t, http.MethodDelete, ruta, "")
}
