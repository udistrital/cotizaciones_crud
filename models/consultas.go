package models

import "fmt"

// recuperarConsulta convierte en error el panic que lanza el ORM cuando query,
// fields o sortby usan un campo inexistente. Se usa con defer en las funciones GetAll.
func recuperarConsulta(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("consulta invalida: %v", r)
	}
}
