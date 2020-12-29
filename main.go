package main

import (
	"collaboration/db_mysql"
	_ "collaboration/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	fmt.Println("hello world")

	fmt.Println("hello liuYi")


	fmt.Println("hi my name is zihanhu")


	fmt.Println("hello luxiaoyan")

	db_mysql.Connect()

	beego.Run()

}

