package app

import (
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func Init()  *gorm.DB {
	dsn := "root:golanggivesmemoney@/usersdb?parseTime=true"
	conn, err := gorm.Open("mysql", dsn )

	if err != nil {
		log.Fatalln(err)
	}
	db = conn
	db.AutoMigrate(&Token{}, &Account{})
	return db
}

func GetDB() *gorm.DB{
	return db
}
