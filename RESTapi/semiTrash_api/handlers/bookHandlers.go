package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"semiTrash_api/models"

	"github.com/gorilla/mux"
)

func GetBookByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	book, ok := models.FindBookByID(id)
	log.Println("Get book with id:", id)

	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "Book with that ID does not exists in data base"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(book)
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new book ...")
	var book models.Book

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		msg := models.Message{Message: "provider json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var newBookID int = len(models.DB) + 1
	book.ID = newBookID
	models.DB = append(models.DB, book)

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(book)
}

func UpdateBookByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating book...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	oldBook, ok := models.FindBookByID(id)
	var newBook models.Book
	if !ok {
		log.Println("book not found in data base id:", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "book with that ID does not exist in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err := json.NewDecoder(request.Body).Decode(&newBook)
	if err != nil {
		msg := models.Message{Message: "provider json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Have to change oldBook on newBook in DB

}

func DeleteBookByID(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Deleting book...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	book, ok := models.FindBookByID(id)
	if !ok {
		log.Println("book not found in data base id:", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "book with that ID does not exist in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
}
