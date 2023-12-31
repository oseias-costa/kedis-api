package repository

import (
	"errors"
	"fmt"
	"main/entities"
	"main/usecases"
	"math/rand"
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
	CreateUser(user entities.User) (entities.User, error)
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

func (*userRepository) CreateUser(user entities.User) (entities.User, error) {
	user.ID = fmt.Sprint(rand.Intn(10000))
	newUser, err := userUseCases.CreateUserUseCase(&user)
	if err != nil {
		return user, errors.New("Error at create user")
	}
	users = append(users, newUser)
	return newUser, nil
}
