// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package helpers

import (
	"net/http"
	"strconv"
	"testing"
)

// IdsNoNumericos son identificadores de ruta que no son enteros.
var IdsNoNumericos = []string{"abc", "12abc"}

// ProbarIdNoNumerico comprueba que la petición (GET, PUT o DELETE) responda 404 ante
// un Id de ruta que no es numérico y que se rechace sin consultar la base de datos (el
// Id no debe convertirse en 0). En la capa de integración la verificación de que no
// hubo consulta no tiene efecto, porque el controlador falso no está registrado.
func ProbarIdNoNumerico(t *testing.T, metodo, ruta, cuerpo string) {
	t.Helper()
	for _, id := range IdsNoNumericos {
		t.Run(metodo+" "+id, func(t *testing.T) {
			LimpiarSentencia()
			DefinirFila(nil)
			VerificarError(t, Peticion(metodo, ruta+"/"+id, cuerpo), http.StatusNotFound)
			VerificarSinConsulta(t)
		})
	}
}

// ProbarPostRechazado comprueba que POST responda 400 con la estructura
// institucional y que la validación rechace el cuerpo sin consultar la base de datos.
func ProbarPostRechazado(t *testing.T, ruta string, cuerpo interface{}) {
	t.Helper()
	LimpiarSentencia()
	VerificarError(t, Peticion(http.MethodPost, ruta+"/", JSON(t, cuerpo)), http.StatusBadRequest)
	VerificarSinConsulta(t)
}

// ProbarPutRechazado comprueba que PUT al registro 1 responda 400 con la estructura
// institucional y que la validación rechace el cuerpo sin consultar la base de datos. El
// cuerpo lleva Id 1, igual al de la ruta, para que el rechazo no se deba a un Id distinto.
func ProbarPutRechazado(t *testing.T, ruta string, cuerpo map[string]interface{}) {
	t.Helper()
	cuerpo["Id"] = 1
	LimpiarSentencia()
	DefinirFila(nil)
	VerificarError(t, Peticion(http.MethodPut, ruta+"/1", JSON(t, cuerpo)), http.StatusBadRequest)
	VerificarSinConsulta(t)
}

// ProbarPostFKInexistente comprueba que POST responda 400 cuando la llave foránea
// campo referencia un registro que no existe, y que no se cree ningún registro.
func ProbarPostFKInexistente(t *testing.T, ruta, tabla string, cuerpo map[string]interface{}, campo string) {
	t.Helper()
	cuerpo[campo] = map[string]interface{}{"Id": IdInexistente}
	antes := Contar(t, tabla)
	VerificarError(t, Peticion(http.MethodPost, ruta+"/", JSON(t, cuerpo)), http.StatusBadRequest)
	VerificarIgual(t, "Registros en "+tabla+" tras un POST con "+campo+" inexistente", Contar(t, tabla), antes)
}

// ProbarPutFKInexistente comprueba que PUT responda 400 cuando la llave foránea
// campo referencia un registro que no existe, y que el registro no cambie.
func ProbarPutFKInexistente(t *testing.T, ruta, tabla string, id int, campo string) {
	t.Helper()
	antes := Fila(t, tabla, id)
	cuerpo := ProbarGetOne(t, ruta, id)
	delete(cuerpo, "FechaCreacion")
	delete(cuerpo, "FechaModificacion")
	cuerpo[campo] = map[string]interface{}{"Id": IdInexistente}
	VerificarError(t, Peticion(http.MethodPut, ruta+"/"+strconv.Itoa(id), JSON(t, cuerpo)), http.StatusBadRequest)
	VerificarIgual(t, "El registro no cambia tras un PUT con "+campo+" inexistente", Fila(t, tabla, id), antes)
}
