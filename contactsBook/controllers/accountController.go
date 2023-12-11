package controllers

import (
	"encoding/json"
	"go/course/contactsbook/models"
	u "go/course/contactsbook/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.BadRequest(w)
		return
	}
	resp := account.CreateAccount()
	if err, ok := resp["error"].(u.Error); ok {
		w.WriteHeader(err.HTTPCode)
	}
	u.Respond(w, resp)
}

var LoginAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.BadRequest(w)
		return
	}
	resp := models.LoginAccount(account.Email, account.Password)
	if err, ok := resp["error"].(u.Error); ok {
		w.WriteHeader(err.HTTPCode)
	}
	u.Respond(w, resp)
}
