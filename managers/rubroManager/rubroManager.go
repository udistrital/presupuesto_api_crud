package rubromanager

import (
	"github.com/astaxie/beego/orm"
	appmessagemanager "github.com/udistrital/presupuesto_crud/managers/appMessageManager"
	"github.com/udistrital/presupuesto_crud/models"
)

func RubroRelationRegistrator(idParent int, Rubro *models.Rubro) {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Insert(Rubro)
	if err != nil {
		o.Rollback()
		panic(appmessagemanager.InsertErrorMessage())
	}
	relation := models.RubroRubro{}
	relation.RubroHijo = Rubro
	relation.RubroPadre = &models.Rubro{}
	relation.RubroPadre.Id = idParent
	_, err = o.Insert(&relation)
	if err != nil {
		o.Rollback()
		panic(appmessagemanager.InsertErrorMessage())
	}
	o.Commit()

}
