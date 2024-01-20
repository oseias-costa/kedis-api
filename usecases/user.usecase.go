package usecases

import (
	"errors"
	"fmt"
	"main/domain"
	"main/infra/repository"
	"net/mail"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Id string
}

type UserUseCaseInterface interface {
	CreateUserUseCase(user domain.User) (domain.User, error)
	LoginUseCase(l domain.Login) (string, error)
	GetUser(id string) (domain.UserResponse, error)
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
	token, err := GenerateToken(res.ID)
	if err != nil {
		return "", err
	}

	sendToken := fmt.Sprintf(`{"token": "%s"}`, token)

	return sendToken, nil
}

func (*userUseCase) GetUser(id string) (domain.UserResponse, error) {
	u, err := repo.GetUserRepo(id)
	if err != nil {
		return u, err
	}

	return u, nil
}

func GenerateToken(u string) (string, error) {
	now := time.Now()
	expires := now.Add(time.Hour * 168).Unix()
	claims := jwt.MapClaims{
		"sub":    u,
		"expire": expires,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	expValue := claims["expire"].(float64)
	expires := int64(expValue)
	if time.Now().Unix() > expires {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
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
