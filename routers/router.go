package routers

import (
    "fmt"
	"apiproj/controllers"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    ns :=
    beego.NewNamespace("/v1",
        beego.NSCond(func(ctx *context.Context) bool {
            if ctx.Input.Domain() == "api.beego.me" {
                return true
            }
            return true
        }),
       // beego.NSBefore(auth),
        beego.NSGet("/notallowed", func(ctx *context.Context){
            ctx.Output.Body([]byte("notAllowed"))
        }),
        beego.NSRouter("/", &controllers.MainController{}),
        beego.NSNamespace("/p",
            beego.NSInclude(
                &controllers.UserController{},
          ),
        ),
    )

    beego.AddNamespace(ns)

    var FilterTet = func(ctx *context.Context){
       // ctx.WriteString("this is filter")
        fmt.Println("print filter at here")
    }
    beego.InsertFilter("/*", beego.BeforeRouter, FilterTet)
}

