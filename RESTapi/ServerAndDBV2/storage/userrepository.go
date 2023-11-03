package storage

import "github.com/killabayte/Go/RESTapi/ServerAndDBV2/internal/app/models"

type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	return nil, nil
}

func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	return nil, false, nil
}

func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	return nil, nil
}
