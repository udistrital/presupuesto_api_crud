package controllers

import (
	"github.com/udistrital/presupuesto_crud/models"

	"github.com/astaxie/beego"
)

// ArbolConceptosController operations for ArbolConceptos
type ArbolConceptosController struct {
	beego.Controller
}

// URLMapping ...
func (c *ArbolConceptosController) URLMapping() {
	c.Mapping("MakeTree", c.MakeTree)
}

// MakeTree ...
// @Title MakeTree
// @Description get Arbol of Conceptos
// @Param	body		body 	models.ConceptoConcepto	true		"body for ConceptoConcepto content"
// @Success 201 {int} models.MakeTreeConcepto
// @Failure 403 body is empty
// @router / [get]
func (c *ArbolConceptosController) MakeTree() {

	l := models.MakeTreeConcepto()
	//fmt.Println(l)

	c.Data["json"] = l
	//Generera el Json con los datos obtenidos
	c.ServeJSON()

}
