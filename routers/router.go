package routers

import (
	"collaboration/controllers"
	"github.com/astaxie/beego"
)

func init() {
<<<<<<< HEAD
	//路由
	beego.Router("/", &controllers.MainController{})
	//用户注册
    beego.Router("/register", &controllers.RegisterController{})
    //用户登录
    beego.Router("/login", &controllers.LoginController{})
    //bitcoin目录
    beego.Router("/directory", &controllers.DirectoryController{})
    //区块数据信息
    beego.Router("/index", &controllers.IndexController{})
=======
    //用户登录
    beego.Router("/login", &controllers.MainController{})
    //用户注册
	beego.Router("/register", &controllers.RegisterController{})
    //目录
	beego.Router("/directory", &controllers.DirectoryController{})
    //主页面
	beego.Router("/index", &controllers.IndexController{})



>>>>>>> remotes/origin/huzihan

}
