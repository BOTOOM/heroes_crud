package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:HeroController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:HeroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:HeroController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:HeroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:HeroController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:HeroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:HeroController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:HeroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:HeroController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:HeroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PowerController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PowerController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PowerController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PowerController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PowerController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PowerController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PowerController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PowerController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PowerController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PowerController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PublisherController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PublisherController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PublisherController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PublisherController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PublisherController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PublisherController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PublisherController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PublisherController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PublisherController"] = append(beego.GlobalControllerRouter["github.com/BOTOOM/heroes_crud/controllers:PublisherController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
