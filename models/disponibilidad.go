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
	"github.com/fatih/structs"
	"github.com/udistrital/utils_oas/formatdata"
)

type Info_disponibilidad_a_anular struct {
	Anulacion                  AnulacionDisponibilidad
	Disponibilidad_apropiacion []*DisponibilidadApropiacion
	Valor                      float64
}
type Disponibilidad struct {
	Id                           int                             `orm:"auto;column(id);pk"`
	Vigencia                     float64                         `orm:"column(vigencia)"`
	NumeroDisponibilidad         float64                         `orm:"column(numero_disponibilidad);null"`
	Responsable                  int                             `orm:"column(responsable);null"`
	FechaRegistro                time.Time                       `orm:"column(fecha_registro);type(date);null"`
	Estado                       *EstadoDisponibilidad           `orm:"column(estado);rel(fk)"`
	Solicitud                    int                             `orm:"column(solicitud)"`
	DisponibilidadApropiacion    []*DisponibilidadApropiacion    `orm:"reverse(many)"`
	DisponibilidadProcesoExterno []*DisponibilidadProcesoExterno `orm:"reverse(many)"`
}

func (t *Disponibilidad) TableName() string {
	return "disponibilidad"
}

func init() {
	orm.RegisterModel(new(Disponibilidad))
}

// totalDisponibilidades retorna total de disponibilidades por vigencia
func GetTotalDisponibilidades(vigencia int, unidadEjecutora int, finicio string, ffin string) (total int, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	if finicio != "" && ffin != "" {
		qb.Select("COUNT(DISTINCT(disponibilidad))").
			From("financiera.disponibilidad").
			InnerJoin("financiera.disponibilidad_apropiacion").
			On("disponibilidad.id = disponibilidad_apropiacion.disponibilidad").
			InnerJoin("financiera.apropiacion").
			On("apropiacion.id = disponibilidad_apropiacion.apropiacion").
			InnerJoin("financiera.rubro").
			On("rubro.id = apropiacion.rubro").
			Where("disponibilidad.vigencia = ?").
			And("fecha_registro >= ?").
			And("fecha_registro <= ?").
			And("unidad_ejecutora = ?")
		err = o.Raw(qb.String(), vigencia, finicio, ffin, unidadEjecutora).QueryRow(&total)
		return
	}
	qb.Select("COUNT(DISTINCT(disponibilidad))").
		From("financiera.disponibilidad").
		InnerJoin("financiera.disponibilidad_apropiacion").
		On("disponibilidad.id = disponibilidad_apropiacion.disponibilidad").
		InnerJoin("financiera.apropiacion").
		On("apropiacion.id = disponibilidad_apropiacion.apropiacion").
		InnerJoin("financiera.rubro").
		On("rubro.id = apropiacion.rubro").
		Where("disponibilidad.vigencia = ?").
		And("unidad_ejecutora = ?")
	err = o.Raw(qb.String(), vigencia, unidadEjecutora).QueryRow(&total)
	return

}

// AddDisponibilidad insert a new Disponibilidad into database and returns
// last inserted Id on success.
func AddDisponibilidad(m map[string]interface{}) (v Disponibilidad, err error) {
	o := orm.NewOrm()
	o.Begin()
	var consecutivo float64
	var afectacion []DisponibilidadApropiacion
	var procesoExterno DisponibilidadProcesoExterno
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COALESCE(MAX(numero_disponibilidad), 0)+1 as consecutivo").
		From("financiera.disponibilidad").
		InnerJoin("financiera.disponibilidad_apropiacion").
		On("disponibilidad.id = disponibilidad_apropiacion.disponibilidad").
		InnerJoin("financiera.apropiacion").
		On("apropiacion.id = disponibilidad_apropiacion.apropiacion").
		InnerJoin("financiera.rubro").
		On("apropiacion.rubro = rubro.id").
		Where("disponibilidad.vigencia = ?").
		And("rubro.unidad_ejecutora = ?")
	err = o.Raw(qb.String(), int(m["Disponibilidad"].(map[string]interface{})["Vigencia"].(float64)),
		int(m["Disponibilidad"].(map[string]interface{})["UnidadEjecutora"].(float64))).QueryRow(&consecutivo)
	err = formatdata.FillStruct(m["Disponibilidad"], &v)
	if err != nil {
		o.Rollback()
		fmt.Println(m["Disponibilidad"])
		return
	}
	v.NumeroDisponibilidad = consecutivo
	if err != nil {
		o.Rollback()
		return
	}
	_, err = o.Insert(&v)
	if err == nil {
		err = formatdata.FillStruct(m["DisponibilidadProcesoExterno"], &procesoExterno)
		if err == nil {
			procesoExterno.Disponibilidad = &v
			_, err = o.Insert(&procesoExterno)
			if err == nil {
				err = formatdata.FillStruct(m["DisponibilidadApropiacion"], &afectacion)
				if err == nil {
					for _, row := range afectacion {
						row.Disponibilidad = &v
						_, err = o.Insert(&row)
						if err != nil {
							o.Rollback()
							return
						}
					}
				} else {
					o.Rollback()
					return
				}
			} else {
				beego.Info(err)
				o.Rollback()
				return
			}
		} else {
			beego.Info("err dprosext")
			o.Rollback()
			return
		}

	} else {
		o.Rollback()
		return
	}
	o.Commit()
	return
}

