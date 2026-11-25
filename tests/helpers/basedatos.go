// Package helpers contiene utilidades compartidas por las pruebas unitarias
// de los controladores.
package helpers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"

	// Registra los modelos en el ORM
	_ "github.com/udistrital/cotizaciones_crud/models"
)

// IdInexistente es un Id que no existe en ninguna tabla.
const IdInexistente = 999999999

// Variables de entorno con la conexión a la base de datos (las mismas de
// environment.env), según el lineamiento de paso de parámetros de la OAS.
const (
	envUsuario    = "COTIZACIONES_CRUD_PGUSER"
	envClave      = "COTIZACIONES_CRUD_PGPASS"
	envServidor   = "COTIZACIONES_CRUD_PGHOST"
	envPuerto     = "COTIZACIONES_CRUD_PGPORT"
	envBaseDatos  = "COTIZACIONES_CRUD_PGDB"
	envEsquema    = "COTIZACIONES_CRUD_PGSCHEMA"
	motivoOmision = "Base de datos no configurada: cargue las variables COTIZACIONES_CRUD_PG* (environment.env) para ejecutar esta prueba"
)

var baseDisponible bool

// IniciarBaseDatos registra la base de datos PostgreSQL indicada en las
// variables de entorno. Si no están definidas, las pruebas se omiten (por
// ejemplo en el pipeline de IC, que no cuenta con base de datos).
func IniciarBaseDatos() error {
	usuario, servidor, base := os.Getenv(envUsuario), os.Getenv(envServidor), os.Getenv(envBaseDatos)
	if usuario == "" || servidor == "" || base == "" {
		return nil
	}
	conexion := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=disable",
		usuario, url.QueryEscape(os.Getenv(envClave)), servidor, os.Getenv(envPuerto), base, os.Getenv(envEsquema))
	if err := orm.RegisterDataBase("default", "postgres", conexion); err != nil {
		return err
	}
	baseDisponible = true
	return nil
}

// BaseDisponible indica si las pruebas cuentan con base de datos.
func BaseDisponible() bool {
	return baseDisponible
}

// RequerirBaseDatos omite la prueba cuando no hay base de datos configurada.
func RequerirBaseDatos(t *testing.T) {
	t.Helper()
	if !baseDisponible {
		t.Skip(motivoOmision)
	}
}

// EliminarAlFinal borra físicamente el registro al terminar la prueba, para
// no dejar datos de prueba en la base. Los registros se eliminan en orden
// inverso a su creación (primero los hijos y luego sus llaves foráneas).
func EliminarAlFinal(t *testing.T, tabla string, id int) {
	t.Helper()
	t.Cleanup(func() {
		if _, err := orm.NewOrm().Raw("DELETE FROM "+tabla+" WHERE id = ?", id).Exec(); err != nil {
			t.Errorf("No se pudo eliminar el registro de prueba %s %d: %v", tabla, id, err)
		}
	})
}

// Fila retorna el registro tal como está guardado en la base de datos.
func Fila(t *testing.T, tabla string, id int) map[string]interface{} {
	t.Helper()
	var texto string
	if err := orm.NewOrm().Raw("SELECT row_to_json(t)::text FROM "+tabla+" t WHERE id = ?", id).QueryRow(&texto); err != nil {
		t.Fatalf("No se pudo leer %s %d de la base de datos: %v", tabla, id, err)
	}
	fila := map[string]interface{}{}
	if err := json.Unmarshal([]byte(texto), &fila); err != nil {
		t.Fatal("No se pudo decodificar el registro:", err)
	}
	return fila
}

// SinColumnas retorna una copia de la fila sin las columnas indicadas.
func SinColumnas(fila map[string]interface{}, columnas ...string) map[string]interface{} {
	copia := map[string]interface{}{}
	for k, v := range fila {
		copia[k] = v
	}
	for _, c := range columnas {
		delete(copia, c)
	}
	return copia
}

// Contar retorna la cantidad de registros de la tabla.
func Contar(t *testing.T, tabla string) int {
	t.Helper()
	var total int
	if err := orm.NewOrm().Raw("SELECT count(*) FROM " + tabla).QueryRow(&total); err != nil {
		t.Fatalf("No se pudo contar los registros de %s: %v", tabla, err)
	}
	return total
}
