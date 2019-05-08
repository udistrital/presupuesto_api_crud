package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type Usuario struct {
	Id int
}

type OrdenPago struct {
	Id               int               `orm:"column(id);pk;auto"`
	Vigencia         float64           `orm:"column(vigencia)"`
	ValorBase        float64           `orm:"column(valor_base)"`
	Convenio         int               `orm:"column(convenio);null"`
	SubTipoOrdenPago *SubTipoOrdenPago `orm:"column(sub_tipo_orden_pago);rel(fk)"`
	UnidadEjecutora  *UnidadEjecutora  `orm:"column(unidad_ejecutora);rel(fk)"`
	Liquidacion      int               `orm:"column(liquidacion);null"`
	EntradaAlmacen   int               `orm:"column(entrada_almacen);null"`
	Consecutivo      int               `orm:"column(consecutivo)"`
	Documento        int               `orm:"column(documento)"`
	FormaPago        *FormaPago        `orm:"column(forma_pago);rel(fk);null"`

	OrdenPagoEstadoOrdenPago      []*OrdenPagoEstadoOrdenPago      `orm:"reverse(many)"`
	OrdenPagoRegistroPresupuestal []*OrdenPagoRegistroPresupuestal `orm:"reverse(many)"`
	OrdenPagoCuentaEspecial       []*OrdenPagoCuentaEspecial       `orm:"reverse(many)"`
}

func (t *OrdenPago) TableName() string {
	return "orden_pago"
}

func init() {
	orm.RegisterModel(new(OrdenPago))
}

// AddOrdenPago insert a new OrdenPago into database and returns
// last inserted Id on success.
func AddOrdenPago(m *OrdenPago) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrdenPagoById retrieves OrdenPago by Id. Returns error if
// Id doesn't exist
func GetOrdenPagoById(id int) (v *OrdenPago, err error) {
	o := orm.NewOrm()
	v = &OrdenPago{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrdenPago retrieves all OrdenPago matches certain condition. Returns empty list if
// no records exist
func GetAllOrdenPago(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrdenPago))
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

	var l []OrdenPago
	qs = qs.OrderBy(sortFields...).Filter("OrdenPagoRegistroPresupuestal__isnull", false).RelatedSel(5).Distinct()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "OrdenPagoEstadoOrdenPago", 5, 1, 0, "-Id")
				o.LoadRelated(&v, "OrdenPagoRegistroPresupuestal", 2, -1, 0, "-Id")
				o.LoadRelated(&v, "OrdenPagoCuentaEspecial", 2, -1, 0, "-Id")
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

