package routers

import (
	"dashboard/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.PodController{})
    beego.Router("/pod", &controllers.PodController{})
    beego.Router("/service", &controllers.ServiceController{})
    beego.Router("/deployment", &controllers.DeploymentController{})
    beego.Router("/ingress", &controllers.IngressController{})
    //beego.AutoRouter(&controllers.MainController{})

}
