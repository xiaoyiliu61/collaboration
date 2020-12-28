package db_mysql

import (

	"database/sql"
	"github.com/astaxie/beego"
)
var Db *sql.DB

/**
 * 数据库连接
 */
func Connect()  {

	config := beego.AppConfig
	dbDriver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	//fmt.Println(dbDriver, dbUser, dbPassword)
	//连接数据库
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db,err:= sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据连接失败，请检查配置")
	}
	Db = db
	//fmt.Println(db)
}


