package http

import (
	"encoding/json"
	"main/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

var useCases = usecases.NewUserUseCase()

type UserHTPP interface {
	GetUserByID(w http.ResponseWriter, r *http.Request)
}

type userHttp struct{}

func NewUserHttp() UserHTPP {
	return &userHttp{}
}

func (*userHttp) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := useCases.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not exists"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
