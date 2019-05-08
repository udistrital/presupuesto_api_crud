package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/optimize"
)

type OrdenPagoEstadoOrdenPago struct {
	Id              int              `orm:"column(id);pk;auto"`
	OrdenPago       *OrdenPago       `orm:"column(orden_pago);rel(fk)"`
	EstadoOrdenPago *EstadoOrdenPago `orm:"column(estado_orden_pago);rel(fk)"`
	FechaRegistro   time.Time        `orm:"column(fecha_registro);type(date)"`
	Usuario         int              `orm:"column(usuario);null"`
}

func (t *OrdenPagoEstadoOrdenPago) TableName() string {
	return "orden_pago_estado_orden_pago"
}

func init() {
	orm.RegisterModel(new(OrdenPagoEstadoOrdenPago))
}

// AddOrdenPagoEstadoOrdenPago insert a new OrdenPagoEstadoOrdenPago into database and returns
// last inserted Id on success.
func AddOrdenPagoEstadoOrdenPago(m *OrdenPagoEstadoOrdenPago) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrdenPagoEstadoOrdenPagoById retrieves OrdenPagoEstadoOrdenPago by Id. Returns error if
// Id doesn't exist
func GetOrdenPagoEstadoOrdenPagoById(id int) (v *OrdenPagoEstadoOrdenPago, err error) {
	o := orm.NewOrm()
	v = &OrdenPagoEstadoOrdenPago{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrdenPagoEstadoOrdenPago retrieves all OrdenPagoEstadoOrdenPago matches certain condition. Returns empty list if
// no records exist
func GetAllOrdenPagoEstadoOrdenPago(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrdenPagoEstadoOrdenPago)).RelatedSel()
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

	var l []OrdenPagoEstadoOrdenPago
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

// UpdateOrdenPagoEstadoOrdenPago updates OrdenPagoEstadoOrdenPago by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrdenPagoEstadoOrdenPagoById(m *OrdenPagoEstadoOrdenPago) (err error) {
	o := orm.NewOrm()
	v := OrdenPagoEstadoOrdenPago{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrdenPagoEstadoOrdenPago deletes OrdenPagoEstadoOrdenPago by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrdenPagoEstadoOrdenPago(id int) (err error) {
	o := orm.NewOrm()
	v := OrdenPagoEstadoOrdenPago{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrdenPagoEstadoOrdenPago{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func WorkFlowOrdenPago(DataWorkFlowOrdenPago map[string]interface{}) (alerta []map[string]interface{}, err error) {
	ouputError := make(map[string]interface{})
	//get data
	ordenPagoData, e := DataWorkFlowOrdenPago["OrdenPago"].([]interface{})
	idEstadoNew, e := DataWorkFlowOrdenPago["NuevoEstado"].(interface{}).(map[string]interface{})["Id"].(float64)
	idUsuario, e := DataWorkFlowOrdenPago["Usuario"].(interface{}).(map[string]interface{})["Id"].(float64)
	if e != true {
		fmt.Println("error de parametros")
		ouputError["Type"] = "error"
		ouputError["Code"] = "E_TRANS_01"
		ouputError["Body"] = ""
		alerta = append(alerta, ouputError)
		return
	}
	var respuesta []map[string]interface{}
	var parametros []interface{}
	parametros = append(parametros, idEstadoNew)
	parametros = append(parametros, idUsuario)

	if ordenPagoData != nil {
		done := make(chan interface{})
		defer close(done)
		resch := optimize.GenChanInterface(ordenPagoData...)
		chlistaLiquidacion := optimize.Digest(done, changeEstadoOP, resch, parametros)
		for dataLiquidacion := range chlistaLiquidacion {
			if dataLiquidacion != nil {
				respuesta = append(respuesta, dataLiquidacion.(map[string]interface{}))
			}
		}
		return respuesta, nil
	} else {
		return nil, nil
	}
}

func changeEstadoOP(ordenPago interface{}, params ...interface{}) (res interface{}) {
	neworden, e := ordenPago.(map[string]interface{})
	alerta := make(map[string]interface{})
	if e {
		o := orm.NewOrm()
		newEstadoOp := OrdenPagoEstadoOrdenPago{}
		newEstadoOp.OrdenPago = &OrdenPago{Id: int(neworden["Id"].(float64))}
		newEstadoOp.EstadoOrdenPago = &EstadoOrdenPago{Id: int(params[0].(float64))}
		newEstadoOp.FechaRegistro = time.Now()
		newEstadoOp.Usuario = int(params[1].(float64))
		_, err := o.Insert(&newEstadoOp)
		if err != nil {
			alerta["Type"] = "error"
			alerta["Code"] = "E_OP_E_ACTUALIZAR"
			alerta["Body"] = err.Error()
			o.Rollback()
			return alerta
		} else {
			alerta["Type"] = "success"
			alerta["Code"] = "S_OP_ESTADO"
			alerta["Body"] = newEstadoOp
			o.Commit()
			return alerta
		}
	}
	return nil
}
