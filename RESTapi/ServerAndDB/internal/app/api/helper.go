package api

import (
	"net/http"

	"github.com/killabayte/Go/RESTapi/ServerAndDB/storage"
	"github.com/sirupsen/logrus"
)

func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! From REST api"))
	})
}

// Configuration for the store field
func (a *API) configureStoreField() error {
	store := storage.New(a.config.Storage)
	//Try to establish connection to the database, if not possible - return error
	if err := store.Open(); err != nil {
		return err
	}
	a.store = store
	return nil
}
