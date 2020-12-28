package routers

import (
	"collaboration/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/login", &controllers.RegisterController{})
    beego.Router("/directory", &controllers.DirectoryController{})
    beego.Router("/index", &controllers.IndexController{})

}
