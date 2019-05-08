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

type VersionTipoTransaccion struct {
	Id                int       `orm:"column(id);pk;auto"`
	Definitiva        bool      `orm:"column(definitiva)"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(date);auto_now_add;type(datetume)"`
	FechaInicio       time.Time `orm:"column(fecha_inicio);type(date)"`
	FechaFin          time.Time `orm:"column(fecha_fin);type(date)"`
	NumeroVersion     int       `orm:"column(numero_version)"`
	Descripcion       string    `orm:"column(descripcion)"`
}

func (t *VersionTipoTransaccion) TableName() string {
	return "version_tipo_transaccion"
}

func init() {
	orm.RegisterModel(new(VersionTipoTransaccion))
}

// AddVersionTipoTransaccion insert a new VersionTipoTransaccion into database and returns
// last inserted Id on success.
func AddVersionTipoTransaccion(m *VersionTipoTransaccion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetVersionTipoTransaccionById retrieves VersionTipoTransaccion by Id. Returns error if
// Id doesn't exist
func GetVersionTipoTransaccionById(id int) (v *VersionTipoTransaccion, err error) {
	o := orm.NewOrm()
	v = &VersionTipoTransaccion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllVersionTipoTransaccion retrieves all VersionTipoTransaccion matches certain condition. Returns empty list if
// no records exist
func GetAllVersionTipoTransaccion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(VersionTipoTransaccion))
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

	var l []VersionTipoTransaccion
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

// UpdateVersionTipoTransaccion updates VersionTipoTransaccion by Id and returns error if
// the record to be updated doesn't exist
func UpdateVersionTipoTransaccionById(m *VersionTipoTransaccion) (err error) {
	o := orm.NewOrm()
	v := VersionTipoTransaccion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteVersionTipoTransaccion deletes VersionTipoTransaccion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteVersionTipoTransaccion(id int) (err error) {
	o := orm.NewOrm()
	v := VersionTipoTransaccion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&VersionTipoTransaccion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetRecordsNumberVersionTipo retrieves quantity of records in version tipo transaccion table
// Id doesn't exist returns 0
func GetRecordsNumberVersionTipo(query map[string]string) (cnt int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(VersionTipoTransaccion))

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

// GetVersionToTipoTransaccion retrieves max numeroVersion given Id. tipo Returns error if
// Id doesn't exist
func GetVersionToTipoTransaccion(id int) (consec float64, err error) {
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("COALESCE(MAX(vtt.numero_version)) + 1").
		From("tipo_transaccion_version ttv").
		InnerJoin("version_tipo_transaccion vtt").On("ttv.version = vtt.id").
		Where("ttv.id = ?")

	sql := qb.String()

	err = o.Raw(sql, id).QueryRow(&consec)

	return consec, err
}

//given a initial date and a final date, retrieves all varsions to a type
//in date range, Returns error it id doesn't exist
func GetVersionInEspecifiedDate(fechaInicio string, fechaFin string, IdTipoTr int) (TipoTranVersiones []TipoTransaccionVersion, err error) {
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("vtt.*").
		From("tipo_transaccion_version ttv ").
		InnerJoin("version_tipo_transaccion vtt").On("ttv.version = vtt.id").
		Where("((vtt.fecha_inicio between ? and ?) or (vtt.fecha_fin between ? and ?))").
		And("ttv.tipo_transaccion = ?")

	sql := qb.String()

	_, err = o.Raw(sql, fechaInicio, fechaFin, fechaInicio, fechaFin, IdTipoTr).QueryRows(&TipoTranVersiones)

	return
}

//given a initial date and a final date, retrieves all varsions to a type
//in date range, Returns error it id doesn't exist
func GetAllDefinitiveVersion(fecha string) (TipoTranVersiones []TipoTransaccionVersion, err error) {
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("ttv.*").
		From("tipo_transaccion_version ttv ").
		InnerJoin("version_tipo_transaccion vtt").On("ttv.version = vtt.id").
		Where("? between vtt.fecha_inicio and vtt.fecha_fin").
		And("vtt.definitiva=true")

	sql := qb.String()

	beego.Error(sql)

	_, err = o.Raw(sql, fecha).QueryRows(&TipoTranVersiones)

	return
}
