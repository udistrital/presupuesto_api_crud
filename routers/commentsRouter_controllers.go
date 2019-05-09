package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AfectacionConceptoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AfectacionConceptoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AfectacionConceptoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AfectacionConceptoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:AfectacionConceptoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ArbolConceptosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ArbolConceptosController"],
        beego.ControllerComments{
            Method: "MakeTree",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CompromisoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CompromisoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CompromisoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CompromisoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CompromisoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoConceptoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoConceptoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoConceptoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoConceptoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoConceptoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoCuentaContableController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoCuentaContableController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoCuentaContableController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoCuentaContableController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoCuentaContableController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoDetalleTipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoDetalleTipoTransaccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoDetalleTipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoDetalleTipoTransaccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoDetalleTipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoDetalleTipoTransaccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoDetalleTipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoDetalleTipoTransaccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoDetalleTipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:ConceptoDetalleTipoTransaccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaBancariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaEspecialController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaEspecialController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaEspecialController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaEspecialController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:CuentaEspecialController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
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
            Method: "TotalDisponibilidades",
            Router: `/TotalDisponibilidades/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "Anular",
            Router: `Anular/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:DisponibilidadController"],
        beego.ControllerComments{
            Method: "SaldoCdp",
            Router: `SaldoCdp/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoSinSituacionFondosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoSinSituacionFondosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoSinSituacionFondosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoSinSituacionFondosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:EstadoIngresoSinSituacionFondosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FormaIngresoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FormaIngresoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FormaIngresoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FormaIngresoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:FormaIngresoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
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

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"],
        beego.ControllerComments{
            Method: "AprobacionContableIngreso",
            Router: `/AprobacionContableIngreso`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"],
        beego.ControllerComments{
            Method: "AprobacionPresupuestalIngreso",
            Router: `/AprobacionPresupuestalIngreso`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"],
        beego.ControllerComments{
            Method: "CreateIngresos",
            Router: `/CreateIngresos`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"],
        beego.ControllerComments{
            Method: "RechazoContableIngreso",
            Router: `/RechazoContableIngreso`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoController"],
        beego.ControllerComments{
            Method: "RechazoPresupuestalIngreso",
            Router: `/RechazoPresupuestalIngreso`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:IngresoSinSituacionFondosEstadoController"],
        beego.ControllerComments{
            Method: "ChangeExistingStates",
            Router: `ChangeExistingStates/`,
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

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"],
        beego.ControllerComments{
            Method: "GetSumMovimientos",
            Router: `/GetSumMovimientos/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:MovimientoContableController"],
        beego.ControllerComments{
            Method: "PostArray",
            Router: `PostArray/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"],
        beego.ControllerComments{
            Method: "GetOrdenPagoByEstado",
            Router: `/GetOrdenPagoByEstado`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"],
        beego.ControllerComments{
            Method: "ActualizarOpProveedor",
            Router: `ActualizarOpProveedor`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"],
        beego.ControllerComments{
            Method: "FechaActual",
            Router: `FechaActual/:formato`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"],
        beego.ControllerComments{
            Method: "RegistrarOpProveedor",
            Router: `RegistrarOpProveedor`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:OrdenPagoController"],
        beego.ControllerComments{
            Method: "ValorTotal",
            Router: `ValorTotal/:id`,
            AllowHTTPMethods: []string{"post"},
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

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoCuentaBancariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoCuentaBancariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoCuentaBancariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoCuentaBancariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:TipoCuentaBancariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:UnidadEjecutoraController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:UnidadEjecutoraController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:UnidadEjecutoraController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:UnidadEjecutoraController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/presupuesto_crud/controllers:UnidadEjecutoraController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
