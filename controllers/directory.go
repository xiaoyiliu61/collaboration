package controllers

import "github.com/astaxie/beego"

type DirectoryController struct {
	beego.Controller
}

<<<<<<< HEAD
func (D *DirectoryController) Post() {
	D.TplName="index.html"
=======
func (D *DirectoryController)Post() {
	D.TplName="bitDirectory.html"
>>>>>>> remotes/origin/huzihan
}
