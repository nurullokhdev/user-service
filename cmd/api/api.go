package main

import (
	"net/http"

	"github.com/myselfBZ/users/internal/store"
)

func NewAPI() *API {
	s := store.NewPostgre()
	return &API{
		store: s,
	}
}

type API struct {
	store store.Store
}

// godoc
// @Summary 	mounts all the routes and handlers
func (a *API) mount() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", makeHTTPHandler(a.registerUser))
	mux.HandleFunc("/login", makeHTTPHandler(a.login))
	mux.HandleFunc("DELETE /users/{id}", makeHTTPHandler(a.deleteAccount))
	mux.HandleFunc("PUT /users/{id}", makeHTTPHandler(a.updateAccount))
	return mux
}

func (a *API) run(port string) error {
	return http.ListenAndServe(port, a.mount())
}
