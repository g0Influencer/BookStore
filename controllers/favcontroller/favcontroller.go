package favcontroller

import (
	"github.com/g0Influencer/BookStore/app"
	"github.com/g0Influencer/BookStore/config"
	"github.com/g0Influencer/BookStore/models/productmodel"
	"html/template"
	"log"
	"net/http"
	"strconv"
)


func Index(w http.ResponseWriter, r *http.Request){
	if app.CurrentUserId == 0{
		http.Redirect(w,r,"/api/user/login",http.StatusSeeOther)

	} else {

		var productModel productmodel.ProductModel
		var fav []app.Fav
		var orders []app.Product

		db := config.GetDB()
		db.Where("user_id = ? ", app.CurrentUserId).Find(&fav)

		for _, c := range fav {
			tmp, _ := productModel.Find(c.ProductId)
			orders = append(orders, tmp)

		}

		data := map[string]interface{}{
			"orders": orders,
		}

		ts, err := template.ParseFiles("../static/html/fav.html")
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

func Add(w http.ResponseWriter, r *http.Request){
	if app.CurrentUserId == 0{
		http.Redirect(w,r,"/api/user/login",http.StatusSeeOther)

	} else {

		var fav app.Fav

		query := r.URL.Query() // получаем данные из url-адреса запроса
		id, _ := strconv.ParseInt(query.Get("id"), 10, 64)

		db := config.GetDB()
		err := db.Where("user_id = ? and `product_id` = ?", app.CurrentUserId, id).Find(&fav).Error

		if err != nil {
			fav = app.Fav{UserId: app.CurrentUserId, ProductId: id}
			db.Create(&fav)
		}

		http.Redirect(w, r, "/fav", http.StatusSeeOther)
	}
}


func Remove(w http.ResponseWriter, r *http.Request){
	if app.CurrentUserId == 0{
		http.Redirect(w,r,"/api/user/login",http.StatusSeeOther)

	} else {
		var fav app.Fav
		query := r.URL.Query() // получаем данные из url-адреса запроса
		id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
		db := config.GetDB()

		db.Where("user_id = ? and `product_id` = ? ", app.CurrentUserId, id).Delete(&fav)

		http.Redirect(w, r, "/fav", http.StatusSeeOther)
	}

}



