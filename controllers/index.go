package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (I *IndexController)Post() {
	I.TplName="index.html"

}