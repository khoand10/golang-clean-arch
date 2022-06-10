package api

import (
	"github.com/gorilla/mux"
)

type Routes struct {
	Users *mux.Router
	Auths *mux.Router
}

func Init() *mux.Router {
	root := mux.NewRouter()
	router := Routes{}
	api := root.PathPrefix("/api/v1").Subrouter()

	router.Users = api.PathPrefix("/users").Subrouter()
	router.InitUser()
	return root
}
