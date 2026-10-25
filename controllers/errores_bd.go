package controllers

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/lib/pq"
)

var (
	// Detail de Postgres en unicidad y llave foranea, por ejemplo
	// "Key (solicitud_id, proveedor_id)=(5, 10) already exists."
	reLlaveDetalle = regexp.MustCompile(`Key \((.+)\)=\((.*)\)`)
	// Error del ORM cuando una relacion obligatoria llega vacia, por ejemplo
	// "field `.../models.SolicitudItem.SolicitudId` cannot be NULL"
	reCampoOrmNulo = regexp.MustCompile("field `[^`]*\\.(\\w+)` cannot be NULL")
)

// mensajeErrorGenerico se responde cuando la causa no es reconocida, para no exponer
// detalles internos de la base de datos. El error completo queda en el log.
const mensajeErrorGenerico = "los datos enviados no son validos"

// abortWithDBError responde un error al crear, actualizar o eliminar un registro con
// el codigo que fija el lineamiento para cada metodo (POST y PUT 400, DELETE 404) y
// la causa concreta del error.
func abortWithDBError(ctrl *beego.Controller, status int, message string, err error) {
	abortWithError(ctrl, status, message+": "+causaErrorBD(err), err)
}

// causaErrorBD describe en lenguaje claro por que la base de datos o el ORM rechazaron la operacion.
func causaErrorBD(err error) string {
	if errors.Is(err, orm.ErrNoRows) {
		return "no existe un registro con el id indicado"
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505": // unique_violation
			if m := reLlaveDetalle.FindStringSubmatch(pqErr.Detail); m != nil {
				return "ya existe un registro activo o inactivo con " + parejasLlave(m[1], m[2])
			}
			return "ya existe un registro con los mismos datos"
		case "23503": // foreign_key_violation
			if m := reLlaveDetalle.FindStringSubmatch(pqErr.Detail); m != nil {
				return "no existe el registro referenciado por " + parejasLlave(m[1], m[2])
			}
			return "el registro referenciado no existe"
		case "23502": // not_null_violation
			return fmt.Sprintf("el campo %s es obligatorio", pqErr.Column)
		case "22001": // string_data_right_truncation
			return "un texto supera la longitud maxima permitida (" + strings.TrimPrefix(pqErr.Message, "value too long for type ") + ")"
		case "22003": // numeric_value_out_of_range
			return "un valor numerico esta fuera del rango permitido"
		}
		return mensajeErrorGenerico
	}

	if m := reCampoOrmNulo.FindStringSubmatch(err.Error()); m != nil {
		return fmt.Sprintf("el campo %s es obligatorio y debe incluir un Id valido", m[1])
	}
	return mensajeErrorGenerico
}

// parejasLlave une columnas y valores del Detail de Postgres, por ejemplo
// "solicitud_id = 7, numero_necesidad = 123". Si un valor contiene comas
// y no se pueden emparejar, los muestra por separado.
func parejasLlave(columnas, valores string) string {
	cols := strings.Split(columnas, ", ")
	vals := strings.Split(valores, ", ")
	if len(cols) != len(vals) {
		return fmt.Sprintf("(%s) = (%s)", columnas, valores)
	}
	parejas := make([]string, len(cols))
	for i := range cols {
		parejas[i] = cols[i] + " = " + vals[i]
	}
	return strings.Join(parejas, ", ")
}
