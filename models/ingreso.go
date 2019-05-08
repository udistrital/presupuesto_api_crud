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

type Ingreso struct {
	Id                   int                     `orm:"column(id);pk;auto"`
	Consecutivo          float64                 `orm:"column(consecutivo)"`
	Vigencia             float64                 `orm:"column(vigencia)"`
	FechaIngreso         time.Time               `orm:"column(fecha_ingreso);type(date)"`
	FechaInicio          time.Time               `orm:"column(fecha_inicio);type(date)"`
	FechaFin             time.Time               `orm:"column(fecha_fin);type(date)"`
	Facultad             int                     `orm:"column(facultad);null"`
	Observaciones        string                  `orm:"column(observaciones);null"`
	FuenteFinanciamiento *FuenteFinanciamiento   `orm:"column(fuente_financiamiento);rel(fk);null"`
	FormaIngreso         *FormaIngreso           `orm:"column(forma_ingreso);rel(fk)"`
	UnidadEjecutora      *UnidadEjecutora        `orm:"column(unidad_ejecutora);rel(fk)"`
	Aportante            int                     `orm:"column(aportante);null"`
	Reviso               int                     `orm:"column(reviso);null"`
	Elaboro              int                     `orm:"column(elaboro)"`
	MotivoRechazo        string                  `orm:"column(motivo_rechazo)"`
	IngresoConcepto      []*IngresoConcepto      `orm:"reverse(many)"`
	IngresoEstadoIngreso []*IngresoEstadoIngreso `orm:"reverse(many)"`
	DocumentoGenerador   *DocumentoGenerador     `orm:"column(documento_generador);rel(fk);null"`
	NumCuenta            string                  `orm:"column(num_cuenta)"`
}

func (t *Ingreso) TableName() string {
	return "ingreso"
}

func init() {
	orm.RegisterModel(new(Ingreso))
}

func RechazoContableIngreso(m map[string]interface{}) (ingreso Ingreso, err error) {
	o := orm.NewOrm()
	o.Begin()
	err = formatdata.FillStruct(m, &ingreso)
	if err != nil {
		o.Rollback()
		return
	}
	_, err = o.Update(&ingreso, "MotivoRechazo")
	if err != nil {
		o.Rollback()
		return
	}
	fmt.Println(ingreso)
	estadoIngreso := EstadoIngreso{Id: 3}
	ingresoEstadoIngreso := IngresoEstadoIngreso{}
	ingresoEstadoIngreso.Ingreso = &ingreso
	ingresoEstadoIngreso.EstadoIngreso = &estadoIngreso
	ingresoEstadoIngreso.FechaRegistro = time.Now()
	_, err = o.Insert(&ingresoEstadoIngreso)
	if err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return
}

func RechazoPresupuestalIngreso(m map[string]interface{}) (ingreso Ingreso, err error) {
	o := orm.NewOrm()
	o.Begin()
	err = formatdata.FillStruct(m, &ingreso)
	if err != nil {
		beego.Info("Debug ", err)
		o.Rollback()
		return
	}
	_, err = o.Update(&ingreso, "MotivoRechazo")
	if err != nil {
		o.Rollback()
		return
	}
	fmt.Println(ingreso)
	estadoIngreso := EstadoIngreso{Id: 5}
	ingresoEstadoIngreso := IngresoEstadoIngreso{}
	ingresoEstadoIngreso.Ingreso = &ingreso
	ingresoEstadoIngreso.EstadoIngreso = &estadoIngreso
	ingresoEstadoIngreso.FechaRegistro = time.Now()
	_, err = o.Insert(&ingresoEstadoIngreso)
	if err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return
}

