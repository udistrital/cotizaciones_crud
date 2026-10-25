package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoErrorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoErrorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoErrorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoErrorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoErrorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoErrorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoErrorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoErrorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoErrorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:CargueMasivoErrorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCiiuController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCiiuController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCiiuController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCiiuController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCiiuController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method:           "GetDetalle",
			Router:           `/:id/detalle`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudItemController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudItemController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudItemController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudItemController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudItemController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudItemController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudNecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudNecesidadController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudPagoParcialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudPagoParcialController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudPagoParcialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudPagoParcialController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudPagoParcialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudPagoParcialController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudPagoParcialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudPagoParcialController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudPagoParcialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudPagoParcialController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudProveedorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudProveedorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudProveedorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudProveedorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudProveedorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudRequisitoDocumentalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudRequisitoDocumentalController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudRequisitoDocumentalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudRequisitoDocumentalController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudRequisitoDocumentalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudRequisitoDocumentalController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudRequisitoDocumentalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudRequisitoDocumentalController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudRequisitoDocumentalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cotizaciones_crud/controllers:SolicitudRequisitoDocumentalController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
