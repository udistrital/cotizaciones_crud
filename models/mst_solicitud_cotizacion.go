package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type MstSolicitudCotizacion struct {
	Id                    int       `orm:"column(id);pk;auto"`
	Activo                bool      `orm:"column(activo)"`
	FechaCreacion         time.Time `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	UsuarioCreacionId     int       `orm:"column(usuario_creacion_id);null"`
	FechaModificacion     time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	UsuarioModificacionId int       `orm:"column(usuario_modificacion_id);null"`
}

func (t *MstSolicitudCotizacion) TableName() string {
	return "mst_solicitud_cotizacion"
}

func init() {
	orm.RegisterModel(new(MstSolicitudCotizacion))
}

// AddMstSolicitudCotizacion insert a new MstSolicitudCotizacion into database and returns
// last inserted Id on success.
func AddMstSolicitudCotizacion(m *MstSolicitudCotizacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMstSolicitudCotizacionById retrieves MstSolicitudCotizacion by Id. Returns error if
// Id doesn't exist
func GetMstSolicitudCotizacionById(id int) (v *MstSolicitudCotizacion, err error) {
	o := orm.NewOrm()
	v = &MstSolicitudCotizacion{}
	// RelatedSel expande las FK, igual que en GetAllMstSolicitudCotizacion
	if err = o.QueryTable(new(MstSolicitudCotizacion)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMstSolicitudCotizacion retrieves all MstSolicitudCotizacion matches certain condition. Returns empty list if
// no records exist
// Total is the number of records matching the query filters, ignoring offset and limit.
func GetAllMstSolicitudCotizacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MstSolicitudCotizacion)).RelatedSel()
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

	// El total se calcula antes de ordenar: PostgreSQL no admite ORDER BY en el COUNT
	total, err = qs.Count()
	if err != nil {
		return nil, 0, err
	}
	var l []MstSolicitudCotizacion
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

// UpdateMstSolicitudCotizacion updates MstSolicitudCotizacion by Id and returns error if
// the record to be updated doesn't exist
// columnas limits the updated columns; by default all except Id, FechaCreacion and UsuarioCreacionId
func UpdateMstSolicitudCotizacionById(m *MstSolicitudCotizacion, columnas ...string) (err error) {
	o := orm.NewOrm()
	v := MstSolicitudCotizacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if len(columnas) == 0 {
			columnas = []string{"Activo", "FechaModificacion", "UsuarioModificacionId"}
		}
		if num, err = o.Update(m, columnas...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMstSolicitudCotizacion deletes MstSolicitudCotizacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMstSolicitudCotizacion(id int) (err error) {
	o := orm.NewOrm()
	v := MstSolicitudCotizacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MstSolicitudCotizacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
