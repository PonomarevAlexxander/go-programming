package models

import (
	u "go/course/contactsbook/utils"
	"net/http"
	"net/mail"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {

	if _, err := mail.ParseAddress(account.Email); err != nil {
		return u.JSONError(u.Error{
			HTTPCode: http.StatusBadRequest,
			Code:     400,
			Message:  "Email is invalid",
		}), false
	}

	if len(account.Password) < 4 {
		return u.JSONError(u.Error{
			HTTPCode: http.StatusBadRequest,
			Code:     400,
			Message:  "Passwords should be not less than 4 symbols",
		}), false
	}

	acc := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(acc).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.JSONError(u.Error{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  "Connection failed.",
		}), false
	}
	if acc.Email != "" {
		return u.JSONError(u.Error{
			HTTPCode: http.StatusBadRequest,
			Code:     400,
			Message:  "Email is already taken",
		}), false
	}

	return u.Message(false, "Check is passed!"), true
}

func (account *Account) CreateAccount() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(pwd)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.JSONError(u.Error{
			HTTPCode: http.StatusBadRequest,
			Code:     400,
			Message:  "Account not found!",
		})
	}

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	account.Password = ""

	response := u.Message(true, "Account has been created!")
	response["account"] = account
	return response
}

func LoginAccount(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.JSONError(u.Error{
				HTTPCode: http.StatusBadRequest,
				Code:     400,
				Message:  "Email address not found",
			})
		}
		return u.JSONError(u.Error{
			HTTPCode: http.StatusInternalServerError,
			Code:     500,
			Message:  "Connection error!",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Password does not match!")
	}
	account.Password = ""
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
