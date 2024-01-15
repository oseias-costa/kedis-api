package repository

import (
	"database/sql"
	"fmt"
	"main/domain"
	"main/infra/persistence"
)

type UserRepositoryInterface interface {
	CreateUserRepo(user domain.User) (domain.User, error)
	LoginUserRepo(email, password string) (sql.Result, error)
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

	fmt.Printf("res: %v\n", res)

	return user, nil
}

func (*userRepo) LoginUserRepo(email, password string) (sql.Result, error) {
	var res sql.Result
	c := persistence.Connect()
	defer c.Close()
	sql := `SELECT * FROM user WHERE email = ?`

	stmt, err := c.Prepare(sql)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	res, err = stmt.Exec(email, password)
	if err != nil {
		return res, err
	}

	return res, nil

}
