package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"main/domain"
	"main/usecases"
	"net/http"
)

var userUseCase = usecases.NewUserUseCase()

type UserControlerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

type userController struct{}

func NewUserController() UserControlerInterface {
	return &userController{}
}

func (*userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

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
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(i))
}

func (*userController) GetUser(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	user, err := userUseCase.GetUser(string(b))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprint(err)))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
