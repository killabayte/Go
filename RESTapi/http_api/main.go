package main

import (
	"fmt"
	"net/http"
)

func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, this is new web-server!")
}

func main() {

}
