package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type Apropiacion struct {
	Id       int                `orm:"column(id);pk;auto"`
	Vigencia float64            `orm:"column(vigencia)"`
	Rubro    *Rubro             `orm:"column(rubro);rel(fk)"`
	Valor    float64            `orm:"column(valor)"`
	Estado   *EstadoApropiacion `orm:"column(estado);rel(fk)"`
}

func (t *Apropiacion) TableName() string {
	return "apropiacion"
}

func init() {
	orm.RegisterModel(new(Apropiacion))
}

// AddApropiacion insert a new Apropiacion into database and returns
// last inserted Id on success.
func AddApropiacion(m *Apropiacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetApropiacionById retrieves Apropiacion by Id. Returns error if
// Id doesn't exist
func GetApropiacionById(id int) (v *Apropiacion, err error) {
	o := orm.NewOrm()
	v = &Apropiacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllApropiacion retrieves all Apropiacion matches certain condition. Returns empty list if
// no records exist
func GetAllApropiacion(query map[string]string, exclude map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Apropiacion))
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

	// exclude k=v
	for k, v := range exclude {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Exclude(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Exclude(k, v)
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

	var l []Apropiacion
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

// UpdateApropiacion updates Apropiacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateApropiacionById(m *Apropiacion) (err error) {
	o := orm.NewOrm()
	v := Apropiacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteApropiacion deletes Apropiacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteApropiacion(id int) (err error) {
	o := orm.NewOrm()
	v := Apropiacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Apropiacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//funcion para comprobar saldo de la apropiacion de un Rubro

func SaldoApropiacion(Id int) (saldo map[string]float64, err error) {
	var valor float64
	saldo = make(map[string]float64)
	valorapr, err := ValorApropiacion(Id)
	if err != nil {
		return
	}
	valorcdpapr, err := ValorCdpPorApropiacion(Id)
	if err != nil {
		return
	}
	valoranuladocdpapr, err := ValorAnuladoCdpPorApropiacion(Id)
	if err != nil {
		return
	}
	valorAdiciones, err := ValorMovimientosPorApropiacion(Id, 3, "cuenta_credito")
	if err != nil {
		return
	}
	valorAdicionesTraslados, err := ValorMovimientosPorApropiacion(Id, 1, "cuenta_credito")
	if err != nil {
		return
	}
	valorReducciones, err := ValorMovimientosPorApropiacion(Id, 2, "cuenta_credito")
	if err != nil {
		return
	}
	valorReduccionesTraslados, err := ValorMovimientosPorApropiacion(Id, 1, "cuenta_contra_credito")
	if err != nil {
		return
	}
	valor = valorapr - valorcdpapr + valoranuladocdpapr + valorAdiciones + valorAdicionesTraslados
	saldo["original"] = valorapr
	saldo["saldo"] = valor
	saldo["comprometido"] = valorcdpapr - valoranuladocdpapr
	saldo["adiciones"] = valorAdiciones
	saldo["traslados"] = valorAdicionesTraslados
	saldo["reducciones"] = valorReducciones + valorReduccionesTraslados
	saldo["comprometido_anulado"] = valoranuladocdpapr
	return
}

func VigenciaApropiacion() (ml []int, err error) {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("DISTINCT vigencia").
		From("" + beego.AppConfig.String("PGschemas") + ".apropiacion")

	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&ml)

	if len(ml) == 0 {
		return nil, err
	}
	return ml, nil
}

//funcion para determinar el valor con traslados de la apropiacion
func ValorApropiacion(Id int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps_valor_tot []orm.Params
	_, err = o.Raw(`SELECT valor
				FROM `+beego.AppConfig.String("PGschemas")+`.apropiacion
				WHERE id= ?`, Id).Values(&maps_valor_tot)
	//fmt.Println("maps: ", len(maps_valor_tot))
	if len(maps_valor_tot) > 0 && err == nil {
		valor, _ = strconv.ParseFloat(maps_valor_tot[0]["valor"].(string), 64)
	} else {
		valor = 0
	}

	return
}

//funcion para determinar el total del valor de los cdp hechos a una apropiacion
func ValorMovimientosPorApropiacion(Id int, tipoMov int, cuenta string) (valor float64, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COALESCE(sum(valor),0) as valor").
		From("" + beego.AppConfig.String("PGschemas") + ".movimiento_apropiacion_disponibilidad_apropiacion").
		InnerJoin("" + beego.AppConfig.String("PGschemas") + ".movimiento_apropiacion").
		On("movimiento_apropiacion.id = movimiento_apropiacion_disponibilidad_apropiacion.movimiento_apropiacion").
		Where(cuenta + " = ?").
		And("tipo_movimiento_apropiacion = ?").
		And("estado_movimiento_apropiacion = 2")
	err = o.Raw(qb.String(), Id, tipoMov).QueryRow(&valor)
	return
}

//funcion para determinar el total del valor de los cdp hechos a una apropiacion
func ValorCdpPorApropiacion(Id int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps_valor_tot []orm.Params
	_, err = o.Raw(`SELECT  disponibilidad_apropiacion.apropiacion,
		COALESCE(sum(disponibilidad_apropiacion.valor),0) AS valor
	   FROM `+beego.AppConfig.String("PGschemas")+`.disponibilidad
		 JOIN `+beego.AppConfig.String("PGschemas")+`.disponibilidad_apropiacion ON disponibilidad_apropiacion.disponibilidad = disponibilidad.id
		 WHERE apropiacion= ?
		 GROUP BY disponibilidad_apropiacion.apropiacion
				`, Id).Values(&maps_valor_tot)
	//fmt.Println("maps: ", len(maps_valor_tot))
	if len(maps_valor_tot) > 0 && err == nil {
		valor, _ = strconv.ParseFloat(maps_valor_tot[0]["valor"].(string), 64)
	} else {
		valor = 0
	}

	return
}

//funcion para determinar el total del valor de los cdp hechos a una apropiacion
func ValorAnuladoCdpPorApropiacion(Id int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps_valor_tot []orm.Params
	_, err = o.Raw(`SELECT anulacion_disponibilidad.estado_anulacion,
								disponibilidad_apropiacion.apropiacion,
								COALESCE(sum(anulacion_disponibilidad_apropiacion.valor),0) AS valor
	   						FROM `+beego.AppConfig.String("PGschemas")+`.anulacion_disponibilidad_apropiacion
		 					JOIN `+beego.AppConfig.String("PGschemas")+`.disponibilidad_apropiacion ON anulacion_disponibilidad_apropiacion.disponibilidad_apropiacion = disponibilidad_apropiacion.id
		 					JOIN `+beego.AppConfig.String("PGschemas")+`.disponibilidad ON disponibilidad_apropiacion.disponibilidad = disponibilidad.id
					 		JOIN `+beego.AppConfig.String("PGschemas")+`.anulacion_disponibilidad ON anulacion_disponibilidad.id = anulacion_disponibilidad_apropiacion.anulacion
							 WHERE apropiacion = ?  AND estado_anulacion = 3
							 GROUP BY  anulacion_disponibilidad.estado_anulacion, disponibilidad_apropiacion.apropiacion
							`, Id).Values(&maps_valor_tot)
	//fmt.Println("maps: ", len(maps_valor_tot))
	if len(maps_valor_tot) > 0 && err == nil {
		valor, _ = strconv.ParseFloat(maps_valor_tot[0]["valor"].(string), 64)
	} else {
		valor = 0
	}

	return
}

//AprobarPresupuesto... Aprobacion de presupuesto (cambio de estado).
func AprobarPresupuesto(UnidadEjecutora int, Vigencia int) (err error) {
	query := make(map[string]string)
	o := orm.NewOrm()
	query["Rubro.UnidadEjecutora"] = strconv.Itoa(UnidadEjecutora)
	query["Vigencia"] = strconv.Itoa(Vigencia)
	fmt.Println(query)
	v, err := GetAllApropiacion(query, nil, nil, nil, nil, 0, -1)
	o.Begin()
	ap := Apropiacion{}
	for _, apropiacion := range v {
		formatdata.FillStruct(apropiacion, &ap)
		ap.Estado.Id = 2
		_, err = o.Update(&ap)
		if err != nil {
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}

//UpdateApropiacionValue... Actualiza la apropiacion inicial de un rubro dado un id
func UpdateApropiacionValue(id int, valor float64) (err error) {
	o := orm.NewOrm()
	apropiacion := &Apropiacion{Id: id, Valor: valor}
	_, err = o.Update(apropiacion, "Valor")
	if err != nil {
		panic(err.Error())
	}
	return
}
