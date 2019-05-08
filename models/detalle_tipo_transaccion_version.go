package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type DetalleTipoTransaccionVersion struct {
	Id                     int                     `orm:"column(id);pk;auto"`
	Nombre                 string                  `orm:"column(nombre)"`
	Descripcion            string                  `orm:"column(descripcion);null"`
	ClaseTransaccion       *TipoConcepto           `orm:"column(clase_transaccion);rel(fk)"`
	CodigoAbreviacion      string                  `orm:"column(codigo_abreviacion);null"`
	NumeroOrden            float64                 `orm:"column(numero_orden)"`
	TipoTransaccionVersion *TipoTransaccionVersion `orm:"column(tipo_transaccion_version);rel(fk)"`
}

func (t *DetalleTipoTransaccionVersion) TableName() string {
	return "detalle_tipo_transaccion_version"
}

func init() {
	orm.RegisterModel(new(DetalleTipoTransaccionVersion))
}

// AddDetalleTipoTransaccionVersion insert a new DetalleTipoTransaccionVersion into database and returns
// last inserted Id on success.
func AddDetalleTipoTransaccionVersion(m *DetalleTipoTransaccionVersion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDetalleTipoTransaccionVersionById retrieves DetalleTipoTransaccionVersion by Id. Returns error if
// Id doesn't exist
func GetDetalleTipoTransaccionVersionById(id int) (v *DetalleTipoTransaccionVersion, err error) {
	o := orm.NewOrm()
	v = &DetalleTipoTransaccionVersion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDetalleTipoTransaccionVersion retrieves all DetalleTipoTransaccionVersion matches certain condition. Returns empty list if
// no records exist
func GetAllDetalleTipoTransaccionVersion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DetalleTipoTransaccionVersion))
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

	var l []DetalleTipoTransaccionVersion
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "ClaseTransaccion")
				o.LoadRelated(&v, "TipoTransaccionVersion")
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

// UpdateDetalleTipoTransaccionVersion updates DetalleTipoTransaccionVersion by Id and returns error if
// the record to be updated doesn't exist
func UpdateDetalleTipoTransaccionVersionById(m *DetalleTipoTransaccionVersion) (err error) {
	o := orm.NewOrm()
	v := DetalleTipoTransaccionVersion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDetalleTipoTransaccionVersion deletes DetalleTipoTransaccionVersion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDetalleTipoTransaccionVersion(id int) (err error) {
	o := orm.NewOrm()
	v := DetalleTipoTransaccionVersion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DetalleTipoTransaccionVersion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//Rollback function for transaction type
func DeleteEntireTipoTransaccion(m map[string]interface{}) (err error) {
	for key, value := range m {
		body := value.(map[string]interface{})["Body"].(map[string]interface{})
		if body["Id"] != nil {
			id := int(body["Id"].(float64))
			switch key {
			case "version_tipo_transaccion":
				err = DeleteVersionTipoTransaccion(id)
			case "tipo_transaccion_version":
				err = DeleteTipoTransaccionVersion(id)
			case "detalle_tipo_transaccion_version":
				err = DeleteDetalleTipoTransaccionVersion(id)
			}
		}
	}
	return
}

// UpdateDetalleTipoTransaccionVersion updates DetalleTipoTransaccionVersion by Id and returns error if
// the record to be updated doesn't exist
func UpdateTipoTransaccionVersion(m map[string]interface{}) (err error) {
	var version VersionTipoTransaccion
	var detalle DetalleTipoTransaccionVersion
	o := orm.NewOrm()
	o.Begin()
	err = formatdata.FillStruct(m["Version"], &version)
	err = formatdata.FillStruct(m["DetalleTipoTransaccion"], &detalle)
	versionRead := VersionTipoTransaccion{Id: version.Id}
	detalleRead := DetalleTipoTransaccionVersion{Id: detalle.Id}
	// ascertain id exists in the database
	if err = o.Read(&versionRead); err == nil {
		var num int64
		if num, err = o.Update(&version); err == nil {
			fmt.Println("Number of records updated in database:", num)
		} else {
			beego.Error("Error ", err)
			o.Rollback()
			return
		}
	}
	if err = o.Read(&detalleRead); err == nil {
		var num int64
		if num, err = o.Update(&detalle); err == nil {
			fmt.Println("Number of records updated in database:", num)
		} else {
			beego.Error("Error ", err)
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}
