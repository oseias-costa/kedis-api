package http

import (
	"main/entities"
	"net/http"
)

type UserHTPP interface { 
   GetUserByID (w http.ResponseWriter, r *http.Request) (*entities.User, error) 
}

type userHttp struct {}

func NewUserHttp() UserHTPP {
  return &userHttp{}
}

func (*userHttp) GetUserByID(w http.ResponseWriter, r *http.Request) (*entities.User, error) {
  
}
  