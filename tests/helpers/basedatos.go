// Package helpers contiene utilidades compartidas por las pruebas unitarias
// de los paquetes models y controllers.
package helpers

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"

	// Registra los modelos en el ORM antes de crear las tablas
	_ "github.com/udistrital/cotizaciones_crud/models"
)

// VariablePostgres activa las pruebas de integración: con el valor true las
// pruebas usan una base de datos PostgreSQL temporal en lugar de SQLite. La
// conexión se toma de las mismas variables COTIZACIONES_CRUD_PG* del API.
const VariablePostgres = "COTIZACIONES_CRUD_PRUEBAS_POSTGRES"

// UsaPostgres indica si las pruebas se ejecutan contra PostgreSQL.
func UsaPostgres() bool {
	return os.Getenv(VariablePostgres) == "true"
}

// IniciarBaseDatos registra como alias "default" la base de datos de pruebas y
// retorna una función que la elimina al finalizar.
//
// Por defecto usa SQLite, que no requiere un servidor pero no aplica los índices
// únicos ni las llaves foráneas del script database/agora_cotizaciones.sql. Con
// COTIZACIONES_CRUD_PRUEBAS_POSTGRES=true crea una base PostgreSQL temporal con
// ese script, para validar las restricciones de integridad reales.
func IniciarBaseDatos() (limpiar func(), err error) {
	if UsaPostgres() {
		return iniciarPostgres()
	}
	return iniciarSQLite()
}

func iniciarSQLite() (limpiar func(), err error) {
	dir, err := os.MkdirTemp("", "cotizaciones_crud_test")
	if err != nil {
		return nil, err
	}
	limpiar = func() { _ = os.RemoveAll(dir) }

	if err = orm.RegisterDriver("sqlite", orm.DRSqlite); err != nil {
		limpiar()
		return nil, err
	}
	dsn := filepath.Join(dir, "pruebas.db")
	if err = orm.RegisterDataBase("default", "sqlite", dsn, orm.MaxOpenConnections(1)); err != nil {
		limpiar()
		return nil, err
	}
	if err = orm.RunSyncdb("default", true, false); err != nil {
		limpiar()
		return nil, err
	}
	return limpiar, nil
}

// iniciarPostgres crea una base de datos con nombre único en el servidor
// configurado, ejecuta en ella database/agora_cotizaciones.sql y la elimina al
// finalizar. La base de datos del API (COTIZACIONES_CRUD_PGDB) no se modifica.
func iniciarPostgres() (limpiar func(), err error) {
	admin, err := sql.Open("postgres", dsnPostgres(os.Getenv("COTIZACIONES_CRUD_PGDB"), ""))
	if err != nil {
		return nil, err
	}
	// En CI el servidor se inicia junto con las pruebas y puede tardar en aceptar conexiones
	for intento := 0; intento < 30; intento++ {
		if err = admin.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		_ = admin.Close()
		return nil, fmt.Errorf("no se pudo conectar al servidor PostgreSQL: %w", err)
	}
	nombre := fmt.Sprintf("cotizaciones_crud_pruebas_%d", time.Now().UnixNano())
	// template0 es la plantilla vacía del servidor; no depende de objetos ni de la
	// versión de collation con que se haya creado template1
	if _, err = admin.Exec("CREATE DATABASE " + nombre + " TEMPLATE template0"); err != nil {
		_ = admin.Close()
		return nil, fmt.Errorf("no se pudo crear la base de datos de pruebas: %w", err)
	}
	limpiar = func() {
		if db, err := orm.GetDB("default"); err == nil {
			_ = db.Close()
		}
		if _, err := admin.Exec("DROP DATABASE IF EXISTS " + nombre + " WITH (FORCE)"); err != nil {
			fmt.Println("No se pudo eliminar la base de datos de pruebas", nombre+":", err)
		}
		_ = admin.Close()
	}

	if err = ejecutarScript(dsnPostgres(nombre, "")); err != nil {
		limpiar()
		return nil, err
	}
	if err = orm.RegisterDataBase("default", "postgres", dsnPostgres(nombre, os.Getenv("COTIZACIONES_CRUD_PGSCHEMA"))); err != nil {
		limpiar()
		return nil, err
	}
	return limpiar, nil
}

// ejecutarScript crea el esquema de la base de datos con database/agora_cotizaciones.sql.
func ejecutarScript(dsn string) error {
	_, archivo, _, _ := runtime.Caller(0)
	ruta := filepath.Join(filepath.Dir(archivo), "..", "..", "database", "agora_cotizaciones.sql")
	script, err := os.ReadFile(ruta)
	if err != nil {
		return err
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()
	if _, err = db.Exec(string(script)); err != nil {
		return fmt.Errorf("no se pudo ejecutar %s: %w", ruta, err)
	}
	return nil
}

// dsnPostgres arma la cadena de conexión con las variables COTIZACIONES_CRUD_PG*.
// esquema, si no está vacío, se usa como search_path.
//
// La sesión usa la zona horaria UTC, como el servidor de producción: el ORM
// convierte las fechas a la zona de la sesión, pero lib/pq lee las columnas
// timestamp without time zone como UTC, por lo que con otra zona (por ejemplo
// America/Bogota en un servidor local) las fechas leídas quedarían desplazadas.
func dsnPostgres(baseDatos, esquema string) string {
	parametros := url.Values{"sslmode": {"disable"}, "timezone": {"UTC"}}
	if esquema != "" {
		parametros.Set("search_path", esquema)
	}
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(os.Getenv("COTIZACIONES_CRUD_PGUSER"), os.Getenv("COTIZACIONES_CRUD_PGPASS")),
		Host:     os.Getenv("COTIZACIONES_CRUD_PGHOST") + ":" + os.Getenv("COTIZACIONES_CRUD_PGPORT"),
		Path:     baseDatos,
		RawQuery: parametros.Encode(),
	}
	return u.String()
}
