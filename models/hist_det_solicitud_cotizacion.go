package models

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type HistDetSolicitudCotizacion struct {
	Id                       int                     `orm:"column(id);pk;auto"`
	DetSolicitudCotizacionId int                     `orm:"column(det_solicitud_cotizacion_id)"`
	MstSolicitudCotizacionId *MstSolicitudCotizacion `orm:"column(mst_solicitud_cotizacion_id);rel(fk)"`
	Nombre                   string                  `orm:"column(nombre)"`
	Descripcion              string                  `orm:"column(descripcion)"`
	TipoBienServicioId       int                     `orm:"column(tipo_bien_servicio_id)"`
	UnidadMedidaId           int                     `orm:"column(unidad_medida_id)"`
	Cantidad                 float64                 `orm:"column(cantidad)"`
	RegistroActual           bool                    `orm:"column(registro_actual)"`
	TipoRegistro             string                  `orm:"column(tipo_registro)"`
	Activo                   bool                    `orm:"column(activo)"`
	FechaCreacion            time.Time               `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	UsuarioCreacionId        int                     `orm:"column(usuario_creacion_id);null"`
	FechaModificacion        time.Time               `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	UsuarioModificacionId    int                     `orm:"column(usuario_modificacion_id);null"`
}

func (t *HistDetSolicitudCotizacion) TableName() string {
	return "hist_det_solicitud_cotizacion"
}

func init() {
	orm.RegisterModel(new(HistDetSolicitudCotizacion))
}

// AddHistDetSolicitudCotizacion insert a new HistDetSolicitudCotizacion into database and returns
// last inserted Id on success.
func AddHistDetSolicitudCotizacion(m *HistDetSolicitudCotizacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHistDetSolicitudCotizacionById retrieves HistDetSolicitudCotizacion by Id. Returns error if
// Id doesn't exist
func GetHistDetSolicitudCotizacionById(id int) (v *HistDetSolicitudCotizacion, err error) {
	o := orm.NewOrm()
	v = &HistDetSolicitudCotizacion{}
	// RelatedSel expande las FK, igual que en GetAllHistDetSolicitudCotizacion
	if err = o.QueryTable(new(HistDetSolicitudCotizacion)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHistDetSolicitudCotizacion retrieves all HistDetSolicitudCotizacion matches certain condition. Returns empty list if
// no records exist
// Total is the number of records matching the query filters, ignoring offset and limit.
func GetAllHistDetSolicitudCotizacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HistDetSolicitudCotizacion)).RelatedSel()
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
	sortFields, err := construirOrden(reflect.TypeOf(HistDetSolicitudCotizacion{}), sortby, order)
	if err != nil {
		return nil, 0, err
	}

	// El total se calcula antes de ordenar: PostgreSQL no admite ORDER BY en el COUNT
	total, err = qs.Count()
	if err != nil {
		return nil, 0, err
	}
	var l []HistDetSolicitudCotizacion
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

// UpdateHistDetSolicitudCotizacion updates HistDetSolicitudCotizacion by Id and returns error if
// the record to be updated doesn't exist
// columnas limits the updated columns; by default all except Id, FechaCreacion and UsuarioCreacionId
func UpdateHistDetSolicitudCotizacionById(m *HistDetSolicitudCotizacion, columnas ...string) (err error) {
	o := orm.NewOrm()
	v := HistDetSolicitudCotizacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if len(columnas) == 0 {
			columnas = []string{"DetSolicitudCotizacionId", "MstSolicitudCotizacionId", "Nombre", "Descripcion", "TipoBienServicioId", "UnidadMedidaId", "Cantidad", "RegistroActual", "TipoRegistro", "Activo", "FechaModificacion", "UsuarioModificacionId"}
		}
		if num, err = o.Update(m, columnas...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHistDetSolicitudCotizacion deletes HistDetSolicitudCotizacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHistDetSolicitudCotizacion(id int) (err error) {
	o := orm.NewOrm()
	v := HistDetSolicitudCotizacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HistDetSolicitudCotizacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
