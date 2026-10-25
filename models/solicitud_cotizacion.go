package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type SolicitudCotizacion struct {
	Id                       int       `orm:"column(id);pk;auto"`
	Activo                   bool      `orm:"column(activo)"`
	FechaCreacion            time.Time `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	NumeroSolicitud          *string   `orm:"column(numero_solicitud);null"`
	Consecutivo              int       `orm:"column(consecutivo);null"`
	Vigencia                 int16     `orm:"column(vigencia)"`
	UnidadEjecutoraId        int       `orm:"column(unidad_ejecutora_id)"`
	TipoProcesoId            int       `orm:"column(tipo_proceso_id);null"`
	EstadoId                 int       `orm:"column(estado_id)"`
	Titulo                   string    `orm:"column(titulo);null"`
	Objetivo                 string    `orm:"column(objetivo);null"`
	Requisitos               string    `orm:"column(requisitos);null"`
	Observaciones            string    `orm:"column(observaciones);null"`
	PlazoEjecucion           string    `orm:"column(plazo_ejecucion);null"`
	CriterioSeleccion        string    `orm:"column(criterio_seleccion);null"`
	PlanAccionId             int       `orm:"column(plan_accion_id);null"`
	DependenciaSolicitanteId int       `orm:"column(dependencia_solicitante_id);null"`
	DependenciaDestinoId     int       `orm:"column(dependencia_destino_id);null"`
	OrdenadorGastoId         int       `orm:"column(ordenador_gasto_id);null"`
	ResponsableTerceroId     int       `orm:"column(responsable_tercero_id)"`
	TipoNecesidadId          int       `orm:"column(tipo_necesidad_id);null"`
	TipoContratoId           int       `orm:"column(tipo_contrato_id);null"`
	MedioPagoId              int       `orm:"column(medio_pago_id);null"`
	TipoFormaPagoId          int       `orm:"column(tipo_forma_pago_id);null"`
	AreaConocimientoId       int       `orm:"column(area_conocimiento_id);null"`
	NucleoBasicoId           int       `orm:"column(nucleo_basico_id);null"`
	FechaApertura            time.Time `orm:"column(fecha_apertura);type(date);null"`
	FechaCierre              time.Time `orm:"column(fecha_cierre);type(date);null"`
	FechaProcesamiento       time.Time `orm:"column(fecha_procesamiento);type(timestamp without time zone);null"`
	FechaModificacion        time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *SolicitudCotizacion) TableName() string {
	return "solicitud_cotizacion"
}

func init() {
	orm.RegisterModel(new(SolicitudCotizacion))
}

// AddSolicitudCotizacion insert a new SolicitudCotizacion into database and returns
// last inserted Id on success.
func AddSolicitudCotizacion(m *SolicitudCotizacion) (id int64, err error) {
	o := orm.NewOrm()
	m.Activo = true
	m.FechaCreacion = time.Now()
	m.FechaModificacion = m.FechaCreacion
	id, err = o.Insert(m)
	return
}

// GetSolicitudCotizacionById retrieves SolicitudCotizacion by Id. Returns error if
// Id doesn't exist
func GetSolicitudCotizacionById(id int) (v *SolicitudCotizacion, err error) {
	o := orm.NewOrm()
	v = &SolicitudCotizacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudCotizacion retrieves all SolicitudCotizacion matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitudCotizacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	defer recuperarConsulta(&err)
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudCotizacion)).RelatedSel()
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

	var l []SolicitudCotizacion
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

// UpdateSolicitudCotizacion updates SolicitudCotizacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudCotizacionById(m *SolicitudCotizacion) (err error) {
	o := orm.NewOrm()
	v := SolicitudCotizacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.FechaCreacion = v.FechaCreacion
		m.FechaModificacion = time.Now()
		_, err = o.Update(m)
	}
	return
}

// DeleteSolicitudCotizacion deletes SolicitudCotizacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudCotizacion(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudCotizacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		v.Activo = false
		v.FechaModificacion = time.Now()
		_, err = o.Update(&v, "activo", "fecha_modificacion")
	}
	return
}
