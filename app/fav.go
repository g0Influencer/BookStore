package app

import "github.com/jinzhu/gorm"

type Fav struct{ // структура избранного
	gorm.Model
	UserId uint `json:"user_id"`
	ProductId int64 `json:"product_id"`
}