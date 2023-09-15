package utils

import (
	handlers "./semiTrash_api/handlers"

	"github.com/gorilla/mux"
)

func BuildBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", GetBookByID).Methods("GET")
	router.HandleFunc(prefix, CreateBook).Methods("POST")
	router.HandleFunc(prefix+"/{id}", UpdateBookByID).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", DeleteBookByID).Methods("DELETE")
}

func BuildBooksResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllBooks).Methods("GET")
}
