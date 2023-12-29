package usecases

import (
	"main/entities"
	"testing"
)

func TestCreateUser(t *testing.T) {
	mockUser := entities.User{
		FirstName: "Oséias",
		LastName:  "Costa",
		Age:       32,
		Password:  "12345",
	}

	expect := "Oséias"

	if mockUser.FirstName != expect {
		t.Errorf("Expected FirstName to be 'Oséias', got %s", mockUser.FirstName)
	}
}
