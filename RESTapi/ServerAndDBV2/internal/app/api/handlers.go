package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	if err != nil {
		api.logger.Info("Error while decoding request body: ")
		msg := Message{
			StatusCode: 400,
			Message:    "Bad request. Check your input data.",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Error while Articles.Create(): ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to access the database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(a)
}

func (api *API) GetArticleById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Get Article by ID GET /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		api.logger.Info("Error while converting {id} parameter: ", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Bad request. Check your input data.",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	article, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Error while Articles.FindArticleById(): ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to access the database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Article not found")
		msg := Message{
			StatusCode: 404,
			Message:    "Article not found in database.",
			IsError:    true,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(article)
}

func (api *API) DeleteArticleById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Delete Article by ID DELETE /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		api.logger.Info("Error while converting {id} parameter: ", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Bad request. Check your input data.",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	_, ok, err := api.storage.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Error while Articles.FindArticleById(): ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to access the database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Article not found")
		msg := Message{
			StatusCode: 404,
			Message:    "Article not found in database.",
			IsError:    true,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return
	}
	_, err = api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Error while deleting article from databease by method Articles.DeleteById(): ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to access the database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Article with id %d was deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(w).Encode(msg)
}

func (api *API) PostUserRegister(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	api.logger.Info("Create user POST /api/v1/users/register")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Error while decoding request body: ")
		msg := Message{
			StatusCode: 400,
			Message:    "Bad request. Check your input data.",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Error while User.FindByLogin(): ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to access the database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if ok {
		api.logger.Info("User already exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User already exists.",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	u, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("Error while User.Create(): ", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to access the database. Try again later.",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(201)
	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User with id %d was created.", u.ID),
		IsError:    false,
	}
	json.NewEncoder(w).Encode(msg)
}
