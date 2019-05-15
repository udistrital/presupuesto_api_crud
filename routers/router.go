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
	)

	beego.AddNamespace(ns)
}
