package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type RegistroAvanceLegalizacion struct {
	Id                       int                     `orm:"column(id);pk"`
	AvanceLegalizacionTipoId *AvanceLegalizacionTipo `orm:"column(avance_legalizacion_tipo_id);rel(fk)"`
	FechaEvento              time.Time               `orm:"column(fecha_evento);type(timestamp without time zone)"`
	TipoDocumentoTerceroId   int                     `orm:"column(tipo_documento_tercero_id)"`
	DocumentoTercero         string                  `orm:"column(documento_tercero)"`
	NumeroFactura            int                     `orm:"column(numero_factura)"`
	TrmFechaCompra           int                     `orm:"column(trm_fecha_compra)"`
	ValorTotal               float64                 `orm:"column(valor_total)"`
	RentaId                  int                     `orm:"column(renta_id)"`
	IcaId                    int                     `orm:"column(ica_id)"`
	IvaId                    int                     `orm:"column(iva_id)"`
	ValorRenta               float64                 `orm:"column(valor_renta)"`
	ValorIca                 float64                 `orm:"column(valor_ica)"`
	ValorIva                 float64                 `orm:"column(valor_iva)"`
	FechaCreacion            time.Time               `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion        time.Time               `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	Activo                   bool                    `orm:"column(activo)"`
}

func (t *RegistroAvanceLegalizacion) TableName() string {
	return "registro_avance_legalizacion"
}

func init() {
	orm.RegisterModel(new(RegistroAvanceLegalizacion))
}

// AddRegistroAvanceLegalizacion insert a new RegistroAvanceLegalizacion into database and returns
// last inserted Id on success.
func AddRegistroAvanceLegalizacion(m *RegistroAvanceLegalizacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRegistroAvanceLegalizacionById retrieves RegistroAvanceLegalizacion by Id. Returns error if
// Id doesn't exist
func GetRegistroAvanceLegalizacionById(id int) (v *RegistroAvanceLegalizacion, err error) {
	o := orm.NewOrm()
	v = &RegistroAvanceLegalizacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRegistroAvanceLegalizacion retrieves all RegistroAvanceLegalizacion matches certain condition. Returns empty list if
// no records exist
func GetAllRegistroAvanceLegalizacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RegistroAvanceLegalizacion))
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

	var l []RegistroAvanceLegalizacion
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

// UpdateRegistroAvanceLegalizacion updates RegistroAvanceLegalizacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateRegistroAvanceLegalizacionById(m *RegistroAvanceLegalizacion) (err error) {
	o := orm.NewOrm()
	v := RegistroAvanceLegalizacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRegistroAvanceLegalizacion deletes RegistroAvanceLegalizacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRegistroAvanceLegalizacion(id int) (err error) {
	o := orm.NewOrm()
	v := RegistroAvanceLegalizacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RegistroAvanceLegalizacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
