package main

import (
	_ "collaboration/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	fmt.Println("hello world")
	fmt.Println("hello liuYi")


}

