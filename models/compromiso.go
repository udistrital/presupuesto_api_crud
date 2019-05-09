package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Compromiso struct {
	Id                    int                    `orm:"column(id);pk;auto"`
	Objeto                string                 `orm:"column(objeto)"`
	Vigencia              float64                `orm:"column(vigencia)"`
	FechaInicio           time.Time              `orm:"column(fecha_inicio);type(date)"`
	FechaFin              time.Time              `orm:"column(fecha_fin);type(date)"`
	FechaModificacion     time.Time              `orm:"column(fecha_modificacion);type(date)"`
	EstadoCompromiso      *EstadoCompromiso      `orm:"column(estado_compromiso);rel(fk)"`
	TipoCompromisoTesoral *TipoCompromisoTesoral `orm:"column(tipo_compromiso_financiero);rel(fk)"`
	UnidadEjecutora       int                    `orm:"column(unidad_ejecutora)"`
}

func (t *Compromiso) TableName() string {
	return "compromiso_financiero"
}

func init() {
	orm.RegisterModel(new(Compromiso))
}

// AddCompromiso insert a new Compromiso into database and returns
// last inserted Id on success.
func AddCompromiso(m *Compromiso) (id int64, err error) {
	o := orm.NewOrm()
	m.FechaModificacion = time.Now()
	id, err = o.Insert(m)
	return
}

// GetCompromisoById retrieves Compromiso by Id. Returns error if
// Id doesn't exist
func GetCompromisoById(id int) (v *Compromiso, err error) {
	o := orm.NewOrm()
	v = &Compromiso{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCompromiso retrieves all Compromiso matches certain condition. Returns empty list if
// no records exist
func GetAllCompromiso(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Compromiso)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.Contains(k, "__in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else if strings.Contains(k, "__not_in") {
			k = strings.Replace(k, "__not_in", "", -1)
			qs = qs.Exclude(k, v)
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

	var l []Compromiso
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

// UpdateCompromiso updates Compromiso by Id and returns error if
// the record to be updated doesn't exist
func UpdateCompromisoById(m *Compromiso) (err error) {
	o := orm.NewOrm()
	v := Compromiso{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCompromiso deletes Compromiso by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCompromiso(id int) (err error) {
	o := orm.NewOrm()
	v := Compromiso{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Compromiso{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
