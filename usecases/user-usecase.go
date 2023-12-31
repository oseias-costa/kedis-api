package usecases

import (
	"errors"
	"main/entities"
)

type UserUseCase interface {
	CreateUserUseCase(user *entities.User) (entities.User, error)
	GetUserUseCase(user entities.User) (*entities.User, error)
}

type userUseCase struct{}

func NewUserUseCase() UserUseCase {
	return &userUseCase{}
}

func (*userUseCase) CreateUserUseCase(user *entities.User) (entities.User, error) {
	if user.FirstName == "" && user.LastName == "" && user.Age == 0 && user.Password == "" {
		err := errors.New("All fields are required")
		return *user, err
	}
	return entities.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Password:  user.Password,
	}, nil
}

func (*userUseCase) GetUserUseCase(user entities.User) (*entities.User, error) {
	if user.FirstName != "" && user.LastName != "" && user.ID != "" {
		return &user, nil
	}

	return &user, errors.New("user is not correct")
}
