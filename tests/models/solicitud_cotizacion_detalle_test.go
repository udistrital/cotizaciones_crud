package models

import (
	"errors"
	"sort"
	"testing"

	"github.com/beego/beego/v2/client/orm"

	"github.com/udistrital/cotizaciones_crud/models"
	"github.com/udistrital/cotizaciones_crud/tests/helpers"
)

// crearCargue crea un cargue masivo de la solicitud indicada.
func crearCargue(t *testing.T, solicitud *models.SolicitudCotizacion) *models.CargueMasivo {
	t.Helper()
	c := helpers.NuevoCargueMasivo(t)
	c.SolicitudId = solicitud
	if _, err := models.AddCargueMasivo(c); err != nil {
		t.Fatal("No se pudo crear el cargue masivo:", err)
	}
	return c
}

// crearErrorCargue crea un error del cargue masivo indicado.
func crearErrorCargue(t *testing.T, cargue *models.CargueMasivo) *models.CargueMasivoError {
	t.Helper()
	e := helpers.NuevoCargueMasivoError(t)
	e.CargueMasivoId = cargue
	if _, err := models.AddCargueMasivoError(e); err != nil {
		t.Fatal("No se pudo crear el error de cargue masivo:", err)
	}
	return e
}

// idsErrores retorna los Id ordenados de los errores de cargue del detalle.
func idsErrores(d *models.SolicitudCotizacionDetalle) []int {
	ids := []int{}
	for _, e := range d.CarguesMasivosErrores {
		ids = append(ids, e.Id)
	}
	sort.Ints(ids)
	return ids
}

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
	for _, err := range altas {
		if err != nil {
			t.Fatal("No se pudieron crear las entidades asociadas:", err)
		}
	}
	crearErrorCargue(t, crearCargue(t, solicitud))
	if err := models.DeleteSolicitudItem(itemInactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el ítem:", err)
	}
	// Una asociada de otra solicitud no debe aparecer en el detalle
	helpers.CrearSolicitudItem(t)

	t.Run("AsociadasActivas", func(t *testing.T) {
		d, err := models.GetSolicitudCotizacionDetalle(solicitud.Id, true)
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
			"CarguesMasivosErrores":  len(d.CarguesMasivosErrores),
		}
		for lista, n := range cantidades {
			if n != 1 {
				t.Errorf("Se esperaba 1 registro activo en %s y se obtuvieron %d", lista, n)
			}
		}
		if len(d.Items) == 1 && (d.Items[0].Id != item.Id || d.Items[0].SolicitudId != nil) {
			t.Error("Se esperaba el ítem activo", item.Id, "sin la solicitud repetida y se obtuvo", d.Items[0])
		}
	})

	t.Run("AsociadasInactivas", func(t *testing.T) {
		d, err := models.GetSolicitudCotizacionDetalle(solicitud.Id, false)
		if err != nil {
			t.Fatal("Se esperaba el detalle de la solicitud y se obtuvo", err)
		}
		if len(d.Items) != 1 || d.Items[0].Id != itemInactivo.Id {
			t.Error("Se esperaba solo el ítem inactivo", itemInactivo.Id, "y se obtuvo", d.Items)
		}
		if n := len(d.Necesidades) + len(d.PagosParciales) + len(d.ActividadesEconomicas) + len(d.Documentos) +
			len(d.RequisitosDocumentales) + len(d.Proveedores) + len(d.CarguesMasivos) + len(d.CarguesMasivosErrores); n != 0 {
			t.Error("Se esperaba que las demás asociadas inactivas estuvieran vacías y se obtuvieron", n, "registros")
		}
	})
	t.Log("TestGetSolicitudCotizacionDetalle Finalizado Correctamente (OK)")
}

