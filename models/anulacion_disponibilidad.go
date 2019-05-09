package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type AnulacionDisponibilidad struct {
	Id                                 int                                   `orm:"column(id);pk;auto"`
	Consecutivo                        int                                   `orm:"column(consecutivo)"`
	Motivo                             string                                `orm:"column(motivo)"`
	FechaRegistro                      time.Time                             `orm:"column(fecha_registro);type(date)"`
	TipoAnulacion                      *TipoAnulacionPresupuestal            `orm:"column(tipo_anulacion);rel(fk)"`
	EstadoAnulacion                    *EstadoAnulacion                      `orm:"column(estado_anulacion);rel(fk)"`
	JustificacionRechazo               string                                `orm:"column(justificacion_rechazo);null"`
	Responsable                        int                                   `orm:"column(responsable)"`
	Solicitante                        int                                   `orm:"column(solicitante)"`
	Expidio                            int                                   `orm:"column(expidio)"`
	AnulacionDisponibilidadApropiacion []*AnulacionDisponibilidadApropiacion `orm:"reverse(many)"`
}

func (t *AnulacionDisponibilidad) TableName() string {
	return "anulacion_disponibilidad"
}

func init() {
	orm.RegisterModel(new(AnulacionDisponibilidad))
}

// totalAnulacionesDisponibilidades retorna total de disponibilidades por vigencia
func GetTotalAnulacionDisponibilidades(vigencia int, unidadEjecutora int) (total int, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COUNT(DISTINCT(financiera.anulacion_disponibilidad.id))").
		From("financiera.anulacion_disponibilidad").
		InnerJoin("financiera.anulacion_disponibilidad_apropiacion").
		On("financiera.anulacion_disponibilidad.id = financiera.anulacion_disponibilidad_apropiacion.anulacion").
		InnerJoin("financiera.disponibilidad_apropiacion").
		On("financiera.disponibilidad_apropiacion.id = financiera.anulacion_disponibilidad_apropiacion.disponibilidad_apropiacion").
		InnerJoin("financiera.disponibilidad").
		On("financiera.disponibilidad.id = financiera.disponibilidad_apropiacion.disponibilidad").
		InnerJoin("financiera.apropiacion").
		On("financiera.disponibilidad_apropiacion.apropiacion = financiera.apropiacion.id").
		InnerJoin("financiera.rubro").
		On("financiera.rubro.id = financiera.apropiacion.rubro").
		Where("financiera.disponibilidad.vigencia = ? and financiera.rubro.unidad_ejecutora = ? ")

	err = o.Raw(qb.String(), vigencia, unidadEjecutora).QueryRow(&total)
	return

}

// AddAnulacionDisponibilidad insert a new AnulacionDisponibilidad into database and returns
// last inserted Id on success.
func AddAnulacionDisponibilidad(m *AnulacionDisponibilidad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAnulacionDisponibilidadById retrieves AnulacionDisponibilidad by Id. Returns error if
// Id doesn't exist
func GetAnulacionDisponibilidadById(id int) (v *AnulacionDisponibilidad, err error) {
	o := orm.NewOrm()
	v = &AnulacionDisponibilidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAnulacionDisponibilidad retrieves all AnulacionDisponibilidad matches certain condition. Returns empty list if
// no records exist
func GetAllAnulacionDisponibilidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AnulacionDisponibilidad))
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

	var l []AnulacionDisponibilidad
	qs = qs.OrderBy(sortFields...).RelatedSel(5).Distinct()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "AnulacionDisponibilidadApropiacion", 5)
				for _, sub := range v.AnulacionDisponibilidadApropiacion {
					o.LoadRelated(sub.DisponibilidadApropiacion.Disponibilidad, "DisponibilidadProcesoExterno", 5, 1, 0, "-Id")
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

// UpdateAnulacionDisponibilidad updates AnulacionDisponibilidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateAnulacionDisponibilidadById(m *AnulacionDisponibilidad, fields ...string) (err error) {
	o := orm.NewOrm()
	v := AnulacionDisponibilidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, fields...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAnulacionDisponibilidad deletes AnulacionDisponibilidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAnulacionDisponibilidad(id int) (err error) {
	o := orm.NewOrm()
	v := AnulacionDisponibilidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AnulacionDisponibilidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
