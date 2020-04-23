package controllers

import (
	"dashboard/models"
	"github.com/astaxie/beego"
)

type  ServiceController struct {
	beego.Controller
}

func (this *ServiceController) Get(){

	servicelist := models.ShowService()
	this.Data["Content"] = models.MakeServiceBlocks(servicelist)
	this.TplName = "home.html"
	
}





