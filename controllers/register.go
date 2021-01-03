package controllers

import (
	"collaboration/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

<<<<<<< HEAD

=======
/**
 * post方法处理用户注册
 */
>>>>>>> remotes/origin/huzihan
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
<<<<<<< HEAD
	if err != nil {fmt.Println(err.Error())
		r.Ctx.WriteString("抱歉，用户注册失败，请重试!")

=======
	if err != nil {
		r.Ctx.WriteString("抱歉，用户注册失败，请重试!")
		fmt.Println(err.Error())
>>>>>>> remotes/origin/huzihan
		return
	}


	fmt.Println("跳转到login页面")
<<<<<<< HEAD

	r.TplName = "login.html"   //


=======
	r.TplName = "login.html"   // 文件上传界面{{.Phone}
>>>>>>> remotes/origin/huzihan
}


