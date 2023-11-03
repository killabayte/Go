package storage

import (
	"fmt"

	"github.com/killabayte/Go/RESTapi/ServerAndDBV2/internal/app/models"
)

type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle string = "articles"
)

func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, context) VALUES ($1, $2, $3) RETURNING id", tableArticle)
	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Context).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}
