package main

import (
	"html/template"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	text := "Here's my template!"
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, nil)
	check(err)

}
