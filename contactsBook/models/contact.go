package models

import (
	"fmt"
	u "go/course/contactsbook/utils"
	"net/http"
	"regexp"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) ValidateContact() (map[string]interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "Name cannot be empty!"), false
	}

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if !re.MatchString(contact.Phone) {
		return u.Message(false, "Phone number is invalid!"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User not found!"), false
	}

	return u.Message(true, "success"), true
}

func (contact *Contact) CreateContact() map[string]interface{} {

	if response, ok := contact.ValidateContact(); !ok {
		return response
	}

	var existing Contact
	err := GetDB().
		Table("contacts").
		Where("id = ?", contact.ID).
		First(existing).
		Error
	if err != nil {
		GetDB().Save(contact)
	} else {
		GetDB().Create(contact)
	}

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func DeleteContact(id uint) map[string]interface{} {

	if GetDB().Table("contacts").
		Delete(&Contact{}, id).
		Error != nil {
		return u.Message(true, "Contact deleted")
	}

	return map[string]interface{}{
		"error": u.Error{
			HTTPCode: http.StatusNoContent,
			Code:     204,
			Message:  "Delete failed",
		},
	}
}

func GetContact(id uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {

	contactsSlice := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contactsSlice).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contactsSlice
}
