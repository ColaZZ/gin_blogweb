package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func InitMysql() {
	fmt.Println("InitMysql...")
	dsn := "root:0312@tcp(127.0.0.1:3306)/gin_blogweb"
	db, _ = sqlx.Connect("mysql", dsn)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)

	CreateTableWithUser()
}

func CreateTableWithUser() {
	sqlStr := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	_, _ = ModifyDB(sqlStr)
}

// 执行sql的exec语句
func ModifyDB(sqlStr string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sqlStr, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//执行sql的查询语句
func QueryRowDB(sqlStr string) *sqlx.Row {
	result := db.QueryRowx(sqlStr)
	return result
}
