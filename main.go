package main

import (
	_ "collaboration/routers"
	"collaboration/utils"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//区块难度值
	//Getdifficulty1:=utils.GetDifficulty();

	/*fmt.Println("当前区块的难度值是",utils.GetDifficulty())
	fmt.Println("当前区块的数量是",utils.GetBlockCount())
	fmt.Println("最新区块的hash是",utils.GetBestBlockHash())*/

	//fmt.Println("获取新的地址",utils.GetNewAddress("123456",utils.LEGACY))
	fmt.Println("根据高度0获取高度为0的hash值",utils.GetBlockHashByHeight(0))

	//连接数据库

	   /* db_mysql.Db.Exec("insert into " +
		"block1(" +
		"id,height,hash1)" +
		"values(" +
		"?,?,?)",
		1,0,utils.GetBlockHashByHeight(0))*/


	//设置静态资源文件映射
	beego.SetStaticPath("/register.html", "./views/register.html")
	beego.SetStaticPath("/directory.html", "./views/bitDirectory.html")
	beego.SetStaticPath("/bitDirectory.html", "./views/bitDirectory.html")
	beego.SetStaticPath("/login.html", "./views/login.html")
	beego.SetStaticPath("/index", "./views/index.html")

	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")



	beego.Run() //阻塞
}

