// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/cotizaciones_crud/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/hist_mst_solicitud_cotizacion_estado_hist",
			beego.NSInclude(
				&controllers.HistMstSolicitudCotizacionEstadoHistController{},
			),
		),

		beego.NSNamespace("/mst_cotizacion",
			beego.NSInclude(
				&controllers.MstCotizacionController{},
			),
		),

		beego.NSNamespace("/mst_solicitud_cotizacion",
			beego.NSInclude(
				&controllers.MstSolicitudCotizacionController{},
			),
		),

		beego.NSNamespace("/forma_pago",
			beego.NSInclude(
				&controllers.FormaPagoController{},
			),
		),

		beego.NSNamespace("/solicitud_cotizacion_actividad_economica",
			beego.NSInclude(
				&controllers.SolicitudCotizacionActividadEconomicaController{},
			),
		),

		beego.NSNamespace("/rel_solicitud_cotizacion_proveedor",
			beego.NSInclude(
				&controllers.RelSolicitudCotizacionProveedorController{},
			),
		),

		beego.NSNamespace("/solicitud_cotizacion_observacion_proveedor",
			beego.NSInclude(
				&controllers.SolicitudCotizacionObservacionProveedorController{},
			),
		),

		beego.NSNamespace("/solicitud_cotizacion_observacion_proveedor_respuesta",
			beego.NSInclude(
				&controllers.SolicitudCotizacionObservacionProveedorRespuestaController{},
			),
		),

		beego.NSNamespace("/validacion_ordenador",
			beego.NSInclude(
				&controllers.ValidacionOrdenadorController{},
			),
		),

		beego.NSNamespace("/cargue_masivo",
			beego.NSInclude(
				&controllers.CargueMasivoController{},
			),
		),

		beego.NSNamespace("/hist_det_cotizacion",
			beego.NSInclude(
				&controllers.HistDetCotizacionController{},
			),
		),

		beego.NSNamespace("/hist_det_solicitud_cotizacion",
			beego.NSInclude(
				&controllers.HistDetSolicitudCotizacionController{},
			),
		),

		beego.NSNamespace("/hist_mst_solicitud_cotizacion",
			beego.NSInclude(
				&controllers.HistMstSolicitudCotizacionController{},
			),
		),

		beego.NSNamespace("/solicitud_cotizacion_especificacion_tecnica",
			beego.NSInclude(
				&controllers.SolicitudCotizacionEspecificacionTecnicaController{},
			),
		),

		beego.NSNamespace("/decision_solicitante",
			beego.NSInclude(
				&controllers.DecisionSolicitanteController{},
			),
		),

		beego.NSNamespace("/hist_mst_cotizacion",
			beego.NSInclude(
				&controllers.HistMstCotizacionController{},
			),
		),

		beego.NSNamespace("/notificacion",
			beego.NSInclude(
				&controllers.NotificacionController{},
			),
		),

		beego.NSNamespace("/solicitud_cotizacion_nucleo_basico",
			beego.NSInclude(
				&controllers.SolicitudCotizacionNucleoBasicoController{},
			),
		),
		beego.NSNamespace("/rel_hist_mst_det_cotizacion",
			beego.NSInclude(
				&controllers.RelHistMstDetCotizacionController{},
			),
		),

		beego.NSNamespace("/rel_hist_mst_det_solicitud_cotizacion",
			beego.NSInclude(
				&controllers.RelHistMstDetSolicitudCotizacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
