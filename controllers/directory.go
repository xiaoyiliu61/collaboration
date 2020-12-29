package controllers

import "github.com/astaxie/beego"

type DirectoryController struct {
	beego.Controller
}

func (D *DirectoryController) Post() {
	D.TplName="index.html"
}
