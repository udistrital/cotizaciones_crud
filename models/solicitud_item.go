package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type SolicitudItem struct {
	Id                 int                  `orm:"column(id);pk;auto"`
	Activo             bool                 `orm:"column(activo)"`
	FechaCreacion      time.Time            `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	SolicitudId        *SolicitudCotizacion `orm:"column(solicitud_id);rel(fk)"`
	TipoBienServicioId int                  `orm:"column(tipo_bien_servicio_id)"`
	Nombre             string               `orm:"column(nombre)"`
	Descripcion        string               `orm:"column(descripcion)"`
	UnidadMedidaId     int                  `orm:"column(unidad_medida_id);null"`
	Cantidad           float64              `orm:"column(cantidad)"`
	CargueMasivoId     *CargueMasivo        `orm:"column(cargue_masivo_id);rel(fk);null"`
	CodigoUnspsc       string               `orm:"column(codigo_unspsc);null"`
	FechaModificacion  time.Time            `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *SolicitudItem) TableName() string {
	return "solicitud_item"
}

func init() {
	orm.RegisterModel(new(SolicitudItem))
}

// AddSolicitudItem insert a new SolicitudItem into database and returns
// last inserted Id on success.
func AddSolicitudItem(m *SolicitudItem) (id int64, err error) {
	o := orm.NewOrm()
	m.Activo = true
	m.FechaCreacion = time.Now()
	m.FechaModificacion = m.FechaCreacion
	id, err = o.Insert(m)
	return
}

// GetSolicitudItemById retrieves SolicitudItem by Id. Returns error if
// Id doesn't exist
func GetSolicitudItemById(id int) (v *SolicitudItem, err error) {
	o := orm.NewOrm()
	v = &SolicitudItem{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudItem retrieves all SolicitudItem matches certain condition. Returns empty list if
// no records exist. total es la cantidad de registros que cumplen los filtros, sin aplicar limit ni offset
func GetAllSolicitudItem(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, total int64, err error) {
	defer recuperarConsulta(&err)
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudItem)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, valorFiltro(new(SolicitudItem), k, v))
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

	var l []SolicitudItem
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

// UpdateSolicitudItem updates SolicitudItem by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudItemById(m *SolicitudItem) (err error) {
	o := orm.NewOrm()
	v := SolicitudItem{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.FechaCreacion = v.FechaCreacion
		m.FechaModificacion = time.Now()
		_, err = o.Update(m)
	}
	return
}

// DeleteSolicitudItem deletes SolicitudItem by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudItem(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudItem{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		v.Activo = false
		v.FechaModificacion = time.Now()
		_, err = o.Update(&v, "activo", "fecha_modificacion")
	}
	return
}
