package models

import (
	"dashboard/utils"
	"strconv"
	"strings"
	"time"
)
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
type IngressStr struct {

	Name string
	Hosts string
	Address string
	Ports string
	Age string
}

func ShowIngress() []IngressStr{

	client, err := utils.CreateK8SClient()

	if err != nil {
		 panic(err)
		 return  nil
	}

	ingressList := metav1.ListOptions{}

	listIngress, err := client.ExtensionsV1beta1().Ingresses("" ).List(ingressList)
	if err != nil {
		panic(err)
		return  nil
	}

	var li []IngressStr
	for _, ingress := range listIngress.Items{

		ing := IngressStr{}
		//获取hosts
		ingressHosts := make([]string, 0, len(ingress.Spec.Rules))
		for _, rule := range ingress.Spec.Rules{
			ingressHosts = append(ingressHosts, rule.Host)
		}
		//获取service端口
		servicePortsSet := make(map[int]struct{})
		for _, rule := range ingress.Spec.Rules{
			for _, path := range rule.IngressRuleValue.HTTP.Paths{
				servicePortsSet[path.Backend.ServicePort.IntValue()]= struct{}{}
			}
		}

		servicePorts := make([]string, 0, len(servicePortsSet))

		for port := range servicePortsSet{
			servicePorts = append(servicePorts, strconv.Itoa(port))
		}

		//获取address
		ingressAddress := make([]string, 0, len(ingress.Spec.Rules))

		for _, ingSta := range ingress.Status.LoadBalancer.Ingress{
			ingressAddress = append(ingressAddress, ingSta.IP)
		}

		ing.Name = ingress.Name
		ing.Hosts = strings.Join(ingressHosts, ",")
		ing.Address = strings.Join(ingressAddress, ",")
		ing.Ports = strings.Join(servicePorts, ",")
		t := utils.FltoStr(time.Now().Sub(ingress.GetCreationTimestamp().Time).Hours())
		ing.Age = t

		li = append(li, ing)
	}
	return  li
}
