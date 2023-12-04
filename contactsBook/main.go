package main

import (
	"fmt"
	app "go/course/contactsbook/authentication"
	"go/course/contactsbook/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/user/login", controllers.LoginAccount).Methods("POST")
	router.HandleFunc("/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/me/contacts", controllers.GetContacts).Methods("GET")
	router.HandleFunc("/me/contacts", controllers.CreateContact).Methods("PUT")
	router.HandleFunc("/me/contacts", controllers.DeleteContact).Methods("DELETE")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
