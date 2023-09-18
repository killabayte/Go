package main

import (
	"log"

	"github.com/killabayte/Go/RESTapi/StandardWebServer/internal/app/api"
)

var ()

func init() {
}

func main() {
	log.Println("It works!")
	server := api.New()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
