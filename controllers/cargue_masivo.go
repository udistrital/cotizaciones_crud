package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/udistrital/cotizaciones_crud/models"

	beego "github.com/beego/beego/v2/server/web"
)

// CargueMasivoController operations for CargueMasivo
type CargueMasivoController struct {
	beego.Controller
}

// URLMapping ...
func (c *CargueMasivoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create CargueMasivo
// @Param	body		body 	models.CargueMasivo	true		"body for CargueMasivo content"
// @Success 201 {object} controllers.APIResponse
// @Failure 400 invalid body
// @router / [post]
func (c *CargueMasivoController) Post() {
	var v models.CargueMasivo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		abortWithError(&c.Controller, 400, "cuerpo de la solicitud invalido", err)
		return
	}

	if _, err := models.AddCargueMasivo(&v); err != nil {
		abortWithError(&c.Controller, 500, "error al crear el registro", err)
		return
	}

	writeSuccess(&c.Controller, 201, "registro creado", v)
}

// GetOne ...
// @Title Get One
// @Description get CargueMasivo by id
// @Param	id		path 	string	true		"The id"
// @Success 200 {object} controllers.APIResponse
// @Failure 400 id invalido
// @router /:id [get]
func (c *CargueMasivoController) GetOne() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		abortWithError(&c.Controller, 400, "el parametro id debe ser numerico", err)
		return
	}

	v, err := models.GetCargueMasivoById(id)
	if err != nil {
		abortWithError(&c.Controller, 404, "registro no encontrado", err)
		return
	}

	writeSuccess(&c.Controller, 200, "consulta exitosa", v)
}

// GetAll ...
// @Title Get All
// @Description get CargueMasivo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Param	vigencia	query	int	false	"Vigencia de la solicitud. e.g. 2026"
// @Param	estado	query	int	false	"Id del estado de la solicitud (parametros)"
// @Param	dependencia	query	int	false	"Id de la dependencia solicitante (Oikos)"
// @Success 200 {object} controllers.APIResponse
// @Failure 400 parametros invalidos
// @router / [get]
func (c *CargueMasivoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	query := make(map[string]string)
	limit := int64(10)
	offset := int64(0)

	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}

	if v := c.GetString("limit"); v != "" {
		parsedLimit, err := strconv.ParseInt(v, 10, 64)
		if err != nil || parsedLimit < 0 {
			abortWithError(&c.Controller, 400, "el parametro limit debe ser un entero mayor o igual a cero", err)
			return
		}
		limit = parsedLimit
	}

	if v := c.GetString("offset"); v != "" {
		parsedOffset, err := strconv.ParseInt(v, 10, 64)
		if err != nil || parsedOffset < 0 {
			abortWithError(&c.Controller, 400, "el parametro offset debe ser un entero mayor o igual a cero", err)
			return
		}
		offset = parsedOffset
	}

	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}

	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}

	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				abortWithError(&c.Controller, 400, "query invalida, use formato campo:valor", nil)
				return
			}
			query[kv[0]] = kv[1]
		}
	}

	if err := agregarFiltrosSolicitud(&c.Controller, query, "SolicitudId."); err != nil {
		abortWithError(&c.Controller, 400, err.Error(), nil)
		return
	}

	l, err := models.GetAllCargueMasivo(query, fields, sortby, order, offset, limit)
	if err != nil {
		abortWithError(&c.Controller, 400, "error al consultar los registros", err)
		return
	}

	writeSuccess(&c.Controller, 200, "consulta exitosa", l)
}

// Put ...
// @Title Put
// @Description update the CargueMasivo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.CargueMasivo	true		"body for CargueMasivo content"
// @Success 200 {object} controllers.APIResponse
// @Failure 400 id/body invalidos
// @router /:id [put]
func (c *CargueMasivoController) Put() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		abortWithError(&c.Controller, 400, "el parametro id debe ser numerico", err)
		return
	}

	var v models.CargueMasivo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		abortWithError(&c.Controller, 400, "cuerpo de la solicitud invalido", err)
		return
	}
	v.Id = id

	if err := models.UpdateCargueMasivoById(&v); err != nil {
		abortWithError(&c.Controller, 500, "error al actualizar el registro", err)
		return
	}

	writeSuccess(&c.Controller, 200, "registro actualizado", v)
}

// Delete ...
// @Title Delete
// @Description delete the CargueMasivo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {object} controllers.APIResponse
// @Failure 400 id invalido
// @router /:id [delete]
func (c *CargueMasivoController) Delete() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		abortWithError(&c.Controller, 400, "el parametro id debe ser numerico", err)
		return
	}

	if err := models.DeleteCargueMasivo(id); err != nil {
		abortWithError(&c.Controller, 500, "error al eliminar el registro", err)
		return
	}

	writeSuccess(&c.Controller, 200, "registro eliminado", nil)
}
