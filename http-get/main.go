package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

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

	if responce.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d: %s\n)", responce.StatusCode, body)
		os.Exit(1)
	}

	var words Words

	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON Parsed\nPage: %s\nWords: %v\n", words.Page, strings.Join(words.Words, ", "))
}
