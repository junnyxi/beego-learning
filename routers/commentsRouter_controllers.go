package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["apiproj/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproj/controllers:UserController"],
		beego.ControllerComments{
			Method: "Info",
			Router: `/user/info`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiproj/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproj/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/user/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["apiproj/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproj/controllers:UserController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/user`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
