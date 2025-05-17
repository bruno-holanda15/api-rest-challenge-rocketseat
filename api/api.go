package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bruno-holanda15/api-rest-challenge-rocketseat/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHTTPHandler() http.Handler {
	r := chi.NewRouter()

	dbLocal := app.NewAppStorage()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Healthy!"))
	})

	r.Post("/api/users", InserUser(dbLocal))

	return r
}

func InserUser(db *app.AppStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user app.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("unprocessable entity"))
		
			return
		}


		db.Insert(user)
		w.WriteHeader(http.StatusCreated)
		fmt.Println(db)
	}
}
