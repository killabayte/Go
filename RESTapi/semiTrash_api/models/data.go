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
}
