package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
	DB   []Pizza
)

type Pizza struct {
	ID       int     `json:"id"`
	Diameter int     `json:"diameter"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
}

func FindPizzaById() {

}

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {}
func GetPizzaById(writer http.ResponseWriter, request *http.Request) {}

func main() {
	log.Println("Trying to start REST API pizza!")
	router := mux.NewRouter()
	router.HandleFunc("/pizza", GetAllPizzas).Methods("GET")
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")
	log.Println("Router configured successfully! Ready to work")
	// err := http.ListenAndServe(":"+port, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Fatal(http.ListenAndServe(":"+port, router))
}
