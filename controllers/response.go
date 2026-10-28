package controllers

import (
	"strconv"

	beeLogger "github.com/beego/bee/v2/logger"
	beego "github.com/beego/beego/v2/server/web"
)

// APIResponse es la estructura de todas las respuestas del API. Status es texto
// para coincidir con las respuestas de error de customerror (utils_oas).
type APIResponse struct {
	Success bool        `json:"Success"`
	Status  string      `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

// APIResponseLista es la respuesta de las consultas GetAll. Total es la cantidad
// de registros que cumplen los filtros, sin aplicar limit ni offset; Data
// contiene solo los registros de la pagina solicitada.
type APIResponseLista struct {
	Success bool        `json:"Success"`
	Status  string      `json:"Status"`
	Message string      `json:"Message"`
	Total   int64       `json:"Total"`
	Data    interface{} `json:"Data"`
}

func writeSuccessLista(ctrl *beego.Controller, message string, total int64, data interface{}) {
	ctrl.Ctx.Output.SetStatus(200)
	ctrl.Data["json"] = APIResponseLista{
		Success: true,
		Status:  "200",
		Message: message,
		Total:   total,
		Data:    data,
	}
	ctrl.ServeJSON()
}

func writeSuccess(ctrl *beego.Controller, status int, message string, data interface{}) {
	ctrl.Ctx.Output.SetStatus(status)
	ctrl.Data["json"] = APIResponse{
		Success: true,
		Status:  strconv.Itoa(status),
		Message: message,
		Data:    data,
	}
	ctrl.ServeJSON()
}

// abortWithError registra el error y delega la respuesta en CustomErrorController,
// que arma el JSON con c.Data["message"] y el codigo del Abort. Solo hay
// respuesta para los codigos que define customerror: 400, 401, 404, 500, 501, 502 y 509.
func abortWithError(ctrl *beego.Controller, status int, message string, err error) {
	if err != nil {
		beeLogger.Log.Error(message + ": " + err.Error())
	} else {
		beeLogger.Log.Error(message)
	}

	ctrl.Data["message"] = message
	ctrl.Abort(strconv.Itoa(status))
}
