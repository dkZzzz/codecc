package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID       int64
	Name     string
	Password string
}

func InitDB() {
	db, err := gorm.Open("mysql", "root:123456@localhost/codecc?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
}
