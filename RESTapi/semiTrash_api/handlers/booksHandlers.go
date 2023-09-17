package handlers

import "net/http"

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application-json")
}

func GetAllBooks(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
}
