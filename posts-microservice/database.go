package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connection() {

	dsn := "user:root@tcp(127.0.0.1:3306)/posts-microservice?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//defer db.Close()
}

func Migrate() {
	db.AutoMigrate(Post{})
}
