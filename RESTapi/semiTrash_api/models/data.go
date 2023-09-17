package models

var db []Book

type Book struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Author         Author `json:"author"`
	YearPublishoed int    `json:"year_published"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {
	book1 := Book{
		ID:             1,
		Title:          "Lord of the rings. Vol.1",
		YearPublishoed: 1937,
		Author: Author{
			Name:     "J.R.R",
			LastName: "Tolkien",
			BornYear: 1892,
		},
	}
}
