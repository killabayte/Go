package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usege: ./http-get <url>")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in invalid format, %s", err)
		os.Exit(1)
	}

	responce, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer responce.Body.Close()
	body, err := io.ReadAll(responce.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP Status code: %d\nBody: %s", responce.StatusCode, body)
}
