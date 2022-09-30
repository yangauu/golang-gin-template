package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 连接
func connect() {
	var err error
	db, err = sql.Open("mysql", "root:@/mapp")

	if err != nil {
		log.Println("连接数据库失败", err)
	}
}
