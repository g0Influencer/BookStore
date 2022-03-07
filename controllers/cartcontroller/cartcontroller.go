package cartcontroller

import (
	"file-share/app"
	"file-share/config"
	"file-share/models/productmodel"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request){

	if app.CurrentUserId == 0{
		http.Redirect(w,r,"/api/user/login",http.StatusSeeOther)

	} else {

		var cart []app.Cart

		db := config.GetDB()
		db.Where("user_id = ? ", app.CurrentUserId).Find(&cart)

		total := Total(cart)
		data := map[string]interface{}{
			"cart":  cart,
			"total": total,
		}

		ts, err := template.ParseFiles("../static/html/cart.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.Execute(w, data)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

func Buy(w http.ResponseWriter, r *http.Request){ // перенаправление в корзину
	if app.CurrentUserId == 0{
		http.Redirect(w,r,"/api/user/login",http.StatusSeeOther)

	} else {
		var cart app.Cart

		query := r.URL.Query() // получаем данные из url-адреса запроса
		id, _ := strconv.ParseInt(query.Get("id"), 10, 64)

		var productModel productmodel.ProductModel
		product, _ := productModel.Find(id)

		db := config.GetDB()
		err := db.Where("user_id = ? and `product_id` = ?", app.CurrentUserId, product.Id).Find(&cart).Error

		if err == nil {
			cart.CartQuantity++
			db.Save(&cart)
		} else {
			cart = app.Cart{UserId: app.CurrentUserId, ProductId: product.Id,
				Price: product.Price, Name: product.Name, Author: product.Author,
				Photo: product.Photo, CartQuantity: 1}
			db.Create(&cart)
		}

		http.Redirect(w, r, "/cart", http.StatusSeeOther)
	}
}

func Remove(w http.ResponseWriter, r *http.Request){
	if app.CurrentUserId == 0{
		http.Redirect(w,r,"/api/user/login",http.StatusSeeOther)

	} else {
		var cart app.Cart

		query := r.URL.Query() // получаем данные из url-адреса запроса
		id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
		db := config.GetDB()
		db.Where("user_id = ? and `product_id` = ? ", app.CurrentUserId, id).Find(&cart)
		if cart.CartQuantity > 1 {
			cart.CartQuantity--
			db.Save(&cart)

		} else {
			db.Where("user_id = ? and `product_id` = ? ", app.CurrentUserId, id).Delete(&cart)
		}

		http.Redirect(w, r, "/cart", http.StatusSeeOther)
	}
}


func Total(cart []app.Cart) float64{ // подсчет итоговой суммы заказа
	var totalSum float64
	for _,s:=range cart{
		totalSum+= float64(s.CartQuantity) * s.Price
	}
	return totalSum
}



