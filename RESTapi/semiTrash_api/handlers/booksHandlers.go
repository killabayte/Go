package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"semiTrash_api/models"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application-json")
}

func GetAllBooks(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Get infos about all books in data base")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.DB)
}
