package models

import (
	"dashboard/utils"
	v1 "k8s.io/api/core/v1"
	"time"
)
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
type NodeStr struct {
	Name string
	Status v1.NodeConditionType
	//Roles string
	Age string
	Version string
}

func ShowNode() []NodeStr{

	client, err := utils.CreateK8SClient()

	if err != nil {
		panic(err)
		return nil
	}
	listOption := metav1.ListOptions{}
	nodeList, err := client.CoreV1().Nodes().List(listOption)

	var ln []NodeStr

	for _, node := range nodeList.Items{

		n := NodeStr{}

		n.Status = node.Status.Conditions[3].Type
		//for _ , t := range node.Status.Conditions{
		//
		//	fmt.Println("bbbbb    :", t)
		//
		//	fmt.Println("aaaaaaaaaaa", t.Type)
		//}
		n.Age = utils.FltoStr(time.Now().Sub(node.GetCreationTimestamp().Time).Hours())  + "h"
		n.Version  = node.Status.NodeInfo.KubeletVersion
		//fmt.Println(node.Status.NodeInfo.KernelVersion)
		n.Name = node.Name
		ln = append(ln, n)
	}
	return ln
}
