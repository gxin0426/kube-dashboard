package controllers

import (
	"dashboard/models"
	"github.com/astaxie/beego"
)

type NodeController struct {

	beego.Controller
}

func (this *NodeController)Get(){

	showjson := models.ShowNode()


		this.Data["json"] = showjson
		this.ServeJSON()


}
