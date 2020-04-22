package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type P struct {
	Name string
	Age int
}
func (c *MainController) Get() {

	p1 := P{"gg", 5}
	c.Data["json"] = &p1
	c.ServeJSON()
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
}
