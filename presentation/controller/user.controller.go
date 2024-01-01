package controller

import (
	"encoding/json"
	"main/entities"
	"main/infra/repository"
	"net/http"

	"github.com/gorilla/mux"
)

var userRepository repository.UserRepository

type UserController interface {
	GetUserByID(w http.ResponseWriter, r *http.Request)
	CreateNewUser(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
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

func (*userController) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"create user error on controller"}`))
	}

	result, errResult := userRepository.CreateUser(user)
	if errResult != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"create user error on controller to repository"}`))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (*userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userRepository.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"the list is empty"}`))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (*userController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"update user error decode json"}`))
	}

	result, errResult := userRepository.UpdateUser(user)
	if errResult != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"update user error repository"}`))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (*userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	isDeleted := userRepository.DeleteUser(id)
	if !isDeleted {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"user not find"}`))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"response":"user deleted"}`))
}
