package models

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type CargueMasivo struct {
	Id                           int                         `orm:"column(id);pk;auto"`
	HistMstSolicitudCotizacionId *HistMstSolicitudCotizacion `orm:"column(hist_mst_solicitud_cotizacion_id);rel(fk)"`
	EnlacePlantilla              *string                     `orm:"column(enlace_plantilla);null"`
	CargadoPorTerceroId          int                         `orm:"column(cargado_por_tercero_id)"`
	TotalRegistros               *int                        `orm:"column(total_registros);null"`
	Aceptados                    *int                        `orm:"column(aceptados);null"`
	Rechazados                   *int                        `orm:"column(rechazados);null"`
	Respuesta                    *string                     `orm:"column(respuesta);type(json);null"`
	Activo                       bool                        `orm:"column(activo)"`
	FechaCreacion                time.Time                   `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	UsuarioCreacionId            int                         `orm:"column(usuario_creacion_id);null"`
	FechaModificacion            time.Time                   `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	UsuarioModificacionId        int                         `orm:"column(usuario_modificacion_id);null"`
}

func (t *CargueMasivo) TableName() string {
	return "cargue_masivo"
}

func init() {
	orm.RegisterModel(new(CargueMasivo))
}

// AddCargueMasivo insert a new CargueMasivo into database and returns
// last inserted Id on success.
func AddCargueMasivo(m *CargueMasivo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCargueMasivoById retrieves CargueMasivo by Id. Returns error if
// Id doesn't exist
func GetCargueMasivoById(id int) (v *CargueMasivo, err error) {
	o := orm.NewOrm()
	v = &CargueMasivo{}
	// RelatedSel expande las FK, igual que en GetAllCargueMasivo
	if err = o.QueryTable(new(CargueMasivo)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCargueMasivo retrieves all CargueMasivo matches certain condition. Returns empty list if
// no records exist
// Total is the number of records matching the query filters, ignoring offset and limit.
func GetAllCargueMasivo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CargueMasivo)).RelatedSel()
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
	sortFields, err := construirOrden(reflect.TypeOf(CargueMasivo{}), sortby, order)
	if err != nil {
		return nil, 0, err
	}

	// El total se calcula antes de ordenar: PostgreSQL no admite ORDER BY en el COUNT
	total, err = qs.Count()
	if err != nil {
		return nil, 0, err
	}
	var l []CargueMasivo
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

// UpdateCargueMasivo updates CargueMasivo by Id and returns error if
// the record to be updated doesn't exist
// columnas limits the updated columns; by default all except Id, FechaCreacion and UsuarioCreacionId
func UpdateCargueMasivoById(m *CargueMasivo, columnas ...string) (err error) {
	o := orm.NewOrm()
	v := CargueMasivo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if len(columnas) == 0 {
			columnas = []string{"HistMstSolicitudCotizacionId", "EnlacePlantilla", "CargadoPorTerceroId",
				"TotalRegistros", "Aceptados", "Rechazados", "Respuesta", "Activo", "FechaModificacion",
				"UsuarioModificacionId"}
		}
		if num, err = o.Update(m, columnas...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCargueMasivo deletes CargueMasivo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCargueMasivo(id int) (err error) {
	o := orm.NewOrm()
	v := CargueMasivo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CargueMasivo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
