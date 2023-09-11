package main

import (
	"log"
	"net/http"
)

var (
	port string = "8080"
)

func main() {
	log.Println("Trying to start REST API pizza!")
	log.Fatal(http.ListenAndServe(":8080"+port, nil))
}
