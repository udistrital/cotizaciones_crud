package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// dirFixtures contiene el código antes (HEAD, sin corregir) y después (corregido a mano,
// usado como oráculo) de notificacion.
const dirFixtures = "testdata"

// copiarArbol copia el árbol src dentro de dst.
func copiarArbol(t *testing.T, src, dst string) {
	t.Helper()
	err := filepath.WalkDir(src, func(ruta string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(src, ruta)
		b, err := os.ReadFile(ruta)
		if err != nil {
			return err
		}
		destino := filepath.Join(dst, rel)
		if err := os.MkdirAll(filepath.Dir(destino), 0o755); err != nil {
			return err
		}
		return os.WriteFile(destino, b, 0o644)
	})
	if err != nil {
		t.Fatal("No se pudo copiar el fixture:", err)
	}
}

// recolector guarda las advertencias que emite el script.
type recolector struct{ avisos []string }

func (r *recolector) advertir(formato string, args ...interface{}) {
	r.avisos = append(r.avisos, fmt.Sprintf(formato, args...))
}

func leerArchivo(t *testing.T, ruta string) []byte {
	t.Helper()
	b, err := os.ReadFile(ruta)
	if err != nil {
		t.Fatal("No se pudo leer:", err)
	}
	return b
}

// TestCorregirReproduceNotificacion: aplicado al código original de notificacion, el
// script produce el mismo código que se corrigió a mano. Es el oráculo del script.
func TestCorregirReproduceNotificacion(t *testing.T) {
	raiz := t.TempDir()
	copiarArbol(t, filepath.Join(dirFixtures, "antes"), raiz)
	var r recolector
	if err := corregirRaiz(raiz, r.advertir); err != nil {
		t.Fatal("El script falló:", err)
	}
	if len(r.avisos) != 0 {
		t.Errorf("No se esperaban advertencias y se obtuvieron: %v", r.avisos)
	}
	for _, ruta := range []string{"models/notificacion.go", "controllers/notificacion.go", "models/orden.go"} {
		esperado := leerArchivo(t, filepath.Join(dirFixtures, "despues", ruta))
		obtenido := leerArchivo(t, filepath.Join(raiz, ruta))
		if !bytes.Equal(obtenido, esperado) {
			t.Errorf("%s no coincide con el código corregido.\n--- obtenido ---\n%s", ruta, obtenido)
		}
	}
}

// TestCorregirEsIdempotente: una segunda ejecución no modifica ningún archivo.
func TestCorregirEsIdempotente(t *testing.T) {
	raiz := t.TempDir()
	copiarArbol(t, filepath.Join(dirFixtures, "antes"), raiz)
	if err := corregirRaiz(raiz, (&recolector{}).advertir); err != nil {
		t.Fatal("La primera ejecución falló:", err)
	}
	primera := leerArchivo(t, filepath.Join(raiz, "models/notificacion.go"))
	primeraCtl := leerArchivo(t, filepath.Join(raiz, "controllers/notificacion.go"))

	var r recolector
	if err := corregirRaiz(raiz, r.advertir); err != nil {
		t.Fatal("La segunda ejecución falló:", err)
	}
	if !bytes.Equal(primera, leerArchivo(t, filepath.Join(raiz, "models/notificacion.go"))) {
		t.Error("La segunda ejecución cambió models/notificacion.go")
	}
	if !bytes.Equal(primeraCtl, leerArchivo(t, filepath.Join(raiz, "controllers/notificacion.go"))) {
		t.Error("La segunda ejecución cambió controllers/notificacion.go")
	}
	if len(r.avisos) != 0 {
		t.Errorf("La segunda ejecución no debía advertir y advirtió: %v", r.avisos)
	}
}

// TestCorregirOmiteNoReconocido: un GetAll sin el bloque de ordenamiento esperado se
// advierte y su archivo no cambia.
func TestCorregirOmiteNoReconocido(t *testing.T) {
	raiz := t.TempDir()
	const original = `package models

type Otro struct {
	Id int ` + "`orm:\"column(id);pk\"`" + `
}

func GetAllOtro(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (ml []interface{}, total int64, err error) {
	return nil, 0, nil
}
`
	if err := os.MkdirAll(filepath.Join(raiz, "models"), 0o755); err != nil {
		t.Fatal(err)
	}
	ruta := filepath.Join(raiz, "models", "otro.go")
	if err := os.WriteFile(ruta, []byte(original), 0o644); err != nil {
		t.Fatal(err)
	}
	var r recolector
	if err := corregirRaiz(raiz, r.advertir); err != nil {
		t.Fatal("El script falló:", err)
	}
	if !bytes.Equal(leerArchivo(t, ruta), []byte(original)) {
		t.Error("El script modificó un GetAll que no reconoce")
	}
	if !strings.Contains(strings.Join(r.avisos, "\n"), "GetAllOtro") {
		t.Errorf("Se esperaba una advertencia que nombre GetAllOtro y se obtuvo: %v", r.avisos)
	}
}

// TestCorregirNoSobrescribeOrden: si models/orden.go ya existe, el script no lo reemplaza.
func TestCorregirNoSobrescribeOrden(t *testing.T) {
	raiz := t.TempDir()
	copiarArbol(t, filepath.Join(dirFixtures, "antes"), raiz)
	propio := []byte("package models\n\n// archivo propio del repositorio\n")
	if err := os.WriteFile(filepath.Join(raiz, "models", "orden.go"), propio, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := corregirRaiz(raiz, (&recolector{}).advertir); err != nil {
		t.Fatal("El script falló:", err)
	}
	if !bytes.Equal(leerArchivo(t, filepath.Join(raiz, "models", "orden.go")), propio) {
		t.Error("El script sobrescribió models/orden.go existente")
	}
}
