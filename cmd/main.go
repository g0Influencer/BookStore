package main

import (
	"github.com/g0Influencer/BookStore/app"
	"github.com/g0Influencer/BookStore/controllers/admincontroller"
	"github.com/g0Influencer/BookStore/controllers/cartcontroller"
	"github.com/g0Influencer/BookStore/controllers/favcontroller"
	"github.com/g0Influencer/BookStore/controllers/productcontroller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main(){
	
	r:=mux.NewRouter()

	r.HandleFunc("/", productcontroller.Home)
	r.HandleFunc("/foreign", productcontroller.Home)
	r.HandleFunc("/kids", productcontroller.Home)
	r.HandleFunc("/fiction", productcontroller.Home)
	r.HandleFunc("/!fiction", productcontroller.Home)
	r.HandleFunc("/manga", productcontroller.Home)

	r.HandleFunc("/product", productcontroller.Home)
	r.HandleFunc("/product/index", productcontroller.Home)
	r.HandleFunc("/cart", cartcontroller.Index)
	r.HandleFunc("/cart/index", cartcontroller.Index)
	r.HandleFunc("/cart/buy", cartcontroller.Buy)
	r.HandleFunc("/cart/remove", cartcontroller.Remove)
	r.HandleFunc("/fav",favcontroller.Index)
	r.HandleFunc("/fav/add", favcontroller.Add)
	r.HandleFunc("/fav/remove", favcontroller.Remove)
	r.HandleFunc("/panel",admincontroller.Panel)
	r.HandleFunc("/panel/add",admincontroller.Add)
	r.HandleFunc("/panel/processadd",admincontroller.ProcessAdd)
	r.HandleFunc("/panel/delete",admincontroller.Delete)
	r.HandleFunc("/panel/edit",admincontroller.Edit)
	r.HandleFunc("/panel/update",admincontroller.Update)



	//добавляем middleware проверки JWT-токена
	r.HandleFunc("/api/user/new", app.CreateAccount)
	app.Init()
	r.HandleFunc("/api/user/login", app.Authenticate)

	r.Use(app.JWTAuthentication)


	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../static")))


	err:=http.ListenAndServe(":8000",r)
	if err!= nil{
		log.Fatal(err)
	}
}