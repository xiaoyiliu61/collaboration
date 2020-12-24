package main

import (
	_ "collaboration/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	fmt.Println("hello world")
	fmt.Println("hi my name is zihanhu")
}

