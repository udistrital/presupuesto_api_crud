package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TipoTransaccionVersion struct {
	Id              int                     `orm:"column(id);pk;auto"`
	TipoTransaccion int                     `orm:"column(tipo_transaccion)"`
	Version         *VersionTipoTransaccion `orm:"column(version);rel(fk)"`
}

func (t *TipoTransaccionVersion) TableName() string {
	return "tipo_transaccion_version"
}

func init() {
	orm.RegisterModel(new(TipoTransaccionVersion))
}

// AddTipoTransaccionVersion insert a new TipoTransaccionVersion into database and returns
// last inserted Id on success.
func AddTipoTransaccionVersion(m *TipoTransaccionVersion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTipoTransaccionVersionById retrieves TipoTransaccionVersion by Id. Returns error if
// Id doesn't exist
func GetTipoTransaccionVersionById(id int) (v *TipoTransaccionVersion, err error) {
	o := orm.NewOrm()
	v = &TipoTransaccionVersion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTipoTransaccionVersion retrieves all TipoTransaccionVersion matches certain condition. Returns empty list if
// no records exist
func GetAllTipoTransaccionVersion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TipoTransaccionVersion))
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

	var l []TipoTransaccionVersion
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "Version")
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

// UpdateTipoTransaccionVersion updates TipoTransaccionVersion by Id and returns error if
// the record to be updated doesn't exist
func UpdateTipoTransaccionVersionById(m *TipoTransaccionVersion) (err error) {
	o := orm.NewOrm()
	v := TipoTransaccionVersion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTipoTransaccionVersion deletes TipoTransaccionVersion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTipoTransaccionVersion(id int) (err error) {
	o := orm.NewOrm()
	v := TipoTransaccionVersion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TipoTransaccionVersion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// AddTipoTransaccionVersion insert a new TipoTransaccionVersion sleecting las value for Type
//into database and returns TipoTransaccionVersion on success.
func AddNewTipoTransaccionVersion(m *VersionTipoTransaccion) (TipoTransV TipoTransaccionVersion, err error) {
	var consec float64
	var id int64
	o := orm.NewOrm()

	o.Begin()
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("COALESCE(MAX(ttv.tipo_transaccion),0)+1").
		From("tipo_transaccion_version ttv")

	sql := qb.String()

	err = o.Raw(sql).QueryRow(&consec)

	if err != nil {
		beego.Error(err)
		o.Rollback()
		return
	}

	TipoTransV.TipoTransaccion = int(consec)
	TipoTransV.Version = m

	id, err = o.Insert(&TipoTransV)

	if err != nil {
		beego.Error(err)
		o.Rollback()
		return
	}
	TipoTransV.Id = int(id)
	o.Commit()
	return
}

// GetRecordsNumberTipoTransaccionVersion retrieves quantity of records in  tipo transaccion version table
// Id doesn't exist returns 0
func GetRecordsNumberTipoTransaccionVersion(query map[string]string) (cnt int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TipoTransaccionVersion))
	cnt = 0
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
	cnt, err = qs.Count()
	return
}
