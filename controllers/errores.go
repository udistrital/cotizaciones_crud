package controllers

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"net"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/lib/pq"
)

// abortarSiBaseNoDisponible responde 500 cuando err indica que la base de datos no está
// disponible (caída, sin conexión o sin recursos). Los demás errores (registro
// inexistente, datos inválidos, restricciones) no se modifican: el llamador responde
// 400 o 404 como antes.
func abortarSiBaseNoDisponible(c *beego.Controller, servicio string, err error) {
	if baseNoDisponible(err) {
		c.Data["message"] = "Error service " + servicio + ": The database is not available"
		c.Abort("500")
	}
}

// baseNoDisponible indica si err es un error de conexión con la base de datos y no de los
// datos de la petición.
func baseNoDisponible(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, driver.ErrBadConn) || errors.Is(err, sql.ErrConnDone) ||
		errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) ||
		errors.Is(err, context.DeadlineExceeded) {
		return true
	}
	// Conexión rechazada, cortada o sin respuesta del servidor
	var errRed net.Error
	if errors.As(err, &errRed) {
		return true
	}
	// Clases SQLSTATE de PostgreSQL: 08 excepción de conexión, 53 recursos insuficientes,
	// 57 intervención del operador (p. ej. la base se está apagando), 58 error del sistema
	// y XX error interno
	var errPq *pq.Error
	if errors.As(err, &errPq) && len(errPq.SQLState()) >= 2 {
		switch errPq.SQLState()[:2] {
		case "08", "53", "57", "58", "XX":
			return true
		}
	}
	return false
}
