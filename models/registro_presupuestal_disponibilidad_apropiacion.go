package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type RegistroPresupuestalDisponibilidadApropiacion struct {
	Id                        int                        `orm:"auto;column(id);pk"`
	RegistroPresupuestal      *RegistroPresupuestal      `orm:"column(registro_presupuestal);rel(fk)"`
	DisponibilidadApropiacion *DisponibilidadApropiacion `orm:"column(disponibilidad_apropiacion);rel(fk)"`
	Valor                     float64                    `orm:"column(valor);null"`
}

func (t *RegistroPresupuestalDisponibilidadApropiacion) TableName() string {
	return "registro_presupuestal_disponibilidad_apropiacion"
}

func init() {
	orm.RegisterModel(new(RegistroPresupuestalDisponibilidadApropiacion))
}

// AddRegistroPresupuestalDisponibilidadApropiacion insert a new RegistroPresupuestalDisponibilidadApropiacion into database and returns
// last inserted Id on success.
func AddRegistroPresupuestalDisponibilidadApropiacion(m *RegistroPresupuestalDisponibilidadApropiacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRegistroPresupuestalDisponibilidadApropiacionById retrieves RegistroPresupuestalDisponibilidadApropiacion by Id. Returns error if
// Id doesn't exist
func GetRegistroPresupuestalDisponibilidadApropiacionById(id int) (v *RegistroPresupuestalDisponibilidadApropiacion, err error) {
	o := orm.NewOrm()
	v = &RegistroPresupuestalDisponibilidadApropiacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRegistroPresupuestalDisponibilidadApropiacion retrieves all RegistroPresupuestalDisponibilidadApropiacion matches certain condition. Returns empty list if
// no records exist
func GetAllRegistroPresupuestalDisponibilidadApropiacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RegistroPresupuestalDisponibilidadApropiacion))
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

	var l []RegistroPresupuestalDisponibilidadApropiacion
	qs = qs.OrderBy(sortFields...).RelatedSel(5)
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

// UpdateRegistroPresupuestalDisponibilidadApropiacion updates RegistroPresupuestalDisponibilidadApropiacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateRegistroPresupuestalDisponibilidadApropiacionById(m *RegistroPresupuestalDisponibilidadApropiacion) (err error) {
	o := orm.NewOrm()
	v := RegistroPresupuestalDisponibilidadApropiacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRegistroPresupuestalDisponibilidadApropiacion deletes RegistroPresupuestalDisponibilidadApropiacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRegistroPresupuestalDisponibilidadApropiacion(id int) (err error) {
	o := orm.NewOrm()
	v := RegistroPresupuestalDisponibilidadApropiacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RegistroPresupuestalDisponibilidadApropiacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
