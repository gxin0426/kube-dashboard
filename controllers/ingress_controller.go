package controllers

import (
	"dashboard/models"
	"github.com/astaxie/beego"
)

type IngressController struct {
	beego.Controller
}

func (this *IngressController) Get(){
	inglist := models.ShowIngress()

	this.Data["Content"] = models.MakeIngressBlocks(inglist)
	this.TplName = "home.html"
}
