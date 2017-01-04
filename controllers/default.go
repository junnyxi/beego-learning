package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type ErrorController struct{
    beego.Controller
}

func (c *ErrorController)Error404(){
    c.Data["json"]= map[string]interface{}{"errno":404,"errmsg":"not found page"}
    c.ServeJSON()
}

func (c *MainController) Get() {
	c.Data["Website"] = "wlr.me"
	c.Data["Email"] = "wlr@gmail.com"
	c.TplName = "index.tpl"
}

