package controller

import (
	"encoding/json"
	"main/infra/repository"
	"net/http"

	"github.com/gorilla/mux"
)

var userRepository repository.UserRepository

type UserController interface {
	GetUserByID(w http.ResponseWriter, r *http.Request)
}

type userController struct{}

func NewUserController(repository repository.UserRepository) UserController {
	userRepository = repository
	return &userController{}
}

func (*userController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := userRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"user not exist"}`))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
