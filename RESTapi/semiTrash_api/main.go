package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	port string
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
	log.Fatal(http.ListenAndServe(":"+port, router))

}
