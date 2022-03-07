package app

import (
	"context"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

var JWTAuthentication = func(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request){
		auth:=[]string {"/fav","/cart", "/cart/buy"} //Список эндпоинтов, для которых требуется авторизация
		requestPath := r.URL.Path // текущий путь запроса
		//проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если она не нужна
		for _,value:=range auth{
			if value != requestPath{
				next.ServeHTTP(w,r)
				return
			}
		}
		response:=make(map[string]interface{})
		tokenHeader:=r.Header.Get("Authorization") // получение токена
		if tokenHeader == ""{ //Токен отсутствует, возвращаем  403 http-код Unauthorized
			response = Message(false, "Missing auth token")

			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-type", "application-json")
			Respond(w,response)
			return
		}
		splitted:=strings.Split(tokenHeader, " ")
		if len(splitted) != 2{
			response = Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-type", "application-json")
			Respond(w,response)
			return
		}
		tokenPart:= splitted[1] // Получаем вторую часть токена
		tk:= &Token{}

		token,err:= jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error){
			return []byte(os.Getenv("token_password")), nil
		})

		if err!=nil{ // Неправильный токен возвращает ошибку 403
			response = Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-type", "application-json")
			Respond(w,response)
			return
		}
		if !token.Valid{ // токен недействителен
			response = Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-type", "application-json")
			Respond(w,response)
			return
		}
		//Всё прошло хорошо, продолжаем выполнение запроса
		fmt.Sprintf("User %", tk.UserId)
		ctx:=context.WithValue(r.Context(),"user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w,r) // передаем управление следующему обработчику
	})
}
