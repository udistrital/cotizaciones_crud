// Código generado por scripts/generar_pruebas. NO EDITAR.
// Para regenerar: go run ./scripts/generar_pruebas

package controladores

import (
	"log"
	"os"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/udistrital/cotizaciones_crud/routers"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
	"github.com/udistrital/utils_oas/v2/customerror"
)

// TestMain prepara el entorno de las pruebas de esta carpeta.
func TestMain(m *testing.M) {
	if err := helpers.ConectarFalso(); err != nil {
		log.Fatal("No se pudo registrar la base de datos falsa: ", err)
	}
	// Equivalen a copyrequestbody = true y autorender = false de conf/app.conf
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	// Mismas respuestas de error (Success/Status/Message/Data) que en main.go
	beego.ErrorController(&customerror.CustomErrorController{})
	os.Exit(m.Run())
}
