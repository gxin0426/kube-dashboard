package controllers

import (
	"dashboard/models"
	"github.com/astaxie/beego"
)

type PodController struct {
	beego.Controller
}

func (this *PodController) Get()  {


	ns := this.Ctx.Input.Param(":ns")
	podlist := []models.PodStr{}
	if ns == "" {
		ns = "default"
	}
		podlist, _ = models.ShowPod(ns)
	//this.Data["json"] = podlist
	//this.ServeJSON()
	this.Data["Content"] = models.MakePodBlocks(podlist)
	this.TplName = "home.html"

}


