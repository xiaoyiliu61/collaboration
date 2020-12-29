package models

import (
	"collaboration/db_mysql"
	"collaboration/utils"
)

type User struct {
	Id       int    `form:"id"`
	Username     string `form:"Username"`
	Password string `form:"password"`


}
/**
将用户信息保存到数据库中
*/
func (u User) AddUser() (int64, error) {
	//脱敏
	u.Password = utils.MD5HashString(u.Password) //把脱敏的密码的md5值重新赋值为密码进行存储

	rs, err := db_mysql.Db.Exec("insert into user(username,password) values(?,?)",
		u.Username, u.Password)
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

	row := db_mysql.Db.QueryRow("select id, username, password from user where username = ? and password = ?",
		u.Username,u.Password)

	err := row.Scan(&u.Id,&u.Username,&u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}