package controllers

import (
	"github.com/udistrital/presupuesto_crud/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/astaxie/beego"
)

// RubroHomologadoController operations for RubroHomologado
type RubroHomologadoController struct {
	beego.Controller
}

// URLMapping ...
func (c *RubroHomologadoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create RubroHomologado
// @Param	body		body 	models.RubroHomologado	true		"body for RubroHomologado content"
// @Success 201 {int} models.RubroHomologado
// @Failure 403 body is empty
// @router / [post]
func (c *RubroHomologadoController) Post() {
	var v models.RubroHomologado
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddRubroHomologado(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = models.Alert{Type:"success", Code: "S_543",Body:v}
		} else {
			alertdb:=structs.Map(err)
			c.Data["json"] = models.Alert{Type:"error",Code:"E_"+alertdb["Code"].(string),Body:err.Error()}
		}
	} else {
		beego.Error("error ",err.Error())
		c.Data["json"] = models.Alert{Type:"error",Code:"E_0458",Body:err.Error()}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get RubroHomologado by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.RubroHomologado
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RubroHomologadoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRubroHomologadoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get RubroHomologado
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.RubroHomologado
// @Failure 403
// @router / [get]
func (c *RubroHomologadoController) GetAll() {
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

	l, err := models.GetAllRubroHomologado(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the RubroHomologado
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.RubroHomologado	true		"body for RubroHomologado content"
// @Success 200 {object} models.RubroHomologado
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RubroHomologadoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.RubroHomologado{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateRubroHomologadoById(&v); err == nil {
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
// @Description delete the RubroHomologado
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RubroHomologadoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteRubroHomologado(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetRecordsNumber...
// @Title Get Records Number RubroHomologado By Id
// @Description get Number of records for homologate item
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.RubroHomologado
// @Failure 403 :id is empty
// @router /GetRecordsNumberRubroHomologadoById/:id [get]
func (c *RubroHomologadoController) GetRecordsNumberRubroHomologadoById() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRecordsNumberRubroHomologadoById(id)
	if err != nil {
		alertdb:=structs.Map(err)
		c.Data["json"] = models.Alert{Type:"error",Code:"E_"+alertdb["Code"].(string),Body:err.Error()}
	} else {
		c.Data["json"] = models.Alert{Type:"success", Code: "S_543",Body:v}
	}
	c.ServeJSON()
}

// GetRecordsNumberRubro...
// @Title Get Records Number RubroHomologado By Id
// @Description get Number of records for item when its homologate
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} interface
// @Failure 403 :id is empty
// @router /GetRecordsNumberRubroHomologadoRubroById/:id [get]
func (c *RubroHomologadoController) GetRecordsNumberRubroHomologadoRubroById() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRecordsNumberRubroHomologadoRubroById(id)
	if err != nil {
		alertdb:=structs.Map(err)
		c.Data["json"] = models.Alert{Type:"error",Code:"E_"+alertdb["Code"].(string),Body:err.Error()}
	} else {
		c.Data["json"] = models.Alert{Type:"success", Code: "S_543",Body:v}
	}
	c.ServeJSON()
}


// GetRecordsNumberByEntity...
// @Title Get Records Number RubroHomologado By Entity
// @Description get Number of records for a entity
// @Param	idEntidad	path 	string	true		"The key for staticblock"
// @Success 200 {object} interface
// @Failure 403 :id is empty
// @router /GetRecordsNumberByEntity [get]
func (c *RubroHomologadoController) GetRecordsNumberByEntity() {
	if idEntidad, err := c.GetInt("idEntidad");err==nil{
		v, err := models.GetRecordsNumberRubroByEntity(idEntidad)
			if err != nil {
				alertdb:=structs.Map(err)
				c.Data["json"] = models.Alert{Type:"error",Code:"E_"+alertdb["Code"].(string),Body:err.Error()}
				} else {
					c.Data["json"] = models.Alert{Type:"success", Code: "S_543",Body:v}
				}
		}else{
			c.Data["json"] = models.Alert{Code: "E_0458", Body: "Not enough parameter", Type: "error"}
		}
	c.ServeJSON()
}



// GetParentHomologation...
// @Title Get Parent Item Homologation
// @Description Identifies if exist homologation for item's parent
// @Param	idRubro	path 	string	true		"The key for staticblock"
// @Success 200 {object} interface
// @Failure 403 Rubro is empty
// @router GetParentHomologation/ [get]
func (c *RubroHomologadoController) GetParentHomologation() {
	
 idRubro := c.GetString("idRubro")
		beego.Error(" GetParentHomologation id rubro",idRubro)
		v, err := models.GetParentHomologation(idRubro)
			if err != nil {
				alertdb:=structs.Map(err)
				c.Data["json"] = models.Alert{Type:"error",Code:"E_"+alertdb["Code"].(string),Body:err.Error()}
				} else {
					c.Data["json"] = models.Alert{Type:"success", Code: "S_543",Body:v}
				}

	c.ServeJSON()
}
