package controllers

import (
	"collaboration/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

/**
 * post方法处理用户注册
 */
func (r *RegisterController) Post() {

	//1、解析用户端提交的请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据解析失败,请重试!")
		return
	}

	//2、将解析到的数据保存到数据库中
	_, err = user.AddUser()
	if err != nil {
		r.Ctx.WriteString("抱歉，用户注册失败，请重试!")
		fmt.Println(err.Error())
		return
	}


	fmt.Println("跳转到login页面")
	r.TplName = "login.html"   // 文件上传界面{{.Phone}
}
