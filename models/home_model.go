package models

import (
	"bytes"
	"fmt"
	"html/template"
)

func MakePodBlocks(podlist []PodStr) template.HTML{
	htmlHome := "<table class= 'home-block-item pod'><tr><th>NAME</th><th>READY</th><th>STATUS</th><th>RESTART</th><th>AGE</th></tr>"

	for _, pod := range podlist{
		p := PodStr{}

		p.Name = pod.Name
		p.Status = pod.Status
		p.Ready = pod.Ready
		p.Restart = pod.Restart
		p.Age = pod.Age

		t, _ := template.ParseFiles("views/block/pod_block.html")
		buffer := bytes.Buffer{}
		t.Execute(&buffer, p)
		htmlHome += buffer.String()
	}
	fmt.Println(template.HTML(htmlHome+"</table>"))
	return template.HTML(htmlHome+"</table>")
}


func MakeServiceBlocks(servicelist []ServiceStr) template.HTML{
	htmlHome := "<table class= 'home-block-item'><tr><th>NAME</th><th>TYPE</th><th>CLUSTER-IP</th><th>EXTERNAL-IP</th><th>PORT</th><th>AGE</th></tr>"

	for _, service := range servicelist{
		s := ServiceStr{}

		s.Name = service.Name
		s.Type = service.Type
		s.Port = service.Port
		s.Externalip = service.Externalip
		s.Clusterip = service.Clusterip
		s.Age = service.Age

		t, _ := template.ParseFiles("views/block/service_block.html")
		buffer := bytes.Buffer{}
		t.Execute(&buffer, s)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome + "</table>")
}

func MakeDeploymentBlocks(deploylist []DeploymentStr) template.HTML{
	htmlHome := "<table class= 'home-block-item'><tr><th>NAME</th><th>DESIRED</th><th>CURRENT</th><th>UP-TO-DATE</th><th>AVAILABLE</th><th>AGE</th></tr>"

	for _, deploy := range deploylist{
		d := DeploymentStr{}

		d.Name = deploy.Name
		d.Desired = deploy.Desired
		d.Current = deploy.Current
		d.Uptodate = deploy.Uptodate
		d.Available = deploy.Available
		d.Age = deploy.Age

		t, _ := template.ParseFiles("views/block/deployment_block.html")
		buffer := bytes.Buffer{}
		t.Execute(&buffer, d)
		htmlHome += buffer.String()
	}
	fmt.Println(template.HTML(htmlHome + "</table>"))
	return template.HTML(htmlHome + "</table>")
}

func MakeIngressBlocks(ingresslist []IngressStr) template.HTML{
	htmlHome := "<table class= 'home-block-item'><tr><th>NAME</th><th>HOSTS</th><th>ADDRESS</th><th>PORT</th><th>AGE</th></tr>"

	for _, ing := range ingresslist{
		i := IngressStr{}

		i.Name = ing.Name
		i.Age = ing.Age
		i.Ports = ing.Ports
		i.Address = ing.Address
		i.Hosts = ing.Hosts

		t, _ := template.ParseFiles("views/block/deployment_block.html")
		buffer := bytes.Buffer{}
		t.Execute(&buffer, i)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome + "</table>")
}
