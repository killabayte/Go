package main

import (
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	placeholder := []byte("signature list goes here")
	_, err := writer.Write(placeholder)
	check(err)
}