// UpdateOrdenPago updates OrdenPago by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrdenPagoById(m *OrdenPago) (err error) {
	o := orm.NewOrm()
	v := OrdenPago{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrdenPago deletes OrdenPago by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrdenPago(id int) (err error) {
	o := orm.NewOrm()
	v := OrdenPago{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrdenPago{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func ConsecutivoOrdenPago(grupoSecuencia string) (StringConsecutivo string, grSecuencia string, outputError map[string]interface{}) {
	if grupoSecuencia != "" {
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("COALESCE(MAX(consecutivo), 0)+1 as consecutivo").
			From("financiera.orden_pago as op").
			InnerJoin("financiera.sub_tipo_orden_pago as sub").On("sub.id = op.sub_tipo_orden_pago").
			And("sub.grupo_secuencia = ?")
		StringConsecutivo, grSecuencia := qb.String(), grupoSecuencia
		return StringConsecutivo, grSecuencia, nil
	}
	outputError = map[string]interface{}{"Code": "E_0458", "Body": "Not enough parameter in ConsecutivoOrdenPago", "Type": "error"}
	return "", "", outputError
}

// personalizado Registrar orden_pago, concepto_ordenpago y transacciones
func RegistrarOpProveedor(DataOpProveedor map[string]interface{}) (alerta Alert) {
	var idOrdenPago int64
	var sqlSecuencia string
	var grupoSecuencia string
	var controlErro map[string]interface{}
	var auxCuentaEspecial []map[string]interface{}
	var auxTipoDocumento TipoDocumentoAfectante
	var consecutivoOp int
	var err error
	o := orm.NewOrm()
	o.Begin()
	// GetData
	ordenPago := OrdenPago{}
	conceptoOrdenPago := []ConceptoOrdenPago{}
	movimientoContable := []MovimientoContable{}
	usuario := Usuario{}
	ordenPagoRegistroPresupuestal := []OrdenPagoRegistroPresupuestal{}
	ordenPagoCuentaEspecial := []OrdenPagoCuentaEspecial{}
	err = formatdata.FillStruct(DataOpProveedor["OrdenPago"], &ordenPago)
	err = formatdata.FillStruct(DataOpProveedor["ConceptoOrdenPago"], &conceptoOrdenPago)
	err = formatdata.FillStruct(DataOpProveedor["MovimientoContable"], &movimientoContable)
	err = formatdata.FillStruct(DataOpProveedor["Usuario"], &usuario)
	err = formatdata.FillStruct(DataOpProveedor["RegistroPresupuestal"], &ordenPagoRegistroPresupuestal)
	err = formatdata.FillStruct(DataOpProveedor["MovimientoContable"], &ordenPagoCuentaEspecial)
	err = formatdata.FillStruct(DataOpProveedor["MovimientoContable"], &auxCuentaEspecial)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPP_01" //error en parametros de entrada
		alerta.Body = "Erro en la estructura de parametro de entrada en RegistrarOpProveedor"
		o.Rollback()
		return
	}

	// Consecutivo
	if sqlSecuencia, grupoSecuencia, controlErro = ConsecutivoOrdenPago(ordenPago.SubTipoOrdenPago.GrupoSecuencia); controlErro != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPP_01"
		alerta.Body = controlErro["Body"]
		o.Rollback()
		return
	}
	o.Raw(sqlSecuencia, grupoSecuencia).QueryRow(&consecutivoOp)
	ordenPago.Consecutivo = consecutivoOp
	// Estado OP
	estadoOpObj := EstadoOrdenPago{CodigoAbreviacion: "EOP_01"}
	err = o.Read(&estadoOpObj, "CodigoAbreviacion")
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPP_01" //en busqueda de estado
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	// Registrar OP
	idOrdenPago, err = o.Insert(&ordenPago)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPP_01"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	// Registrar estado OP
	newEstadoOp := OrdenPagoEstadoOrdenPago{}
	newEstadoOp.OrdenPago = &OrdenPago{Id: int(idOrdenPago)}
	newEstadoOp.EstadoOrdenPago = &EstadoOrdenPago{Id: int(estadoOpObj.Id)}
	newEstadoOp.FechaRegistro = time.Now()
	newEstadoOp.Usuario = usuario.Id

	_, err = o.Insert(&newEstadoOp)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPP_01"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}

	// Data MovimientoContable
	estadoMovimientoContable := EstadoMovimientoContable{CodigoAbreviacion: "AF"} //Afectado
	err = o.Read(&estadoMovimientoContable, "CodigoAbreviacion")
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPN_02"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	// TipoDocumentoAfectante
	tipoDocumentoAfectante, err := GetTipoDocumentoAfectanteByCode("DA-OP")     //documento Orden Pago
	tipoDocumentoAfectanteGiro, err := GetTipoDocumentoAfectanteByCode("DA-GI") //documento Orden Pago
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPN_02"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}

	//== Insertar data Conceptos
	for i := 0; i < len(conceptoOrdenPago); i++ {
		conceptoOrdenPago[i].OrdenDePago = &OrdenPago{Id: int(idOrdenPago)}
		_, err = o.Insert(&conceptoOrdenPago[i])
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_02"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
	}
	for i := 0; i < len(ordenPagoRegistroPresupuestal); i++ {
		ordenPagoRegistroPresupuestalData := OrdenPagoRegistroPresupuestal{
			OrdenPago:            &OrdenPago{Id: int(idOrdenPago)},
			RegistroPresupuestal: &RegistroPresupuestal{Id: int(ordenPagoRegistroPresupuestal[i].Id)},
			FechaRegistro:        time.Now(),
		}
		_, err = o.Insert(&ordenPagoRegistroPresupuestalData)
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_01"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
	}
	// Insertar data Movimientos Contables
	for i := 0; i < len(movimientoContable); i++ {
		movimientoContableData := MovimientoContable{
			Debito:                   movimientoContable[i].Debito,
			Credito:                  movimientoContable[i].Credito,
			Fecha:                    time.Now(),
			Concepto:                 movimientoContable[i].Concepto,
			CuentaContable:           movimientoContable[i].CuentaContable,
			TipoDocumentoAfectante:   &TipoDocumentoAfectante{Id: int(tipoDocumentoAfectante.Id)}, //documento afectante tipo op
			CodigoDocumentoAfectante: int(idOrdenPago),
			EstadoMovimientoContable: &EstadoMovimientoContable{Id: int(estadoMovimientoContable.Id)},
		}
		if movimientoContable[i].CuentaEspecial != nil {
			movimientoContableData.CuentaEspecial = movimientoContable[i].CuentaEspecial
			ordenPagoCuentaEspecialData := OrdenPagoCuentaEspecial{
				OrdenPago:      &OrdenPago{Id: int(idOrdenPago)},
				CuentaEspecial: &CuentaEspecial{Id: int(movimientoContable[i].CuentaEspecial.Id)},
				FormaPago:      &FormaPago{Id: int(ordenPagoCuentaEspecial[i].FormaPago.Id)},
				ValorBase:      ordenPagoCuentaEspecial[i].ValorBase,
				FechaRegistro:  time.Now(),
				Usuario:        usuario.Id,
			}
			_, err = o.Insert(&ordenPagoCuentaEspecialData)
			if err != nil {
				alerta.Type = "errorCuentaEspecial"
				alerta.Code = "E_OPP_04"
				alerta.Body = err.Error()
				o.Rollback()
				return
			}
		}
		_, err = o.Insert(&movimientoContableData)
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_03"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
		if auxCuentaEspecial[i]["TipoCuentaEspecial"] != nil {
			if auxCuentaEspecial[i]["TipoCuentaEspecial"].(map[string]interface{})["Id"].(float64) == 2 {
				auxTipoDocumento = tipoDocumentoAfectanteGiro
				movimientoContableData := MovimientoContable{
					Debito:                   movimientoContable[i].Debito,
					Credito:                  movimientoContable[i].Credito,
					Fecha:                    time.Now(),
					Concepto:                 movimientoContable[i].Concepto,
					CuentaContable:           movimientoContable[i].CuentaContable,
					TipoDocumentoAfectante:   &TipoDocumentoAfectante{Id: int(auxTipoDocumento.Id)}, //documento afectante tipo op
					CodigoDocumentoAfectante: int(idOrdenPago),
					EstadoMovimientoContable: &EstadoMovimientoContable{Id: int(estadoMovimientoContable.Id)},
				}
				_, err = o.Insert(&movimientoContableData)
				if err != nil {
					alerta.Type = "error"
					alerta.Code = "E_OPP_03"
					alerta.Body = err.Error()
					o.Rollback()
					return
				}

			}
		}

	}
	alerta = Alert{Type: "success", Code: "S_OPP_01", Body: consecutivoOp}
	o.Commit()
	return
}

// personalizado Actualiza orden_pago, concepto_ordenpago y movimeintos contalbes
func ActualizarOpProveedor(DataActualizarOpProveedor map[string]interface{}) (alerta Alert, err error, consecutivoOp int) {
	o := orm.NewOrm()
	o.Begin()
	// GetData
	var auxCuentaEspecial []map[string]interface{}
	var auxTipoDocumento TipoDocumentoAfectante
	ordenPago := OrdenPago{}
	conceptoOrdenPago := []ConceptoOrdenPago{}
	movimientoContable := []MovimientoContable{}
	usuario := Usuario{}

	err = formatdata.FillStruct(DataActualizarOpProveedor["OrdenPago"], &ordenPago)
	err = formatdata.FillStruct(DataActualizarOpProveedor["ConceptoOrdenPago"], &conceptoOrdenPago)
	err = formatdata.FillStruct(DataActualizarOpProveedor["MovimientoContable"], &movimientoContable)
	err = formatdata.FillStruct(DataActualizarOpProveedor["Usuario"], &usuario)
	err = formatdata.FillStruct(DataActualizarOpProveedor["MovimientoContable"], &auxCuentaEspecial)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPP_UPD_01" //error en parametros de entrada
		alerta.Body = err.Error()
		o.Rollback()
		return
	}

	// Actualizar datos de la Orden
	orden := OrdenPago{Id: ordenPago.Id}
	if o.Read(&orden) == nil {
		//orden.Iva = ordenPago.Iva
		orden.Documento = ordenPago.Documento
		orden.SubTipoOrdenPago = ordenPago.SubTipoOrdenPago
		orden.FormaPago = ordenPago.FormaPago
		orden.ValorBase = ordenPago.ValorBase
		if _, err = o.Update(&orden); err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_UPD_01"
			alerta.Body = err.Error()
			o.Rollback()
			return
		} else {
			consecutivoOp = orden.Consecutivo
		}
	}

	// Data MovimientoContable
	estadoMovimientoContable := EstadoMovimientoContable{CodigoAbreviacion: "AF"} //Afectado
	err = o.Read(&estadoMovimientoContable, "CodigoAbreviacion")
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPN_02"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	// TipoDocumentoAfectante
	tipoDocumentoAfectante, err := GetTipoDocumentoAfectanteByCode("DA-OP")     //documento Orden Pago
	tipoDocumentoAfectanteGiro, err := GetTipoDocumentoAfectanteByCode("DA-GI") //documento Giro Impuestos
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPN_02"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}

	// Eliminar Conceptos Orden de Pagos y Movimientos contables
	if len(conceptoOrdenPago) > 0 {
		_, err = o.Raw("DELETE FROM financiera.concepto_orden_pago where orden_de_pago = ?", ordenPago.Id).Exec()
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_UPD_02"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
	}
	if len(movimientoContable) > 0 {
		_, err = o.Raw("DELETE FROM financiera.movimiento_contable where codigo_documento_afectante = ?", ordenPago.Id).Exec()
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_UPD_03"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
	}
	// Insertar Nueva Data Conceptos Orden de Pagos y Movimientos contables
	//Conceptos
	for i := 0; i < len(conceptoOrdenPago); i++ {
		conceptoOrdenPago[i].OrdenDePago = &OrdenPago{Id: int(ordenPago.Id)}
		_, err = o.Insert(&conceptoOrdenPago[i])
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_UPD_04"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
	}
	//Movimientos
	for i := 0; i < len(movimientoContable); i++ {
		movimientoContableData := MovimientoContable{
			Debito:                   movimientoContable[i].Debito,
			Credito:                  movimientoContable[i].Credito,
			Fecha:                    time.Now(),
			Concepto:                 movimientoContable[i].Concepto,
			CuentaContable:           movimientoContable[i].CuentaContable,
			TipoDocumentoAfectante:   &TipoDocumentoAfectante{Id: int(tipoDocumentoAfectante.Id)}, //documento afectante tipo op
			CodigoDocumentoAfectante: int(ordenPago.Id),
			EstadoMovimientoContable: &EstadoMovimientoContable{Id: int(estadoMovimientoContable.Id)},
		}
		_, err = o.Insert(&movimientoContableData)
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_UPD_05"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
		if auxCuentaEspecial[i]["TipoCuentaEspecial"] != nil {
			if auxCuentaEspecial[i]["TipoCuentaEspecial"].(map[string]interface{})["Id"].(float64) == 2 {
				auxTipoDocumento = tipoDocumentoAfectanteGiro
				movimientoContableData := MovimientoContable{
					Debito:                   movimientoContable[i].Debito,
					Credito:                  movimientoContable[i].Credito,
					Fecha:                    time.Now(),
					Concepto:                 movimientoContable[i].Concepto,
					CuentaContable:           movimientoContable[i].CuentaContable,
					TipoDocumentoAfectante:   &TipoDocumentoAfectante{Id: int(auxTipoDocumento.Id)}, //documento afectante tipo op
					CodigoDocumentoAfectante: int(ordenPago.Id),
					EstadoMovimientoContable: &EstadoMovimientoContable{Id: int(estadoMovimientoContable.Id)},
				}
				_, err = o.Insert(&movimientoContableData)
				if err != nil {
					alerta.Type = "error"
					alerta.Code = "E_OPP_UPD_05"
					alerta.Body = err.Error()
					o.Rollback()
					return
				}

			}
		}
	}
	o.Commit()
	return
}

