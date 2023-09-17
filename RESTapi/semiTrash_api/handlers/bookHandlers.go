package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBookByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
}
func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
}
func UpdateBookByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
}
func DeleteBookByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
}
