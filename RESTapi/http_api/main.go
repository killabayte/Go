package main

import (
	"fmt"
	"log"
	"net/http"
)

func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, this is new web-server!")
}

func RequestHandler() {
	http.HandleFunc("/", GetGreet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	RequestHandler()
}
