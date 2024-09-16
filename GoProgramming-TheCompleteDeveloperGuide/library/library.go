package main

import (
	"fmt"
	"time"
)

type Member struct {
	name string
}
type Book struct {
	name     string
	checkOut time.Time
	checkIn  time.Time
}
type Library struct {
	members []Member
	books   []Book
}

func checkOut(b *Book, l *Library, m Member) {
	b.checkOut = time.Now()
	fmt.Println(m.name, "Is Checking Out next book:", b.name, "at the time:", b.checkOut)
}

func checkIn(b *Book, l *Library, m Member) {
	b.checkIn = time.Now()
	fmt.Println(m.name, "Is Checking In next book:", b.name, "at the time:", b.checkIn)
}

func printLibraryInfo(l *Library) {
	for i := 0; i < len(l.books); i++ {
		book := l.books[i]
		fmt.Printf("Book: %s\n", book.name)
		if !book.checkOut.IsZero() {
			fmt.Printf("  Check Out: %s\n", book.checkOut.Format(time.RFC3339))
		} else {
			fmt.Println("  Check Out: Not checked out")
		}
		if !book.checkIn.IsZero() {
			fmt.Printf("  Check In: %s\n", book.checkIn.Format(time.RFC3339))
		} else {
			fmt.Println("  Check In: Not checked in")
		}
	}
	for i := 0; i < len(l.members); i++ {
		fmt.Println("Member:", l.members[i].name)
	}
}

func main() {
	newLibrary := Library{}
	book1 := Book{name: "Harry Potter and philosopher's stone"}
	book2 := Book{name: "Harry Potter and the Chamber of Secrets"}
	book3 := Book{name: "Harry Potter and the Prisoner of Azkaban"}
	book4 := Book{name: "Harry Potter and the Goblet of Fire"}
	book5 := Book{name: "Harry Potter and the Order of the Phoenix"}
	book6 := Book{name: "Harry Potter and the Half-Blood Prince"}
	book7 := Book{name: "Harry Potter and the Deathly Hallows"}

	newLibrary.books = append(newLibrary.books, book1, book2, book3, book4, book5, book6, book7)
	fmt.Println(newLibrary.books)
	newLibrary.members = append(newLibrary.members, Member{"Harry"}, Member{"Ron"}, Member{"Hermione"})
	checkOut(&newLibrary.books[1], &newLibrary, newLibrary.members[0])
	checkIn(&newLibrary.books[0], &newLibrary, newLibrary.members[1])
	printLibraryInfo(&newLibrary)
}
