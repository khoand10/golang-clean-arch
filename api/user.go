package api

import (
	"github.com/gorilla/mux"
	"golang-clean-arch/modules/user/service"
	"golang-clean-arch/utils/web"
	"net/http"
)

type UserHandler struct {
	//*ErrorHandler
	userService service.UserService
	//log             bot.Logger
	//permissions     *service.PermissionsService
}

func NewUserHandler(router *mux.Router, userService service.UserServiceImpl) *UserHandler {
	handler := &UserHandler{
		userService: userService,
	}
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", handler.GetUsers).Methods(http.MethodGet)

	return handler
}

//func (router *Routes) InitUser() {
//	router.Users.HandleFunc("", GetUsers).Methods(http.MethodGet)
//}

type (
	User struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	user, err := h.userService.GetUsers()
	if err != nil {
		web.ReturnJSON(w, nil, http.StatusBadRequest)
	}
	web.ReturnJSON(w, user, http.StatusOK)
}
