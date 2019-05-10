package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/udistrital/utils_oas/formatdata"
)

type MovimientoApropiacion struct {
	Id                                             int                                               `orm:"auto;column(id);pk"`
	NumeroMovimiento                               int                                               `orm:"column(numero_movimiento)"`
	Vigencia                                       int                                               `orm:"column(vigencia)"`
	FechaMovimiento                                time.Time                                         `orm:"column(fecha_movimiento);type(date)"`
	Noficio                                        int                                               `orm:"column(n_oficio)"`
	Foficio                                        time.Time                                         `orm:"column(f_oficio);type(date)"`
	Descripcion                                    string                                            `orm:"column(descripcion);null"`
	MovimientoApropiacionDisponibilidadApropiacion []*MovimientoApropiacionDisponibilidadApropiacion `orm:"reverse(many)"`
	EstadoMovimientoApropiacion                    *EstadoMovimientoApropiacion                      `orm:"column(estado_movimiento_apropiacion);rel(fk)"`
	UnidadEjecutora                                int                                               `orm:"column(unidad_ejecutora)"`
}

type MovimientosPorApropiacion struct {
	NumeroDisponibilidad float64   `orm:"column(numero_disponibilidad)"`
	NumeroMovimiento     float64   `orm:"column(numero_movimiento)"`
	CuentaContraCredito  string    `orm:"column(cuenta_contra_credito)"`
	CuentaCredito        string    `orm:"column(cuenta_credito)"`
	Valor                float64   `orm:"column(valor)"`
	Tipo                 string    `orm:"column(tipo)"`
	Noficio              float64   `orm:"column(n_oficio)"`
	Foficio              time.Time `orm:"column(f_oficio);type(date)"`
}

func (t *MovimientoApropiacion) TableName() string {
	return "movimiento_apropiacion"
}

func init() {
	orm.RegisterModel(new(MovimientoApropiacion))
}

