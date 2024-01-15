package repository

// import (
// 	"errors"
// 	"fmt"
// 	"main/entities"
// 	"main/infra/persistence"
// 	"main/usecases"
// 	// "slices"
// )

// var userUseCases usecases.UserUseCase

// var users = []entities.User{
// 	{
// 		ID:        "1",
// 		FirstName: "Oséias",
// 		LastName:  "Costa",
// 		Email:     "oseiasc2@gmail.com",
// 		Password:  "12345",
// 	},
// 	{
// 		ID:        "2",
// 		FirstName: "Rangel",
// 		LastName:  "Amilton Costa",
// 		Email:     "rangel@email.com",
// 		Password:  "54312",
// 	},
// }

// type UserRepository interface {
// 	GetUser(id string) (entities.User, error)
// 	CreateUser(user entities.User) (entities.User, error)
// 	GetAllUsers() ([]entities.User, error)
// 	UpdateUser(u entities.User) (entities.User, error)
// 	DeleteUser(id string) bool
// }

// type userRepository struct{}

// func NewUserRepository(usecase usecases.UserUseCase) UserRepository {
// 	userUseCases = usecase
// 	return &userRepository{}
// }

// func (*userRepository) GetUser(id string) (entities.User, error) {
// 	for _, user := range users {
// 		if user.ID == id {
// 			verifyUser, err := userUseCases.GetUserUseCase(user)
// 			if err != nil {
// 				return user, err
// 			}
// 			return *verifyUser, nil
// 		}
// 	}

// 	return entities.User{}, errors.New("User not exist")
// }

// func (*userRepository) CreateUser(user entities.User) (entities.User, error) {
// 	newUser, err := userUseCases.CreateUserUseCase(user)
// 	if err != nil {
// 		return user, errors.New("Error at create user")
// 	}

// 	c := persistence.Connect()
// 	sql := `INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)`

// 	smtp, err := c.Prepare(sql)
// 	if err != nil {
// 		return user, err
// 	}

// 	res, err := smtp.Exec(newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password)
// 	if err != nil {
// 		return user, err
// 	}

// 	fmt.Printf("res: %v\n", res)
// 	defer smtp.Close()
// 	defer c.Close()

// 	return newUser, nil
// }

// func (*userRepository) GetAllUsers() ([]entities.User, error) {
// 	allUsers, err := userUseCases.GetAllUsers(users)
// 	if err != nil {
// 		return users, errors.New("User not find")
// 	}

// 	return allUsers, nil
// }

// func (*userRepository) UpdateUser(u entities.User) (entities.User, error) {
// 	for key, user := range users {
// 		if user.ID == u.ID {
// 			updateUser, err := userUseCases.UpdateUser(u)
// 			if err != nil {
// 				return user, errors.New("user not update")
// 			}

// 			users[key] = updateUser
// 			return user, nil
// 		}
// 	}
// 	return u, errors.New("user not found")
// }

// func (*userRepository) DeleteUser(id string) bool {
// 	for _, user := range users {
// 		idDelete := userUseCases.DeleteUser(user, id)
// 		if idDelete {
// 			// users = slices.Delete(users, key, key+1)
// 			return true
// 		}
// 	}
// 	return false
// }
