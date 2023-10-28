package storage

import "honnef.co/go/tools/config"

//Instance of storage
type Storage struct {
	config *Config
}

func New(config *Config) *Storage {
	return &Storage{
		config: config
	}
}	
