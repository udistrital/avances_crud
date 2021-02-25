// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/avances_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/avance_legalizacion_tipo",
			beego.NSInclude(
				&controllers.AvanceLegalizacionTipoController{},
			),
		),

		beego.NSNamespace("/especificacion_tipo_avance",
			beego.NSInclude(
				&controllers.EspecificacionTipoAvanceController{},
			),
		),

		beego.NSNamespace("/estado_legalizacion_avance_legalizacion",
			beego.NSInclude(
				&controllers.EstadoLegalizacionAvanceLegalizacionController{},
			),
		),

		beego.NSNamespace("/especificacion_solicitud_tipo_avance",
			beego.NSInclude(
				&controllers.EspecificacionSolicitudTipoAvanceController{},
			),
		),

		beego.NSNamespace("/movimiento",
			beego.NSInclude(
				&controllers.MovimientoController{},
			),
		),

		beego.NSNamespace("/orden_pago_avance_legalizacion",
			beego.NSInclude(
				&controllers.OrdenPagoAvanceLegalizacionController{},
			),
		),

		beego.NSNamespace("/norma",
			beego.NSInclude(
				&controllers.NormaController{},
			),
		),

		beego.NSNamespace("/reintegro",
			beego.NSInclude(
				&controllers.ReintegroController{},
			),
		),

		beego.NSNamespace("/solicitud_avance",
			beego.NSInclude(
				&controllers.SolicitudAvanceController{},
			),
		),

		beego.NSNamespace("/avance_legalizacion",
			beego.NSInclude(
				&controllers.AvanceLegalizacionController{},
			),
		),

		beego.NSNamespace("/registro_avance_legalizacion",
			beego.NSInclude(
				&controllers.RegistroAvanceLegalizacionController{},
			),
		),

		beego.NSNamespace("/tipo_avance",
			beego.NSInclude(
				&controllers.TipoAvanceController{},
			),
		),

		beego.NSNamespace("/practica_academica",
			beego.NSInclude(
				&controllers.PracticaAcademicaController{},
			),
		),

		beego.NSNamespace("/informacion_estudiante",
			beego.NSInclude(
				&controllers.InformacionEstudianteController{},
			),
		),

		beego.NSNamespace("/requisito_tipo_avance",
			beego.NSInclude(
				&controllers.RequisitoTipoAvanceController{},
			),
		),

		beego.NSNamespace("/solicitud_tipo_avance",
			beego.NSInclude(
				&controllers.SolicitudTipoAvanceController{},
			),
		),

		beego.NSNamespace("/solicitud_requisito_tipo_avance",
			beego.NSInclude(
				&controllers.SolicitudRequisitoTipoAvanceController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
