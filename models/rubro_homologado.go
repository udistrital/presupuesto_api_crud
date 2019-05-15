package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type RubroHomologado struct {
	Id                int      `orm:"column(id);pk;auto"`
	CodigoHomologado  string   `orm:"column(codigo_homologado)"`
	NombreHomologado  string   `orm:"column(nombre_homologado)"`
	Organizacion 			int			 `orm:"column(organizacion)"`
	Vigencia          float64  `orm:"column(vigencia)"`
}

type RubroPadreHomol struct {
	RubroPadre string
	CntHomologacion int64
}

func (t *RubroHomologado) TableName() string {
	return "rubro_homologado"
}

func init() {
	orm.RegisterModel(new(RubroHomologado))
}

// AddRubroHomologado insert a new RubroHomologado into database and returns
// last inserted Id on success.
func AddRubroHomologado(m *RubroHomologado) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRubroHomologadoById retrieves RubroHomologado by Id. Returns error if
// Id doesn't exist
func GetRubroHomologadoById(id int) (v *RubroHomologado, err error) {
	o := orm.NewOrm()
	v = &RubroHomologado{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRubroHomologado retrieves all RubroHomologado matches certain condition. Returns empty list if
// no records exist
func GetAllRubroHomologado(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RubroHomologado))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []RubroHomologado
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateRubroHomologado updates RubroHomologado by Id and returns error if
// the record to be updated doesn't exist
func UpdateRubroHomologadoById(m *RubroHomologado) (err error) {
	o := orm.NewOrm()
	v := RubroHomologado{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRubroHomologado deletes RubroHomologado by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRubroHomologado(id int) (err error) {
	o := orm.NewOrm()
	v := RubroHomologado{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RubroHomologado{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


// GetRubroHomologadoRubroById retrieves RubroHomologadoRubro by Id. Returns error if
// Id doesn't exist
func GetRecordsNumberRubroByEntity(idEntidad int) (cnt int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RubroHomologado))
	qs = qs.Filter("organizacion", idEntidad)
	cnt, err = qs.Count()
	return
}
//validate if item parent's has got homologation
//Return true if has it
func GetParentHomologation (idRubro interface{})(res interface{},err error){
 beego.Error("GetParentHomologation idRubro ",idRubro)
//var qb QueryBuilder
var respRubroPHomol RubroPadreHomol

qb, err := orm.NewQueryBuilder("tidb")

if (err != nil ){
	return
}

if (idRubro != nil) {
		qb.Select("rr.rubro_padre",
						"count(rh.id)").
			From("financiera.rubro_rubro rr").
			LeftJoin("financiera.rubro r").On("rr.rubro_padre = r.id").
			LeftJoin("financiera.rubro_homologado rh").On("rh.codigo_homologado = r.codigo").
			Where("rr.rubro_hijo = ?").
			GroupBy("rr.rubro_padre")
}


		sql := qb.String()
		beego.Error("query",sql)
		o := orm.NewOrm()
		err = o.Raw(sql, idRubro.(string)).QueryRow(&respRubroPHomol)
		beego.Error("rspuesta query ",respRubroPHomol,"error ",err)
		if err == nil {

			}else if err == orm.ErrNoRows{
				qb, err = orm.NewQueryBuilder("tidb")

				qb.Select("null as rubro_padre",
								"count(1) ").
					From("financiera.rubro_homologado_rubro rh").
					Where("rh.rubro = ?")

					sql := qb.String()
					beego.Error("query",sql)

				 err = o.Raw(sql, idRubro.(string)).QueryRow(&respRubroPHomol)
				 if err != nil {
					 beego.Error(err)
					 return
				 }
			}
		if (respRubroPHomol.CntHomologacion == 0 ){
			res = false
			return
		}
		if (respRubroPHomol.RubroPadre != ""){
			 beego.Error(" call function again ")
			 GetParentHomologation(respRubroPHomol.RubroPadre)
		}
		return
	}