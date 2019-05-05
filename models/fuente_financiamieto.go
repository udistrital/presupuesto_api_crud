package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/manucorporat/try"
	"github.com/udistrital/utils_oas/formatdata"
)

type FuenteFinanciamiento struct {
	Id                       int                       `orm:"column(id);pk;auto"`
	Descripcion              string                    `orm:"column(descripcion);null"`
	Nombre                   string                    `orm:"column(nombre)"`
	Codigo                   string                    `orm:"column(codigo)"`
	TipoFuenteFinanciamiento *TipoFuenteFinanciamiento `orm:"column(tipo_fuente_financiamiento);rel(fk)"`
}

func (t *FuenteFinanciamiento) TableName() string {
	return "fuente_financiamiento"
}

func init() {
	orm.RegisterModel(new(FuenteFinanciamiento))
}

// AddFuenteFinanciamiento insert a new FuenteFinanciamiento into database and returns
// last inserted Id on success.
func AddFuenteFinanciamiento(m *FuenteFinanciamiento) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFuenteFinanciamientoById retrieves FuenteFinanciamiento by Id. Returns error if
// Id doesn't exist
func GetFuenteFinanciamientoById(id int) (v *FuenteFinanciamiento, err error) {
	o := orm.NewOrm()
	v = &FuenteFinanciamiento{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFuenteFinanciamiento retrieves all FuenteFinanciamiento matches certain condition. Returns empty list if
// no records exist
func GetAllFuenteFinanciamiento(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FuenteFinanciamiento)).RelatedSel()
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

	var l []FuenteFinanciamiento
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

// UpdateFuenteFinanciamiento updates FuenteFinanciamiento by Id and returns error if
// the record to be updated doesn't exist
func UpdateFuenteFinanciamientoById(m *FuenteFinanciamiento) (err error) {
	o := orm.NewOrm()
	v := FuenteFinanciamiento{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFuenteFinanciamiento deletes FuenteFinanciamiento by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFuenteFinanciamiento(id int) (err error) {
	o := orm.NewOrm()
	v := FuenteFinanciamiento{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FuenteFinanciamiento{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// AddFuenteFinanciamientoTr insert a new FuenteFinanciamiento into database and returns
// last inserted Id on success.
func AddFuenteFinanciamientoTr(m map[string]interface{}) (res map[string]interface{}, err error) {
	o := orm.NewOrm()
	var FuenteData = FuenteFinanciamiento{}
	var response = make(map[string]interface{})
	o.Begin()
	try.This(func() {
		if errAux := formatdata.FillStruct(m["FuenteFinanciamiento"], &FuenteData); errAux != nil {
			panic(errAux.Error())
		}
		var AfectacionFuenteData = m["AfectacionFuente"].([]interface{})
		if idFuente, errAux := AddFuenteFinanciamiento(&FuenteData); errAux == nil {
			for _, afectacionIntfc := range AfectacionFuenteData {
				afectacion := FuenteFinanciamientoApropiacion{}
				if errAux := formatdata.FillStruct(afectacionIntfc, &afectacion); errAux != nil {
					panic(errAux.Error())
				}
				afectacion.FuenteFinanciamiento.Id = int(idFuente)
				if idAfectacion, errAux := AddFuenteFinanciamientoApropiacion(&afectacion); errAux == nil {
					afectacion.MovimientoFuenteFinanciamientoApropiacion[0].FuenteFinanciamientoApropiacion.Id = int(idAfectacion)
					afectacion.MovimientoFuenteFinanciamientoApropiacion[0].Fecha = time.Now()
					if _, errAux = AddMovimientoFuenteFinanciamientoApropiacion(afectacion.MovimientoFuenteFinanciamientoApropiacion[0]); errAux != nil {
						fmt.Println("Error3: ", errAux.Error())
						errAux = errors.New("error afectacion 2")
						panic(errAux.Error())
					}
				} else {
					fmt.Println("Error2: ", errAux.Error())
					errAux = errors.New("error afectacion 1")
					panic(errAux.Error())
				}

			}
			FuenteData.Id = int(idFuente)
			response["FuenteFinanciamiento"] = FuenteData
			response["AfectacionFuente"] = AfectacionFuenteData
		} else {
			fmt.Println("Error1: ", errAux.Error())
			panic(errAux.Error())
		}
	}).Catch(func(e try.E) {
		fmt.Println("Err ", e)
		o.Rollback()
		err = errors.New("transaction error !")
	})
	o.Commit()
	return response, err
}

// AddMovimientoFuenteFinanciamientoTr insert a new MovimientoFuenteFinanciamientoTr into database and returns
// last inserted Id on success.
func AddMovimientoFuenteFinanciamientoTr(arr []map[string]interface{}) (res interface{}, err error) {
	o := orm.NewOrm()
	o.Begin()
	var afectData []interface{}
	try.This(func() {
		for _, m := range arr {
			afectacion := FuenteFinanciamientoApropiacion{}
			if errAux := formatdata.FillStruct(m, &afectacion); errAux != nil {
				panic(errAux.Error())
			}
			if idAfectacion, errAux := AddFuenteFinanciamientoApropiacion(&afectacion); errAux == nil {
				afectacion.Id = int(idAfectacion)
				afectacion.MovimientoFuenteFinanciamientoApropiacion[0].FuenteFinanciamientoApropiacion.Id = int(idAfectacion)
				afectacion.MovimientoFuenteFinanciamientoApropiacion[0].Fecha = time.Now()
				if _, errAux = AddMovimientoFuenteFinanciamientoApropiacion(afectacion.MovimientoFuenteFinanciamientoApropiacion[0]); errAux != nil {
					fmt.Println("Error3: ", errAux.Error())
					errAux = errors.New("error afectacion 2")
					panic(errAux.Error())
				}
			} else {
				fmt.Println("Error2: ", errAux.Error())
				errAux = errors.New("error afectacion 1")
				panic(errAux.Error())
			}
			afectData = append(afectData, afectacion)
		}
	}).Catch(func(e try.E) {
		fmt.Println("Err ", e)
		o.Rollback()
		err = errors.New("transaction error !")
	})
	o.Commit()
	return afectData, err
}

func DeleteMovimientoFuenteFinanciamientoTr(id int) (err error) {
	o := orm.NewOrm()
	o.Begin()

	var maps []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id as \"Id\"").
		From("" + beego.AppConfig.String("PGschemas") + "fuente_financiamiento_apropiacion").
		Where("fuente_financiamiento = ?")
	if _, err = o.Raw(qb.String(), id).Values(&maps); err != nil {
		o.Rollback()
		return
	}

	for _, data := range maps {
		if idFuenteApr, err := strconv.Atoi(data["Id"].(string)); err == nil {
			// eliminar el dato de movimiento antes de la relacion fuente - apropiacion
			var maps2 []orm.Params
			qb2, _ := orm.NewQueryBuilder("mysql")
			qb2.Select("id as \"Id\"").
				From("" + beego.AppConfig.String("PGschemas") + "movimiento_fuente_financiamiento_apropiacion").
				Where("fuente_financiamiento_apropiacion = ?")
			if _, errAux := o.Raw(qb2.String(), idFuenteApr).Values(&maps2); errAux != nil {
				o.Rollback()
				return errAux
			}
			for _, dataMov := range maps2 {
				if idMovFuenteApr, err := strconv.Atoi(dataMov["Id"].(string)); err == nil {
					if _, err := o.Delete(&MovimientoFuenteFinanciamientoApropiacion{Id: idMovFuenteApr}); err != nil {
						o.Rollback()
						return err
					}
				} else {
					o.Rollback()
					return err
				}

			}
			if _, err := o.Delete(&FuenteFinanciamientoApropiacion{Id: idFuenteApr}); err != nil {
				o.Rollback()
				return err
			}
		} else {
			o.Rollback()
			return err
		}
	}
	o.Commit()
	return

}

func DeleteModificacionFuenteFinanciamiento(data []map[string]interface{}) (err error) {
	//Delete the mofication data ...
	// VARIABLES
	o := orm.NewOrm()
	o.Begin()
	try.This(func() {
		// Eliminar Relacion FUente Apropiacion  Creada en la Modificacion.
		for _, m := range data {
			afectacion := FuenteFinanciamientoApropiacion{}
			if err = formatdata.FillStruct(m, &afectacion); err != nil {
				panic(err.Error())
			}
			var movimientos []MovimientoFuenteFinanciamientoApropiacion
			if err = formatdata.FillStruct(m["MovimientoFuenteFinanciamientoApropiacion"], &movimientos); err != nil {
				panic(err.Error())
			}

			// Eliminar Movimientos realizados en la Modificacion de la Fuente.
			for _, v := range movimientos {
				if _, err = o.Delete(&v); err != nil {
					panic(err.Error())
				}
			}
			if _, err = o.Delete(&afectacion); err != nil {
				panic(err.Error())
			}

		}

	}).Catch(func(e try.E) {
		o.Rollback()
		err = errors.New("Transaction Error")
	})
	o.Commit()
	return err
}
