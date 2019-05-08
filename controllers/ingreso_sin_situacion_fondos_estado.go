package controllers

import (
	"encoding/json"
	"errors"
	"github.com/udistrital/presupuesto_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/astaxie/beego"
)

// IngresoSinSituacionFondosEstadoController operations for IngresoSinSituacionFondosEstado
type IngresoSinSituacionFondosEstadoController struct {
	beego.Controller
}

// URLMapping ...
func (c *IngresoSinSituacionFondosEstadoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create IngresoSinSituacionFondosEstado
// @Param	body		body 	models.IngresoSinSituacionFondosEstado	true		"body for IngresoSinSituacionFondosEstado content"
// @Success 201 {int} models.IngresoSinSituacionFondosEstado
// @Failure 403 body is empty
// @router / [post]
func (c *IngresoSinSituacionFondosEstadoController) Post() {
	var v models.IngresoSinSituacionFondosEstado
	var code string
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddIngresoSinSituacionFondosEstado(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			alert := models.Alert{Type: "success", Code: "S_543", Body: v}
			c.Data["json"] = alert
		} else {
			alertdb := structs.Map(err)
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type:"error",Code:"E_" + code,Body:err}
			c.Data["json"] = alert
		}

	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get IngresoSinSituacionFondosEstado by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.IngresoSinSituacionFondosEstado
// @Failure 403 :id is empty
// @router /:id [get]
func (c *IngresoSinSituacionFondosEstadoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetIngresoSinSituacionFondosEstadoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get IngresoSinSituacionFondosEstado
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.IngresoSinSituacionFondosEstado
// @Failure 403
// @router / [get]
func (c *IngresoSinSituacionFondosEstadoController) GetAll() {
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

	l, err := models.GetAllIngresoSinSituacionFondosEstado(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the IngresoSinSituacionFondosEstado
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.IngresoSinSituacionFondosEstado	true		"body for IngresoSinSituacionFondosEstado content"
// @Success 200 {object} models.IngresoSinSituacionFondosEstado
// @Failure 403 :id is not int
// @router /:id [put]
func (c *IngresoSinSituacionFondosEstadoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.IngresoSinSituacionFondosEstado{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateIngresoSinSituacionFondosEstadoById(&v); err == nil {
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
// @Description delete the IngresoSinSituacionFondosEstado
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *IngresoSinSituacionFondosEstadoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteIngresoSinSituacionFondosEstado(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// ChangeExistingStates ...
// @Title ChangeExistingStates
// @Description Change Existing States for IngresoSinSituacionFondosEstado
// @Param	body		body 	models.IngresoSinSituacionFondosEstado	true		"body for IngresoSinSituacionFondosEstado content"
// @Success 201 {int} models.IngresoSinSituacionFondosEstado
// @Failure 403 body is empty
// @router ChangeExistingStates/ [post]
func (c *IngresoSinSituacionFondosEstadoController) ChangeExistingStates() {

	var ingresoEstado models.IngresoSinSituacionFondosEstado
	var code string
	beego.Error("going on here....")
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ingresoEstado); err == nil {
			if err = models.ChangeExistingStates(ingresoEstado.IngresoSinSituacionFondos.Id);err == nil{
				if estadoIngreso, err := models.AddIngresoSinSituacionFondosEstado(&ingresoEstado); err == nil {
					c.Ctx.Output.SetStatus(201)
					alert := models.Alert{Type: "success", Code: "S_543", Body: estadoIngreso}
					c.Data["json"] = alert
				}else{
					alertdb := structs.Map(err)
					formatdata.FillStruct(alertdb["Code"], &code)
					alert := models.Alert{Type:"error",Code:"E_" + code,Body:err}
					beego.Error(err.Error())
					c.Data["json"] = alert
				}
			}else{
				alertdb := structs.Map(err)
				formatdata.FillStruct(alertdb["Code"], &code)
				alert := models.Alert{Type:"error",Code:"E_" + code,Body:err}
				beego.Error(err.Error())
				c.Data["json"] = alert
			}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = alert
		beego.Error(err.Error())
	}
	c.ServeJSON()
}
