package main

import "net/http"

func main() {
	resp, err := http.Get("https://google.com")
}
