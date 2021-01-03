package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

<<<<<<< HEAD

=======
>>>>>>> remotes/origin/huzihan
func (I *IndexController)Post() {
	I.TplName="index.html"

}