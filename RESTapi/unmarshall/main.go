package main

type Users struct {
	Users []User
}

type User struct {
	Name   string
	Type   string
	Age    int
	Socila Social
}

type Social struct {
	Twitter string
	Meta    string
}
