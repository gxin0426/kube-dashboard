package models

import (
	"dashboard/utils"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"time"
)

type DeploymentStr struct {
	Name string
	Desired string
	Current string
	Uptodate string
	Available string
	Age string
}


//展示deployment
func ShowDeployment(ns string) []DeploymentStr{

	client, err := utils.CreateK8SClient()

	if err != nil {
		 panic(err)
		 return nil
	}

	fmt.Println("link success")

	listDeploy := metav1.ListOptions{}

	deploymentList, err := client.AppsV1().Deployments(ns).List(listDeploy)

	if err != nil{
		panic(err)
		return nil
	}

	var ld []DeploymentStr

	for _, deployment := range deploymentList.Items{
		d := DeploymentStr{}

		d.Name = deployment.Name
		d.Desired = strconv.Itoa(int(*deployment.Spec.Replicas))
		d.Current = strconv.Itoa(int(deployment.Status.Replicas))
		d.Uptodate = strconv.Itoa(int(deployment.Status.UpdatedReplicas))
		d.Available = strconv.Itoa(int(deployment.Status.AvailableReplicas))
		d.Age = utils.FltoStr(time.Now().Sub(deployment.GetCreationTimestamp().Time).Hours())

		ld = append(ld, d)
	}

	return  ld

}
