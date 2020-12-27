package routers

import (
	"collaboration/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //用户注册接口
    beego.Router("login",&controllers.RegisterController{})
}
