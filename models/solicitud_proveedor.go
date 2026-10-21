package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type SolicitudProveedor struct {
	Id                int                  `orm:"column(id);pk;auto"`
	Activo            bool                 `orm:"column(activo)"`
	FechaCreacion     time.Time            `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	SolicitudId       *SolicitudCotizacion `orm:"column(solicitud_id);rel(fk)"`
	ProveedorId       int                  `orm:"column(proveedor_id)"`
	FechaConvocatoria time.Time            `orm:"column(fecha_convocatoria);type(timestamp without time zone)"`
	FechaModificacion time.Time            `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *SolicitudProveedor) TableName() string {
	return "solicitud_proveedor"
}

func init() {
	orm.RegisterModel(new(SolicitudProveedor))
}

// AddSolicitudProveedor insert a new SolicitudProveedor into database and returns
// last inserted Id on success.
func AddSolicitudProveedor(m *SolicitudProveedor) (id int64, err error) {
	o := orm.NewOrm()
	m.Activo = true
	m.FechaCreacion = time.Now()
	id, err = o.Insert(m)
	return
}

// GetSolicitudProveedorById retrieves SolicitudProveedor by Id. Returns error if
// Id doesn't exist
func GetSolicitudProveedorById(id int) (v *SolicitudProveedor, err error) {
	o := orm.NewOrm()
	v = &SolicitudProveedor{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudProveedor retrieves all SolicitudProveedor matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitudProveedor(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudProveedor)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SolicitudProveedor
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
		return ml, nil
	}
	return nil, err
}

// UpdateSolicitudProveedor updates SolicitudProveedor by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudProveedorById(m *SolicitudProveedor) (err error) {
	o := orm.NewOrm()
	v := SolicitudProveedor{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.FechaCreacion = v.FechaCreacion
		m.FechaModificacion = time.Now()
		_, err = o.Update(m)
	}
	return
}

// DeleteSolicitudProveedor deletes SolicitudProveedor by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudProveedor(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudProveedor{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		v.Activo = false
		v.FechaModificacion = time.Now()
		_, err = o.Update(&v, "activo", "fecha_modificacion")
	}
	return
}
