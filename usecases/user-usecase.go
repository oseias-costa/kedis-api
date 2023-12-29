package usecases

import (
	"errors"
	"main/entities"
)

var users = []entities.User{
	{
		ID:        "1",
		FirstName: "Oséias",
		LastName:  "Costa",
		Age:       32,
		Password:  "12345",
	},
	{
		ID:        "2",
		FirstName: "Outro",
		LastName:  "Costa",
		Age:       34,
		Password:  "34567",
	},
}

type UserUseCase interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetUser(id string) (*entities.User, error)
}

type userUseCase struct{}

func NewUserUseCase() UserUseCase {
	return &userUseCase{}
}

func (*userUseCase) CreateUser(user *entities.User) (*entities.User, error) {
	if user.FirstName == "" && user.LastName == "" && user.Age == 0 && user.Password == "" {
		err := errors.New("All fields are required")
		return user, err
	}
	return &entities.User{
		ID:        "1",
		FirstName: "Oséias",
		LastName:  "Costa",
		Age:       32,
		Password:  "12345",
	}, nil
}

func (*userUseCase) GetUser(id string) (*entities.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return &entities.User{}, errors.New("User not found")
}
