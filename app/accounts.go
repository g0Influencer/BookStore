package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

var CurrentUserId = Account{}.ID


type Token struct{
	UserId uint
	jwt.StandardClaims
}
type Account struct{
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool){
	if !strings.Contains(account.Email, "@") && account.Email != "admin"{
		return Message(false, "Email address is required"), false
	}
	if len(account.Password) < 6 && account.Email != "admin"{
		return Message(false, "Password is required"), false
	}
	temp:= &Account{}

	err:=GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err!= nil && err!= gorm.ErrRecordNotFound{
		return Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != ""{
		return Message(false, "Email address already in use by another user."), false
	}
	return Message(false,"Requirement passed" ), true
}

func (account *Account) Create() (map[string]interface{}){
	if resp,ok:=account.Validate();!ok{
		return resp
	}
	hashedPassword,_:=bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <=0{
		return Message(false, "Failed to create account, connection error.")
	}
	//Создать новый токен JWT для новой зарегистрированной учётной записи
	tk:=&Token{UserId: account.ID}
	token:= jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString,_:=token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" // удалить пароль

	response:=Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) (map[string]interface{}){
	account:=&Account{}
	err:=GetDB().Table("accounts").Where("email = ?", email).First(account).Error

	if err!= nil{
		if err == gorm.ErrRecordNotFound{
			return Message(false, "Email address not found")
		}
		return Message(false, "Connection error. Please retry")
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err!= nil && err == bcrypt.ErrMismatchedHashAndPassword{ // Пароль не совпадает
		return Message(false, "Invalid login credentials. Please try again")
	}
	//Работает! Войти в систему
	account.Password = ""

	//Создать токен JWT
	tk:= &Token{UserId: account.ID}
	token:=jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString,_:=token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString // Сохраняем токен в ответе

	resp:=Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(email string) *Account{
	acc:=&Account{}
	GetDB().Table("accounts").Where("email = ?",email).First(acc)
	if acc.Email == ""{ //Пользователь не найден
		return nil
	}
	acc.Password = ""
	return acc


}