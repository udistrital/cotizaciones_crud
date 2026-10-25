package controllers

import (
	"fmt"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

// filtrosSolicitud son los parametros de consulta por vigencia, estado y dependencia
// y el campo de solicitud_cotizacion al que corresponde cada uno.
var filtrosSolicitud = []struct {
	parametro string
	campo     string
}{
	{"vigencia", "Vigencia"},
	{"estado", "EstadoId"},
	{"dependencia", "DependenciaSolicitanteId"},
}

// agregarFiltrosSolicitud lee los parametros vigencia, estado y dependencia y los agrega
// a query. prefijo es la ruta desde la entidad consultada hasta su solicitud, por
// ejemplo "" para solicitud_cotizacion o "SolicitudId." para sus entidades asociadas.
func agregarFiltrosSolicitud(ctrl *beego.Controller, query map[string]string, prefijo string) error {
	for _, f := range filtrosSolicitud {
		v := ctrl.GetString(f.parametro)
		if v == "" {
			continue
		}
		if _, err := strconv.Atoi(v); err != nil {
			return fmt.Errorf("el parametro %s debe ser numerico", f.parametro)
		}
		query[prefijo+f.campo] = v
	}
	return nil
}
