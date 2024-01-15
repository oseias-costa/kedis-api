package usecases

// import (
// 	"errors"
// 	"main/entities"
// )

// type UserUseCase interface {
// 	CreateUserUseCase(user entities.User) (entities.User, error)
// 	GetUserUseCase(user entities.User) (*entities.User, error)
// 	GetAllUsers(users []entities.User) ([]entities.User, error)
// 	UpdateUser(user entities.User) (entities.User, error)
// 	DeleteUser(u entities.User, id string) bool
// }

// type userUseCase struct{}

// func NewUserUseCase() UserUseCase {
// 	return &userUseCase{}
// }

// func (*userUseCase) CreateUserUseCase(user entities.User) (entities.User, error) {
// 	if user.FirstName == "" && user.LastName == "" && user.Password == "" {
// 		err := errors.New("All fields are required")
// 		return user, err
// 	}

// 	return entities.User{
// 		FirstName: user.FirstName,
// 		LastName:  user.LastName,
// 		Email:     user.Email,
// 		Password:  user.Password,
// 	}, nil
// }

// func (*userUseCase) GetUserUseCase(user entities.User) (*entities.User, error) {
// 	if user.FirstName != "" && user.LastName != "" && user.ID != "" {
// 		return &user, nil
// 	}

// 	return &user, errors.New("user is not correct")
// }

// func (*userUseCase) GetAllUsers(users []entities.User) ([]entities.User, error) {
// 	return users, nil
// }

// func (*userUseCase) UpdateUser(user entities.User) (entities.User, error) {
// 	if user.FirstName == "" && user.LastName == "" && user.Password == "" {
// 		err := errors.New("All fields are required")
// 		return user, err
// 	}
// 	return entities.User{
// 		FirstName: user.FirstName,
// 		LastName:  user.LastName,
// 		Email:     user.Email,
// 		Password:  user.Password,
// 	}, nil
// }

// func (*userUseCase) DeleteUser(u entities.User, id string) bool {
// 	if u.ID == id {
// 		return true
// 	}
// 	return false
// }
