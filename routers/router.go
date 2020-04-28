package routers

import (
	"dashboard/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//首页
	beego.Router("/", &controllers.PodController{})

	//pod
    beego.Router("/pod/:ns", &controllers.PodController{})
	beego.Router("/pod/", &controllers.PodController{})
	//service
    beego.Router("/service", &controllers.ServiceController{})
	beego.Router("/service/:ns", &controllers.ServiceController{})
	//deployment
    beego.Router("/deployment", &controllers.DeploymentController{})
	beego.Router("/deployment/:ns", &controllers.DeploymentController{})

	//ingress
    //beego.Router("/ingress", &controllers.IngressController{})

    //daemonset
	beego.Router("/daemonset", &controllers.DaemonsetController{})
	beego.Router("/daemonset/:ns", &controllers.DaemonsetController{})

	//node
    beego.Router("/node", &controllers.NodeController{})
	//log
	beego.Router("/log/:ns/:podname", &controllers.ShowPodLogController{})



}
