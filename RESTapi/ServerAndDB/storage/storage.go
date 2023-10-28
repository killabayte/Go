package storage

import "honnef.co/go/tools/config"

type Storage struct {
	config *Config
}

func New(config *Config) *Storage {
	return &Storage{
		config: config
	}
}	
