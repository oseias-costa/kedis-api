package usecases

import (
	"errors"
	"fmt"
	"main/domain"
	"main/infra/repository"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCaseInterface interface {
	CreateUserUseCase(user domain.User) (domain.User, error)
	LoginUseCase(l domain.Login) (string, error)
	GetUserById(token string) (domain.User, error)
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

func (*userUseCase) LoginUseCase(l domain.Login) (string, error) {
	res, err := repo.LoginUserRepo(l)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(l.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": res.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return "", err
	// }

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	sendToken := fmt.Sprintf(`"token STRING": "%s"`, tokenString)

	return sendToken, nil
}

func (*userUseCase) GetUserById(token string) (domain.User, error) {
	var user domain.User
	tokenTest := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDgxNjkwNjcsInN1YiI6ImUzYWUwZGJkLTU3NDUtNGY4Ny1hN2E1LWU4ZjJhZDU3NjMzZiJ9.i6LP1EWmPHz9dPrh1Rlxfbf6wgSbBXWQEjwaA9euNliylsDYCgA3s_0B-VF56L3WJOQzsNJ6EUAzYo3zNW_gzg"

	fmt.Println(tokenTest)
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
