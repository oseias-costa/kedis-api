package repository

import (
	"encoding/json"
	"fmt"
	"main/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

var useCases = usecases.NewUserUseCase()

type UserRepository interface {
	GetUserByID(w http.ResponseWriter, r *http.Request)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (*userRepository) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println("id is:", id)
	user, err := useCases.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not exists"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
