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

// TestMain prepara el entorno de las pruebas: conecta la base de datos de las
// variables COTIZACIONES_CRUD_PG* y configura beego igual que main.go.
func TestMain(m *testing.M) {
	if err := helpers.IniciarBaseDatos(); err != nil {
		log.Fatal("No se pudo conectar a la base de datos de pruebas: ", err)
	}
	if !helpers.BaseDisponible() {
		log.Println("AVISO: variables COTIZACIONES_CRUD_PG* no definidas; las pruebas de controladores se omiten")
	}
	// Equivalen a copyrequestbody = true y autorender = false de conf/app.conf
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	// Mismas respuestas de error (Success/Status/Message/Data) que en main.go
	beego.ErrorController(&customerror.CustomErrorController{})
	os.Exit(m.Run())
}
