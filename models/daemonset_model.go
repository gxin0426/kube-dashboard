package models

import (
	"dashboard/utils"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"time"
)


type DaemonsetStr struct {

	Name string
	Current string
	Ready string
	Available string
	Age string
}

func ShowDaemonset(ns string) []DaemonsetStr{

	client, err := utils.CreateK8SClient()
	fmt.Println()
	if err != nil {
		 panic(err)
		 return nil
	}

	listOptions := metav1.ListOptions{}

	daemonsetList, err := client.AppsV1().DaemonSets(ns).List(listOptions)

	ld := []DaemonsetStr{}

	for _, daemon := range daemonsetList.Items{

		d := DaemonsetStr{}

		d.Name = daemon.Name
		d.Available = strconv.Itoa(int(daemon.Status.NumberAvailable))
		d.Ready = strconv.Itoa(int(daemon.Status.NumberReady))
		d.Current = strconv.Itoa(int(daemon.Status.CurrentNumberScheduled))
		d.Age = utils.FltoStr(time.Now().Sub(daemon.GetCreationTimestamp().Time).Hours())
		ld = append(ld, d)
	}
	return ld
}
