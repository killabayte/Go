package api

import (
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *API {
	return &API{
		config: config,
	}
}

func (a *API) Start() error {
	return nil
}
