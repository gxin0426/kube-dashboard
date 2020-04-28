package models

import (
	"dashboard/utils"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"time"
)

type PodStr struct {
	Name string
	Ready string
	Status v1.PodPhase
	Restart string
	Age string
}

func ShowPod(ns string)([]PodStr, *v1.NamespaceList){

	client, err := utils.CreateK8SClient()

	if err != nil{
		panic(err)
	}
	fmt.Println("连接k8s success")

	listOption := metav1.ListOptions{}

	//获取namespace
	namespace, err := client.CoreV1().Namespaces().List(listOption)

	var nstrList []string

	for _, nms := range namespace.Items{
		nstrList = append(nstrList, nms.Name)
	}

	fmt.Println(nstrList)

	podList, err := client.CoreV1().Pods(ns).List(listOption)
	var lp []PodStr

	for _, pod := range podList.Items{

		containerAllCount := len(pod.Status.ContainerStatuses)
		containerReadyCount := 0
		for _, cs := range pod.Status.ContainerStatuses {
			if cs.State.Running != nil {
				containerReadyCount++
			}
		}
		p := PodStr{}

		p.Name = pod.Name
		p.Ready = strconv.Itoa(containerReadyCount)+"/"+strconv.Itoa(containerAllCount)
		p.Status = pod.Status.Phase
		p.Restart = strconv.Itoa(int(pod.Status.ContainerStatuses[0].RestartCount))
		s := utils.FltoStr(time.Now().Sub(pod.Status.StartTime.Time).Hours())
		p.Age = s

		lp = append(lp, p)

	}

	return  lp, namespace
}

