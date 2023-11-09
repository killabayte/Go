package api

import (
	"encoding/json"
	"net/http"

	"github.com/killabayte/Go/RESTapi/ServerAndDBV2/internal/app/models"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Get All Articles GET /api/v1/articles")
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info("Error while Articles.SelectAll(): ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to access the database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(articles)
}

func (api *API) PostArticle(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Post Article POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)

}

func (api *API) GetArticleById(w http.ResponseWriter, r *http.Request)    {}
func (api *API) DeleteArticleById(w http.ResponseWriter, r *http.Request) {}
func (api *API) PostUserRegister(w http.ResponseWriter, r *http.Request)  {}
