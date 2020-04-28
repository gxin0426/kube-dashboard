package controllers

import (
	"dashboard/models"
	"github.com/astaxie/beego"
)

type  ServiceController struct {
	beego.Controller
}

func (this *ServiceController) Get(){
	ns := this.Ctx.Input.Param(":ns")

	if ns == ""{
		ns = "default"
	}

	servicelist := models.ShowService(ns)
	this.Data["Content"] = models.MakeServiceBlocks(servicelist)
	this.TplName = "home.html"
	
}





