package usecases

import (
	"main/entities"
	"testing"
)

func TestUserUseCase(t *testing.T) {
	var usecase = NewUserUseCase()
	mockUser := entities.User{
		ID:        "1",
		FirstName: "Oséias",
		LastName:  "Costa",
		Email:     "oseiasc2@gmail.com",
		Password:  "12345",
	}

	t.Run("Shold create a user", func(t *testing.T) {
		t.Helper()
		newuser, _ := usecase.CreateUserUseCase(mockUser)
		expect := "Oséias"

		if newuser.FirstName != expect {
			t.Errorf("Expected FirstName to be 'Oséias', got %s", mockUser.FirstName)
		}
	})

	t.Run("Shold Get user by id", func(t *testing.T) {
		t.Helper()
		getUser, _ := usecase.GetUserUseCase(mockUser)
		expect := "oseiasc2@gmail.com"

		if getUser.Email != expect {
			t.Errorf("Expect Age %v, but got %v", getUser.Age, expect)
		}
	})

	t.Run("Shoud update user", func(t *testing.T) {
		t.Helper()
		updateUser, _ := usecase.UpdateUser(mockUser)
		expect := "12345"

		if updateUser.Password != expect {
			t.Errorf("Expect password %s, but got %s", expect, updateUser.Password)
		}
	})

	t.Run("Shoud delete user", func(t *testing.T) {
		t.Helper()
		deleteUser := usecase.DeleteUser(mockUser, "1")

		if !deleteUser {
			t.Error("The user is not ok to delete")
		}
	})
}
