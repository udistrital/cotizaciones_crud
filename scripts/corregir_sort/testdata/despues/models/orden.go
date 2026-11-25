package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// ErrSortInvalido indica que los parámetros sortby u order de una consulta GetAll son
// inválidos: un campo que no existe en el modelo, un order distinto de asc o desc, o
// cantidades incompatibles. Los controladores lo traducen a HTTP 400.
var ErrSortInvalido = errors.New("parámetro de ordenamiento inválido")

// construirOrden retorna las expresiones ORDER BY de Beego para los campos de sortby.
// Si order está vacío, cada campo se ordena ascendente. Si order tiene un solo valor,
// aplica a todos los campos de sortby. Sin sortby no se ordena. Todo error de
// validación envuelve ErrSortInvalido.
func construirOrden(modelo reflect.Type, sortby, order []string) ([]string, error) {
	if len(sortby) == 0 {
		if len(order) != 0 {
			return nil, fmt.Errorf("%w: se envió order sin sortby", ErrSortInvalido)
		}
		return nil, nil
	}
	if len(order) == 0 {
		order = []string{"asc"}
	}
	if len(order) != 1 && len(order) != len(sortby) {
		return nil, fmt.Errorf("%w: sortby y order deben tener la misma cantidad, o order un solo valor", ErrSortInvalido)
	}
	var expresiones []string
	for i, campo := range sortby {
		if err := validarCampoOrden(modelo, campo); err != nil {
			return nil, err
		}
		sentido := order[0]
		if len(order) > 1 {
			sentido = order[i]
		}
		switch sentido {
		case "asc":
			expresiones = append(expresiones, campo)
		case "desc":
			expresiones = append(expresiones, "-"+campo)
		default:
			return nil, fmt.Errorf("%w: order %q debe ser asc o desc", ErrSortInvalido, sentido)
		}
	}
	return expresiones, nil
}

// validarCampoOrden verifica que campo sea un campo del modelo, por nombre de campo o de
// columna. Un campo de relación se recorre con "__" (por ejemplo MstCotizacionId__Id) y
// cada segmento se valida contra el modelo relacionado.
func validarCampoOrden(modelo reflect.Type, campo string) error {
	if modelo.Kind() != reflect.Struct {
		return fmt.Errorf("%w: %q no se puede ordenar", ErrSortInvalido, campo)
	}
	nombre, resto, anidado := strings.Cut(campo, "__")
	f, ok := buscarCampoOrden(modelo, nombre)
	if !ok {
		return fmt.Errorf("%w: %q no es un campo de %s", ErrSortInvalido, campo, modelo.Name())
	}
	if !anidado {
		return nil
	}
	tipo := f.Type
	for tipo.Kind() == reflect.Ptr {
		tipo = tipo.Elem()
	}
	if tipo.Kind() != reflect.Struct {
		return fmt.Errorf("%w: %q no es una relación", ErrSortInvalido, nombre)
	}
	return validarCampoOrden(tipo, resto)
}

// buscarCampoOrden retorna el campo del struct cuyo nombre o columna orm es nombre.
func buscarCampoOrden(modelo reflect.Type, nombre string) (reflect.StructField, bool) {
	if nombre == "" {
		return reflect.StructField{}, false
	}
	for i := 0; i < modelo.NumField(); i++ {
		f := modelo.Field(i)
		if f.Name == nombre || columnaORM(f.Tag.Get("orm")) == nombre {
			return f, true
		}
	}
	return reflect.StructField{}, false
}

// columnaORM retorna el valor de column(...) en una etiqueta orm, o "" si no lo tiene.
func columnaORM(etiqueta string) string {
	for _, p := range strings.Split(etiqueta, ";") {
		if strings.HasPrefix(p, "column(") && strings.HasSuffix(p, ")") {
			return p[len("column(") : len(p)-1]
		}
	}
	return ""
}
