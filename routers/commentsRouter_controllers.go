package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionDisponibilidadController"],
        beego.ControllerComments{
            Method: "TotalAnulacionDisponibilidad",
            Router: `/TotalAnulacionDisponibilidad/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AnulacionRegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "TotalAnulacionRegistroPresupuestal",
            Router: `/TotalAnulacionRegistroPresupuestal/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "AprobarPresupuesto",
            Router: `/AprobacionAsignacionInicial/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "UpdateApropiacionValue",
            Router: `/UpdateApropiacionValue/:id/:valor`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "VigenciaApropiaciones",
            Router: `/VigenciaApropiaciones`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:Date"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:Date"],
        beego.ControllerComments{
            Method: "FechaActual",
            Router: `/FechaActual/:formato`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadApropiacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadApropiacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadApropiacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadApropiacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadApropiacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "Anular",
            Router: `/Anular`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "AprobarAnulacionDisponibilidad",
            Router: `/AprobarAnulacion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "DeleteDisponibilidadData",
            Router: `/DeleteDisponibilidadData/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "DeleteDisponibilidadMovimiento",
            Router: `/DeleteDisponibilidadMovimiento/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "GetPrincDisponibilidadInfo",
            Router: `/GetPrincDisponibilidadInfo/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "SaldoCdp",
            Router: `/SaldoCdp`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "TotalDisponibilidades",
            Router: `/TotalDisponibilidades/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "DeleteModificacionFuenteFinanciamientoTr",
            Router: `/DeleteModificacionFuenteFinanciamientoTr`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "DeleteMovimientoFuenteFinanciamientoTr",
            Router: `/DeleteMovimientoFuenteFinanciamientoTr/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "MovimientoFuenteFinanciamientoTr",
            Router: `/MovimientoFuenteFinanciamientoTr`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FuenteFinanciamientoController"],
        beego.ControllerComments{
            Method: "RegistrarFuenteFinanciamientoTr",
            Router: `/RegistrarFuenteFinanciamientoTr`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "AprobarMovimietnoApropiacion",
            Router: `/AprobarMovimietnoApropiacion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "RegistroSolicitudMovimientoApropiacion",
            Router: `/RegistroSolicitudMovimientoApropiacion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "TotalMovimientosApropiacion",
            Router: `/TotalMovimientosApropiacion/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "GetMovimientosApropiacionByApropiacion",
            Router: `GetMovimientosApropiacionByApropiacion/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoController"],
        beego.ControllerComments{
            Method: "TotalProductos",
            Router: `/TotalProductos/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"],
        beego.ControllerComments{
            Method: "AddProductoRubrotr",
            Router: `/AddProductoRubrotr/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ProductoRubroController"],
        beego.ControllerComments{
            Method: "SetVariacionProducto",
            Router: `/SetVariacionProducto/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "AprobarAnulacionRp",
            Router: `/AprobarAnulacion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "DeleteRpData",
            Router: `/DeleteRpData/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "TotalRp",
            Router: `/TotalRp/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "Anular",
            Router: `Anular/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "SaldoRp",
            Router: `SaldoRp/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalController"],
        beego.ControllerComments{
            Method: "ValorActualRp",
            Router: `ValorActualRp/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:RubroRubroController"],
        beego.ControllerComments{
            Method: "DeleteRubroRelation",
            Router: `/DeleteRubroRelation/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoAnulacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoAnulacionPresupuestalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoAnulacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoAnulacionPresupuestalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoAnulacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoAnulacionPresupuestalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoAnulacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoAnulacionPresupuestalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoAnulacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoAnulacionPresupuestalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoMovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoMovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoMovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoMovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoMovimientoApropiacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
