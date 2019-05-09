package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/udistrital/presupuesto_crud/models"
	"github.com/udistrital/utils_oas/formatdata"

	"github.com/astaxie/beego"
)

// CompromisoController operations for Compromiso
type CompromisoController struct {
	beego.Controller
}

// URLMapping ...
func (c *CompromisoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Compromiso
// @Param	body		body 	models.Compromiso	true		"body for Compromiso content"
// @Success 201 {int} models.Compromiso
// @Failure 403 body is empty
// @router / [post]
func (c *CompromisoController) Post() {
	var v models.Compromiso
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.FechaModificacion = time.Now()
		if _, err = models.AddCompromiso(&v); err == nil {
			alert := models.Alert{Type: "success", Code: "S_543", Body: v.Id} //codigo de registro exitoso
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alert
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
			c.Data["json"] = alert
		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Compromiso by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Compromiso
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CompromisoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCompromisoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Compromiso
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Compromiso
// @Failure 403
// @router / [get]
func (c *CompromisoController) GetAll() {
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

	l, err := models.GetAllCompromiso(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description El compromiso unicamente actualiza el objeto y las fechas de inicio y fin
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Compromiso	true		"body for Compromiso content"
// @Success 200 {object} models.Compromiso
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CompromisoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	uv := models.Compromiso{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &uv); err == nil {
		if v, err := models.GetCompromisoById(uv.Id); err == nil {
			if uv.EstadoCompromiso.Id == 1 {
				v.Objeto = uv.Objeto
				v.FechaInicio = uv.FechaInicio
				v.FechaFin = uv.FechaFin
				v.FechaModificacion = time.Now()
				v.Vigencia = uv.Vigencia
				//v.Vigencia = float64(uv.FechaInicio.Year())
				if err = models.UpdateCompromisoById(v); err == nil {
					alert := models.Alert{Type: "success", Code: "S_542", Body: v.Id} //codigo de registro exitoso
					c.Ctx.Output.SetStatus(201)
					c.Data["json"] = alert
				} else {
					alertdb := structs.Map(err)
					var code string
					formatdata.FillStruct(alertdb["Code"], &code)
					alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
					c.Data["json"] = alert
				}
			} else {
				alert := models.Alert{Type: "error", Code: "E_0471", Body: "No es posible realizar actualizacion debido al estado del registro"}
				c.Data["json"] = alert
			}

		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
			c.Data["json"] = alert
		}

	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Compromiso
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CompromisoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if v, err := models.GetCompromisoById(id); err == nil {
		v.EstadoCompromiso.Id = 4 // Id compromiso cancelado
		if err = models.UpdateCompromisoById(v); err == nil {
			alert := models.Alert{Type: "success", Code: "S_5412", Body: v.Id} //codigo de cambio exitoso
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alert
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
			c.Data["json"] = alert
		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
		c.Data["json"] = alert
	}
	/*if err := models.DeleteCompromiso(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}*/
	c.ServeJSON()
}
