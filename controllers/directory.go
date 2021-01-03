package controllers

import (
	"collaboration/utils"
	"github.com/astaxie/beego"
)

type DirectoryController struct {
	beego.Controller
}



func (D *DirectoryController)Post() {

/*	fmt.Println("当前区块的数量是",utils.GetBlockCount())
	fmt.Println("最新区块的hash是",utils.GetBestBlockHash())
	fmt.Println("根据高度0获取高度为0的hash值",utils.GetBlockHashByHeight(0))
	fmt.Println("获取新的地址",utils.GetNewAddress("123456",utils.LEGACY))*/


		D.TplName="bitDirectory.html"
	D.Data["d"] =utils.GetDifficulty()
	D.Data["b"] =utils.GetBestBlockHash()
	D.Data["h"] =utils.GetBlockHashByHeight(0)
	D.Data["c"] =utils.GetBlockCount()



}

func (D *DirectoryController)Get() {

	/*	fmt.Println("当前区块的数量是",utils.GetBlockCount())
		fmt.Println("最新区块的hash是",utils.GetBestBlockHash())
		fmt.Println("根据高度0获取高度为0的hash值",utils.GetBlockHashByHeight(0))
		fmt.Println("获取新的地址",utils.GetNewAddress("123456",utils.LEGACY))*/

	D.Data["d"] =utils.GetDifficulty()
	D.Data["b"] =utils.GetBestBlockHash()
	D.Data["h"] =utils.GetBlockHashByHeight(0)
	D.Data["c"] =utils.GetBlockCount()

	//D.TplName="login.html"


}

