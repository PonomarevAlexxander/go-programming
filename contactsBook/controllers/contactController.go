package controllers

import (
	"encoding/json"
	"go/course/contactsbook/models"
	u "go/course/contactsbook/utils"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.ServerError(w)
		return
	}

	contact.UserId = user
	resp := contact.CreateContact()
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.ServerError(w)
		return
	}
	if contact.Id <= 0 {
		u.BadRequest(w)
		return
	}

	contact.UserId = user
	data := models.DeleteContact(contact.Id)
	if err, ok := data["error"].(u.Error); ok {
		w.WriteHeader(err.HTTPCode)
	}
	u.Respond(w, data)
}
