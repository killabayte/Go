package api

import (
	"github.com/killabayte/Go/RESTapi/JWTAuth/storage"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
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
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")
}

// Configuration for the store field
func (a *API) configureStoreField() error {
	storage := storage.New(a.config.Storage)
	//Try to establish connection to the database, if not possible - return error
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
