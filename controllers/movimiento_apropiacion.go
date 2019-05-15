package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/fatih/structs"
	"github.com/udistrital/presupuesto_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/resposeformat"
)

// MovimientoApropiacionController operations for MovimientoApropiacion
type MovimientoApropiacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *MovimientoApropiacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// TotalMovimientosApropiacion ...
// @Title TotalMovimientosApropiacion
// @Description numero de movimientos segun vigencia
// @Param	vigencia		query 	string	true		"vigencia para la consulta del total de disponibilidades"
// @Param	UnidadEjecutora	query	string	false	"unidad ejecutora de las solicitudes a consultar"
// @Success 201 {int} total
// @Failure 403 vigencia is empty
// @router /TotalMovimientosApropiacion/:vigencia [get]
func (c *MovimientoApropiacionController) TotalMovimientosApropiacion() {
	vigenciaStr := c.Ctx.Input.Param(":vigencia")
	vigencia, err := strconv.Atoi(vigenciaStr)
	UnidadEjecutora, err2 := c.GetInt("UnidadEjecutora")
	if err == nil && err2 == nil {
		total, err := models.GetTotalMovimientosApropiacion(vigencia, UnidadEjecutora)
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

// RegistroSolicitudMovimientoApropiacion ...
// @Title RegistroSolicitudMovimientoApropiacion
// @Description create MovimientoApropiacion
// @Param	body		body 	models.MovimientoApropiacion	true		"body for MovimientoApropiacion content"
// @Success 201 {int} models.MovimientoApropiacion
// @Failure 403 body is empty
// @router /RegistroSolicitudMovimientoApropiacion [post]
func (c *MovimientoApropiacionController) RegistroSolicitudMovimientoApropiacion() {
	var v map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if res, err := models.RegistrarMovimietnoApropiaciontr(v); err == nil {
			resposeformat.SetResponseFormat(&c.Controller, res.Body, res.Code, 200)
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			resposeformat.SetResponseFormat(&c.Controller, err, "E_"+code, 500)
		}
	} else {
		resposeformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
}

// AprobarMovimietnoApropiacion ...
// @Title AprobarMovimietnoApropiacion
// @Description create MovimientoApropiacion
// @Param	body		body 	models.MovimientoApropiacion	true		"body for MovimientoApropiacion content"
// @Success 201 {int} models.MovimientoApropiacion
// @Failure 403 body is empty
// @router /AprobarMovimietnoApropiacion [post]
func (c *MovimientoApropiacionController) AprobarMovimietnoApropiacion() {
	var v models.MovimientoApropiacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil && v.Vigencia == time.Now().Year() {
		if res, err := models.AprobarMovimietnoApropiaciontr(&v); err == nil && res != nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = res
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			var alert []models.Alert
			alt := models.Alert{Type: "error", Code: "E_" + code, Body: err}
			alert = append(alert, alt)
			c.Data["json"] = alert
		}
	} else {
		var alert []models.Alert
		alt := models.Alert{}
		alt.Code = "E_0458"
		alt.Body = err
		alt.Type = "error"
		alert = append(alert, alt)
		c.Data["json"] = alert
	}
	c.ServeJSON()
}

// Post ...
// @Title Post
// @Description create MovimientoApropiacion
// @Param	body		body 	models.MovimientoApropiacion	true		"body for MovimientoApropiacion content"
// @Success 201 {int} models.MovimientoApropiacion
// @Failure 403 body is empty
// @router / [post]
func (c *MovimientoApropiacionController) Post() {
	var v models.MovimientoApropiacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddMovimientoApropiacion(&v); err == nil {
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
// @Description get MovimientoApropiacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.MovimientoApropiacion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MovimientoApropiacionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetMovimientoApropiacionById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetMovimientosApropiacionByApropiacion ...
// @Title GetMovimientosApropiacionByApropiacion
// @Description get MovimientoApropiacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.MovimientoApropiacion
// @Failure 403 :id is empty
// @router GetMovimientosApropiacionByApropiacion/:id [get]
func (c *MovimientoApropiacionController) GetMovimientosApropiacionByApropiacion() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.MovimientosByApropiacion(id)
	if err != nil {
		beego.Info(err)
		alertdb := structs.Map(err)
		var code string
		formatdata.FillStruct(alertdb["Code"], &code)
		alt := models.Alert{Type: "error", Code: "E_" + code, Body: err}
		c.Data["json"] = alt
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get MovimientoApropiacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.MovimientoApropiacion
// @Failure 403
// @router / [get]
func (c *MovimientoApropiacionController) GetAll() {
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

	l, err := models.GetAllMovimientoApropiacion(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the MovimientoApropiacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	fields		query 	string	true		"The fields you want to update"
// @Param	body		body 	models.MovimientoApropiacion	true		"body for MovimientoApropiacion content"
// @Success 200 {object} models.MovimientoApropiacion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MovimientoApropiacionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	var fields []string
	if f := c.GetString("fields"); f != "" {
		fields = strings.Split(f, ",")
	}
	id, _ := strconv.Atoi(idStr)
	v := models.MovimientoApropiacion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateMovimientoApropiacionById(&v, fields); err == nil {
			alt := models.Alert{}
			alt.Code = "S_MODP004"
			alt.Body = v
			alt.Type = "success"
			c.Data["json"] = alt
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alt := models.Alert{Type: "error", Code: "E_" + code, Body: err}
			c.Data["json"] = alt
		}
	} else {
		alt := models.Alert{}
		alt.Code = "E_0458"
		alt.Body = err
		alt.Type = "error"
		c.Data["json"] = alt
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the MovimientoApropiacion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MovimientoApropiacionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteMovimientoApropiacion(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
