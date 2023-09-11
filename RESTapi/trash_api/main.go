package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
)

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {

}

func GetPizzaById(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	log.Println("Trying to start REST API pizza!")
	router := mux.NewRouter()
	router.HandleFunc("/pizza", GetAllPizzas).Method("GET")
	router.HandleFunc("/pizza/{id}", GetPizzaById).Method("GET")
	log.Println("Router configured successfully! Ready to work")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}