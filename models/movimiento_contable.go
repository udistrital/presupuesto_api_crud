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

type MovimientoContable struct {
	Id                       int                       `orm:"column(id);pk;auto"`
	Debito                   int64                     `orm:"column(debito)"`
	Credito                  int64                     `orm:"column(credito)"`
	Fecha                    time.Time                 `orm:"column(fecha);type(timestamp without time zone)"`
	Concepto                 *Concepto                 `orm:"column(concepto_tesoral);rel(fk)"`
	CuentaContable           *CuentaContable           `orm:"column(cuenta_contable);rel(fk)"`
	TipoDocumentoAfectante   *TipoDocumentoAfectante   `orm:"column(tipo_documento_afectante);rel(fk)"`
	CodigoDocumentoAfectante int                       `orm:"column(codigo_documento_afectante)"`
	EstadoMovimientoContable *EstadoMovimientoContable `orm:"column(estado_movimiento_contable);rel(fk);null"`
	CuentaEspecial           *CuentaEspecial           `orm:"column(cuenta_especial);rel(fk);null"`
}

func (t *MovimientoContable) TableName() string {
	return "movimiento_contable"
}

func init() {
	orm.RegisterModel(new(MovimientoContable))
}

// AddMovimientoContable insert a new MovimientoContable into database and returns
// last inserted Id on success.
func AddMovimientoContable(m *MovimientoContable) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//AddMovimientoContableArray insert an array from MovimientoContable into database
func AddMovimientoContableArray(m *[]MovimientoContable) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.InsertMulti(100, m)
	return
}

// GetMovimientoContableById retrieves MovimientoContable by Id. Returns error if
// Id doesn't exist
func GetMovimientoContableById(id int) (v *MovimientoContable, err error) {
	o := orm.NewOrm()
	v = &MovimientoContable{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMovimientoContable retrieves all MovimientoContable matches certain condition. Returns empty list if
// no records exist
func GetAllMovimientoContable(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MovimientoContable)).RelatedSel()
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

	var l []MovimientoContable
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

// UpdateMovimientoContable updates MovimientoContable by Id and returns error if
// the record to be updated doesn't exist
func UpdateMovimientoContableById(m *MovimientoContable) (err error) {
	o := orm.NewOrm()
	v := MovimientoContable{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMovimientoContable deletes MovimientoContable by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMovimientoContable(id int) (err error) {
	o := orm.NewOrm()
	v := MovimientoContable{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MovimientoContable{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetSumMovimientos(tipoDocumento int, codigoDocumento int) (totalesMov []orm.Params, alerta Alert) {
	o := orm.NewOrm()
	o.Begin()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("sum(debito) as debito, sum(credito) as credito").
		From("financiera.movimiento_contable").
		Where("tipo_documento_afectante = ?").
		And("codigo_documento_afectante = ?")
	_, err := o.Raw(qb.String(), tipoDocumento, codigoDocumento).Values(&totalesMov)
	beego.Info(totalesMov)
	beego.Info(qb.String())
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_GetSumMovimientos_01"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	o.Commit()
	return
}
