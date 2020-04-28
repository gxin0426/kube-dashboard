package controllers

import (
	"dashboard/models"
	"github.com/astaxie/beego"
)

type DeploymentController struct{

	beego.Controller
}


func (this *DeploymentController) Get(){

	ns := this.Ctx.Input.Param(":ns")

	if ns == ""{
		ns = "default"
	}


	deploylist := models.ShowDeployment(ns)
	this.Data["Content"] = models.MakeDeploymentBlocks(deploylist)
	this.TplName = "home.html"
}


