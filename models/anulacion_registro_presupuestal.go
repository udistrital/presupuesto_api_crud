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

type AnulacionRegistroPresupuestal struct {
	Id                   int                        `orm:"auto;column(id);pk"`
	Consecutivo          int                        `orm:"column(consecutivo)"`
	Motivo               string                     `orm:"column(motivo)"`
	FechaRegistro        time.Time                  `orm:"column(fecha_registro);type(date)"`
	TipoAnulacion        *TipoAnulacionPresupuestal `orm:"column(tipo_anulacion);rel(fk)"`
	EstadoAnulacion      *EstadoAnulacion           `orm:"column(estado_anulacion);rel(fk)"`
	JustificacionRechazo string                     `orm:"column(justificacion_rechazo);null"`
	Responsable          int                        `orm:"column(responsable)"`
	Solicitante          int                        `orm:"column(solicitante)"`
	Expidio              int                        `orm:"column(expidio)"`

	AnulacionRegistroPresupuestalDisponibilidadApropiacion []*AnulacionRegistroPresupuestalDisponibilidadApropiacion `orm:"reverse(many)"`
}

func (t *AnulacionRegistroPresupuestal) TableName() string {
	return "anulacion_registro_presupuestal"
}

func init() {
	orm.RegisterModel(new(AnulacionRegistroPresupuestal))
}

// totalAnulacionesDisponibilidades retorna total de disponibilidades por vigencia
func GetTotalAnulacionRegistroPresupuestal(vigencia int, unidadEjecutora int) (total int, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COUNT(DISTINCT(" + beego.AppConfig.String("PGschemas") + ".anulacion_registro_presupuestal.id))").
		From("" + beego.AppConfig.String("PGschemas") + ".anulacion_registro_presupuestal").
		InnerJoin("" + beego.AppConfig.String("PGschemas") + ".anulacion_registro_presupuestal_disponibilidad_apropiacion").
		On("" + beego.AppConfig.String("PGschemas") + ".anulacion_registro_presupuestal.id = " + beego.AppConfig.String("PGschemas") + ".anulacion_registro_presupuestal_disponibilidad_apropiacion.anulacion_registro_presupuestal").
		InnerJoin("" + beego.AppConfig.String("PGschemas") + ".registro_presupuestal_disponibilidad_apropiacion").
		On("" + beego.AppConfig.String("PGschemas") + ".registro_presupuestal_disponibilidad_apropiacion.id = " + beego.AppConfig.String("PGschemas") + ".anulacion_registro_presupuestal_disponibilidad_apropiacion.registro_presupuestal_disponibilidad_apropiacion").
		InnerJoin("" + beego.AppConfig.String("PGschemas") + ".registro_presupuestal").
		On("" + beego.AppConfig.String("PGschemas") + ".registro_presupuestal.id = " + beego.AppConfig.String("PGschemas") + ".registro_presupuestal_disponibilidad_apropiacion.registro_presupuestal").
		InnerJoin("" + beego.AppConfig.String("PGschemas") + ".disponibilidad_apropiacion").
		On("" + beego.AppConfig.String("PGschemas") + ".disponibilidad_apropiacion.id = " + beego.AppConfig.String("PGschemas") + ".registro_presupuestal_disponibilidad_apropiacion.disponibilidad_apropiacion").
		InnerJoin("" + beego.AppConfig.String("PGschemas") + ".disponibilidad").
		On("" + beego.AppConfig.String("PGschemas") + ".disponibilidad.id = " + beego.AppConfig.String("PGschemas") + ".disponibilidad_apropiacion.disponibilidad").
		InnerJoin("" + beego.AppConfig.String("PGschemas") + ".apropiacion").
		On("" + beego.AppConfig.String("PGschemas") + ".disponibilidad_apropiacion.apropiacion = " + beego.AppConfig.String("PGschemas") + ".apropiacion.id").
		InnerJoin("" + beego.AppConfig.String("PGschemas") + ".rubro").
		On("" + beego.AppConfig.String("PGschemas") + ".rubro.id = " + beego.AppConfig.String("PGschemas") + ".apropiacion.rubro").
		Where("" + beego.AppConfig.String("PGschemas") + ".registro_presupuestal.vigencia = ? and " + beego.AppConfig.String("PGschemas") + ".rubro.unidad_ejecutora = ? ")

	err = o.Raw(qb.String(), vigencia, unidadEjecutora).QueryRow(&total)
	return

}

// AddAnulacionRegistroPresupuestal insert a new AnulacionRegistroPresupuestal into database and returns
// last inserted Id on success.
func AddAnulacionRegistroPresupuestal(m *AnulacionRegistroPresupuestal) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAnulacionRegistroPresupuestalById retrieves AnulacionRegistroPresupuestal by Id. Returns error if
// Id doesn't exist
func GetAnulacionRegistroPresupuestalById(id int) (v *AnulacionRegistroPresupuestal, err error) {
	o := orm.NewOrm()
	v = &AnulacionRegistroPresupuestal{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAnulacionRegistroPresupuestal retrieves all AnulacionRegistroPresupuestal matches certain condition. Returns empty list if
// no records exist
func GetAllAnulacionRegistroPresupuestal(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AnulacionRegistroPresupuestal))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		//beego.Info(k)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.Contains(k, "__in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else if strings.Contains(k, "__not_in") {
			//beego.Info(k)
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

	var l []AnulacionRegistroPresupuestal
	qs = qs.OrderBy(sortFields...).RelatedSel(5).Distinct()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "AnulacionRegistroPresupuestalDisponibilidadApropiacion", 5)
				for _, sub := range v.AnulacionRegistroPresupuestalDisponibilidadApropiacion {
					o.LoadRelated(sub.RegistroPresupuestalDisponibilidadApropiacion.DisponibilidadApropiacion.Disponibilidad, "DisponibilidadProcesoExterno", 5, 1, 0, "-Id")
				}
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

// UpdateAnulacionRegistroPresupuestal updates AnulacionRegistroPresupuestal by Id and returns error if
// the record to be updated doesn't exist
func UpdateAnulacionRegistroPresupuestalById(m *AnulacionRegistroPresupuestal, fields ...string) (err error) {
	o := orm.NewOrm()
	v := AnulacionRegistroPresupuestal{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, fields...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAnulacionRegistroPresupuestal deletes AnulacionRegistroPresupuestal by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAnulacionRegistroPresupuestal(id int) (err error) {
	o := orm.NewOrm()
	v := AnulacionRegistroPresupuestal{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AnulacionRegistroPresupuestal{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