func AprobacionContableIngreso(m map[string]interface{}) (ingreso Ingreso, err error) {
	o := orm.NewOrm()
	o.Begin()
	err = formatdata.FillStruct(m["Ingreso"], &ingreso)
	fmt.Println(ingreso)
	estadoIngreso := EstadoIngreso{Id: 2}
	ingresoEstadoIngreso := IngresoEstadoIngreso{}
	ingresoEstadoIngreso.Ingreso = &ingreso
	ingresoEstadoIngreso.EstadoIngreso = &estadoIngreso
	ingresoEstadoIngreso.FechaRegistro = time.Now()
	_, err = o.Insert(&ingresoEstadoIngreso)
	if err != nil {
		o.Rollback()
		return
	}
	var mov []MovimientoContable
	err = formatdata.FillStruct(m["Movimientos"], &mov)
	if err != nil {
		o.Rollback()
		return
	}
	for _, element := range mov {
		element.EstadoMovimientoContable.Id = 2
		_, err = o.Update(&element, "EstadoMovimientoContable")
		if err != nil {
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}

func AprobacionPresupuestalIngreso(m map[string]interface{}) (ingreso Ingreso, err error) {
	o := orm.NewOrm()
	o.Begin()
	err = formatdata.FillStruct(m["Ingreso"], &ingreso)
	fmt.Println(ingreso)
	estadoIngreso := EstadoIngreso{Id: 4}
	ingresoEstadoIngreso := IngresoEstadoIngreso{}
	ingresoEstadoIngreso.Ingreso = &ingreso
	ingresoEstadoIngreso.EstadoIngreso = &estadoIngreso
	ingresoEstadoIngreso.FechaRegistro = time.Now()
	_, err = o.Insert(&ingresoEstadoIngreso)
	if err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return
}

// AddIngreso insert a new Ingreso into database and returns
// last inserted Id on success.
func AddIngresotr(m map[string]interface{}) (ingreso Ingreso, err error) {
	var id int64
	var idDocgenerador int64
	var docGen DocumentoGenerador
	var voidDocGen DocumentoGenerador
	err = formatdata.FillStruct(m["DocumentoGenerador"], &docGen)
	o := orm.NewOrm()
	o.Begin()
	if err == nil && !reflect.DeepEqual(docGen, voidDocGen) {
		idDocgenerador, err = o.Insert(&docGen)
		if err != nil {
			beego.Info(err)
			o.Rollback()
			return
		}
	}
	err = formatdata.FillStruct(m["Ingreso"], &ingreso)
	if err == nil {
		ingreso.FechaIngreso = time.Now()
		ingreso.Vigencia = float64(time.Now().Year())
		var consecutivo float64
		o.Raw(`SELECT COALESCE(MAX(consecutivo), 0)+1  as consecutivo
						FROM financiera.ingreso WHERE vigencia = ?`, ingreso.Vigencia).QueryRow(&consecutivo)
		ingreso.Consecutivo = consecutivo
		if !reflect.DeepEqual(docGen, voidDocGen) {
			ingreso.DocumentoGenerador = &DocumentoGenerador{Id: int(idDocgenerador)}
		}
		//insert ingreso
		id, err = o.Insert(&ingreso)
		if err == nil {
			ingreso.Id = int(id)
			//crear el rompimiento para registrar el estado del ingreso.
			estadoIngreso := EstadoIngreso{Id: 1}
			ingresoEstadoIngreso := IngresoEstadoIngreso{}
			ingresoEstadoIngreso.Ingreso = &ingreso
			ingresoEstadoIngreso.EstadoIngreso = &estadoIngreso
			ingresoEstadoIngreso.FechaRegistro = time.Now()
			if _, err = o.Insert(&ingresoEstadoIngreso); err != nil {
				beego.Error("Error", err.Error())
				o.Rollback()
				return
			}
		} else {
			beego.Error("Error", err.Error())
			o.Rollback()
			return
		}

		//insert MovimientoContable
		var mov []MovimientoContable
		err = formatdata.FillStruct(m["Movimientos"], &mov)
		for _, element := range mov {
			element.Fecha = time.Now()
			element.TipoDocumentoAfectante = &TipoDocumentoAfectante{Id: 2}
			element.CodigoDocumentoAfectante = int(id)
			element.EstadoMovimientoContable = &EstadoMovimientoContable{Id: 1}
			_, err = o.Insert(&element)
			if err != nil {
				beego.Error("Error", err.Error())
				o.Rollback()
				return
			}
		}

		if err != nil {
			beego.Info(err.Error())
			o.Rollback()
			return
		} else {

			var ingresos float64
			err = formatdata.FillStruct(m["IngresoBanco"], &ingresos)
			if err == nil {
				concepto := &Concepto{}
				fmt.Println("concepto ", m["Concepto"])
				err = formatdata.FillStruct(m["Concepto"], concepto)
				if err == nil {
					ingreso_concepto := &IngresoConcepto{ValorAgregado: ingresos,
						Ingreso:  &ingreso,
						Concepto: concepto}
					_, err = o.Insert(ingreso_concepto)
					if err != nil {
						beego.Info(err.Error())
						o.Rollback()
						return
					}

				} else {
					beego.Info(err.Error())
					o.Rollback()
					return
				}

			} else {
				beego.Info(err.Error())
				o.Rollback()
				return
			}

			o.Commit()
			return
		}
	} else {
		beego.Info(err)
		o.Rollback()
		return
	}
	o.Rollback()
	return
}

// AddIngreso insert a new Ingreso into database and returns
// last inserted Id on success.
func AddIngreso(m *Ingreso) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetIngresoById retrieves Ingreso by Id. Returns error if
// Id doesn't exist
func GetIngresoById(id int) (v *Ingreso, err error) {
	o := orm.NewOrm()
	v = &Ingreso{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllIngreso retrieves all Ingreso matches certain condition. Returns empty list if
// no records exist
func GetAllIngreso(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Ingreso)).RelatedSel()
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

	var l []Ingreso
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "IngresoConcepto", 5)
				o.LoadRelated(&v, "IngresoEstadoIngreso", 5, 1, 0, "-Id")
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

// UpdateIngreso updates Ingreso by Id and returns error if
// the record to be updated doesn't exist
func UpdateIngresoById(m *Ingreso) (err error) {
	o := orm.NewOrm()
	v := Ingreso{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteIngreso deletes Ingreso by Id and returns error if
// the record to be deleted doesn't exist
func DeleteIngreso(id int) (err error) {
	o := orm.NewOrm()
	v := Ingreso{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Ingreso{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
