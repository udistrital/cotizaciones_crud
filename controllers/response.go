package controllers

import (
	"strconv"

	beeLogger "github.com/beego/bee/v2/logger"
	beego "github.com/beego/beego/v2/server/web"
)

type APIResponse struct {
	Success bool        `json:"Success"`
	Status  string      `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
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

func abortWithError(ctrl *beego.Controller, status int, message string, err error) {
	if err != nil {
		beeLogger.Log.Error(message + ": " + err.Error())
	} else {
		beeLogger.Log.Error(message)
	}

	ctrl.Data["message"] = message
	ctrl.Abort(strconv.Itoa(status))
}
