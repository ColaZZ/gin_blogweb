package models

import (
	"fmt"
	"gin_blogweb/database"
	)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int
	CreateTime int64
}

// 插入新用户
func InsertUser(user User) (int64, error){
	row, _ := database.ModifyDB("insert into user(username, password, status, createtime) values(?,?,?,?)",
		user.Id, user.Username, user.Password, user.CreateTime)
	return row, nil
}


// 按条件查询用户
func QueryUserWithCon(con string) int {
	sqlStr := fmt.Sprintf("select id from user ? ", con)
	fmt.Println(sqlStr)
	row := database.QueryRowDB(sqlStr)
	id := 0
	_ = row.Scan(&id)
	return id
}

// 按用户名查询用户
func QueryUserWithUsername(username string) int {
	sqlStr := fmt.Sprintf("where username = ?", username)
	return QueryUserWithCon(sqlStr)
}

//根据用户名和密码，查询id
func QueryUserWithParams(username string, password string) int {
	sqlStr := fmt.Sprintf("where username=? and pasword=?", username, password)
	return QueryUserWithCon(sqlStr)
}
