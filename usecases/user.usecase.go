package usecases

import (
	"errors"
	"fmt"
	"main/domain"
	"main/infra/repository"
	"net/mail"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCaseInterface interface {
	CreateUserUseCase(user domain.User) (domain.User, error)
}

type userUseCase struct{}

var repo = repository.NewUserRepository()

func NewUserUseCase() UserUseCaseInterface {
	return &userUseCase{}
}

func (*userUseCase) CreateUserUseCase(user domain.User) (domain.User, error) {
	id, errUiid := uuid.NewRandom()
	if errUiid != nil {
		return user, errUiid
	}
	user.ID = id.String()

	_, err := verifyUser(user)
	if err != nil {
		return user, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return user, err
	}
	user.Password = string(hash)

	u, err := repo.CreateUserRepo(user)
	if err != nil {
		return user, err
	}

	fmt.Printf("test %v", u)

	return user, nil
}

func verifyUser(user domain.User) (bool, error) {
	if user.ID == "" {
		return false, errors.New("ID is required")
	}

	if user.FirstName == "" {
		return false, errors.New("Name is required")
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return false, errors.New("Email is not valid")
	}

	return true, nil
}
