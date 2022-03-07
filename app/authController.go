package app

import (
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request){
	account:=&Account{}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "../static/html/registration.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			Respond(w, Message(false, "Invalid request"))
			return
		}
		email:= r.FormValue("email")
		password:= r.FormValue("password")
		account.Email = email
		account.Password = password
		resp:=account.Create()
		if resp["status"] == true{
			http.Redirect(w, r, "/", http.StatusFound)
		}else{
			http.Redirect(w,r,"/api/user/new",http.StatusFound)
		}
	}

}

var Authenticate = func(w http.ResponseWriter, r * http.Request){

	account:=&Account{}


	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "../static/html/login.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			Respond(w, Message(false, "Invalid request"))
			return
		}
		email := r.FormValue("email")
		password := r.FormValue("password")
		account.Email = email
		account.Password = password
		resp := Login(account.Email, account.Password)
		if email == "admin" && resp["status"] == true {
			http.Redirect(w, r, "/panel", http.StatusFound)
		} else {
			if resp["status"] == true {
				CurrentUserId = GetUser(email).ID
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				http.Redirect(w, r, "/api/user/login", http.StatusFound)

			}

		}
	}
}