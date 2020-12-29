package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

<<<<<<< HEAD
func (I *IndexController) Post() {
   I.TplName="index.html"
=======
func (I *IndexController)Post() {
	I.TplName="index.html"
>>>>>>> remotes/origin/huzihan
}