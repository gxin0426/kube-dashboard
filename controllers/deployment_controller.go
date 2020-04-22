package controllers

import (
	"dashboard/models"
	"github.com/astaxie/beego"
)

type DeploymentController struct{

	beego.Controller
}


func (this *DeploymentController) Get(){

	deploylist := models.ShowDeployment()
	this.Data["Content"] = models.MakeDeploymentBlocks(deploylist)
	this.TplName = "home.html"
}


