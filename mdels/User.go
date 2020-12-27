package mdels

import (
	"collaboration/db_mysql"
)

/**
 *@Id：账号
 *@Password：密码
 */
type User struct {
	Id       int    `form:"id"`
	Password string `form:"password"`
}

/**
 * 保存用户信息的方法：保存用户信息到数据库中
 */
func (u User) AddUser() (int64, error) {
	//1、密码脱敏处理
	//u.Password = util.MD5HashString(u.Password)

	//2、执行数据库操作
	row, err := db_mysql.Db.Exec("insert into user(id, password)"+
		" values(?,?) ", u.Id, u.Password)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}
