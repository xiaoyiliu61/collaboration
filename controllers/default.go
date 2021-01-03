package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "login.html"
}

func (c *MainController) Post() {
	c.TplName = "login.html"

}
