package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bruno-holanda15/api-rest-challenge-rocketseat/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func NewHTTPHandler() http.Handler {
	r := chi.NewRouter()

	dbLocal := app.NewAppStorage()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Healthy!"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Post("/users", InserUser(dbLocal))
		r.Get("/users/{id}", GetUser(dbLocal))
	})

	return r
}

type ResponseInsertUser struct {
	ID string `json:"id"`
}

func InserUser(db *app.AppStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user app.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("unprocessable entity"))
		
			return
		}


		id := db.Insert(user)
		response := ResponseInsertUser{ID: id.String()}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error encoding response"))
		
			return
		}
		
		w.WriteHeader(http.StatusCreated)
		fmt.Println(db)
	}
}

func GetUser(db *app.AppStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("user not found"))
			return
		}
		
		uuid, err := uuid.Parse(id)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("id provided not uuid valid"))
			return
		}

		user, err := db.FindById(app.ID(uuid))
		if err != nil {
			if errors.Is(err, app.ErrUserNotFound) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("user not found"))
			return
			}
	
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("id provided not uuid valid"))
			return
		}

		if err := json.NewEncoder(w).Encode(user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error encoding response"))
		
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}