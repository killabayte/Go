package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                string
	bookResourcePrefix  string = apiPrefix + "/book"
	booksResourcePrefix string = apiPrefix + "/books"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.Router()
	log.Println("Router initialization is succesfull. Ready to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))

}
