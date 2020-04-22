package main

import (
	_ "dashboard/routers"
	"github.com/astaxie/beego"
)
type people struct {
	id int
	name string
}
func main() {
	beego.SetStaticPath("/down", "down")
	beego.Run()

}

