package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/udistrital/presupuesto_crud/models"

	"github.com/astaxie/beego"
)

// OrdenPagoController operations for OrdenPago
type OrdenPagoController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrdenPagoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("RegistrarOpProveedor", c.RegistrarOpProveedor)
	c.Mapping("ActualizarOpProveedor", c.ActualizarOpProveedor)
	c.Mapping("FechaActual", c.FechaActual)
	c.Mapping("ValorTotal", c.ValorTotal)
}

// Post ...
// @Title Post
// @Description create OrdenPago
// @Param	body		body 	models.OrdenPago	true		"body for OrdenPago content"
// @Success 201 {int} models.OrdenPago
// @Failure 403 body is empty
// @router / [post]
func (c *OrdenPagoController) Post() {
	var v models.OrdenPago
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddOrdenPago(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get OrdenPago by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.OrdenPago
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OrdenPagoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOrdenPagoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get OrdenPago
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.OrdenPago
// @Failure 403
// @router / [get]
func (c *OrdenPagoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllOrdenPago(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the OrdenPago
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.OrdenPago	true		"body for OrdenPago content"
// @Success 200 {object} models.OrdenPago
// @Failure 403 :id is not int
// @router /:id [put]
func (c *OrdenPagoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.OrdenPago{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateOrdenPagoById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the OrdenPago
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OrdenPagoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOrdenPago(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// RegistrarOpProveedor ...
// @Title RegistrarOpProveedor
// @Description Registrar orden_pago de proveedor, concepto_ordenpago, mivimientos contables
// @Param	body		body 	models.OrdenPago	true		"body for OrdenPago content"
// @Success 201 {int} models.OrdenPago
// @Failure 403 body is empty
// @router RegistrarOpProveedor [post]
func (c *OrdenPagoController) RegistrarOpProveedor() {
	var v interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		m := v.(map[string]interface{})
		mensaje := models.RegistrarOpProveedor(m)
		if mensaje.Type != "success" {
			c.Data["json"] = mensaje
		} else {
			c.Ctx.Output.SetStatus(201)
			//alert := models.Alert{Type: mensaje.Type, Code: mensaje.Code, Body: consecutivoOp}
			c.Data["json"] = mensaje
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

// ActualizarOpProveedor ...
// @Title ActualizarOpProveedor
// @Description Actualiza orden_pago de proveedor, concepto_ordenpago, mivimientos contables
// @Param	body		body 	models.OrdenPago	true		"body for OrdenPago content"
// @Success 201 {int} models.OrdenPago
// @Failure 403 body is empty
// @router ActualizarOpProveedor [post]
func (c *OrdenPagoController) ActualizarOpProveedor() {
	var v interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		m := v.(map[string]interface{})
		alerta, err, consecutivoOp := models.ActualizarOpProveedor(m)
		if err != nil {
			c.Data["json"] = alerta
		} else {
			alert := models.Alert{Type: "success", Code: "S_OPP_02", Body: consecutivoOp}
			c.Data["json"] = alert
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

// FechaActual ...
// @Title FechaActual
// @Description retorba fecga del servidor
// @Param	formato		formato como quiere obtener la fecha
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router FechaActual/:formato
func (c *OrdenPagoController) FechaActual() {
	formatoInput := c.Ctx.Input.Param(":formato")
	fechaActual, err := models.FechaActual(formatoInput)
	if err == nil {
		c.Data["json"] = fechaActual
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// ValorTotal ...
// @Title ValorTotal
// @Description Valor Total of the OrdenPago
// @Param	id		path 	string	true		"The id of orden pago"
// @Param	body		body 	models.OrdenPago	true		"body for OrdenPago content"
// @Success 200 {object} models.OrdenPago
// @Failure 403 :id is not int
// @router ValorTotal/:id [post]
func (c *OrdenPagoController) ValorTotal() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if valorTotal, err := models.ValorTotal(id); err == nil {
		c.Data["json"] = valorTotal
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOrdenPagoByEstado ...
// @Title GetOrdenPagoByEstado
// @Description return OrdenPago by Estado
// @Param	codigoEstado		query string true		"code estado OrdenPago for search"
// @Param	vigencia		query string true		"vigencia of OrdenPago for search"
// @Param	tipoOp		query string true		"tipo OrdenPago for search"
// @Param	formaPago		query string true		"forma pago OrdenPago for search"
// @Success 201 {int} models.OrdenPago
// @Failure 403 body is empty
// @router /GetOrdenPagoByEstado [get]
func (c *OrdenPagoController) GetOrdenPagoByEstado() {
	estadoOrden := c.GetString("codigoEstado")
	vigenciaOrden := c.GetString("vigencia")
	tipoOrden := c.GetString("tipoOp")
	formaPago := c.GetString("formaPago")
	if estadoOrden != "" && vigenciaOrden != "" && tipoOrden != "" && formaPago != "" {
		Ordenes, mensaje := models.GetOrdenPagoByEstado(estadoOrden, vigenciaOrden, tipoOrden, formaPago)
		if mensaje.Type == "success" {
			c.Data["json"] = Ordenes
		} else {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = mensaje
		}
	} else {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = models.Alert{Code: "E_0458", Body: "Not enough parameter", Type: "error"}
	}
	c.ServeJSON()
}
