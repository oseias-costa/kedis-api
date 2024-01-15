package repository

import (
	"database/sql"
	"fmt"
	"main/domain"
	"main/infra/persistence"
)

type UserRepositoryInterface interface {
	CreateUserRepo(user domain.User) (domain.User, error)
	LoginUserRepo(l domain.Login) (sql.Result, error)
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

	fmt.Println(res)

	return user, nil
}

func (*userRepo) LoginUserRepo(l domain.Login) (sql.Result, error) {
	var res sql.Result
	c := persistence.Connect()
	defer c.Close()
	sql := `SELECT * FROM user WHERE email = ?`

	r, err := c.Query(sql)
	if err != nil {
		return res, err
	}
	defer r.Close()

	// r, err = stmt.Exec(l.Email)
	// if err != nil {
	// 	return res, err
	// }

	var users []domain.User

	for r.Next() {
		var user domain.User
		err := r.Scan(
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.ID,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	fmt.Println("esse é o retorno", users)

	return res, nil

}
