package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/presupuesto_crud/models"
)

// ApropiacionController operations for Apropiacion
type ApropiacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ApropiacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Apropiacion
// @Param	body		body 	models.Apropiacion	true		"body for Apropiacion content"
// @Success 201 {int} models.Apropiacion
// @Failure 403 body is empty
// @router / [post]
func (c *ApropiacionController) Post() {
	var v models.Apropiacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddApropiacion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Apropiacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Apropiacion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ApropiacionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetApropiacionById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Apropiacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Apropiacion
// @Failure 403
// @router / [get]
func (c *ApropiacionController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var exclude = make(map[string]string)
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

	// exclude: k:v,k:v
	if v := c.GetString("exclude"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid exclude key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			exclude[k] = v
		}
	}

	l, err := models.GetAllApropiacion(query, exclude, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
}

// Put ...
// @Title Put
// @Description update the Apropiacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Apropiacion	true		"body for Apropiacion content"
// @Success 200 {object} models.Apropiacion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ApropiacionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Apropiacion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateApropiacionById(&v); err == nil {
			c.Data["json"] = v
		} else {
			c.Data["json"] = err
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Apropiacion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ApropiacionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteApropiacion(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// AprobarPresupuesto ...
// @Title AprobarPresupuesto
// @Description aprueba la asignacion inicial de presupuesto
// @Param	Vigencia		query 	string	true		"vigencia a comprobar"
// @Param	UnidadEjecutora		query 	string	true		"unidad ejecutora de los rubros a comprobar"
// @Success 200 {string} resultado
// @Failure 403
// @router /AprobacionAsignacionInicial/ [put]
func (c *ApropiacionController) AprobarPresupuesto() {
	var (
		vigencia int
		ue       int
		err      error
	)

	defer func() {
		if r := recover(); r != nil {
			c.Data["json"] = r
		}
		c.ServeJSON()
	}()

	if vigencia, err = c.GetInt("Vigencia"); err != nil {
		panic(err.Error())
	}
	if ue, err = c.GetInt("UnidadEjecutora"); err != nil {
		panic(err.Error())
	}
	beego.Debug(vigencia, ue)
}
