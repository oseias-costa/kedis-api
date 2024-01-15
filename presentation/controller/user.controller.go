package controller

import (
	"encoding/json"
	"fmt"
	"main/domain"
	"main/usecases"
	"net/http"
)

var userUseCase = usecases.NewUserUseCase()

type UserControlerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}

type userController struct{}

func NewUserController() UserControlerInterface {
	return &userController{}
}

func (*userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "user not decoded"}`))
	}

	u, err := userUseCase.CreateUserUseCase(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprint(err)))
	}
	fmt.Println("return controller", u)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (*userController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var login domain.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "email or password not valid"}`))
	}

	i, newError := userUseCase.LoginUseCase(login)
	if newError != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprint(newError)))
	}

	fmt.Println("return controller", i)
}