// personalizado Retrona la fecha actual del servidor
func FechaActual(formato string) (fechaActual string, err error) {
	hoy := time.Now()
	fechaActual = hoy.Format(formato)
	return
}

func ValorTotal(m int) (valorTotal int, err error) {
	o := orm.NewOrm()
	err = o.Raw("SELECT SUM(valor) FROM concepto_orden_pago WHERE orden_de_pago = ?", m).QueryRow(&valorTotal)
	return
}

func GetEstadoOrdenPago(CodeEstado string) (outputIdEstado EstadoOrdenPago, alerta Alert) {
	o := orm.NewOrm()
	o.Begin()
	outputIdEstado = EstadoOrdenPago{CodigoAbreviacion: CodeEstado}
	err := o.Read(&outputIdEstado, "CodigoAbreviacion")
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPP_01" //en busqueda de estado
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	return
}

// personalizado consultar orden_pago en su estado actual
func GetOrdenPagoByEstado(codeEstdoOrdenPago, vigencia, tipoOp, formaPago string) (outputOrdenes []interface{}, alerta Alert) {
	var ordenes []OrdenPago
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("DISTINCT OP.*").
		From("orden_pago OP,orden_pago_estado_orden_pago OPEOP, estado_orden_pago OPE, tipo_orden_pago TOP, sub_tipo_orden_pago STOP").
		Where("OP.id = OPEOP.orden_pago").
		And("OPEOP.id=(SELECT MAX(opeop2.id) from orden_pago_estado_orden_pago opeop2 where opeop2.orden_pago = OP.id group by opeop2.orden_pago)").
		And("OPEOP.estado_orden_pago = OPE.id").
		And("TOP.id = STOP.tipo_orden_pago").
		And("STOP.id = OP.sub_tipo_orden_pago").
		And("OPE.codigo_abreviacion = ?").
		And("OP.vigencia = ?").
		And("TOP.id = ?").
		And("OP.forma_pago = ?")
	_, err := o.Raw(qb.String(), codeEstdoOrdenPago, vigencia, tipoOp, formaPago).QueryRows(&ordenes)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPP_01"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	for i := 0; i < len(ordenes); i++ {
		/* println("Id: ", ordenes[i].Id) */
		var query = make(map[string]string)
		query["Id"] = strconv.Itoa(ordenes[i].Id)
		v, err := GetAllOrdenPago(query, nil, nil, nil, 0, 10)
		if err == nil {
			outputOrdenes = append(outputOrdenes, v[0])
		}
	}
	alerta = Alert{Type: "success", Code: "S_OPP_01", Body: "OK"}
	//
	return
}
