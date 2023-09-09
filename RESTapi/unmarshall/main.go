package main

import (
	"fmt"
	"log"
	"os"
)

// Struct for representation total slice
type Users struct {
	Users []User `json:"users"`
}

// Interanal user representation
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Socila Social `json:"social"`
}

// Social block representation
type Social struct {
	Twitter  string `json:"twitter"`
	Facebook string `json:"meta"`
}

func main() {
	jsonFile, err := os.Open("../users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	fmt.Println("File descriptor sucessfully created")
}
