package app

import "github.com/jinzhu/gorm"

type Cart struct{ // структура корзины
	gorm.Model
	ProductId int64 `json:"product_id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Author string `json:"author"`
	Photo string `json:"photo"`
	UserId uint `json:"user_id"`
	CartQuantity int `json:"cart_quantity"`
}