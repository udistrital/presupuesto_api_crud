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
		beego.NSNamespace("/rubro_rubro",
			beego.NSInclude(
				&controllers.RubroRubroController{},
			),
		),
		beego.NSNamespace("/apropiacion",
			beego.NSInclude(
				&controllers.ApropiacionController{},
			),
		),
		beego.NSNamespace("/producto",
			beego.NSInclude(
				&controllers.ProductoController{},
			),
		),
		beego.NSNamespace("/producto_rubro",
			beego.NSInclude(
				&controllers.ProductoRubroController{},
			),
		),
		beego.NSNamespace("/date",
			beego.NSInclude(
				&controllers.Date{},
			),
		),
		beego.NSNamespace("/tipo_movimiento_apropiacion",
			beego.NSInclude(
				&controllers.TipoMovimientoApropiacionController{},
			),
		),
		beego.NSNamespace("/movimiento_apropiacion",
			beego.NSInclude(
				&controllers.MovimientoApropiacionController{},
			),
		),
		beego.NSNamespace("/disponibilidad",
			beego.NSInclude(
				&controllers.DisponibilidadController{},
			),
		),
		beego.NSNamespace("/fuente_financiamiento",
			beego.NSInclude(
				&controllers.FuenteFinanciamientoController{},
			),
		),
		beego.NSNamespace("/disponibilidad_apropiacion",
			beego.NSInclude(
				&controllers.DisponibilidadApropiacionController{},
			),
		),
		beego.NSNamespace("/tipo_anulacion_presupuestal",
			beego.NSInclude(
				&controllers.TipoAnulacionPresupuestalController{},
			),
		),
		beego.NSNamespace("/anulacion_disponibilidad",
			beego.NSInclude(
				&controllers.AnulacionDisponibilidadController{},
			),
		),
		beego.NSNamespace("/registro_presupuestal",
			beego.NSInclude(
				&controllers.RegistroPresupuestalController{},
			),
		),
		beego.NSNamespace("/registro_presupuestal_disponibilidad_apropiacion",
			beego.NSInclude(
				&controllers.RegistroPresupuestalDisponibilidadApropiacionController{},
			),
		),
		beego.NSNamespace("/anulacion_registro_presupuestal",
			beego.NSInclude(
				&controllers.AnulacionRegistroPresupuestalController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
