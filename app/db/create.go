package db

import (
	"go-template/app/conf"
	"log"
)

func createUserTable() {
	if !db.Migrator().HasTable(&conf.User{}) {
		db.AutoMigrate(&conf.User{})
	} else {
		log.Println("createUserTable表已创建")
	}
}

func createAllTable() {
	createUserTable()
}
