package controllers

import (
	"dashboard/models"
	"github.com/astaxie/beego"
)

type DaemonsetController struct {
	beego.Controller
}

func (c *DaemonsetController) Get(){

	ns := c.Ctx.Input.Param(":ns")

	if ns == ""{
		ns = "default"
	}

	daemonsetList := models.ShowDaemonset(ns)

	c.Data["json"] = daemonsetList
	c.ServeJSON()

}
