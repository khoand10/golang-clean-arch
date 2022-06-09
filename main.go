package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

const (
	PORT = 8080
)

type (
	User struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func main() {
	root := mux.NewRouter()
	api := root.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}).Methods(http.MethodGet)
	//authRouter := api.PathPrefix("/auth").Subrouter()
	usersRouter := api.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", GetUsers)
	log.Fatal(http.ListenAndServe(":8080", root))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	user := &User{
		FirstName: "Khoa",
		LastName:  "Nguyen",
		Email:     "khoand@gmail.com",
	}
	ReturnJSON(w, user, http.StatusOK)
}

// message as JSON in a response with the provided code.
func HandleErrorWithCode(w http.ResponseWriter, code int, publicErrorMsg string, internalErr error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	details := ""
	if internalErr != nil {
		details = internalErr.Error()
	}

	logrus.Errorf("public error message: %v; internal details: %v", publicErrorMsg, details)

	responseMsg, _ := json.Marshal(struct {
		Error string `json:"error"` // A public facing message providing details about the error.
	}{
		Error: publicErrorMsg,
	})
	_, _ = w.Write(responseMsg)
}

// ReturnJSON writes the given pointerToObject as json with the provided httpStatus
func ReturnJSON(w http.ResponseWriter, pointerToObject interface{}, httpStatus int) {
	jsonBytes, err := json.Marshal(pointerToObject)
	if err != nil {
		logrus.Warnf("Unable to marshall JSON. Error details: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	if _, err = w.Write(jsonBytes); err != nil {
		logrus.Warnf("Unable to write to http.ResponseWriter. Error details: %s", err.Error())
		return
	}
}
