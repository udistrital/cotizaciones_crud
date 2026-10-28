package controllers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

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

// leerActivo lee el parametro activo de las consultas: true (predeterminado) para
// registros activos o false para registros inactivos (borrado logico).
func leerActivo(ctrl *beego.Controller) (bool, error) {
	switch ctrl.GetString("activo") {
	case "", "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, errors.New("el parametro activo debe ser true o false")
}

// agregarFiltroActivo agrega a query el filtro por Activo de las consultas GetAll.
// El parametro activo tiene prioridad; si no se envia, se respeta Activo dentro del
// parametro query y, si tampoco esta, se consultan solo los registros activos.
func agregarFiltroActivo(ctrl *beego.Controller, query map[string]string) error {
	if ctrl.GetString("activo") == "" {
		for k := range query {
			if strings.EqualFold(k, "Activo") {
				return nil
			}
		}
	}
	activo, err := leerActivo(ctrl)
	if err != nil {
		return err
	}
	for k := range query {
		if strings.EqualFold(k, "Activo") {
			delete(query, k)
		}
	}
	query["Activo"] = strconv.FormatBool(activo)
	return nil
}
