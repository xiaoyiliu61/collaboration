package main

import (
	_ "collaboration/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	fmt.Println("hello world")
	fmt.Println("hello liuYi")
	beego.Run()

}

