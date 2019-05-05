package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/fatih/structs"
	"github.com/udistrital/presupuesto_api_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
)

// RegistroPresupuestalController operations for RegistroPresupuestal
type RegistroPresupuestalController struct {
	beego.Controller
}

// URLMapping ...
func (c *RegistroPresupuestalController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("SaldoRp", c.SaldoRp)
	c.Mapping("Anular", c.Anular)
}

// Post ...
// @Title Post
// @Description create RegistroPresupuestal
// @Param	body		body 	models.RegistroPresupuestal	true		"body for RegistroPresupuestal content"
// @Success 201 {int} models.RegistroPresupuestal
// @Failure 403 body is empty
// @router / [post]
func (c *RegistroPresupuestalController) Post() {
	var v models.DatosRegistroPresupuestal
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddRegistoPresupuestal(&v); err == nil {
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
// @Description get RegistroPresupuestal by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.RegistroPresupuestal
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RegistroPresupuestalController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRegistroPresupuestalById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get RegistroPresupuestal
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.RegistroPresupuestal
// @Failure 403
// @router / [get]
func (c *RegistroPresupuestalController) GetAll() {
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

	l, err := models.GetAllRegistroPresupuestal(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the RegistroPresupuestal
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.RegistroPresupuestal	true		"body for RegistroPresupuestal content"
// @Success 200 {object} models.RegistroPresupuestal
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RegistroPresupuestalController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.RegistroPresupuestal{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateRegistroPresupuestalById(&v); err == nil {
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
// @Description delete the RegistroPresupuestal
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RegistroPresupuestalController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteRegistroPresupuestal(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// SaldoRp ...
// @Title SaldoRp
// @Description create RegistroPresupuestal
// @Param	body		body 	models.RegistroPresupuestal	true		"body for RegistroPresupuestal content"
// @Success 201 {int} models.RegistroPresupuestal
// @Failure 403 body is empty
// @router SaldoRp/ [post]
func (c *RegistroPresupuestalController) SaldoRp() {
	var v models.DatosSaldoRp
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		var ff int
		if v.FuenteFinanciacion != nil {
			ff = v.FuenteFinanciacion.Id
		} else {
			ff = 0
		}
		saldo, comprometido, anulado, err := models.SaldoRp(v.Rp.Id, v.Apropiacion.Id, ff)
		if err != nil {
			c.Data["json"] = err
		} else {
			var m map[string]float64
			m = make(map[string]float64)
			m["saldo"] = saldo
			m["comprometido"] = comprometido
			m["anulado"] = anulado
			c.Data["json"] = m
		}
	} else {
		c.Data["json"] = err
		fmt.Println("error: ", err)
	}

	c.ServeJSON()
}

// Anular ...
// @Title Anular
// @Description create RegistroPresupuestal
// @Param	body		body 	models.RegistroPresupuestal	true		"body for RegistroPresupuestal content"
// @Success 201 {int} models.RegistroPresupuestal
// @Failure 403 body is empty
// @router Anular/ [post]
func (c *RegistroPresupuestalController) Anular() {
	var v models.Info_rp_a_anular
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.Anulacion.TipoAnulacion.Id == 2 || v.Anulacion.TipoAnulacion.Id == 3 {
			alertas, err := models.AnulacionTotalRp(&v)
			if err != nil {
				c.Data["json"] = err
			} else {
				c.Data["json"] = alertas
			}
		} else if v.Anulacion.TipoAnulacion.Id == 1 {
			alertas, err := models.AnulacionParcialRp(&v)
			if err != nil {
				c.Data["json"] = err
			} else {
				c.Data["json"] = alertas
			}
		} else {
			c.Data["json"] = "No se pudo cargar el tipo de la anulacion"
		}

	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

// ValorTotalRp ...
// @Title Valor Total Rp
// @Description retorna valor total del RegistroPresupuestal por id
// @Param	body		body 	models.RegistroPresupuestal	true		"body for RegistroPresupuestal content"
// @Success 201 {int} suma valor de los RegistroPresupuestal por id
// @Failure 403 body is empty
// @router ValorTotalRp/:id [get]

func (c *RegistroPresupuestalController) ValorTotalRp() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetValorTotalRp(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// ValorActualRp ...
// @Title Valor Actual Rp
// @Description retorna valor actual del RegistroPresupuestal por id
// @Param	body		body 	models.RegistroPresupuestal	true		"body for RegistroPresupuestal content"
// @Success 201 {int} suma valor de los RegistroPresupuestal por id
// @Failure 403 body is empty
// @router ValorActualRp/:id [get]
func (c *RegistroPresupuestalController) ValorActualRp() {
	idStr := c.Ctx.Input.Param(":id")
	if idStr != "" {
		id, _ := strconv.Atoi(idStr)
		v, err := models.GetValorActualRp(id)
		if err != nil {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
			c.Data["json"] = alert
		} else {
			c.Data["json"] = v
		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: "No Id defined"}
		c.Data["json"] = alert
	}

	c.ServeJSON()
}

// AprobarAnulacionRegistroresupuestal ...
// @Title AprobarAnulacionRegistroresupuestal
// @Description aprueba la anulacion de un rp ya sea total o parcial
// @Param	body		body 	models.AnulacionRegistroPresupuestal	true		"body for AnulacionRegistroPresupuestal content"
// @Success 201 {int} models.AnulacionRegistroPresupuestal
// @Failure 403 body is empty
// @router /AprobarAnulacion [post]
func (c *RegistroPresupuestalController) AprobarAnulacionRp() {
	var v models.AnulacionRegistroPresupuestal
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if alert, err := models.AprobacionAnulacionRp(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alert
		} else {
			c.Data["json"] = alert
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// TotalRp ...
// @Title TotalRp
// @Description numero de rp segun vigencia o rango de fechas
// @Param	vigencia		query 	string	true		"vigencia para la consulta del total de rp"
// @Param	UnidadEjecutora	query	string	false	"unidad ejecutora de las solicitudes a consultar"
// @Param	rangoinicio		query 	string	true		"opcional, fecha inicio de consulta de rp"
// @Param	rangofin		query 	string	true		"opcional, fecha fin de consulta de rp"
// @Success 201 {int} total
// @Failure 403 vigencia is empty
// @router /TotalRp/:vigencia [get]
func (c *RegistroPresupuestalController) TotalRp() {
	vigenciaStr := c.Ctx.Input.Param(":vigencia")
	vigencia, err := strconv.Atoi(vigenciaStr)
	var startrange string
	var endrange string
	if r := c.GetString("rangoinicio"); r != "" {
		startrange = r

	}
	if r := c.GetString("rangofin"); r != "" {
		endrange = r
	}
	UnidadEjecutora, err2 := c.GetInt("UnidadEjecutora")
	if err == nil && err2 == nil {
		total, err := models.GetTotalRp(vigencia, UnidadEjecutora, startrange, endrange)
		if err == nil {
			c.Data["json"] = total
		} else {
			c.Data["json"] = models.Alert{Code: "E_0458", Body: err.Error(), Type: "error"}
		}
	} else {
		c.Data["json"] = models.Alert{Code: "E_0458", Body: "Not enough parameter", Type: "error"}
	}

	c.ServeJSON()
}

// DeleteRpData ...
// @Title DeleteRpData
// @Description delete the Disponibilidad
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /DeleteRpData/:id [delete]
func (c *RegistroPresupuestalController) DeleteRpData() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteRpData(id); err == nil {
		c.Data["json"] = models.Alert{Code: "S_554", Body: nil, Type: "success"}
	} else {
		c.Data["json"] = models.Alert{Code: "E_0458", Body: err.Error(), Type: "error"}
	}
	c.ServeJSON()
}
