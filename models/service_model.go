package models

import (
	"dashboard/utils"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	"time"
)

type ServiceStr struct {
	Name string
	Type string
	Clusterip string
	Externalip string
	Port string
	Age string
}


func ShowService() []ServiceStr {

	clientK8s, err := utils.CreateK8SClient()

	if err != nil{
		panic(err)
	}

	fmt.Println("连接kubernetes success")

	listOption := metav1.ListOptions{}
	serviceList, err := clientK8s.CoreV1().Services("").List(listOption)
	if err != nil{
		panic(err)
	}

	var ls []ServiceStr

	for _, service := range serviceList.Items{

			service2 := ServiceStr{}

			servicePorts := make([]string, 0, len(service.Spec.Ports))

			for _, p := range service.Spec.Ports{
				servicePorts = append(servicePorts, fmt.Sprintf("%d:%d/%s", p.Port, p.NodePort, p.Protocol))

			}

			externalIPs := make([]string, 0, len(service.Spec.ExternalIPs))

			for _, ip := range service.Spec.ExternalIPs{
				externalIPs = append(externalIPs, ip)
			}

			externalIPsStr := ""
			if len(externalIPs) > 0{
				externalIPsStr = strings.Join(externalIPs, ",")
			}

			service2.Name = service.Name
			service2.Type = string(service.Spec.Type)
			service2.Clusterip = service.Spec.ClusterIP
			service2.Externalip = externalIPsStr
			service2.Port = strings.Join(servicePorts, ",")
			t := utils.FltoStr(time.Now().Sub(service.GetCreationTimestamp().Time).Hours())
			service2.Age = t

			ls = append(ls, service2)
	}
			return  ls
}