// TestGetSolicitudCotizacionDetalleErroresCargue verifica la colección
// CarguesMasivosErrores: es independiente de CarguesMasivos y un error se
// considera activo solo si él y su cargue lo están.
func TestGetSolicitudCotizacionDetalleErroresCargue(t *testing.T) {
	solicitud := helpers.CrearSolicitudCotizacion(t)
	cargueActivo := crearCargue(t, solicitud)
	errorActivo := crearErrorCargue(t, cargueActivo)
	errorInactivo := crearErrorCargue(t, cargueActivo)
	cargueInactivo := crearCargue(t, solicitud)
	errorDeCargueInactivo := crearErrorCargue(t, cargueInactivo)
	if err := models.DeleteCargueMasivoError(errorInactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el error:", err)
	}
	if err := models.DeleteCargueMasivo(cargueInactivo.Id); err != nil {
		t.Fatal("No se pudo inactivar el cargue:", err)
	}
	// Los errores de cargues de otra solicitud no deben aparecer
	helpers.CrearCargueMasivoError(t)

	t.Run("PorDefectoSoloErroresActivosDeCarguesActivos", func(t *testing.T) {
		d, err := models.GetSolicitudCotizacionDetalle(solicitud.Id, true)
		if err != nil {
			t.Fatal("Se esperaba el detalle de la solicitud y se obtuvo", err)
		}
		if ids := idsErrores(d); len(ids) != 1 || ids[0] != errorActivo.Id {
			t.Fatal("Se esperaba solo el error", errorActivo.Id, "y se obtuvo", ids)
		}
		e := d.CarguesMasivosErrores[0]
		if e.CargueMasivoId == nil || e.CargueMasivoId.Id != cargueActivo.Id || e.CargueMasivoId.SolicitudId != nil {
			t.Error("Se esperaba el error con su cargue", cargueActivo.Id, "sin la solicitud repetida y se obtuvo", e.CargueMasivoId)
		}
		if len(d.CarguesMasivos) != 1 || d.CarguesMasivos[0].Id != cargueActivo.Id {
			t.Error("Se esperaba solo el cargue activo", cargueActivo.Id, "y se obtuvo", d.CarguesMasivos)
		}
	})

	t.Run("InactivosIncluyenErroresDeCarguesInactivos", func(t *testing.T) {
		d, err := models.GetSolicitudCotizacionDetalle(solicitud.Id, false)
		if err != nil {
			t.Fatal("Se esperaba el detalle de la solicitud y se obtuvo", err)
		}
		esperados := []int{errorInactivo.Id, errorDeCargueInactivo.Id}
		sort.Ints(esperados)
		if ids := idsErrores(d); len(ids) != 2 || ids[0] != esperados[0] || ids[1] != esperados[1] {
			t.Error("Se esperaban los errores", esperados, "y se obtuvo", ids)
		}
		if len(d.CarguesMasivos) != 1 || d.CarguesMasivos[0].Id != cargueInactivo.Id {
			t.Error("Se esperaba solo el cargue inactivo", cargueInactivo.Id, "y se obtuvo", d.CarguesMasivos)
		}
	})
	t.Log("TestGetSolicitudCotizacionDetalleErroresCargue Finalizado Correctamente (OK)")
}

func TestGetSolicitudCotizacionDetalleSinAsociadas(t *testing.T) {
	solicitud := helpers.CrearSolicitudCotizacion(t)
	for _, activo := range []bool{true, false} {
		d, err := models.GetSolicitudCotizacionDetalle(solicitud.Id, activo)
		if err != nil {
			t.Fatal("Se esperaba el detalle de la solicitud y se obtuvo", err)
		}
		// Las listas vacías se responden como [] y no como null
		if d.Necesidades == nil || d.Items == nil || d.PagosParciales == nil || d.ActividadesEconomicas == nil ||
			d.Documentos == nil || d.RequisitosDocumentales == nil || d.Proveedores == nil || d.CarguesMasivos == nil ||
			d.CarguesMasivosErrores == nil {
			t.Error("Se esperaban listas vacías y no nulas con activo", activo, "y se obtuvo", d)
		}
		if len(d.CarguesMasivosErrores) != 0 {
			t.Error("Se esperaba una solicitud sin errores de cargue y se obtuvo", d.CarguesMasivosErrores)
		}
	}
	if _, err := models.GetSolicitudCotizacionDetalle(helpers.IdInexistente, true); !errors.Is(err, orm.ErrNoRows) {
		t.Error("Se esperaba orm.ErrNoRows al consultar un Id inexistente y se obtuvo", err)
	}
	t.Log("TestGetSolicitudCotizacionDetalleSinAsociadas Finalizado Correctamente (OK)")
}
