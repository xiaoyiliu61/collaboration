package models

import (
	"collaboration/db_mysql"
	"collaboration/utils"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Username     string `form:"Username"`
	Password string `form:"password"`


}

func (u User) AddUser() (int64, error) {
	//脱敏
	u.Password = utils.MD5HashString(u.Password) //把脱敏的密码的md5值重新赋值为密码进行存储

	rs, err := db_mysql.Db.Exec("insert into user(phone) values(?)",
		u.Phone)
	//错误早发现早解决
	if err != nil { //保存数据遇到错误
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil { //保存数据遇到错误
		return -1, err
	}
	//id值代表的是此次数据操作影响的行数,id是一个整数int64类型
	return id, nil
}


/**
 * 查询用户信息
 */
func (u User) QueryUser() (*User, error) {

	//把脱敏的密码的md5值重新赋值为密码进行存储
	u.Password = utils.MD5HashString(u.Password)

	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?",
		u.Phone, u.Password)

	err := row.Scan(&u.Phone)
	if err != nil {
		return nil, err
	}
	return &u, nil
}