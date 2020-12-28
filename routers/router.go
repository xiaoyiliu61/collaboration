package routers

import (
	"collaboration/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//用户注册
    beego.Router("/", &controllers.MainController{})
    //用户登录
    beego.Router("/login", &controllers.RegisterController{})
    //
    beego.Router("/directory", &controllers.DirectoryController{})
    //
    beego.Router("/index", &controllers.IndexController{})

}
