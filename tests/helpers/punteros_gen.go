// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package helpers

// Ptr retorna un puntero al valor recibido. Construye los valores de ejemplo de
// los campos escalares opcionales (punteros) en los registros generados.
func Ptr[T any](v T) *T {
	return &v
}
