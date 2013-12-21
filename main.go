package main

import (
	"github.com/astaxie/beego"
	"push/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/push/create", &controllers.IndexController{})
	beego.Router("/push/test", &controllers.TestController{})
	beego.Router("/push/query", &controllers.QueryController{})
	beego.Run()
}
