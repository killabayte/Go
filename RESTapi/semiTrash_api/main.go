package main

import (
	"log"

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
}

func main() {

}
