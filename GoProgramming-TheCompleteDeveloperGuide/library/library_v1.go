package main

import (
	"fmt"
	"time"
)

type Member_v1 struct {
	name string
}
type Book struct {
	name     string
	checkOut string
	checkIn  string
}
type Library_v1 struct {
	members []Member_v1
	books   []Book
}

func checkOut(b *Book, l *Library_v1, m Member_v1) {
	b.checkOut = time.Now().Format(time.RFC3339)
	fmt.Println(m.name, "Is Checking Out next book:", b.name, "at the time:", b.checkOut)
}

func checkIn(b *Book, l *Library_v1, m Member_v1) {
	b.checkIn = time.Now().Format(time.RFC3339)
	fmt.Println(m.name, "Is Checking In next book:", b.name, "at the time:", b.checkIn)
}

func printLibraryInfo(l *Library_v1) {
	for i := 0; i < len(l.books); i++ {
		book := l.books[i]
		fmt.Printf("Book: %s\n", book)
	}
	for i := 0; i < len(l.members); i++ {
		fmt.Println("Member:", l.members[i].name)
	}
}

func main_v1() {
	newLibrary := Library_v1{}
	book1 := Book{name: "Harry Potter and philosopher's stone"}
	book2 := Book{name: "Harry Potter and the Chamber of Secrets"}
	book3 := Book{name: "Harry Potter and the Prisoner of Azkaban"}
	book4 := Book{name: "Harry Potter and the Goblet of Fire"}
	book5 := Book{name: "Harry Potter and the Order of the Phoenix"}
	book6 := Book{name: "Harry Potter and the Half-Blood Prince"}
	book7 := Book{name: "Harry Potter and the Deathly Hallows"}

	newLibrary.books = append(newLibrary.books, book1, book2, book3, book4, book5, book6, book7)
	fmt.Println(newLibrary.books)
	newLibrary.members = append(newLibrary.members, Member_v1{"Harry"}, Member_v1{"Ron"}, Member_v1{"Hermione"})
	checkOut(&newLibrary.books[1], &newLibrary, newLibrary.members[0])
	checkIn(&newLibrary.books[0], &newLibrary, newLibrary.members[1])
	printLibraryInfo(&newLibrary)
}
