package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/utils_oas/time_bogota"

	beeLogger "github.com/beego/bee/v2/logger"
	beego "github.com/beego/beego/v2/server/web"
)

// SolicitudCotizacionActividadEconomicaController operations for SolicitudCotizacionActividadEconomica
type SolicitudCotizacionActividadEconomicaController struct {
	beego.Controller
}

// URLMapping ...
func (c *SolicitudCotizacionActividadEconomicaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// obligatoriosSolicitudCotizacionActividadEconomica son los campos de columnas NOT NULL que no son llave foránea: en POST y PUT
// no pueden faltar, ser null ni, si son texto, estar vacíos.
var obligatoriosSolicitudCotizacionActividadEconomica = []string{
	"ActividadEconomicaId",
}

// Post ...
// @Title Post
// @Description create SolicitudCotizacionActividadEconomica
// @Param	body		body 	models.SolicitudCotizacionActividadEconomica	true		"body for SolicitudCotizacionActividadEconomica content"
// @Success 201 {int} models.SolicitudCotizacionActividadEconomica
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *SolicitudCotizacionActividadEconomicaController) Post() {
	var v models.SolicitudCotizacionActividadEconomica
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		// Columnas NOT NULL que no son llave foránea: no pueden faltar, ser null ni estar vacías
		if err := validarObligatorios(c.Ctx.Input.RequestBody, obligatoriosSolicitudCotizacionActividadEconomica...); err != nil {
			c.Data["message"] = "Error service POST: " + err.Error()
			c.Abort("400")
		}
		// El Id lo asigna la base de datos (serial): se ignora el enviado por el cliente
		v.Id = 0
		v.Activo = true
		// usuario_creacion_id es obligatorio en el POST; no lo asigna el servidor
		if v.UsuarioCreacionId <= 0 {
			c.Data["message"] = "Error service POST: usuario_creacion_id is required"
			c.Abort("400")
		}
		v.FechaCreacion = time_bogota.Tiempo_bogota()
		v.FechaModificacion = time_bogota.Tiempo_bogota()
		if _, err := models.AddSolicitudCotizacionActividadEconomica(&v); err == nil {
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
// @Description get SolicitudCotizacionActividadEconomica by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SolicitudCotizacionActividadEconomica
// @Failure 404 not found resource
// @router /:id [get]
func (c *SolicitudCotizacionActividadEconomicaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	// El Id de la ruta debe ser numérico: no se convierte en 0 ni se consulta la base
	id, err := strconv.Atoi(idStr)
	if err != nil {
		beeLogger.Log.Error(err.Error())
		c.Data["message"] = "Error service GetOne: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	}
	v, err := models.GetSolicitudCotizacionActividadEconomicaById(id)
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
// @Description get SolicitudCotizacionActividadEconomica
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SolicitudCotizacionActividadEconomica
// @Failure 404 not found resource
// @router / [get]
func (c *SolicitudCotizacionActividadEconomicaController) GetAll() {
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

	l, total, err := models.GetAllSolicitudCotizacionActividadEconomica(query, fields, sortby, order, offset, limit)
	if err != nil {
		beeLogger.Log.Error(err.Error())
		if errors.Is(err, models.ErrSortInvalido) {
			c.Data["message"] = "Error service GetAll: The request contains an invalid sortby or order parameter"
			c.Abort("400")
		} else {
			c.Data["message"] = "Error service GetAll: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
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
// @Description update the SolicitudCotizacionActividadEconomica
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SolicitudCotizacionActividadEconomica	true		"body for SolicitudCotizacionActividadEconomica content"
// @Success 200 {object} models.SolicitudCotizacionActividadEconomica
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *SolicitudCotizacionActividadEconomicaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	// El Id de la ruta debe ser numérico: no se convierte en 0 ni se consulta la base
	id, err := strconv.Atoi(idStr)
	if err != nil {
		beeLogger.Log.Error(err.Error())
		c.Data["message"] = "Error service Put: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	}
	v := models.SolicitudCotizacionActividadEconomica{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		// Columnas NOT NULL que no son llave foránea: no pueden faltar, ser null ni estar vacías
		if err := validarObligatorios(c.Ctx.Input.RequestBody, obligatoriosSolicitudCotizacionActividadEconomica...); err != nil {
			c.Data["message"] = "Error service Put: " + err.Error()
			c.Abort("400")
		}
		if v.Id != id {
			c.Data["message"] = "Error service Put: The id cannot be modified"
			c.Abort("400")
		}
		if v.Id == 0 {
			v.Id = id
		}
		// usuario_creacion_id no se modifica en PUT: se conserva el valor almacenado
		actual, errActual := models.GetSolicitudCotizacionActividadEconomicaById(id)
		if errActual != nil {
			beeLogger.Log.Error(errActual.Error())
			c.Data["message"] = "Error service Put: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
		v.UsuarioCreacionId = actual.UsuarioCreacionId
		v.FechaModificacion = time_bogota.Tiempo_bogota()
		if err := models.UpdateSolicitudCotizacionActividadEconomicaById(&v); err == nil {
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
// @Description delete the SolicitudCotizacionActividadEconomica
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *SolicitudCotizacionActividadEconomicaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	// El Id de la ruta debe ser numérico: no se convierte en 0 ni se consulta la base
	id, err := strconv.Atoi(idStr)
	if err != nil {
		beeLogger.Log.Error(err.Error())
		c.Data["message"] = "Error service Delete: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	}
	v, err := models.GetSolicitudCotizacionActividadEconomicaById(id)
	if err != nil {
		beeLogger.Log.Error(err.Error())
		c.Data["message"] = "Error service Delete: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	}
	v.Activo = false
	v.FechaModificacion = time_bogota.Tiempo_bogota()
	// Borrado lógico: solo cambian activo y fecha_modificacion
	if err := models.UpdateSolicitudCotizacionActividadEconomicaById(v, "Activo", "FechaModificacion"); err == nil {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Update successful", "Data": v}
	} else {
		beeLogger.Log.Error(err.Error())
		c.Data["message"] = "Error service Put: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}
