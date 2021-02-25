package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionTipoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionTipoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionTipoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionTipoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:AvanceLegalizacionTipoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionSolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionSolicitudTipoAvanceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionSolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionSolicitudTipoAvanceController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionSolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionSolicitudTipoAvanceController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionSolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionSolicitudTipoAvanceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionSolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionSolicitudTipoAvanceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionTipoAvanceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionTipoAvanceController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionTipoAvanceController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionTipoAvanceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EspecificacionTipoAvanceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EstadoLegalizacionAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EstadoLegalizacionAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EstadoLegalizacionAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EstadoLegalizacionAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EstadoLegalizacionAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EstadoLegalizacionAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EstadoLegalizacionAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EstadoLegalizacionAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EstadoLegalizacionAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:EstadoLegalizacionAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:InformacionEstudianteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:InformacionEstudianteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:InformacionEstudianteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:InformacionEstudianteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:InformacionEstudianteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:InformacionEstudianteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:InformacionEstudianteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:InformacionEstudianteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:InformacionEstudianteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:InformacionEstudianteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:NormaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:NormaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:NormaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:NormaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:NormaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:NormaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:NormaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:NormaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:NormaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:NormaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:OrdenPagoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:OrdenPagoAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:OrdenPagoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:OrdenPagoAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:OrdenPagoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:OrdenPagoAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:OrdenPagoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:OrdenPagoAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:OrdenPagoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:OrdenPagoAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:PracticaAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:PracticaAcademicaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:PracticaAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:PracticaAcademicaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:PracticaAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:PracticaAcademicaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:PracticaAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:PracticaAcademicaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:PracticaAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:PracticaAcademicaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RegistroAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RegistroAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RegistroAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RegistroAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RegistroAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RegistroAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RegistroAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RegistroAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RegistroAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RegistroAvanceLegalizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:ReintegroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:ReintegroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:ReintegroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:ReintegroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:ReintegroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RequisitoTipoAvanceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RequisitoTipoAvanceController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RequisitoTipoAvanceController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RequisitoTipoAvanceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:RequisitoTipoAvanceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudAvanceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudAvanceController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudAvanceController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudAvanceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudAvanceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudRequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudRequisitoTipoAvanceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudRequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudRequisitoTipoAvanceController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudRequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudRequisitoTipoAvanceController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudRequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudRequisitoTipoAvanceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudRequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudRequisitoTipoAvanceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudTipoAvanceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudTipoAvanceController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudTipoAvanceController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudTipoAvanceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:SolicitudTipoAvanceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:TipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:TipoAvanceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:TipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:TipoAvanceController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:TipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:TipoAvanceController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:TipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:TipoAvanceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:TipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/avances_crud/controllers:TipoAvanceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
