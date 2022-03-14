package cartmodel

import (
	"github.com/g0Influencer/BookStore/app"
	"github.com/g0Influencer/BookStore/config"
)

type CartModel struct{

}

func (*CartModel) FindCart(userId uint) (app.Cart,error){ // поиск корзины конкретного юзера

	var cart app.Cart
	rows, err :=config.Init().Raw("select * from carts where user_id = ?",userId).Rows()
	if err!=nil{
		return app.Cart{},err
	}
	for rows.Next(){
		rows.Scan(&cart.CartQuantity,&cart.Name, &cart.ProductId, &cart.Photo,&cart.Price,
			&cart.Author, &cart.UserId)

	}

	return cart,nil
}