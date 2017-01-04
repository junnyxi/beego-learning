package controllers

import (
    "github.com/astaxie/beego"
)

type UserController struct {
    beego.Controller
}

type user struct {
    Username interface{} `form:"username"`
    Email    interface{} `form:"email"`
}

// @router /user/info [get]
func (c *UserController) Info(){
    beego.Informational(">>> get user info")
    var username string
    c.Ctx.Input.Bind(&username, "username")
    var email string
    c.Ctx.Input.Bind(&email, "email")
   // username := c.GetString("username")
    user := map[string]interface{}{"username":username, "email":email}
    c.Data["json"] = user
    c.ServeJSON()
}

// @router /user/login [get]
func (c *UserController) Login(){
    c.Ctx.WriteString("This is login method")
    c.StopRun()
}

// @router /user [get]
func (c *UserController)Index(){
    c.Ctx.WriteString("user index")
    c.StopRun()
}
