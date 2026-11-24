package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DecisionSolicitanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DecisionSolicitanteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DecisionSolicitanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DecisionSolicitanteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DecisionSolicitanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DecisionSolicitanteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DecisionSolicitanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DecisionSolicitanteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DecisionSolicitanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DecisionSolicitanteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:FormaPagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:FormaPagoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:FormaPagoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:FormaPagoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:FormaPagoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetCotizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetCotizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetCotizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetCotizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetCotizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistDetSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstCotizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstCotizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstCotizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstCotizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstCotizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionEstadoHistController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionEstadoHistController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionEstadoHistController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionEstadoHistController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionEstadoHistController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionEstadoHistController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionEstadoHistController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionEstadoHistController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionEstadoHistController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:HistMstSolicitudCotizacionEstadoHistController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstCotizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstCotizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstCotizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstCotizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstCotizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:MstSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetCotizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetCotizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetCotizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetCotizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetCotizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetSolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelHistMstDetSolicitudCotizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelSolicitudCotizacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelSolicitudCotizacionProveedorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelSolicitudCotizacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelSolicitudCotizacionProveedorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelSolicitudCotizacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelSolicitudCotizacionProveedorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelSolicitudCotizacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelSolicitudCotizacionProveedorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelSolicitudCotizacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:RelSolicitudCotizacionProveedorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionActividadEconomicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionActividadEconomicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionActividadEconomicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionActividadEconomicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionActividadEconomicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionActividadEconomicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionActividadEconomicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionActividadEconomicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionActividadEconomicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionActividadEconomicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionEspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionEspecificacionTecnicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionEspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionEspecificacionTecnicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionEspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionEspecificacionTecnicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionEspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionEspecificacionTecnicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionEspecificacionTecnicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionEspecificacionTecnicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionNucleoBasicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionNucleoBasicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionNucleoBasicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionNucleoBasicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionNucleoBasicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorRespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorRespuestaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorRespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorRespuestaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorRespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorRespuestaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorRespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorRespuestaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorRespuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionObservacionProveedorRespuestaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:ValidacionOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:ValidacionOrdenadorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:ValidacionOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:ValidacionOrdenadorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:ValidacionOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:ValidacionOrdenadorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:ValidacionOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:ValidacionOrdenadorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:ValidacionOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:ValidacionOrdenadorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
