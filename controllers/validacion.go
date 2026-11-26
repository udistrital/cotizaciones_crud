package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// validarObligatorios comprueba que el cuerpo JSON incluya cada campo indicado (columnas
// NOT NULL que no son llave foránea) con un valor: el campo no puede faltar, ser null ni,
// si es texto, estar vacío o tener solo espacios. Sin esta validación, json.Unmarshal deja
// el valor cero de Go ("", 0, false o fecha cero) y la base de datos lo acepta.
// Los valores 0 y false enviados explícitamente son válidos.
func validarObligatorios(cuerpo []byte, campos ...string) error {
	var recibido map[string]json.RawMessage
	if err := json.Unmarshal(cuerpo, &recibido); err != nil {
		return err
	}
	for _, campo := range campos {
		valor, ok := buscarCampo(recibido, campo)
		if !ok || bytes.Equal(bytes.TrimSpace(valor), []byte("null")) {
			return fmt.Errorf("the field %s is required and cannot be null", campo)
		}
		var texto string
		if json.Unmarshal(valor, &texto) == nil && strings.TrimSpace(texto) == "" {
			return fmt.Errorf("the field %s cannot be empty", campo)
		}
	}
	return nil
}

// buscarCampo retorna el valor del campo en el cuerpo. Como json.Unmarshal, prefiere la
// coincidencia exacta y, si no existe, acepta el nombre sin distinguir mayúsculas.
func buscarCampo(recibido map[string]json.RawMessage, campo string) (json.RawMessage, bool) {
	if valor, ok := recibido[campo]; ok {
		return valor, true
	}
	for clave, valor := range recibido {
		if strings.EqualFold(clave, campo) {
			return valor, true
		}
	}
	return nil, false
}
