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

type Producto struct {
	Id            int              `orm:"column(id);pk;auto"`
	Nombre        string           `orm:"column(nombre)"`
	Descripcion   string           `orm:"column(descripcion);null"`
	FechaRegistro time.Time        `orm:"column(fecha_registro);type(date)"`
	ProductoRubro []*ProductoRubro `orm:"reverse(many)"`
}

func (t *Producto) TableName() string {
	return "producto"
}

func init() {
	orm.RegisterModel(new(Producto))
}

// GetTotalProductos get number of Producto into database and returns
// integer number.
func GetTotalProductos() (total int, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COUNT(id)").
		From("" + beego.AppConfig.String("PGschemas") + ".producto")
	err = o.Raw(qb.String()).QueryRow(&total)
	return
}

// AddProducto insert a new Producto into database and returns
// last inserted Id on success.
func AddProducto(m *Producto) (id int64, err error) {
	o := orm.NewOrm()
	m.FechaRegistro = time.Now().Local()
	id, err = o.Insert(m)
	return
}

// GetProductoById retrieves Producto by Id. Returns error if
// Id doesn't exist
func GetProductoById(id int) (v *Producto, err error) {
	o := orm.NewOrm()
	v = &Producto{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProducto retrieves all Producto matches certain condition. Returns empty list if
// no records exist
func GetAllProducto(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Producto))
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

	var l []Producto
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "ProductoRubro", 5, 1, 0, "-Id")
				/*for _, sub := range v.RegistroPresupuestalDisponibilidadApropiacion {
					o.LoadRelated(sub.DisponibilidadApropiacion.Disponibilidad, "DisponibilidadProcesoExterno", 5, 1, 0, "-Id")
				}*/
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

// UpdateProducto updates Producto by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductoById(m *Producto) (err error) {
	o := orm.NewOrm()
	v := Producto{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProducto deletes Producto by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProducto(id int) (err error) {
	o := orm.NewOrm()
	v := Producto{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Producto{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
