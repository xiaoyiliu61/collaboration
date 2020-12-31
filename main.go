package main

import (
	"collaboration/db_mysql"
	_ "collaboration/routers"
	"collaboration/utils"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	Getdifficulty1:=utils.GetDifficulty();
	fmt.Println("当前区块的难度值是",Getdifficulty1)
	//连接数据库
	db_mysql.Connect()




	//设置静态资源文件映射
	beego.SetStaticPath("/register.html", "./views/register.html")
	beego.SetStaticPath("/directory.html", "./views/bitDirectory.html")
	beego.SetStaticPath("/bitDirectory.html", "./views/bitDirectory.html")
	beego.SetStaticPath("/login.html", "./views/login.html")

	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")



	beego.Run() //阻塞
}

