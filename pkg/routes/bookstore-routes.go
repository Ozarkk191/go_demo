package routes

import (
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", bookcontroller.CreateBook).Methods("POST")
}
