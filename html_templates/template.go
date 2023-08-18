package main

import (
	"log"
)

type Part struct {
	Name  string
	Count int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	templateText := "Name: {{.Name}}\nCount: {{.Count}}\n"

}
