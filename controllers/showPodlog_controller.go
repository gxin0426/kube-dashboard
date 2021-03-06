package controllers

import (

	"bytes"
	"dashboard/utils"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	v1 "k8s.io/api/core/v1"

)

type ShowPodLogController struct {
	beego.Controller
}

func (this *ShowPodLogController) Get(){

	client, err := utils.CreateK8SClient()

	if err != nil{
		this.Data["json"] = string(err.Error())
		this.ServeJSON()
	}
	fmt.Println("连接k8s success")

	getlogoption := v1.PodLogOptions{}
	podName := this.Ctx.Input.Param(":podname")
	ns := this.Ctx.Input.Param(":ns")
	fmt.Println(podName)
	req := client.CoreV1().Pods(ns).GetLogs(podName, &getlogoption)

	podLogs, err := req.Stream()

	if err != nil {
		this.Data["json"] = string(err.Error())
		this.ServeJSON()
	}

	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil{
		this.Data["json"] = string(err.Error())
	}

	str := buf.String()

	//fmt.Println(str)
	this.Data["json"] = str
	this.ServeJSON()

}
