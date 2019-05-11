package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

type ProductoRubro struct {
	Id                int       `orm:"column(id);pk;auto"`
	Rubro             *Rubro    `orm:"column(rubro);rel(fk)"`
	Producto          *Producto `orm:"column(producto);rel(fk)"`
	ValorDistribucion float64   `orm:"column(valor_distribucion)"`
	Activo            bool      `orm:"column(activo);default(true)"`
	FechaRegistro     time.Time `orm:"column(fecha_registro);auto_now_add"`
}

func (t *ProductoRubro) TableName() string {
	return "producto_rubro"
}

func init() {
	orm.RegisterModel(new(ProductoRubro))
}

// AddProductoRubro insert a new ProductoRubro into database and returns
// last inserted Id on success.
func AddProductoRubro(m *ProductoRubro) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProductoRubroById retrieves ProductoRubro by Id. Returns error if
// Id doesn't exist
func GetProductoRubroById(id int) (v *ProductoRubro, err error) {
	o := orm.NewOrm()
	v = &ProductoRubro{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProductoRubro retrieves all ProductoRubro matches certain condition. Returns empty list if
// no records exist
func GetAllProductoRubro(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductoRubro))
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

	var l []ProductoRubro
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

// UpdateProductoRubro updates ProductoRubro by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductoRubroById(m *ProductoRubro) (err error) {
	o := orm.NewOrm()
	v := ProductoRubro{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Activo"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductoRubro deletes ProductoRubro by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProductoRubro(id int) (err error) {
	o := orm.NewOrm()
	v := ProductoRubro{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductoRubro{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// SetVariacionProducto set a new record in the database
// with a new variation for the rubro
func SetVariacionProducto(m *ProductoRubro) (total float64, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	o.Begin()
	qb.Select("SUM(valor_distribucion) as total").
		From("" + beego.AppConfig.String("PGschemas") + ".producto_rubro").
		Where("rubro = ?").
		And("activo = true").
		And("id NOT IN (?)")
	err = o.Raw(qb.String(), m.Rubro.Id, m.Id).QueryRow(&total)

	if err != nil {
		o.Rollback()
		return
	}

	total = total + m.ValorDistribucion

	if total > 1 {
		o.Rollback()
		return
	}

	m.Activo = false
	_, err = o.Update(m, "Activo")
	if err != nil {
		o.Rollback()
		return
	}
	m.Id = 0
	m.Activo = true
	m.FechaRegistro = time.Now().Local()
	_, err = o.Insert(m)
	if err != nil {
		o.Rollback()
		return
	}

	o.Commit()

	return
}

//AddProductoRubrotr Add ProductoRubro Realiton to Rubro
func AddProductoRubrotr(m *ProductoRubro) (total float64, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	o.Begin()
	qb.Select("SUM(valor_distribucion) as total").
		From("" + beego.AppConfig.String("PGschemas") + ".producto_rubro").
		Where("rubro = ?").
		And("activo = true")
	err = o.Raw(qb.String(), m.Rubro.Id).QueryRow(&total)
	if err != nil {
		o.Rollback()
		return
	}
	m.ValorDistribucion = m.ValorDistribucion / 100
	m.FechaRegistro = time.Now().Local()
	total = total + (m.ValorDistribucion)
	if total > 1 {
		o.Rollback()
		return
	}
	if m.ValorDistribucion <= 0 {
		o.Rollback()
		total = 2
		return
	}
	m.Activo = true
	_, err = o.Insert(m)
	if err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return

}
