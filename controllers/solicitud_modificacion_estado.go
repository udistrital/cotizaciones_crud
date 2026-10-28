package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/udistrital/cotizaciones_crud/models"

	beego "github.com/beego/beego/v2/server/web"
)

// SolicitudModificacionEstadoController operations for SolicitudModificacionEstado
type SolicitudModificacionEstadoController struct {
	beego.Controller
}

// URLMapping ...
func (c *SolicitudModificacionEstadoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create SolicitudModificacionEstado
// @Param	body		body 	models.SolicitudModificacionEstado	true		"body for SolicitudModificacionEstado content"
// @Success 201 {object} controllers.APIResponse
// @Failure 400 body o datos invalidos
// @router / [post]
func (c *SolicitudModificacionEstadoController) Post() {
	var v models.SolicitudModificacionEstado
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		abortWithError(&c.Controller, 400, "cuerpo de la solicitud invalido", err)
		return
	}

	if _, err := models.AddSolicitudModificacionEstado(&v); err != nil {
		abortWithDBError(&c.Controller, 400, "error al crear el registro", err)
		return
	}

	writeSuccess(&c.Controller, 201, "registro creado", v)
}

// GetOne ...
// @Title Get One
// @Description get SolicitudModificacionEstado by id
// @Param	id		path 	string	true		"The id"
// @Success 200 {object} controllers.APIResponse
// @Failure 404 id invalido o registro inexistente
// @router /:id [get]
func (c *SolicitudModificacionEstadoController) GetOne() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		abortWithError(&c.Controller, 404, "el parametro id debe ser numerico", err)
		return
	}

	v, err := models.GetSolicitudModificacionEstadoById(id)
	if err != nil {
		abortWithError(&c.Controller, 404, "registro no encontrado", err)
		return
	}

	writeSuccess(&c.Controller, 200, "consulta exitosa", v)
}

// GetAll ...
// @Title Get All
// @Description get SolicitudModificacionEstado. Total es la cantidad de registros que cumplen los filtros, sin limit ni offset; Data contiene los registros de la pagina
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Param	vigencia	query	int	false	"Vigencia de la solicitud. e.g. 2026"
// @Param	estado	query	int	false	"Id del estado de la solicitud (parametros)"
// @Param	dependencia	query	int	false	"Id de la dependencia solicitante (Oikos)"
// @Success 200 {object} controllers.APIResponseLista
// @Failure 400 parametros invalidos
// @Failure 404 error al consultar los registros
// @router / [get]
func (c *SolicitudModificacionEstadoController) GetAll() {
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

	if err := agregarFiltrosSolicitud(&c.Controller, query, "SolicitudModificacionId.SolicitudId."); err != nil {
		abortWithError(&c.Controller, 400, err.Error(), nil)
		return
	}

	l, total, err := models.GetAllSolicitudModificacionEstado(query, fields, sortby, order, offset, limit)
	if err != nil {
		abortWithError(&c.Controller, 404, "error al consultar los registros", err)
		return
	}

	if l == nil {
		l = append(l, map[string]interface{}{})
	}

	writeSuccessLista(&c.Controller, "consulta exitosa", total, l)
}

// Put ...
// @Title Put
// @Description update the SolicitudModificacionEstado
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SolicitudModificacionEstado	true		"body for SolicitudModificacionEstado content"
// @Success 200 {object} controllers.APIResponse
// @Failure 400 id, body o datos invalidos, o registro inexistente
// @router /:id [put]
func (c *SolicitudModificacionEstadoController) Put() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		abortWithError(&c.Controller, 400, "el parametro id debe ser numerico", err)
		return
	}

	var v models.SolicitudModificacionEstado
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		abortWithError(&c.Controller, 400, "cuerpo de la solicitud invalido", err)
		return
	}
	v.Id = id

	if err := models.UpdateSolicitudModificacionEstadoById(&v); err != nil {
		abortWithDBError(&c.Controller, 400, "error al actualizar el registro", err)
		return
	}

	writeSuccess(&c.Controller, 200, "registro actualizado", v)
}

// Delete ...
// @Title Delete
// @Description delete the SolicitudModificacionEstado
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {object} controllers.APIResponse
// @Failure 404 id invalido o registro inexistente
// @router /:id [delete]
func (c *SolicitudModificacionEstadoController) Delete() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		abortWithError(&c.Controller, 404, "el parametro id debe ser numerico", err)
		return
	}

	if err := models.DeleteSolicitudModificacionEstado(id); err != nil {
		abortWithDBError(&c.Controller, 404, "error al eliminar el registro", err)
		return
	}

	writeSuccess(&c.Controller, 200, "registro eliminado", map[string]interface{}{"Id": id})
}