// GetDisponibilidadById retrieves Disponibilidad by Id. Returns error if
// Id doesn't exist
func GetDisponibilidadById(id int) (v *Disponibilidad, err error) {
	o := orm.NewOrm()
	v = &Disponibilidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDisponibilidad retrieves all Disponibilidad matches certain condition. Returns empty list if
// no records exist
func GetAllDisponibilidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Disponibilidad))
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

	var l []Disponibilidad
	qs = qs.OrderBy(sortFields...).RelatedSel(5).Distinct()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "DisponibilidadApropiacion", 5)
				o.LoadRelated(&v, "DisponibilidadProcesoExterno", 5)
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

// UpdateDisponibilidad updates Disponibilidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateDisponibilidadById(m *Disponibilidad) (err error) {
	o := orm.NewOrm()
	v := Disponibilidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDisponibilidad deletes Disponibilidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDisponibilidad(id int) (err error) {
	o := orm.NewOrm()
	v := Disponibilidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Disponibilidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func AnulacionTotal(m *Info_disponibilidad_a_anular) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
	m.Anulacion.FechaRegistro = time.Now()
	var consecutivo int
	o.Raw(`SELECT COALESCE(MAX(consecutivo), 0)+1  as consecutivo
						FROM financiera.anulacion_disponibilidad
						JOIN
						financiera.anulacion_disponibilidad_apropiacion as ada
						ON
						ada.anulacion = anulacion_disponibilidad.id
						JOIN
						financiera.disponibilidad_apropiacion
						ON
						disponibilidad_apropiacion.id = ada.disponibilidad_apropiacion
						JOIN
						financiera.disponibilidad
						ON
						disponibilidad.id = disponibilidad_apropiacion.disponibilidad
						WHERE vigencia = ?`, m.Disponibilidad_apropiacion[0].Disponibilidad.Vigencia).QueryRow(&consecutivo)
	m.Anulacion.Consecutivo = consecutivo
	id_anulacion_cdp, err1 := o.Insert(&m.Anulacion)
	fmt.Println("error")
	if err1 != nil {
		alerta[0] = "error"
		alerta = append(alerta, "No se pudo registrar el detalle de la anulacion")
		err = err1
		o.Rollback()
		return
	}
	var acumCdp float64
	acumCdp = 0
	for i := 0; i < len(m.Disponibilidad_apropiacion); i++ {
		var saldoCDP float64
		var err2 error
		if m.Disponibilidad_apropiacion[i].FuenteFinanciamiento != nil {
			saldoCDP, _, _, err2 = SaldoCdp(m.Disponibilidad_apropiacion[i].Disponibilidad.Id, m.Disponibilidad_apropiacion[i].Apropiacion.Id, m.Disponibilidad_apropiacion[i].FuenteFinanciamiento.Id)

		} else {
			saldoCDP, _, _, err2 = SaldoCdp(m.Disponibilidad_apropiacion[i].Disponibilidad.Id, m.Disponibilidad_apropiacion[i].Apropiacion.Id, 0)

		}
		if err2 != nil {
			alerta[0] = "error"
			alerta = append(alerta, "No se pudo cargar el saldo del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo)
			err = err2
			o.Rollback()
			return
		}
		acumCdp = acumCdp + saldoCDP
		if saldoCDP > 0 {
			anulacion_apropiacion := AnulacionDisponibilidadApropiacion{
				DisponibilidadApropiacion: m.Disponibilidad_apropiacion[i],
				Anulacion:                 &AnulacionDisponibilidad{Id: int(id_anulacion_cdp)},
				Valor:                     saldoCDP,
			}
			_, err3 := o.Insert(&anulacion_apropiacion)
			if err3 != nil {
				alerta[0] = "error"
				alerta = append(alerta, "No se pudo registrar la anulacion del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo)
				err = err3
				o.Rollback()
				return
			} else {
				alerta = append(alerta, "Se expidio la solicitud N°"+strconv.Itoa(m.Anulacion.Consecutivo)+" de anulación para el CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo+" la suma de $"+strconv.FormatFloat(saldoCDP, 'f', -1, 64))

			}
		} else {
			alerta[0] = "error"
			alerta = append(alerta, "El CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo+" tiene saldo 0")
			o.Rollback()
			return
		}

	}
	/*if acumCdp > 0 {
		m.Disponibilidad_apropiacion[0].Disponibilidad.Estado = &EstadoDisponibilidad{Id: 3}
		o.Update(m.Disponibilidad_apropiacion[0].Disponibilidad)

	} else {
		o.Rollback()
	}*/
	if m.Anulacion.TipoAnulacion.Id == 3 {
		args := []string{"estado"}
		m.Disponibilidad_apropiacion[0].Disponibilidad.Estado = &EstadoDisponibilidad{Id: 3}
		_, err = o.Update(m.Disponibilidad_apropiacion[0].Disponibilidad, args...)
		if err != nil {
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}
func AprobacionAnulacion(m *AnulacionDisponibilidad) (alert Alert, err error) {
	o := orm.NewOrm()
	o.Begin()
	args := []string{"estado_anulacion", "solicitante", "responsable"}
	_, err = o.Update(m, args...)
	if err != nil {
		o.Rollback()
		alertdb := structs.Map(err)
		var code string
		formatdata.FillStruct(alertdb["Code"], &code)
		alert = Alert{Type: "error", Code: "E_" + code, Body: err}
		return
	}
	var acumCDP float64
	acumCDP = 0

	for i := 0; i < len(m.AnulacionDisponibilidadApropiacion); i++ {
		var saldoCDP float64
		if m.AnulacionDisponibilidadApropiacion[i].DisponibilidadApropiacion.FuenteFinanciamiento != nil {
			saldoCDP, _, _, err = SaldoCdp(m.AnulacionDisponibilidadApropiacion[i].DisponibilidadApropiacion.Disponibilidad.Id, m.AnulacionDisponibilidadApropiacion[i].DisponibilidadApropiacion.Apropiacion.Id, m.AnulacionDisponibilidadApropiacion[i].DisponibilidadApropiacion.FuenteFinanciamiento.Id)
		} else {
			saldoCDP, _, _, err = SaldoCdp(m.AnulacionDisponibilidadApropiacion[i].DisponibilidadApropiacion.Disponibilidad.Id, m.AnulacionDisponibilidadApropiacion[i].DisponibilidadApropiacion.Apropiacion.Id, 0)

		}
		if saldoCDP < m.AnulacionDisponibilidadApropiacion[i].Valor {
			o.Rollback()
			alert = Alert{Type: "error", Code: "E_A12", Body: m}
			return
		}
		if err != nil {
			o.Rollback()
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert = Alert{Type: "error", Code: "E_" + code, Body: err}
			return
		}
		acumCDP = acumCDP + saldoCDP - m.AnulacionDisponibilidadApropiacion[i].Valor
		fmt.Println("acum: ", acumCDP)
	}
	if acumCDP == 0 && m.EstadoAnulacion.Id == 3 {
		m.AnulacionDisponibilidadApropiacion[0].DisponibilidadApropiacion.Disponibilidad.Estado = &EstadoDisponibilidad{Id: 3}
		o.Update(m.AnulacionDisponibilidadApropiacion[0].DisponibilidadApropiacion.Disponibilidad)
	}

	if err != nil {
		o.Rollback()
		alertdb := structs.Map(err)
		var code string
		formatdata.FillStruct(alertdb["Code"], &code)
		alert = Alert{Type: "error", Code: "E_" + code, Body: err}
		return
	}
	o.Commit()
	alert = Alert{Type: "success", Code: "S_A12", Body: m}
	return
}
func AnulacionParcial(m *Info_disponibilidad_a_anular) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
	m.Anulacion.FechaRegistro = time.Now()
	var consecutivo int
	o.Raw(`SELECT COALESCE(MAX(consecutivo), 0)+1  as consecutivo
						FROM financiera.anulacion_disponibilidad
						JOIN
						financiera.anulacion_disponibilidad_apropiacion as ada
						ON
						ada.anulacion = anulacion_disponibilidad.id
						JOIN
						financiera.disponibilidad_apropiacion
						ON
						disponibilidad_apropiacion.id = ada.disponibilidad_apropiacion
						JOIN
						financiera.disponibilidad
						ON
						disponibilidad.id = disponibilidad_apropiacion.disponibilidad
						WHERE vigencia = ?`, m.Disponibilidad_apropiacion[0].Disponibilidad.Vigencia).QueryRow(&consecutivo)
	m.Anulacion.Consecutivo = consecutivo
	id_anulacion_cdp, err1 := o.Insert(&m.Anulacion)
	if err1 != nil {
		alerta[0] = "error"
		alerta = append(alerta, "No se pudo registrar el detalle de la anulacion")
		err = err1
		o.Rollback()
		return
	}
	for i := 0; i < len(m.Disponibilidad_apropiacion); i++ {
		var saldoCDP float64
		var err2 error
		if m.Disponibilidad_apropiacion[i].FuenteFinanciamiento != nil {
			saldoCDP, _, _, err2 = SaldoCdp(m.Disponibilidad_apropiacion[i].Disponibilidad.Id, m.Disponibilidad_apropiacion[i].Apropiacion.Id, m.Disponibilidad_apropiacion[i].FuenteFinanciamiento.Id)

		} else {
			saldoCDP, _, _, err2 = SaldoCdp(m.Disponibilidad_apropiacion[i].Disponibilidad.Id, m.Disponibilidad_apropiacion[i].Apropiacion.Id, 0)

		}
		if err2 != nil {
			alerta[0] = "error"
			alerta = append(alerta, "No se pudo cargar el saldo del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo)
			err = err2
			o.Rollback()
			return
		}
		fmt.Println("saldo: ", saldoCDP)
		if saldoCDP < m.Valor {
			alerta[0] = "error"
			alerta = append(alerta, "Valor a anular supera el saldo del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo)
			o.Rollback()
			return
		} else {
			anulacion_apropiacion := AnulacionDisponibilidadApropiacion{
				DisponibilidadApropiacion: m.Disponibilidad_apropiacion[i],
				Anulacion:                 &AnulacionDisponibilidad{Id: int(id_anulacion_cdp)},
				Valor:                     m.Valor,
			}
			_, err3 := o.Insert(&anulacion_apropiacion)
			if err3 != nil {
				alerta[0] = "error"
				alerta = append(alerta, "No se pudo registrar la anulación del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo)
				err = err3
				o.Rollback()
				return
			} else {
				alerta = append(alerta, "Se expidio la solicitud N°"+strconv.Itoa(m.Anulacion.Consecutivo)+" de anulación para el CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo+" la suma de $"+strconv.FormatFloat(m.Valor, 'f', -1, 64))

			}
		}

	}
	o.Commit()
	/*var acumCDP float64
	acumCDP = 0

	for i := 0; i < len(m.Disponibilidad_apropiacion); i++ {
		var saldoCDP float64
		if m.Disponibilidad_apropiacion[i].FuenteFinanciamiento != nil {
			saldoCDP, err = GetValorActualCDP(m.Disponibilidad_apropiacion[i].Disponibilidad.Id)
		} else {
			saldoCDP, err = GetValorActualCDP(m.Disponibilidad_apropiacion[i].Disponibilidad.Id)

		}
		if err != nil {
			o.Rollback()
			alerta[0] = "error"
			alerta = append(alerta, "No se pudo registrar la anulacion del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo)
			fmt.Println("alerta ", alerta)
			return
		}
		acumCDP = acumCDP + saldoCDP
	}
	if acumCDP == 0 {
		m.Disponibilidad_apropiacion[0].Disponibilidad.Estado = &EstadoDisponibilidad{Id: 3}
		o.Update(m.Disponibilidad_apropiacion[0].Disponibilidad)
	}*/
	return
}

//----------------------------------------
//funcion para obtener saldo restante del cdp
func SaldoCdp(id_cdp int, id_apropiacion int, id_fuente int) (saldo float64, comprometido float64, anulado float64, err error) {
	/*o := orm.NewOrm()
	var maps []orm.Params
	o.Raw(`SELECT * FROM financiera.saldo_cdp WHERE id = ? AND apropiacion = ? `, id_cdp, id_apropiacion).Values(&maps)
	fmt.Println("maps: ", maps)
	if maps[0]["valor"] == nil {
		valor = 0
	} else {
		valor, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}*/
	valorcdp, err := ValorCdp(id_cdp, id_apropiacion, id_fuente)
	comprometidocdp, err := ComprometidoCdp(id_cdp, id_apropiacion, id_fuente)
	anuladocdp, err := AnuladoCdp(id_cdp, id_apropiacion, id_fuente)
	anuladorp, err := AnuladoRpPorCDP(id_cdp, id_apropiacion, id_fuente)
	comprometido = comprometidocdp - anuladorp
	anulado = anuladocdp
	saldo = valorcdp + anuladorp - anuladocdp - comprometidocdp

	return
}

//valor original del CDP.
func ValorCdp(id_cdp int, id_apropiacion int, id_fuente int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw(`SELECT * FROM (SELECT disponibilidad.id,
			            disponibilidad_apropiacion.apropiacion,
			            COALESCE(disponibilidad_apropiacion.fuente_financiamiento, 0) as fuente_financiamiento,
			            COALESCE(sum(disponibilidad_apropiacion.valor),0) AS valor
			           FROM financiera.disponibilidad
			             JOIN financiera.disponibilidad_apropiacion ON disponibilidad_apropiacion.disponibilidad = disponibilidad.id
			          GROUP BY disponibilidad.id, disponibilidad_apropiacion.apropiacion,disponibilidad_apropiacion.fuente_financiamiento) as saldo
								WHERE id = ? AND apropiacion= ? AND fuente_financiamiento = ?;`, id_cdp, id_apropiacion, id_fuente).Values(&maps)
	if maps == nil {
		valor = 0
	} else {
		valor, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}

	return
}

//saldo comprometido del cdp.
func ComprometidoCdp(id_cdp int, id_apropiacion int, id_fuente int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw(`SELECT * FROM (SELECT disponibilidad.id,
				            disponibilidad_apropiacion.apropiacion,
				            COALESCE(disponibilidad_apropiacion.fuente_financiamiento, 0) as fuente_financiamiento,
				            COALESCE(sum(registro_presupuestal_disponibilidad_apropiacion.valor),0) AS valor
				           FROM financiera.disponibilidad
				             JOIN financiera.disponibilidad_apropiacion ON disponibilidad_apropiacion.disponibilidad = disponibilidad.id
				             JOIN financiera.registro_presupuestal_disponibilidad_apropiacion ON registro_presupuestal_disponibilidad_apropiacion.disponibilidad_apropiacion = disponibilidad_apropiacion.id
				          GROUP BY disponibilidad.id, disponibilidad_apropiacion.apropiacion, disponibilidad_apropiacion.fuente_financiamiento) as saldo
									WHERE id = ? AND apropiacion=? AND fuente_financiamiento = ?;`, id_cdp, id_apropiacion, id_fuente).Values(&maps)
	if maps == nil {
		valor = 0
	} else {
		valor, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}

	return
}

//valor anulaciones del cdp.
func AnuladoCdp(id_cdp int, id_apropiacion int, id_fuente int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw(`SELECT * FROM(SELECT disponibilidad.id,
											anulacion_disponibilidad.estado_anulacion,
					            disponibilidad_apropiacion.apropiacion,
					            COALESCE(disponibilidad_apropiacion.fuente_financiamiento,0) as fuente_financiamiento,
					            COALESCE(sum(anulacion_disponibilidad_apropiacion.valor),0) AS valor
					           FROM financiera.anulacion_disponibilidad_apropiacion
					             JOIN financiera.disponibilidad_apropiacion ON anulacion_disponibilidad_apropiacion.disponibilidad_apropiacion = disponibilidad_apropiacion.id
					             JOIN financiera.disponibilidad ON disponibilidad_apropiacion.disponibilidad = disponibilidad.id
											 JOIN financiera.anulacion_disponibilidad ON anulacion_disponibilidad.id = anulacion_disponibilidad_apropiacion.anulacion
					          GROUP BY disponibilidad.id, anulacion_disponibilidad.estado_anulacion, disponibilidad_apropiacion.apropiacion,disponibilidad_apropiacion.fuente_financiamiento) as saldo
										WHERE id = ? AND apropiacion = ? AND fuente_financiamiento = ? AND estado_anulacion = 3`, id_cdp, id_apropiacion, id_fuente).Values(&maps)
	if maps == nil {
		valor = 0
	} else {
		valor, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}

	return
}

func AnuladoRpPorCDP(id_disponibilidad int, id_apropiacion int, id_fuente int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw(`SELECT * FROM(SELECT disponibilidad.id,
					            disponibilidad_apropiacion.apropiacion,
					            COALESCE(disponibilidad_apropiacion.fuente_financiamiento) as fuente_financiamiento,
					            COALESCE(sum(anulacion_registro_presupuestal_disponibilidad_apropiacion.valor),0) AS valor
					           FROM financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion
					             JOIN financiera.registro_presupuestal_disponibilidad_apropiacion ON anulacion_registro_presupuestal_disponibilidad_apropiacion.registro_presupuestal_disponibilidad_apropiacion = registro_presupuestal_disponibilidad_apropiacion.id
					             JOIN financiera.disponibilidad_apropiacion ON registro_presupuestal_disponibilidad_apropiacion.disponibilidad_apropiacion = disponibilidad_apropiacion.id
					             JOIN financiera.disponibilidad ON disponibilidad_apropiacion.disponibilidad = disponibilidad.id
					          GROUP BY disponibilidad.id, disponibilidad_apropiacion.apropiacion, disponibilidad_apropiacion.fuente_financiamiento) as saldo
										WHERE id = ? AND apropiacion = ? AND fuente_financiamiento = ?;`, id_disponibilidad, id_apropiacion, id_fuente).Values(&maps)
	fmt.Println("maps: ", maps)
	if maps == nil {
		valor = 0
	} else {
		valor, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}

	return
}

//----------------------------------------

//funcion GetValorTotalRp
func GetValorTotalCDP(cdp_id int) (total float64, err error) {
	o := orm.NewOrm()
	var totalSql float64
	err = o.Raw("select sum(valor) from financiera.disponibilidad_apropiacion where disponibilidad = ?", cdp_id).QueryRow(&totalSql)
	if err == nil {
		fmt.Println("total val: ", totalSql)
		return totalSql, nil
	}
	fmt.Println("total comp: ", err)
	return 0, nil
}
func GetValorTotalComprometidoRpPorCDP(cdp_id int) (total float64, err error) {
	o := orm.NewOrm()
	var totalSql float64
	err = o.Raw(`SELECT valor FROM (SELECT disponibilidad.id,
				            disponibilidad_apropiacion.apropiacion,
				            COALESCE(disponibilidad_apropiacion.fuente_financiamiento, 0) as fuente_financiamiento,
				            COALESCE(sum(registro_presupuestal_disponibilidad_apropiacion.valor),0) AS valor
				           FROM financiera.disponibilidad
				             JOIN financiera.disponibilidad_apropiacion ON disponibilidad_apropiacion.disponibilidad = disponibilidad.id
				             JOIN financiera.registro_presupuestal_disponibilidad_apropiacion ON registro_presupuestal_disponibilidad_apropiacion.disponibilidad_apropiacion = disponibilidad_apropiacion.id
				          GROUP BY disponibilidad.id, disponibilidad_apropiacion.apropiacion, disponibilidad_apropiacion.fuente_financiamiento) as saldo
									WHERE id = ? `, cdp_id).QueryRow(&totalSql)
	if err == nil {
		fmt.Println("total comp: ", totalSql)
		return totalSql, nil
	}
	fmt.Println("total comp: ", err)
	return 0, nil

}

func GetValorTotalAnuladoRpPorCDP(cdp_id int) (total float64, err error) {
	o := orm.NewOrm()
	var totalSql float64
	err = o.Raw(`SELECT valor FROM (SELECT disponibilidad.id,
				            disponibilidad_apropiacion.apropiacion,
				            COALESCE(disponibilidad_apropiacion.fuente_financiamiento, 0) as fuente_financiamiento,
				            COALESCE(sum(registro_presupuestal_disponibilidad_apropiacion.valor),0) AS valor
				           FROM financiera.disponibilidad
				             JOIN financiera.disponibilidad_apropiacion ON disponibilidad_apropiacion.disponibilidad = disponibilidad.id
				             JOIN financiera.registro_presupuestal_disponibilidad_apropiacion ON registro_presupuestal_disponibilidad_apropiacion.disponibilidad_apropiacion = disponibilidad_apropiacion.id
				          GROUP BY disponibilidad.id, disponibilidad_apropiacion.apropiacion, disponibilidad_apropiacion.fuente_financiamiento) as saldo
									WHERE id = ? `, cdp_id).QueryRow(&totalSql)
	if err == nil {
		fmt.Println("total a rp: ", totalSql)
		return totalSql, nil
	}
	fmt.Println("total a rp: ", err)
	return 0, nil

}

func GetValorTotalAnuladoCDP(cdp_id int) (total float64, err error) {
	o := orm.NewOrm()
	var totalSql float64
	err = o.Raw(`SELECT valor FROM(SELECT disponibilidad.id,
					            disponibilidad_apropiacion.apropiacion,
					            COALESCE(disponibilidad_apropiacion.fuente_financiamiento,0) as fuente_financiamiento,
					            COALESCE(sum(anulacion_disponibilidad_apropiacion.valor),0) AS valor
					           FROM financiera.anulacion_disponibilidad_apropiacion
					             JOIN financiera.disponibilidad_apropiacion ON anulacion_disponibilidad_apropiacion.disponibilidad_apropiacion = disponibilidad_apropiacion.id
					             JOIN financiera.disponibilidad ON disponibilidad_apropiacion.disponibilidad = disponibilidad.id
					          GROUP BY disponibilidad.id, disponibilidad_apropiacion.apropiacion,disponibilidad_apropiacion.fuente_financiamiento) as saldo
										WHERE id = ?`, cdp_id).QueryRow(&totalSql)
	if err == nil {
		fmt.Println("++++++++++++++++++++++total A: ", totalSql)
		return totalSql, nil
	}
	fmt.Println("total A: ", err)
	return 0, nil

}

func GetValorActualCDP(cdp_id int) (total float64, err error) {
	valor, err := GetValorTotalCDP(cdp_id)
	comprometido, err := GetValorTotalComprometidoRpPorCDP(cdp_id)
	anulado, err := GetValorTotalAnuladoCDP(cdp_id)
	anulado_rp, err := GetValorTotalAnuladoRpPorCDP(cdp_id)
	total = valor - comprometido - anulado - anulado_rp
	return
}

//GetPrincDisponibilidadInfo... Obtiene la informacion principal de una disponibilidad
//afectacion
func GetPrincDisponibilidadInfo(id int) (interface{}, error) {
	o := orm.NewOrm()
	var maps []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("apropiacion as \"Apropiacion\", disponibilidad_apropiacion.valor as \"Valor\", rubro.codigo as \"Rubro\" , unidad_ejecutora as \"UnidadEjecutora\", fuente_financiamiento.codigo as \"FuenteCodigo\",  fuente_financiamiento.nombre as \"FuenteNombre\"").
		From("financiera.disponibilidad_apropiacion").
		InnerJoin("financiera.apropiacion").
		On("apropiacion.Id = disponibilidad_apropiacion.apropiacion").
		InnerJoin("financiera.rubro").
		On("rubro.id = apropiacion.rubro").
		LeftJoin("financiera.fuente_financiamiento").
		On("disponibilidad_apropiacion.fuente_financiamiento = fuente_financiamiento.id").
		Where("disponibilidad = ?")

	_, err := o.Raw(qb.String(), id).Values(&maps)
	maps[0]["Valor"], err = strconv.ParseFloat(maps[0]["Valor"].(string), 64)

	return maps, err
}

//DeleteDisponibilidadData... Elimina la disponibilidad dado su id
// y todos los datos que esta representa.
func DeleteDisponibilidadData(id int) (err error) {
	o := orm.NewOrm()
	o.Begin()

	var maps []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id as \"Id\"").
		From("financiera.disponibilidad_proceso_externo").
		Where("disponibilidad = ?")
	if _, err = o.Raw(qb.String(), id).Values(&maps); err != nil {
		o.Rollback()
		return
	}

	for _, data := range maps {
		if idDispExt, err := strconv.Atoi(data["Id"].(string)); err == nil {
			if _, err := o.Delete(&DisponibilidadProcesoExterno{Id: idDispExt}); err != nil {
				o.Rollback()
				return err
			}
		} else {
			o.Rollback()
			return err
		}

	}

	var dispApr []orm.Params
	qb, _ = orm.NewQueryBuilder("mysql")
	qb.Select("id as \"Id\"").
		From("financiera.disponibilidad_apropiacion").
		Where("disponibilidad = ?")
	if _, err = o.Raw(qb.String(), id).Values(&dispApr); err != nil {
		o.Rollback()
		return
	}

	for _, data := range dispApr {
		if idDispExt, err := strconv.Atoi(data["Id"].(string)); err == nil {
			if _, err := o.Delete(&DisponibilidadApropiacion{Id: idDispExt}); err != nil {
				o.Rollback()
				return err
			}

		} else {
			o.Rollback()
			return err
		}

	}

	if _, err = o.Delete(&Disponibilidad{Id: id}); err != nil {
		o.Rollback()
		return
	}

	o.Commit()
	return
}

//DeleteDisponibilidadMovimiento... Elimina la disponibilidad dado su id
// y todos los datos que esta representa.
func DeleteDisponibilidadMovimiento(id int) (err error) {
	o := orm.NewOrm()
	o.Begin()

	var maps []orm.Params
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id as \"Id\"").
		From("financiera.disponibilidad_proceso_externo").
		Where("disponibilidad = ?")
	if _, err = o.Raw(qb.String(), id).Values(&maps); err != nil {
		o.Rollback()
		return
	}

	for _, data := range maps {
		if idDispExt, err := strconv.Atoi(data["Id"].(string)); err == nil {
			if _, err := o.Delete(&DisponibilidadProcesoExterno{Id: idDispExt}); err != nil {
				o.Rollback()
				return err
			}
		} else {
			o.Rollback()
			return err
		}

	}

	var dispApr []orm.Params
	qb, _ = orm.NewQueryBuilder("mysql")
	qb.Select("id as \"Id\"").
		From("financiera.disponibilidad_apropiacion").
		Where("disponibilidad = ?")
	if _, err = o.Raw(qb.String(), id).Values(&dispApr); err != nil {
		o.Rollback()
		return
	}

	for _, data := range dispApr {
		if idDispExt, err := strconv.Atoi(data["Id"].(string)); err == nil {
			if _, err := o.Delete(&DisponibilidadApropiacion{Id: idDispExt}); err != nil {
				o.Rollback()
				return err
			}

		} else {
			o.Rollback()
			return err
		}

	}

	var dispMov []orm.Params
	qb, _ = orm.NewQueryBuilder("mysql")
	qb.Select("id as \"Id\"").
		From("financiera.movimiento_apropiacion_disponibilidad_apropiacion").
		Where("disponibilidad = ?")
	if _, err = o.Raw(qb.String(), id).Values(&dispMov); err != nil {
		o.Rollback()
		return
	}

	for _, data := range dispMov {
		if idDispExt, err := strconv.Atoi(data["Id"].(string)); err == nil {
			if _, err := o.Update(&MovimientoApropiacionDisponibilidadApropiacion{Id: idDispExt, Disponibilidad: nil}, "Disponibilidad"); err != nil {
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
