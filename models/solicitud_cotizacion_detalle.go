package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// SolicitudCotizacionDetalle es la solicitud con sus entidades asociadas, activas o
// inactivas segun la consulta. En cada asociada SolicitudId va vacio porque es la
// misma solicitud consultada.
type SolicitudCotizacionDetalle struct {
	*SolicitudCotizacion
	Necesidades            []SolicitudNecesidad
	Items                  []SolicitudItem
	PagosParciales         []SolicitudPagoParcial
	ActividadesEconomicas  []SolicitudCiiu
	Documentos             []DocumentoSolicitud
	RequisitosDocumentales []SolicitudRequisitoDocumental
	Proveedores            []SolicitudProveedor
	CarguesMasivos         []CargueMasivo
	// CarguesMasivosErrores son los errores de los cargues masivos de la solicitud.
	// Es una coleccion independiente: cargue_masivo_error se relaciona con
	// cargue_masivo y conserva su propio modelo y CRUD.
	CarguesMasivosErrores []CargueMasivoError
}

// GetSolicitudCotizacionDetalle retorna la solicitud con sus entidades asociadas activas
// (activo=true) o inactivas (activo=false). La solicitud se retorna este activa o no.
// Retorna orm.ErrNoRows si la solicitud no existe.
func GetSolicitudCotizacionDetalle(id int, activo bool) (*SolicitudCotizacionDetalle, error) {
	solicitud, err := GetSolicitudCotizacionById(id)
	if err != nil {
		return nil, err
	}

	d := &SolicitudCotizacionDetalle{
		SolicitudCotizacion:    solicitud,
		Necesidades:            []SolicitudNecesidad{},
		Items:                  []SolicitudItem{},
		PagosParciales:         []SolicitudPagoParcial{},
		ActividadesEconomicas:  []SolicitudCiiu{},
		Documentos:             []DocumentoSolicitud{},
		RequisitosDocumentales: []SolicitudRequisitoDocumental{},
		Proveedores:            []SolicitudProveedor{},
		CarguesMasivos:         []CargueMasivo{},
		CarguesMasivosErrores:  []CargueMasivoError{},
	}

	o := orm.NewOrm()
	consultas := []struct {
		modelo  interface{}
		destino interface{}
	}{
		{new(SolicitudNecesidad), &d.Necesidades},
		{new(SolicitudItem), &d.Items},
		{new(SolicitudPagoParcial), &d.PagosParciales},
		{new(SolicitudCiiu), &d.ActividadesEconomicas},
		{new(DocumentoSolicitud), &d.Documentos},
		{new(SolicitudRequisitoDocumental), &d.RequisitosDocumentales},
		{new(SolicitudProveedor), &d.Proveedores},
		{new(CargueMasivo), &d.CarguesMasivos},
	}
	for _, q := range consultas {
		if _, err := o.QueryTable(q.modelo).Filter("SolicitudId", id).Filter("Activo", activo).OrderBy("Id").All(q.destino); err != nil {
			return nil, err
		}
	}

	for i := range d.Necesidades {
		d.Necesidades[i].SolicitudId = nil
	}
	for i := range d.Items {
		d.Items[i].SolicitudId = nil
	}
	for i := range d.PagosParciales {
		d.PagosParciales[i].SolicitudId = nil
	}
	for i := range d.ActividadesEconomicas {
		d.ActividadesEconomicas[i].SolicitudId = nil
	}
	for i := range d.Documentos {
		d.Documentos[i].SolicitudId = nil
	}
	for i := range d.RequisitosDocumentales {
		d.RequisitosDocumentales[i].SolicitudId = nil
	}
	for i := range d.Proveedores {
		d.Proveedores[i].SolicitudId = nil
	}
	for i := range d.CarguesMasivos {
		d.CarguesMasivos[i].SolicitudId = nil
	}

	// Un error esta activo solo si el y su cargue lo estan, para no mostrar por
	// defecto errores de un cargue eliminado. Con activo=false se retornan los
	// errores inactivos o los de cargues inactivos.
	errores := orm.NewCondition().And("CargueMasivoId__SolicitudId__Id", id)
	if activo {
		errores = errores.And("Activo", true).And("CargueMasivoId__Activo", true)
	} else {
		errores = errores.AndCond(orm.NewCondition().Or("Activo", false).Or("CargueMasivoId__Activo", false))
	}
	if _, err := o.QueryTable(new(CargueMasivoError)).SetCond(errores).RelatedSel("CargueMasivoId").
		OrderBy("Id").All(&d.CarguesMasivosErrores); err != nil {
		return nil, err
	}
	for i := range d.CarguesMasivosErrores {
		d.CarguesMasivosErrores[i].CargueMasivoId.SolicitudId = nil
	}

	return d, nil
}
