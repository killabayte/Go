package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/killabayte/Go/RESTapi/ServerAndDB/storage"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//Add storage field
	store *storage.Storage
}

func New(config *Config) *API {

	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("Starting API server at port ", api.config.BindAddr)
	//Configure router
	api.configureRouterField()
	//Configure storage
	api.configureStoreField()

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
