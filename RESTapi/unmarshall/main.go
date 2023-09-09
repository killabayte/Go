package main

// Struct for representation total slice
type Users struct {
	Users []User
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
	Twitter string
	Meta    string
}
