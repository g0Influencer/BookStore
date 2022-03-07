package config

import (
	"file-share/app"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func Init() *gorm.DB {
	dsn := "root:golanggivesmemoney@/productsdb?parseTime=true"
	conn, err := gorm.Open("mysql", dsn )

	if err != nil {
		log.Fatalln(err)
	}
	db = conn
	db.AutoMigrate(&app.Product{},&app.Cart{},&app.Fav{})

	return db
}
func GetDB() *gorm.DB{
	return db
}
