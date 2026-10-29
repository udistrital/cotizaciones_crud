package models

import (
	"fmt"
	"reflect"
	"strings"
)

// recuperarConsulta convierte en error el panic que lanza el ORM cuando query,
// fields o sortby usan un campo inexistente. Se usa con defer en las funciones GetAll.
func recuperarConsulta(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("consulta invalida: %v", r)
	}
}

// valorFiltro retorna el valor de un filtro del parametro query con el tipo del
// campo del modelo. El ORM envia los textos sin convertir, por lo que un campo
// booleano (por ejemplo Activo o SolicitudId__Activo) se compara con true o false
// en lugar de con el texto "true" o "false", que no todos los motores convierten.
// modelo es un puntero al modelo consultado y campo usa la notacion Campo__Campo.
func valorFiltro(modelo interface{}, campo string, valor string) interface{} {
	tipo := reflect.TypeOf(modelo)
	for _, nombre := range strings.Split(campo, "__") {
		for tipo.Kind() == reflect.Ptr {
			tipo = tipo.Elem()
		}
		if tipo.Kind() != reflect.Struct {
			return valor
		}
		f, ok := tipo.FieldByNameFunc(func(n string) bool { return strings.EqualFold(n, nombre) })
		if !ok {
			return valor
		}
		tipo = f.Type
	}
	if tipo.Kind() == reflect.Bool {
		return valor == "true" || valor == "1"
	}
	return valor
}
