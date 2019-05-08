package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CuentaContable struct {
	Id                 int                 `orm:"column(id);pk;auto"`
	Saldo              int64               `orm:"column(saldo)"`
	Nombre             string              `orm:"column(nombre)"`
	Naturaleza         string              `orm:"column(naturaleza)"`
	Descripcion        string              `orm:"column(descripcion);null"`
	Codigo             string              `orm:"column(codigo)"`
	NivelClasificacion *NivelClasificacion `orm:"column(nivel_clasificacion_cuenta_contable);rel(fk)"`
}

func (t *CuentaContable) TableName() string {
	return "cuenta_contable"
}

func init() {
	orm.RegisterModel(new(CuentaContable))
}

// AddCuentaContable insert a new CuentaContable into database and returns
// last inserted Id on success.
func AddCuentaContable(m *CuentaContable) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCuentaContableById retrieves CuentaContable by Id. Returns error if
// Id doesn't exist
func GetCuentaContableById(id int) (v *CuentaContable, err error) {
	o := orm.NewOrm()
	v = &CuentaContable{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCuentaContable retrieves all CuentaContable matches certain condition. Returns empty list if
// no records exist
func GetAllCuentaContable(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CuentaContable)).RelatedSel()
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

	var l []CuentaContable
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				today := time.Now()
				var year, month int
				if int(today.Month()) == 1 {
					year = today.Year() - 1
					month = 12
				} else {
					year = today.Year()
					month = int(today.Month()) - 1
				}
				firstdate := time.Date(today.Year(), today.Month(), 1, 23, 0, 0, 0, time.UTC)
				var rul string
				if v.Naturaleza == "debito" {
					rul = "sum(debito) - sum(credito)"
				} else {
					rul = "sum(credito) - sum(debito)"
				}
				o.Raw(`select sum(saldo) from (
							select saldo from financiera.saldo_cuenta_contable where cuenta_contable=? and anio = ? and mes = ?
							union
							select `+rul+` saldo from financiera.movimiento_contable
							 where cuenta_contable=? and fecha >= ?::DATE group by cuenta_contable ) a`, v.Id, year, month, v.Id, firstdate).QueryRow(&v.Saldo)
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

// UpdateCuentaContable updates CuentaContable by Id and returns error if
// the record to be updated doesn't exist
func UpdateCuentaContableById(m *CuentaContable) (err error) {
	o := orm.NewOrm()
	v := CuentaContable{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCuentaContable deletes CuentaContable by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCuentaContable(id int) (err error) {
	o := orm.NewOrm()
	v := CuentaContable{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CuentaContable{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
