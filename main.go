package main

import (
	"collaboration/db_mysql"
	_ "collaboration/routers"
	"github.com/astaxie/beego"
)

func main() {

	//连接数据库
	db_mysql.Connect()

	//设置静态资源文件映射
	beego.SetStaticPath("/register.html", "./views/register.html")
	beego.SetStaticPath("/login1.html", "./views/login1.html")

	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")






	beego.Run()



}

