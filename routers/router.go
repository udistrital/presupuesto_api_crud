// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/presupuesto_crud/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/rubro",
			beego.NSInclude(
				&controllers.RubroController{},
			),
		),
		beego.NSNamespace("/afectacion_concepto",
			beego.NSInclude(
				&controllers.AfectacionConceptoController{},
			),
		),
		beego.NSNamespace("/apropiacion",
			beego.NSInclude(
				&controllers.ApropiacionController{},
			),
		),
		beego.NSNamespace("/arbol_conceptos",
			beego.NSInclude(
				&controllers.ArbolConceptosController{},
			),
		),
		beego.NSNamespace("/concepto",
			beego.NSInclude(
				&controllers.ConceptoController{},
			),
		),
		beego.NSNamespace("/concepto_cuenta_contable",
			beego.NSInclude(
				&controllers.ConceptoCuentaContableController{},
			),
		),
		beego.NSNamespace("/concepto_detalle_tipo_transaccion",
			beego.NSInclude(
				&controllers.ConceptoDetalleTipoTransaccionController{},
			),
		),
		beego.NSNamespace("/cuenta_bancaria",
			beego.NSInclude(
				&controllers.CuentaBancariaController{},
			),
		),
		beego.NSNamespace("/cuenta_especial",
			beego.NSInclude(
				&controllers.CuentaEspecialController{},
			),
		),
		beego.NSNamespace("/estado_ingreso",
			beego.NSInclude(
				&controllers.EstadoIngresoController{},
			),
		),
		beego.NSNamespace("/estado_ingreso_sin_situacion_fondos",
			beego.NSInclude(
				&controllers.EstadoIngresoSinSituacionFondosController{},
			),
		),
		beego.NSNamespace("/forma_ingreso",
			beego.NSInclude(
				&controllers.FormaIngresoController{},
			),
		),
		beego.NSNamespace("/forma_pago",
			beego.NSInclude(
				&controllers.FormaIngresoController{},
			),
		),
		beego.NSNamespace("/fuente_financiamiento",
			beego.NSInclude(
				&controllers.FuenteFinanciamientoController{},
			),
		),
		beego.NSNamespace("/ingreso",
			beego.NSInclude(
				&controllers.IngresoController{},
			),
		),
		beego.NSNamespace("/ingreso_sin_situacion_fondos",
			beego.NSInclude(
				&controllers.IngresoSinSituacionFondosController{},
			),
		),
		beego.NSNamespace("/ingreso_sin_situacion_fondos_estado",
			beego.NSInclude(
				&controllers.IngresoSinSituacionFondosEstadoController{},
			),
		),
		beego.NSNamespace("/movimiento_contable",
			beego.NSInclude(
				&controllers.MovimientoContableController{},
			),
		),
		beego.NSNamespace("/orden_pago",
			beego.NSInclude(
				&controllers.OrdenPagoController{},
			),
		),
		beego.NSNamespace("/tipo_cuenta_bancaria",
			beego.NSInclude(
				&controllers.TipoCuentaBancariaController{},
			),
		),
		beego.NSNamespace("/unidad_ejecutora",
			beego.NSInclude(
				&controllers.UnidadEjecutoraController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
