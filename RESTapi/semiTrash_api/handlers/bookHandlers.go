package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "command-line-arguments/Users/ykostiannikov/workspace/Go/RESTapi/semiTrash_api/models/helper.go"

	"github.com/gorilla/mux"
)

func GetBookByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err := nil{
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
	}
	book, ok := models.FindBookById(id)
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
