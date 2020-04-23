package controllers

import (
	"dashboard/models"
	"github.com/astaxie/beego"
)

type PodController struct {
	beego.Controller
}

func (this *PodController) Get()  {


	podlist:= models.ShowPod()

	this.Data["json"] = podlist
	this.ServeJSON()
	//this.Data["Content"] = models.MakePodBlocks(podlist)
	//this.TplName = "home.html"

}


