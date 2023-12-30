package repository

import (
	"errors"
	"main/entities"
	"main/usecases"
)

var userUseCases usecases.UserUseCase

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
		FirstName: "Rangel",
		LastName:  "Amilton Costa",
		Age:       35,
		Password:  "54312",
	},
}

type UserRepository interface {
	GetUser(id string) (entities.User, error)
}

type userRepository struct{}

func NewUserRepository(usecase usecases.UserUseCase) UserRepository {
	userUseCases = usecase
	return &userRepository{}
}

func (*userRepository) GetUser(id string) (entities.User, error) {
	for _, user := range users {
		if user.ID == id {
			verifyUser, err := userUseCases.GetUserUseCase(user)
			if err != nil {
				return user, err
			}
			return *verifyUser, nil
		}
	}

	return entities.User{}, errors.New("User not exist")
}
