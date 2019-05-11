package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

type Rubro struct {
	Id              int              `orm:"auto;column(id);pk"`
	Entidad         int              `orm:"column(entidad)"`
	Codigo          string           `orm:"column(codigo)"`
	Descripcion     string           `orm:"column(descripcion);null"`
	UnidadEjecutora int16            `orm:"column(unidad_ejecutora)"`
	Nombre          string           `orm:"column(nombre);null"`
	ProductoRubro   []*ProductoRubro `orm:"reverse(many)"`
}

func (t *Rubro) TableName() string {
	return "rubro"
}

func init() {
	orm.RegisterModel(new(Rubro))
}

// AddRubro insert a new Rubro into database and returns
// last inserted Id on success.
func AddRubro(m *Rubro) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRubroById retrieves Rubro by Id. Returns error if
// Id doesn't exist
func GetRubroById(id int) (v *Rubro, err error) {
	o := orm.NewOrm()
	v = &Rubro{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRubro retrieves all Rubro matches certain condition. Returns empty list if
// no records exist
func GetAllRubro(query map[string]string, group []string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Rubro))
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
	var l []Rubro

	qs = qs.OrderBy(sortFields...).RelatedSel(5)

	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				//o.LoadRelated(&v, "ProductoRubro", 5, 0, 0, "-Activo")
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

// UpdateRubro updates Rubro by Id and returns error if
// the record to be updated doesn't exist
func UpdateRubroById(m *Rubro) (err error) {
	o := orm.NewOrm()
	v := Rubro{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRubro deletes Rubro by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRubro(id int) (err error) {
	o := orm.NewOrm()
	v := Rubro{Id: id}
	var apropiaciones []int
	var rubrorubro []int
	// ascertain id exists in the database
	o.Begin()
	if err = o.Read(&v); err == nil {
		var num int64
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("id").
			From("" + beego.AppConfig.String("PGschemas") + ".apropiacion").
			Where("rubro=?")
		if _, err = o.Raw(qb.String(), id).QueryRows(&apropiaciones); err == nil {

			if len(apropiaciones) == 0 {
				qb, _ = orm.NewQueryBuilder("mysql")
				qb.Select("id").
					From("" + beego.AppConfig.String("PGschemas") + ".rubro_rubro").
					//Where("rubro_padre=?").
					Where("rubro_hijo=?")
				if _, err = o.Raw(qb.String(), id).QueryRows(&rubrorubro); err == nil {
					for _, idx := range rubrorubro {
						if _, err = o.Delete(&RubroRubro{Id: idx}); err == nil {

						} else {
							o.Rollback()
						}
					}
				}
			} else {
				o.Rollback()
			}

		} else {
			fmt.Println("Error 1 ", err)
			o.Rollback()
		}
		if num, err = o.Delete(&Rubro{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
			o.Commit()
		} else {
			fmt.Println("Error 2 ", err)

			o.Rollback()
		}
	}
	return
}
