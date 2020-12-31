package controllers

import (
	"github.com/astaxie/beego"
)

type DirectoryController struct {
	beego.Controller
}



func (D *DirectoryController)Post() {
	/*var user1 models.User1
	a:=user1.Num2
	b:=user1.Num
    fmt.Println(a)
    fmt.Println(b)

	if a==b {*/
		D.TplName="login.html"
	/*}
	D.TplName="bitDirectory.html"*/

}

