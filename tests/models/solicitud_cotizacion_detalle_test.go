package models

import (
	"errors"
	"testing"

	"github.com/beego/beego/v2/client/orm"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

func TestGetSolicitudCotizacionDetalle(t *testing.T) {
	solicitud := helpers.CrearSolicitudCotizacion(t)

	// Una asociada activa de cada tipo, asociada a la misma solicitud
	necesidad := helpers.NuevoSolicitudNecesidad(t)
	necesidad.SolicitudId = solicitud
	item := helpers.NuevoSolicitudItem(t)
	item.SolicitudId = solicitud
	itemInactivo := helpers.NuevoSolicitudItem(t)
	itemInactivo.SolicitudId = solicitud
	pago := helpers.NuevoSolicitudPagoParcial(t)
	pago.SolicitudId = solicitud
	ciiu := helpers.NuevoSolicitudCiiu(t)
	ciiu.SolicitudId = solicitud
	documento := helpers.NuevoDocumentoSolicitud(t)
	documento.SolicitudId = solicitud
	requisito := helpers.NuevoSolicitudRequisitoDocumental(t)
	requisito.SolicitudId = solicitud
	proveedor := helpers.NuevoSolicitudProveedor(t)
	proveedor.SolicitudId = solicitud
	cargue := helpers.NuevoCargueMasivo(t)
	cargue.SolicitudId = solicitud

	altas := []error{}
	_, err := models.AddSolicitudNecesidad(necesidad)
	altas = append(altas, err)
	_, err = models.AddSolicitudItem(item)
	altas = append(altas, err)
	_, err = models.AddSolicitudItem(itemInactivo)
	altas = append(altas, err)
	_, err = models.AddSolicitudPagoParcial(pago)
	altas = append(altas, err)
	_, err = models.AddSolicitudCiiu(ciiu)
	altas = append(altas, err)
	_, err = models.AddDocumentoSolicitud(documento)
	altas = append(altas, err)
	_, err = models.AddSolicitudRequisitoDocumental(requisito)
	altas = append(altas, err)
	_, err = models.AddSolicitudProveedor(proveedor)
	altas = append(altas, err)
	_, err = models.AddCargueMasivo(cargue)
	altas = append(altas, err)
	for _, err := range altas {
		if err != nil {
			t.Fatal("No se pudieron crear las entidades asociadas:", err)
		}
	}
	if err := models.DeleteSolicitudItem(itemInactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el ítem:", err)
	}
	// Una asociada de otra solicitud no debe aparecer en el detalle
	helpers.CrearSolicitudItem(t)

	d, err := models.GetSolicitudCotizacionDetalle(solicitud.Id)
	if err != nil || d.Id != solicitud.Id || d.Titulo != solicitud.Titulo {
		t.Fatal("Se esperaba el detalle de la solicitud", solicitud.Id, "y se obtuvo", d, err)
	}

	cantidades := map[string]int{
		"Necesidades":            len(d.Necesidades),
		"Items":                  len(d.Items),
		"PagosParciales":         len(d.PagosParciales),
		"ActividadesEconomicas":  len(d.ActividadesEconomicas),
		"Documentos":             len(d.Documentos),
		"RequisitosDocumentales": len(d.RequisitosDocumentales),
		"Proveedores":            len(d.Proveedores),
		"CarguesMasivos":         len(d.CarguesMasivos),
	}
	for lista, n := range cantidades {
		if n != 1 {
			t.Errorf("Se esperaba 1 registro activo en %s y se obtuvieron %d", lista, n)
		}
	}
	if len(d.Items) == 1 && (d.Items[0].Id != item.Id || d.Items[0].SolicitudId != nil) {
		t.Error("Se esperaba el ítem activo", item.Id, "sin la solicitud repetida y se obtuvo", d.Items[0])
	}
	t.Log("TestGetSolicitudCotizacionDetalle Finalizado Correctamente (OK)")
}

func TestGetSolicitudCotizacionDetalleSinAsociadas(t *testing.T) {
	solicitud := helpers.CrearSolicitudCotizacion(t)
	d, err := models.GetSolicitudCotizacionDetalle(solicitud.Id)
	if err != nil {
		t.Fatal("Se esperaba el detalle de la solicitud y se obtuvo", err)
	}
	// Las listas vacías se responden como [] y no como null
	if d.Necesidades == nil || d.Items == nil || d.PagosParciales == nil || d.ActividadesEconomicas == nil ||
		d.Documentos == nil || d.RequisitosDocumentales == nil || d.Proveedores == nil || d.CarguesMasivos == nil {
		t.Error("Se esperaban listas vacías y no nulas y se obtuvo", d)
	}
	if _, err := models.GetSolicitudCotizacionDetalle(helpers.IdInexistente); !errors.Is(err, orm.ErrNoRows) {
		t.Error("Se esperaba orm.ErrNoRows al consultar un Id inexistente y se obtuvo", err)
	}
	t.Log("TestGetSolicitudCotizacionDetalleSinAsociadas Finalizado Correctamente (OK)")
}
