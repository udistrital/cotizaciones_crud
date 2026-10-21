// @APIVersion 1.0.0
// @Title API CRUD Cotizaciones Agora
// @Description Servicio CRUD para la gestion de cotizaciones del modulo Agora en Arka II
// @Contact soporte@udistrital.edu.co
// @TermsOfServiceUrl https://www.udistrital.edu.co/
// @License EUPL-1.2
// @LicenseUrl https://interoperable-europe.ec.europa.eu/collection/eupl/eupl-text-eupl-12
package routers

import (
	"github.com/udistrital/cotizaciones_crud/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/cargue_masivo_error",
			beego.NSInclude(
				&controllers.CargueMasivoErrorController{},
			),
		),

		beego.NSNamespace("/solicitud_item",
			beego.NSInclude(
				&controllers.SolicitudItemController{},
			),
		),

		beego.NSNamespace("/solicitud_proveedor",
			beego.NSInclude(
				&controllers.SolicitudProveedorController{},
			),
		),

		beego.NSNamespace("/cargue_masivo",
			beego.NSInclude(
				&controllers.CargueMasivoController{},
			),
		),

		beego.NSNamespace("/documento_solicitud",
			beego.NSInclude(
				&controllers.DocumentoSolicitudController{},
			),
		),

		beego.NSNamespace("/solicitud_ciiu",
			beego.NSInclude(
				&controllers.SolicitudCiiuController{},
			),
		),

		beego.NSNamespace("/solicitud_cotizacion",
			beego.NSInclude(
				&controllers.SolicitudCotizacionController{},
			),
		),

		beego.NSNamespace("/solicitud_necesidad",
			beego.NSInclude(
				&controllers.SolicitudNecesidadController{},
			),
		),

		beego.NSNamespace("/solicitud_pago_parcial",
			beego.NSInclude(
				&controllers.SolicitudPagoParcialController{},
			),
		),

		beego.NSNamespace("/solicitud_requisito_documental",
			beego.NSInclude(
				&controllers.SolicitudRequisitoDocumentalController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
