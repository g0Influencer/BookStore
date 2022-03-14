package productmodel

import (
	"github.com/g0Influencer/BookStore/app"
	"github.com/g0Influencer/BookStore/config"
	"net/http"
)

type ProductModel struct{

}

func (*ProductModel) Find(id int64) (app.Product,error){ // поиск товара по id
	var product app.Product
	rows, err :=config.Init().Raw("select * from products where id = ?",id).Rows()
	if err!=nil{
		return app.Product{},err
	}
	for rows.Next(){
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity,
			&product.Description, &product.Author,&product.Photo)
	}
	return product,nil
}


func FindByPath(w http.ResponseWriter, r *http.Request) ([]app.Product,error){
	var products []app.Product
	var product app.Product
	rows, _ :=config.Init().Raw("select * from products where id is null").Rows()

	if r.URL.Path == "/" || r.URL.Path == "/panel"{
		rows, _ =config.Init().Raw("select * from products").Rows()
	} else if r.URL.Path == "/foreign"{
		rows, _ =config.Init().Raw("select * from products where id = 5 or id = 15").Rows()
	} else if r.URL.Path == "/kids"{
		rows, _ =config.Init().Raw("select * from products where id = 1 or id = 17 or id = 2").Rows()
	} else if r.URL.Path == "/manga"{
		rows, _ =config.Init().Raw("select * from products where id = 13 or id = 16").Rows()
	} else if r.URL.Path == "/fiction"{
		rows, _ =config.Init().Raw("select * from products where id in (1,2,6,7,8,9,10,11,12)").Rows()
	} else if r.URL.Path == "/!fiction"{
		rows, _ =config.Init().Raw("select * from products where id in (3,4,14,18)").Rows()
	}

	for rows.Next(){
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity,
			&product.Description, &product.Author, &product.Photo)
		products = append(products,product)
	}
	return products,nil
}

func Delete(id int64){
	var product app.Product
	db:=config.GetDB()
	db.Where("id = ?",id).Delete(&product)
}

func Update(product *app.Product){
	db:=config.GetDB()
	db.Model(&product).Updates(app.Product{Name: product.Name, Id: product.Id,
		Author: product.Author,Price: product.Price,Description: product.Description,
		Quantity: product.Quantity,Photo: product.Photo})

}
