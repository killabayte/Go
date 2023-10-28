package storage

import "honnef.co/go/tools/config"

//Instance of storage
type Storage struct {
	config *Config
	//Database file descriptor
	db *sql.DB
}

//Storage constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config
	}
}	

//Open connection method
func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DataBaseURL)
	if err != nil {
		return err
	}
	return nil
}

//Close connection method
func (s *Storage) Close() error {

}