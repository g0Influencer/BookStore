package admincontroller

import (
	"github.com/g0Influencer/BookStore/app"
	"github.com/g0Influencer/BookStore/config"
	"github.com/g0Influencer/BookStore/models/productmodel"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Panel(w http.ResponseWriter, r *http.Request){

	products,_:=productmodel.FindByPath(w,r)
	data:=map[string]interface{}{
		"products":products, // all products
	}

	ts, err := template.ParseFiles("../static/html/admin.html")
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

func Add(w http.ResponseWriter, r *http.Request){
	ts, err := template.ParseFiles("../static/html/add.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func ProcessAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id ,_:= strconv.ParseInt(r.Form.Get("id"),10,64)
	name:= r.Form.Get("name")
	author := r.Form.Get("author")
	price,_:= strconv.ParseFloat(r.Form.Get("price"),64)
	quantity ,_:= strconv.ParseInt(r.Form.Get("id"),10,64)
	description:=r.Form.Get("description")
	photo := r.Form.Get("photo")

	product:=app.Product{Id:id, Name:name,Author: author, Price: price,
		Quantity: quantity,Description: description,  Photo: photo}

	db:=config.GetDB()
	db.Create(&product)
	http.Redirect(w,r,"/panel",http.StatusSeeOther)
}

func Delete(w  http.ResponseWriter, r *http.Request){
	query:=r.URL.Query()
	id,_:=strconv.ParseInt(query.Get("id"),10,64)
	productmodel.Delete(id)
	http.Redirect(w,r,"/panel",http.StatusSeeOther)

}

func Edit(w  http.ResponseWriter, r *http.Request){
	query:=r.URL.Query()
	id,_:=strconv.ParseInt(query.Get("id"),10,64)
	var productModel productmodel.ProductModel
	product,_:=productModel.Find(id)
	data:=map[string]interface{}{
		"product":product,
	}

	ts, err := template.ParseFiles("../static/html/edit.html")
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

func Update(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id ,_:= strconv.ParseInt(r.Form.Get("id"),10,64)
	name:= r.Form.Get("name")
	author := r.Form.Get("author")
	price,_:= strconv.ParseFloat(r.Form.Get("price"),64)
	quantity ,_:= strconv.ParseInt(r.Form.Get("id"),10,64)
	description:=r.Form.Get("description")
	photo := r.Form.Get("photo")

	product:=app.Product{Id:id, Name:name,Author: author, Price: price,
		Quantity: quantity,Description: description,  Photo: photo}

	productmodel.Update(&product)
	http.Redirect(w,r,"/panel",http.StatusSeeOther)
}


