package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type RubroHomologadoRubro struct {
	Id              int              `orm:"column(id);pk;auto"`
	RubroHomologado *RubroHomologado `orm:"column(rubro_homologado);rel(fk)"`
	Rubro           *Rubro           `orm:"column(rubro);rel(fk)"`
}

func (t *RubroHomologadoRubro) TableName() string {
	return "rubro_homologado_rubro"
}

func init() {
	orm.RegisterModel(new(RubroHomologadoRubro))
}

// AddRubroHomologadoRubro insert a new RubroHomologadoRubro into database and returns
// last inserted Id on success.
func AddRubroHomologadoRubro(m *RubroHomologadoRubro) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRubroHomologadoRubroById retrieves RubroHomologadoRubro by Id. Returns error if
// Id doesn't exist
func GetRubroHomologadoRubroById(id int) (v *RubroHomologadoRubro, err error) {
	o := orm.NewOrm()
	v = &RubroHomologadoRubro{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


// GetRubroHomologadoRubroById retrieves RubroHomologadoRubro by Id. Returns error if
// Id doesn't exist
func GetRecordsNumberRubroHomologadoById(id int) (cnt int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RubroHomologadoRubro))
	qs = qs.Filter("rubro_homologado", id)
	cnt, err = qs.Count()
	return
}

// GetRecordsNumberRubroHomologadoRubroById retrieves number of recor for a item
//in RubroHomologadoRubro. Returns error if
// Id doesn't exist
func GetRecordsNumberRubroHomologadoRubroById(id int) (cnt int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RubroHomologadoRubro))
	qs = qs.Filter("rubro", id)
	cnt, err = qs.Count()
	return
}
// GetAllRubroHomologadoRubro retrieves all RubroHomologadoRubro matches certain condition. Returns empty list if
// no records exist
func GetAllRubroHomologadoRubro(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RubroHomologadoRubro)).RelatedSel()
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

	var l []RubroHomologadoRubro
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

// UpdateRubroHomologadoRubro updates RubroHomologadoRubro by Id and returns error if
// the record to be updated doesn't exist
func UpdateRubroHomologadoRubroById(m *RubroHomologadoRubro) (err error) {
	o := orm.NewOrm()
	v := RubroHomologadoRubro{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRubroHomologadoRubro deletes RubroHomologadoRubro by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRubroHomologadoRubro(id int) (err error) {
	o := orm.NewOrm()
	v := RubroHomologadoRubro{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RubroHomologadoRubro{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}