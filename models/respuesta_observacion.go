package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type RespuestaObservacion struct {
	Id                int          `orm:"column(id);pk;auto"`
	Activo            bool         `orm:"column(activo)"`
	FechaCreacion     time.Time    `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	ObservacionId     *Observacion `orm:"column(observacion_id);rel(fk)"`
	AutorTerceroId    int          `orm:"column(autor_tercero_id)"`
	Texto             string       `orm:"column(texto)"`
	Publica           bool         `orm:"column(publica);null"`
	FechaModificacion time.Time    `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *RespuestaObservacion) TableName() string {
	return "respuesta_observacion"
}

func init() {
	orm.RegisterModel(new(RespuestaObservacion))
}

// AddRespuestaObservacion insert a new RespuestaObservacion into database and returns
// last inserted Id on success.
func AddRespuestaObservacion(m *RespuestaObservacion) (id int64, err error) {
	o := orm.NewOrm()
	m.Activo = true
	m.FechaCreacion = time.Now()
	m.FechaModificacion = m.FechaCreacion
	id, err = o.Insert(m)
	return
}

// GetRespuestaObservacionById retrieves RespuestaObservacion by Id. Returns error if
// Id doesn't exist
func GetRespuestaObservacionById(id int) (v *RespuestaObservacion, err error) {
	o := orm.NewOrm()
	v = &RespuestaObservacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRespuestaObservacion retrieves all RespuestaObservacion matches certain condition. Returns empty list if
// no records exist. total es la cantidad de registros que cumplen los filtros, sin aplicar limit ni offset
func GetAllRespuestaObservacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, total int64, err error) {
	defer recuperarConsulta(&err)
	o := orm.NewOrm()
	qs := o.QueryTable(new(RespuestaObservacion)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, valorFiltro(new(RespuestaObservacion), k, v))
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("Error: unused 'order' fields")
		}
	}

	// total de registros que cumplen los filtros, antes de aplicar limit y offset
	if total, err = qs.Count(); err != nil {
		return nil, 0, err
	}

	var l []RespuestaObservacion
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, total, nil
	}
	return nil, 0, err
}

// UpdateRespuestaObservacion updates RespuestaObservacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateRespuestaObservacionById(m *RespuestaObservacion) (err error) {
	o := orm.NewOrm()
	v := RespuestaObservacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.FechaCreacion = v.FechaCreacion
		m.FechaModificacion = time.Now()
		_, err = o.Update(m)
	}
	return
}

// DeleteRespuestaObservacion deletes RespuestaObservacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRespuestaObservacion(id int) (err error) {
	o := orm.NewOrm()
	v := RespuestaObservacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		v.Activo = false
		v.FechaModificacion = time.Now()
		_, err = o.Update(&v, "activo", "fecha_modificacion")
	}
	return
}
