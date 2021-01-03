package routers

import (
	"collaboration/controllers"
	"github.com/astaxie/beego"
)

func init() {

    //用户登录
    beego.Router("/login", &controllers.MainController{})
    //用户注册
	beego.Router("/register", &controllers.RegisterController{})

    //发送验证码
    beego.Router("/sendsms", &controllers.SentSmsController{})
    //手机号验证登录
    beego.Router("/login_sms", &controllers.LoginSmsController{})
    //目录
	beego.Router("/directory", &controllers.DirectoryController{})
    //主页面
	beego.Router("/index", &controllers.IndexController{})


}
