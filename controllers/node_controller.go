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

	for _, node := range showjson{
		this.Data["json"] = node
		this.ServeJSON()
	}

}
