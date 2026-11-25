package models

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type HistDetCotizacion struct {
	Id                       int            `orm:"column(id);pk;auto"`
	DetCotizacionId          int            `orm:"column(det_cotizacion_id)"`
	MstCotizacionId          *MstCotizacion `orm:"column(mst_cotizacion_id);rel(fk)"`
	DetSolicitudCotizacionId int            `orm:"column(det_solicitud_cotizacion_id)"`
	ValorUnitario            *float64       `orm:"column(valor_unitario);null"`
	Iva                      float64        `orm:"column(iva)"`
	FichaTecnica             string         `orm:"column(ficha_tecnica)"`
	RegistroActual           bool           `orm:"column(registro_actual)"`
	Activo                   bool           `orm:"column(activo)"`
	FechaCreacion            time.Time      `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	UsuarioCreacionId        int            `orm:"column(usuario_creacion_id);null"`
	FechaModificacion        time.Time      `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	UsuarioModificacionId    int            `orm:"column(usuario_modificacion_id);null"`
}

func (t *HistDetCotizacion) TableName() string {
	return "hist_det_cotizacion"
}

func init() {
	orm.RegisterModel(new(HistDetCotizacion))
}

// AddHistDetCotizacion insert a new HistDetCotizacion into database and returns
// last inserted Id on success.
func AddHistDetCotizacion(m *HistDetCotizacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHistDetCotizacionById retrieves HistDetCotizacion by Id. Returns error if
// Id doesn't exist
func GetHistDetCotizacionById(id int) (v *HistDetCotizacion, err error) {
	o := orm.NewOrm()
	v = &HistDetCotizacion{}
	// RelatedSel expande las FK, igual que en GetAllHistDetCotizacion
	if err = o.QueryTable(new(HistDetCotizacion)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHistDetCotizacion retrieves all HistDetCotizacion matches certain condition. Returns empty list if
// no records exist
// Total is the number of records matching the query filters, ignoring offset and limit.
func GetAllHistDetCotizacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HistDetCotizacion)).RelatedSel()
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
	// order by: sortby sin order se ordena ascendente; los campos inválidos y el order
	// incorrecto se rechazan con ErrSortInvalido antes de consultar la base de datos
	sortFields, err := construirOrden(reflect.TypeOf(HistDetCotizacion{}), sortby, order)
	if err != nil {
		return nil, 0, err
	}

	// El total se calcula antes de ordenar: PostgreSQL no admite ORDER BY en el COUNT
	total, err = qs.Count()
	if err != nil {
		return nil, 0, err
	}
	var l []HistDetCotizacion
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

// UpdateHistDetCotizacion updates HistDetCotizacion by Id and returns error if
// the record to be updated doesn't exist
// columnas limits the updated columns; by default all except Id, FechaCreacion and UsuarioCreacionId
func UpdateHistDetCotizacionById(m *HistDetCotizacion, columnas ...string) (err error) {
	o := orm.NewOrm()
	v := HistDetCotizacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if len(columnas) == 0 {
			columnas = []string{"DetCotizacionId", "MstCotizacionId", "DetSolicitudCotizacionId", "ValorUnitario", "Iva", "FichaTecnica", "RegistroActual", "Activo", "FechaModificacion", "UsuarioModificacionId"}
		}
		if num, err = o.Update(m, columnas...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHistDetCotizacion deletes HistDetCotizacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHistDetCotizacion(id int) (err error) {
	o := orm.NewOrm()
	v := HistDetCotizacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HistDetCotizacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
