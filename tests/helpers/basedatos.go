// Package helpers contiene utilidades compartidas por las pruebas unitarias
// de los paquetes models y controllers.
package helpers

import (
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/client/orm"
	_ "modernc.org/sqlite"

	// Registra los modelos en el ORM antes de crear las tablas
	_ "github.com/udistrital/cotizaciones_crud/models"
)

// IniciarBaseDatos registra como alias "default" una base de datos SQLite
// temporal y crea en ella las tablas de todos los modelos, de modo que las
// pruebas no dependan de un servidor PostgreSQL. Retorna una función que
// elimina la base de datos al finalizar las pruebas.
//
// SQLite no aplica los índices únicos ni las llaves foráneas del script
// database/agora_cotizaciones.sql; esas restricciones se prueban contra
// PostgreSQL en las pruebas de integración.
func IniciarBaseDatos() (limpiar func(), err error) {
	dir, err := os.MkdirTemp("", "cotizaciones_crud_test")
	if err != nil {
		return nil, err
	}
	limpiar = func() { os.RemoveAll(dir) }

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
