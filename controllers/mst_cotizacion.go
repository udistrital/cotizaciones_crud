package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/utils_oas/time_bogota"

	beeLogger "github.com/beego/bee/v2/logger"
	beego "github.com/beego/beego/v2/server/web"
)

// MstCotizacionController operations for MstCotizacion
type MstCotizacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *MstCotizacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create MstCotizacion
// @Param	body		body 	models.MstCotizacion	true		"body for MstCotizacion content"
// @Success 201 {int} models.MstCotizacion
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *MstCotizacionController) Post() {
	var v models.MstCotizacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		// El Id lo asigna la base de datos (serial): se ignora el enviado por el cliente
		v.Id = 0
		v.Activo = true
		v.FechaCreacion = time_bogota.Tiempo_bogota()
		v.FechaModificacion = time_bogota.Tiempo_bogota()
		if _, err := models.AddMstCotizacion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Registration successful", "Data": v}
		} else {
			beeLogger.Log.Error(err.Error())
			c.Data["message"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
			c.Abort("400")
		}
	} else {
		beeLogger.Log.Error(err.Error())
		c.Data["message"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get MstCotizacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.MstCotizacion
// @Failure 404 not found resource
// @router /:id [get]
func (c *MstCotizacionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetMstCotizacionById(id)
	if err != nil {
		beeLogger.Log.Error(err.Error())
		c.Data["message"] = "Error service GetOne: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get MstCotizacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.MstCotizacion
// @Failure 404 not found resource
// @router / [get]
func (c *MstCotizacionController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Ctx.Output.SetStatus(400)
				c.Data["json"] = map[string]interface{}{"Success": false, "Status": "400", "Message": "Error service GetAll: invalid query parameter. Expected format: query=campo:valor,campo2:valor2", "Data": nil}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, total, err := models.GetAllMstCotizacion(query, fields, sortby, order, offset, limit)
	if err != nil {
		beeLogger.Log.Error(err.Error())
		c.Data["message"] = "Error service GetAll: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		if l == nil {
			l = []interface{}{}
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": l, "Total": total}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the MstCotizacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.MstCotizacion	true		"body for MstCotizacion content"
// @Success 200 {object} models.MstCotizacion
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *MstCotizacionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.MstCotizacion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.Id != id {
			c.Data["message"] = "Error service Put: The id cannot be modified"
			c.Abort("400")
		}
		if v.Id == 0 {
			v.Id = id
		}
		v.FechaModificacion = time_bogota.Tiempo_bogota()
		if err := models.UpdateMstCotizacionById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Update successful", "Data": v}
		} else {
			beeLogger.Log.Error(err.Error())
			c.Data["message"] = "Error service Put: The request contains an incorrect data type or an invalid parameter"
			c.Abort("400")
		}
	} else {
		beeLogger.Log.Error(err.Error())
		c.Data["message"] = "Error service Put: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the MstCotizacion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *MstCotizacionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetMstCotizacionById(id)
	if err != nil {
		beeLogger.Log.Error(err.Error())
		c.Data["message"] = "Error service Delete: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	}
	v.Activo = false
	v.FechaModificacion = time_bogota.Tiempo_bogota()
	// Borrado lógico: solo cambian activo y fecha_modificacion
	if err := models.UpdateMstCotizacionById(v, "Activo", "FechaModificacion"); err == nil {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Update successful", "Data": v}
	} else {
		beeLogger.Log.Error(err.Error())
		c.Data["message"] = "Error service Put: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}
