package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type HistMstSolicitudCotizacion struct {
	Id                       int                     `orm:"column(id);pk;auto"`
	MstSolicitudCotizacionId *MstSolicitudCotizacion `orm:"column(mst_solicitud_cotizacion_id);rel(fk)"`
	NumeroSolicitud          *string                 `orm:"column(numero_solicitud);null"`
	Vigencia                 *int16                  `orm:"column(vigencia);null"`
	Consecutivo              *int                    `orm:"column(consecutivo);null"`
	FechaProcesamiento       time.Time               `orm:"column(fecha_procesamiento);type(timestamp without time zone)"`
	EstadoId                 int                     `orm:"column(estado_id)"`
	UnidadEjecutoraId        int                     `orm:"column(unidad_ejecutora_id)"`
	TipoNecesidadId          int                     `orm:"column(tipo_necesidad_id)"`
	Titulo                   string                  `orm:"column(titulo)"`
	FechaApertura            time.Time               `orm:"column(fecha_apertura);type(date)"`
	FechaCierre              time.Time               `orm:"column(fecha_cierre);type(date)"`
	Objetivo                 string                  `orm:"column(objetivo)"`
	Requisitos               string                  `orm:"column(requisitos)"`
	Observaciones            string                  `orm:"column(observaciones)"`
	OrdenadorGastoId         int                     `orm:"column(ordenador_gasto_id)"`
	DependenciaSolicitanteId int                     `orm:"column(dependencia_solicitante_id)"`
	JustificacionSeleccion   string                  `orm:"column(justificacion_seleccion)"`
	MedioPagoId              int                     `orm:"column(medio_pago_id)"`
	PlanAccionId             *int                    `orm:"column(plan_accion_id);null"`
	ResponsableTerceroId     int                     `orm:"column(responsable_tercero_id)"`
	MetodoSeleccionId        *int                    `orm:"column(metodo_seleccion_id);null"`
	NumeroNecesidad          *float64                `orm:"column(numero_necesidad);null"`
	DependenciaDestinoId     int                     `orm:"column(dependencia_destino_id)"`
	TipoContratoId           *int                    `orm:"column(tipo_contrato_id);null"`
	CriterioSeleccion        *string                 `orm:"column(criterio_seleccion);null"`
	PlazoEjecucion           *string                 `orm:"column(plazo_ejecucion);null"`
	RegistroActual           bool                    `orm:"column(registro_actual)"`
	Activo                   bool                    `orm:"column(activo)"`
	FechaCreacion            time.Time               `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	UsuarioCreacionId        int                     `orm:"column(usuario_creacion_id);null"`
	FechaModificacion        time.Time               `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	UsuarioModificacionId    int                     `orm:"column(usuario_modificacion_id);null"`
}

func (t *HistMstSolicitudCotizacion) TableName() string {
	return "hist_mst_solicitud_cotizacion"
}

func init() {
	orm.RegisterModel(new(HistMstSolicitudCotizacion))
}

// AddHistMstSolicitudCotizacion insert a new HistMstSolicitudCotizacion into database and returns
// last inserted Id on success.
func AddHistMstSolicitudCotizacion(m *HistMstSolicitudCotizacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHistMstSolicitudCotizacionById retrieves HistMstSolicitudCotizacion by Id. Returns error if
// Id doesn't exist
func GetHistMstSolicitudCotizacionById(id int) (v *HistMstSolicitudCotizacion, err error) {
	o := orm.NewOrm()
	v = &HistMstSolicitudCotizacion{}
	// RelatedSel expande las FK, igual que en GetAllHistMstSolicitudCotizacion
	if err = o.QueryTable(new(HistMstSolicitudCotizacion)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHistMstSolicitudCotizacion retrieves all HistMstSolicitudCotizacion matches certain condition. Returns empty list if
// no records exist
// Total is the number of records matching the query filters, ignoring offset and limit.
func GetAllHistMstSolicitudCotizacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HistMstSolicitudCotizacion)).RelatedSel()
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
	var l []HistMstSolicitudCotizacion
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

// UpdateHistMstSolicitudCotizacion updates HistMstSolicitudCotizacion by Id and returns error if
// the record to be updated doesn't exist
// columnas limits the updated columns; by default all except Id, FechaCreacion and UsuarioCreacionId
func UpdateHistMstSolicitudCotizacionById(m *HistMstSolicitudCotizacion, columnas ...string) (err error) {
	o := orm.NewOrm()
	v := HistMstSolicitudCotizacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if len(columnas) == 0 {
			columnas = []string{"MstSolicitudCotizacionId", "NumeroSolicitud", "Vigencia", "Consecutivo",
				"FechaProcesamiento", "EstadoId", "UnidadEjecutoraId", "TipoNecesidadId", "Titulo",
				"FechaApertura", "FechaCierre", "Objetivo", "Requisitos", "Observaciones",
				"OrdenadorGastoId", "DependenciaSolicitanteId", "JustificacionSeleccion", "MedioPagoId",
				"PlanAccionId", "ResponsableTerceroId", "MetodoSeleccionId", "NumeroNecesidad",
				"DependenciaDestinoId", "TipoContratoId", "CriterioSeleccion", "PlazoEjecucion", "RegistroActual",
				"Activo", "FechaModificacion", "UsuarioModificacionId"}
		}
		if num, err = o.Update(m, columnas...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHistMstSolicitudCotizacion deletes HistMstSolicitudCotizacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHistMstSolicitudCotizacion(id int) (err error) {
	o := orm.NewOrm()
	v := HistMstSolicitudCotizacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HistMstSolicitudCotizacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
