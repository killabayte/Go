package main

import (
	"log"
	"os"

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
	port = os.Getenv("app_port")
}

func main() {

}
