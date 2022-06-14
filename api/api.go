package api

import (
	"github.com/gorilla/mux"
	"golang-clean-arch/modules/user/service"
	"net/http"
)

type Routes struct {
	Users *mux.Router
	Auths *mux.Router
}

//func Init() *mux.Router {
//	root := mux.NewRouter()
//	root.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("hello =))"))
//	}).Methods("GET")
//
//	router := Routes{}
//	api := root.PathPrefix("/api/v1").Subrouter()
//
//	router.Users = api.PathPrefix("/users").Subrouter()
//	router.InitUser()
//	return root
//}

// Handler Root API handler.
type Handler struct {
	//*ErrorHandler
	APIRouter *mux.Router
	root      *mux.Router
}

// NewHandler constructs a new handler.
func NewHandler() *Handler {
	handler := &Handler{}

	root := mux.NewRouter()
	root.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello =))"))
	}).Methods("GET")

	api := root.PathPrefix("/api/v1").Subrouter()
	//api.Use(MattermostAuthorizationRequired)

	api.Handle("{anything:.*}", http.NotFoundHandler())
	api.NotFoundHandler = http.NotFoundHandler()

	handler.APIRouter = api
	handler.root = root

	userService := service.NewUserService()
	NewUserHandler(handler.APIRouter, userService)

	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.root.ServeHTTP(w, r)
}
