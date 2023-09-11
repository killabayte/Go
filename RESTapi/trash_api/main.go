package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
)

func main() {
	log.Println("Trying to start REST API pizza!")
	router := mux.NewRouter()
	router.HandleFunc("/pizza", GetAllPizzas).Method("GET")
	router.HandleFunc("/pizza/{id}", GetPizzaById).Method("GET")
	log.Println("Router configured successfully! Ready to work")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
