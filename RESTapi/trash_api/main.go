package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
	db   []Pizza
)

func init() {
	pizza1 := Pizza{
		ID:       1,
		Diameter: 28,
		Price:    17.5,
		Title:    "Margarita",
	}
	pizza2 := Pizza{
		ID:       2,
		Diameter: 35,
		Price:    27.99,
		Title:    "Double Cheese",
	}
	pizza3 := Pizza{
		ID:       3,
		Diameter: 18,
		Price:    21.30,
		Title:    "BBQ",
	}
	db = append(db, pizza1, pizza2, pizza3)
}

type Pizza struct {
	ID       int     `json:"id"`
	Diameter int     `json:"diameter"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
}

func FindPizzaById(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool
	for _, p := range db {
		if p.ID == id {
			pizza = p
			found = true
			break
		}
	}
	return pizza, found
}

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Println("Get info about all Pizzas in DB")
	writer.WriteHeader(200) // Status code
	json.NewEncoder(writer).Encode(db)
}
func GetPizzaById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request) // { "id": "42" }
	id, err := strconv.Atoi()
}

func main() {
	log.Println("Trying to start REST API pizza!")
	router := mux.NewRouter()
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")
	log.Println("Router configured successfully! Ready to work")
	// err := http.ListenAndServe(":"+port, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Fatal(http.ListenAndServe(":"+port, router))
}
