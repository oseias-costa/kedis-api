package repository

import (
	"main/entities"
	"main/usecases"
	"testing"
)

var usersMock = []entities.User{
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

var usecase = usecases.NewUserUseCase()
var userRepo = NewUserRepository(usecase)

func TestUserRepository(t *testing.T) {
	mockUser := entities.User{
		ID:        "3",
		FirstName: "Milton",
		LastName:  "P Costa",
		Age:       69,
		Password:  "224422",
	}

	t.Run("Should create user", func(t *testing.T) {
		newUser, _ := userRepo.CreateUser(mockUser)
		expect := "Milton"

		if newUser.FirstName != expect {
			t.Errorf("Expect %s, but got %s", expect, newUser.FirstName)
		}
	})

	t.Run("Shoul create user", func(t *testing.T) {

	})

	t.Run("Shoul create user", func(t *testing.T) {

	})

	t.Run("Shoul create user", func(t *testing.T) {

	})
}
