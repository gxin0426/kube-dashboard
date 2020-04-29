package controllers

import (
	"dashboard/utils"
	"github.com/astaxie/beego"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Nscontroller struct {
	beego.Controller
}

func (c Nscontroller) Get(){

	client, err := utils.CreateK8SClient()

	if err != nil {
		 panic(err)
		return
	}

	listOption := metav1.ListOptions{}

	nslist, err := client.CoreV1().Namespaces().List(listOption)

	var nsslice []string

	for _, ns2 := range nslist.Items{
		nsslice = append(nsslice, ns2.Name)
	}

	c.Data["json"] = nsslice
	c.ServeJSON()

}
