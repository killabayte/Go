package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Instance of storage
type Storage struct {
	config *Config
	//Database file descriptor
	db                *sql.DB
	userRepository    *UserRepository
	articleRepository *ArticleRepository
}

// Storage constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open connection method
func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DataBaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	log.Println("Database connection established successfully")
	return nil
}

// Close connection method
func (s *Storage) Close() error {
	s.db.Close()
	return nil
}

func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{}
	return nil
}

func (s *Storage) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{}
	return nil
}
