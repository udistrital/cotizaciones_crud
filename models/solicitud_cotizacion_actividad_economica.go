package models

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type SolicitudCotizacionActividadEconomica struct {
	Id                       int                     `orm:"column(id);pk;auto"`
	MstSolicitudCotizacionId *MstSolicitudCotizacion `orm:"column(mst_solicitud_cotizacion_id);rel(fk)"`
	ActividadEconomicaId     int                     `orm:"column(actividad_economica_id)"`
	Activo                   bool                    `orm:"column(activo)"`
	FechaCreacion            time.Time               `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	UsuarioCreacionId        int                     `orm:"column(usuario_creacion_id);null"`
	FechaModificacion        time.Time               `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	UsuarioModificacionId    int                     `orm:"column(usuario_modificacion_id);null"`
}

func (t *SolicitudCotizacionActividadEconomica) TableName() string {
	return "solicitud_cotizacion_actividad_economica"
}

func init() {
	orm.RegisterModel(new(SolicitudCotizacionActividadEconomica))
}

// AddSolicitudCotizacionActividadEconomica insert a new SolicitudCotizacionActividadEconomica into database and returns
// last inserted Id on success.
func AddSolicitudCotizacionActividadEconomica(m *SolicitudCotizacionActividadEconomica) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSolicitudCotizacionActividadEconomicaById retrieves SolicitudCotizacionActividadEconomica by Id. Returns error if
// Id doesn't exist
func GetSolicitudCotizacionActividadEconomicaById(id int) (v *SolicitudCotizacionActividadEconomica, err error) {
	o := orm.NewOrm()
	v = &SolicitudCotizacionActividadEconomica{}
	// RelatedSel expande las FK, igual que en GetAllSolicitudCotizacionActividadEconomica
	if err = o.QueryTable(new(SolicitudCotizacionActividadEconomica)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudCotizacionActividadEconomica retrieves all SolicitudCotizacionActividadEconomica matches certain condition. Returns empty list if
// no records exist
// Total is the number of records matching the query filters, ignoring offset and limit.
func GetAllSolicitudCotizacionActividadEconomica(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudCotizacionActividadEconomica)).RelatedSel()
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
	sortFields, err := construirOrden(reflect.TypeOf(SolicitudCotizacionActividadEconomica{}), sortby, order)
	if err != nil {
		return nil, 0, err
	}

	// El total se calcula antes de ordenar: PostgreSQL no admite ORDER BY en el COUNT
	total, err = qs.Count()
	if err != nil {
		return nil, 0, err
	}
	var l []SolicitudCotizacionActividadEconomica
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

// UpdateSolicitudCotizacionActividadEconomica updates SolicitudCotizacionActividadEconomica by Id and returns error if
// the record to be updated doesn't exist
// columnas limits the updated columns; by default all except Id, FechaCreacion and UsuarioCreacionId
func UpdateSolicitudCotizacionActividadEconomicaById(m *SolicitudCotizacionActividadEconomica, columnas ...string) (err error) {
	o := orm.NewOrm()
	v := SolicitudCotizacionActividadEconomica{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if len(columnas) == 0 {
			columnas = []string{"MstSolicitudCotizacionId", "ActividadEconomicaId", "Activo",
				"FechaModificacion", "UsuarioModificacionId"}
		}
		if num, err = o.Update(m, columnas...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSolicitudCotizacionActividadEconomica deletes SolicitudCotizacionActividadEconomica by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudCotizacionActividadEconomica(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudCotizacionActividadEconomica{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SolicitudCotizacionActividadEconomica{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
