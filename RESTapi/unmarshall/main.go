package main

// Struct for representation total slice
type Users struct {
	Users []User
}

// Interanal user representation
type User struct {
	Name   string
	Type   string
	Age    int
	Socila Social
}

// Social block representation
type Social struct {
	Twitter string
	Meta    string
}
