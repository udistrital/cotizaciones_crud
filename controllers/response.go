package controllers

import (
	"strconv"

	beeLogger "github.com/beego/bee/v2/logger"
	beego "github.com/beego/beego/v2/server/web"
)

type APIResponse struct {
	Success bool        `json:"Success"`
	Status  int         `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

func writeSuccess(ctrl *beego.Controller, status int, message string, data interface{}) {
	ctrl.Ctx.Output.SetStatus(status)
	ctrl.Data["json"] = APIResponse{
		Success: true,
		Status:  status,
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

	ctrl.Ctx.Output.SetStatus(status)
	ctrl.Data["json"] = APIResponse{
		Success: false,
		Status:  status,
		Message: message,
		Data:    nil,
	}
	ctrl.ServeJSON()
	ctrl.Abort(strconv.Itoa(status))
}