// AddMovimientoApropiacion insert a new MovimientoApropiacion into database and returns
// last inserted Id on success.
func AddMovimientoApropiacion(m *MovimientoApropiacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// totalMovimientos retorna total de movimientos por vigencia
func GetTotalMovimientosApropiacion(vigencia int, unidadEjecutora int) (total int, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("COUNT(DISTINCT(movimiento_apropiacion))").
		From("" + beego.AppConfig.String("PGschemas") + ".movimiento_apropiacion").
		Where("movimiento_apropiacion.vigencia = ?").
		And("unidad_ejecutora = ?")
	err = o.Raw(qb.String(), vigencia, unidadEjecutora).QueryRow(&total)
	return

}

// Registra un MovimientoApropiacion
// retorna structura de alerta
func RegistrarMovimietnoApropiaciontr(movimiento map[string]interface{}) (alert Alert, err error) {
	var movimientoapr MovimientoApropiacion
	var desgrMovimientoApr []MovimientoApropiacionDisponibilidadApropiacion
	if err = formatdata.FillStruct(movimiento["MovimientoApropiacion"], &movimientoapr); err == nil {
		if err = formatdata.FillStruct(movimiento["MovimientoApropiacionDisponibilidadApropiacion"], &desgrMovimientoApr); err == nil {
			o := orm.NewOrm()
			o.Begin()
			var consecutivo int
			vigencia := time.Now().Year()
			o.Raw(`SELECT COALESCE(MAX(numero_movimiento), 0)+1  as consecutivo
						FROM `+beego.AppConfig.String("PGschemas")+`.movimiento_apropiacion 
						WHERE vigencia = ?`, vigencia).QueryRow(&consecutivo)

			movimientoapr.FechaMovimiento = time.Now().Local() //asignacion de la fecha de la solicitud del movimiento
			movimientoapr.Vigencia = vigencia
			movimientoapr.EstadoMovimientoApropiacion = &EstadoMovimientoApropiacion{Id: 1}
			movimientoapr.NumeroMovimiento = consecutivo
			_, err = o.Insert(&movimientoapr)
			if err != nil {
				o.Rollback()
				return
			}

			for _, datDesgrMov := range desgrMovimientoApr {
				datDesgrMov.MovimientoApropiacion = &movimientoapr
				_, err = o.Insert(&datDesgrMov)
				if err != nil {
					o.Rollback()
					return
				}

			}

			o.Commit()

		} else {
			alert.Code = "E_0458"
			alert.Body = err
			alert.Type = "error"
			return
		}
	} else {
		alert.Code = "E_0458"
		alert.Body = err
		alert.Type = "error"
		return
	}
	alert.Code = "S_MODP001"
	alert.Body = map[string]interface{}{"MovimientoApropiacion": movimientoapr, "MovimientoApropiacionDisponibilidadApropiacion": desgrMovimientoApr}
	alert.Type = "success"
	return
}
func aprobacionMovimientoPresupuestalDispatcher(tipo *TipoMovimientoApropiacion) (f func(data *MovimientoApropiacionDisponibilidadApropiacion, o *orm.Ormer) (alert Alert, err error)) {
	switch os := tipo.Disponibilidad; os {
	case true:
		return registroModificacionPresupuestalCDP
	default:
		return nil
	}
}

func registroModificacionPresupuestalCDP(movimiento *MovimientoApropiacionDisponibilidadApropiacion, o *orm.Ormer) (alert Alert, err error) {
	valorCDP := movimiento.Valor
	var saldoApr map[string]float64
	movDestino := &Apropiacion{}
	if movimiento.CuentaContraCredito != nil {
		saldoApr, err = SaldoApropiacion(movimiento.CuentaContraCredito.Id)
		movDestino = movimiento.CuentaContraCredito
	} else {
		saldoApr, err = SaldoApropiacion(movimiento.CuentaCredito.Id)
		movDestino = movimiento.CuentaCredito
	}
	or := orm.NewOrm()
	if err == nil {
		if valorCDP <= saldoApr["saldo"] {
			disponibilidad := make(map[string]interface{})
			disponibilidad["Vigencia"] = float64(movimiento.MovimientoApropiacion.Vigencia)
			disponibilidad["FechaRegistro"] = time.Now().Local()
			disponibilidad["Estado"] = map[string]interface{}{"Id": 1}
			//disponibilidad["Solicitud"] = int(solicitud["SolicitudDisponibilidad"].(map[string]interface{})["Id"].(float64))
			disponibilidad["Responsable"] = 876543216
			disponibilidad["UnidadEjecutora"] = float64(1)
			DisponibilidadProcesoExterno := map[string]interface{}{"ProcesoExterno": movimiento.MovimientoApropiacion.Id}
			TipoDisponibilidad := map[string]interface{}{"Id": 2}
			DisponibilidadProcesoExterno["TipoDisponibilidad"] = TipoDisponibilidad

			var afectacion []interface{}

			disponibilidadApropiacion := make(map[string]interface{})
			disponibilidadApropiacion["Apropiacion"] = movDestino
			disponibilidadApropiacion["Valor"] = movimiento.Valor
			disponibilidadApropiacion["FuenteFinanciamiento"] = map[string]interface{}{"Id": 0}
			afectacion = append(afectacion, disponibilidadApropiacion)

			infoDisponibilidad := make(map[string]interface{})
			infoDisponibilidad["Disponibilidad"] = disponibilidad
			infoDisponibilidad["DisponibilidadApropiacion"] = afectacion
			infoDisponibilidad["DisponibilidadProcesoExterno"] = DisponibilidadProcesoExterno

			dispexp, err1 := AddDisponibilidad(infoDisponibilidad)
			movimiento.Disponibilidad = &dispexp
			_, err2 := or.Update(movimiento, "Disponibilidad")
			if err1 == nil && err2 == nil {
				alert.Type = "success"
				alert.Code = "S_MODP003"
				alert.Body = map[string]interface{}{"Movimiento": movimiento.MovimientoApropiacion, "Disponibilidad": dispexp.NumeroDisponibilidad, "Apropiacion": movDestino.Rubro.Codigo}
				return
			} else {
				alert.Type = "error"
				alert.Code = "E_MODP006"
				alert.Body = map[string]interface{}{"Movimiento": movimiento.MovimientoApropiacion, "Disponibilidad": 0, "Apropiacion": movDestino.Rubro.Codigo}
				return alert, err1
			}

		} else {
			alert.Type = "error"
			alert.Code = "E_MODP004"
			alert.Body = map[string]interface{}{"Movimiento": movimiento.MovimientoApropiacion, "Disponibilidad": 0, "Apropiacion": movDestino.Rubro.Codigo}
			err = errors.New("E_MODP004")
			return
		}

	} else {
		alert.Type = "error"
		alert.Code = "E_MODP005"
		alert.Body = map[string]interface{}{"Movimiento": movimiento.MovimientoApropiacion, "Disponibilidad": 0, "Apropiacion": movDestino.Rubro.Codigo}
		return
	}
	return
}

// Aprueba un MovimientoApropiacion
// retorna structura de alerta
func AprobarMovimietnoApropiaciontr(movimiento *MovimientoApropiacion) (alert []Alert, err error) {
	o := orm.NewOrm()
	o.Begin()
	for _, desgrMov := range movimiento.MovimientoApropiacionDisponibilidadApropiacion {
		f := aprobacionMovimientoPresupuestalDispatcher(desgrMov.TipoMovimientoApropiacion)
		if f != nil {
			alt, err1 := f(desgrMov, &o)
			fmt.Println("err ", err1)
			if err1 != nil {
				o.Rollback()
				alert = append(alert, alt)
				return alert, err
			}
			alert = append(alert, alt)
		} else {

		}
	}
	movimiento.EstadoMovimientoApropiacion.Id = 2
	_, err = o.Update(movimiento, "EstadoMovimientoApropiacion")
	if err != nil {
		alt := Alert{}
		alt.Type = "error"
		alt.Code = "E_MODP007"
		alt.Body = map[string]interface{}{"Movimiento": movimiento, "Disponibilidad": 0, "Apropiacion": 0}
		alert = append(alert, alt)
		o.Rollback()
	} else {
		alt := Alert{}
		alt.Type = "success"
		alt.Code = "S_MODP002"
		alt.Body = map[string]interface{}{"Movimiento": movimiento, "Disponibilidad": 0, "Apropiacion": 0}
		alert = append(alert, alt)
		o.Commit()
	}

	return
}

// Consulta los movimientos de una apropiacion segun su id
// retorna los movimientos de una apropiacion
func MovimientosByApropiacion(apropiacionId int) (res []MovimientosPorApropiacion, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("numero_movimiento,numero_disponibilidad,rb_ccr.codigo as cuenta_contra_credito," +
		"rb_cr.codigo as cuenta_credito," +
		"movimiento_apropiacion_disponibilidad_apropiacion.valor," +
		"tipo_movimiento_apropiacion.nombre as tipo," +
		"n_oficio," +
		"f_oficio").
		From("" + beego.AppConfig.String("PGschemas") + ".movimiento_apropiacion").
		LeftJoin("" + beego.AppConfig.String("PGschemas") + ".movimiento_apropiacion_disponibilidad_apropiacion").
		On("movimiento_apropiacion = movimiento_apropiacion.id").
		LeftJoin("" + beego.AppConfig.String("PGschemas") + ".tipo_movimiento_apropiacion").
		On("tipo_movimiento_apropiacion.id = tipo_movimiento_apropiacion").
		LeftJoin("" + beego.AppConfig.String("PGschemas") + ".apropiacion as apr_cr").
		On("cuenta_credito = apr_cr.id").
		LeftJoin("" + beego.AppConfig.String("PGschemas") + ".rubro as rb_cr").
		On("rb_cr.id = apr_cr.rubro").
		LeftJoin("" + beego.AppConfig.String("PGschemas") + ".apropiacion as apr_ccr").
		On("movimiento_apropiacion_disponibilidad_apropiacion.cuenta_contra_credito = apr_ccr.id").
		LeftJoin("" + beego.AppConfig.String("PGschemas") + ".rubro as rb_ccr").
		On("rb_ccr.id = apr_ccr.rubro").
		LeftJoin("" + beego.AppConfig.String("PGschemas") + ".disponibilidad").
		On("disponibilidad.id = movimiento_apropiacion_disponibilidad_apropiacion.disponibilidad").
		Where("movimiento_apropiacion.estado_movimiento_apropiacion = 2 AND (movimiento_apropiacion_disponibilidad_apropiacion.cuenta_contra_credito = ? OR cuenta_credito = ?)")
	_, err = o.Raw(qb.String(), apropiacionId, apropiacionId).QueryRows(&res)
	return
}

// GetMovimientoApropiacionById retrieves MovimientoApropiacion by Id. Returns error if
// Id doesn't exist
func GetMovimientoApropiacionById(id int) (v *MovimientoApropiacion, err error) {
	o := orm.NewOrm()
	v = &MovimientoApropiacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMovimientoApropiacion retrieves all MovimientoApropiacion matches certain condition. Returns empty list if
// no records exist
func GetAllMovimientoApropiacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MovimientoApropiacion))
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

	var l []MovimientoApropiacion
	qs = qs.OrderBy(sortFields...).RelatedSel(5)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "MovimientoApropiacionDisponibilidadApropiacion", 5)
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

// UpdateMovimientoApropiacion updates MovimientoApropiacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateMovimientoApropiacionById(m *MovimientoApropiacion, fields []string) (err error) {
	o := orm.NewOrm()
	v := MovimientoApropiacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, fields...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMovimientoApropiacion deletes MovimientoApropiacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMovimientoApropiacion(id int) (err error) {
	o := orm.NewOrm()
	v := MovimientoApropiacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MovimientoApropiacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
