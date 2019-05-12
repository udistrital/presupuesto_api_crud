package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

type Date struct {
	beego.Controller
}

// FechaActual ...
// @Title FechaActual
// @Description retorba fecga del servidor
// @Param	formato		formato como quiere obtener la fecha
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /FechaActual/:formato
func (c *Date) FechaActual() {
	formato := c.Ctx.Input.Param(":formato")
	hoy := time.Now()
	fechaActual := hoy.Format(formato)

	c.Data["json"] = fechaActual

	c.ServeJSON()
}
