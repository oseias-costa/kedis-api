package repository

import (
	"fmt"
	"main/domain"
	"main/infra/persistence"
)

type UserRepositoryInterface interface {
	CreateUserRepo(user domain.User) (domain.User, error)
	LoginUserRepo(l domain.Login) (domain.User, error)
	GetUserRepo(id string) (domain.UserResponse, error)
}

type userRepo struct{}

func NewUserRepository() UserRepositoryInterface {
	return &userRepo{}
}

func (*userRepo) CreateUserRepo(user domain.User) (domain.User, error) {
	c := persistence.Connect()
	sql := `INSERT INTO user (id, firstName, lastName, password, email) values (?, ?, ?, ?, ?)`
	defer c.Close()

	stmt, err := c.Prepare(sql)
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.ID, user.FirstName, user.LastName, user.Password, user.Email)
	if err != nil {
		return user, err
	}

	fmt.Println("res createUser", res)
	return user, nil
}

func (*userRepo) LoginUserRepo(l domain.Login) (domain.User, error) {
	var user domain.User

	c := persistence.Connect()
	sql := `SELECT * FROM user WHERE email = ?`

	r, err := c.Query(sql, l.Email)
	if err != nil {
		return user, err
	}

	for r.Next() {
		err := r.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return user, err
		}
	}

	defer c.Close()
	defer r.Close()

	return user, nil
}

func (*userRepo) GetUserRepo(id string) (domain.UserResponse, error) {
	var user domain.UserResponse
	c := persistence.Connect()

	r, err := c.Query(`SELECT * FROM user WHERE id = ?`, id)
	if err != nil {
		return user, err
	}

	for r.Next() {
		err := r.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return user, err
		}
	}
	defer c.Close()
	defer r.Close()

	return user, nil
}
