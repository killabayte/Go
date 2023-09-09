package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	Facebook string `json:"facebook"`
}

func PrintUser(u *User) {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Social: Twitter - %s and Facebook - %s\n", u.Socila.Twitter, u.Socila.Facebook)
}

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	fmt.Println("File descriptor sucessfully created")

	var users Users
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteValue, &users)
	fmt.Println(users)
}
