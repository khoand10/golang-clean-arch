package api

import (
	"golang-clean-arch/utils/web"
	"net/http"
)

func (router *Routes) InitUser() {
	router.Users.HandleFunc("", GetUsers).Methods("GET")
}

type (
	User struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	user := &User{
		FirstName: "Khoa",
		LastName:  "Nguyen",
		Email:     "khoand@gmail.com",
	}
	web.ReturnJSON(w, user, http.StatusOK)
}
