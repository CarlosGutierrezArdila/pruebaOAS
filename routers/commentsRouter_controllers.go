package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ActividadController"] = append(beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ActividadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ActividadController"] = append(beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ActividadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ActividadController"] = append(beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ActividadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ActividadController"] = append(beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ActividadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ActividadController"] = append(beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ActividadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ResponsableController"] = append(beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ResponsableController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ResponsableController"] = append(beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ResponsableController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ResponsableController"] = append(beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ResponsableController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ResponsableController"] = append(beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ResponsableController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ResponsableController"] = append(beego.GlobalControllerRouter["github.com/carlosgutierrezardila/pruebaOAS/controllers:ResponsableController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
