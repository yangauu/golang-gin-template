package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// 连接
func connect() {
	var err error
	db, err = gorm.Open(mysql.Open("root:@(localhost:3306)/mapp"), &gorm.Config{})
	if err != nil {
		log.Println("连接数据库失败", err)
		return
	}
	log.Println("连接数据库成功")
	// createAllTable()
}
