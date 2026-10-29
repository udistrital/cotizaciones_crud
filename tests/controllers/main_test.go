package controllers

import (
	"log"
	"os"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/udistrital/utils_oas/v2/customerror"

	// Registra las rutas /v1/... de los controladores
	_ "github.com/udistrital/cotizaciones_crud/routers"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestMain(m *testing.M) {
	limpiar, err := helpers.IniciarBaseDatos()
	if err != nil {
		log.Fatal("No se pudo iniciar la base de datos de pruebas: ", err)
	}
	// Equivalen a copyrequestbody = true y autorender = false de conf/app.conf
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	// Mismo controlador de errores que registra main.go
	beego.ErrorController(&customerror.CustomErrorController{})
	codigo := m.Run()
	limpiar()
	os.Exit(codigo)
}
