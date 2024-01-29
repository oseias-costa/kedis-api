package repository

import (
	"errors"
	"fmt"
	"main/domain"
	"main/infra/persistence"
)

type UserRepositoryInterface interface {
	CreateUserRepo(user domain.User) (domain.User, error)
	LoginUserRepo(l domain.Login) (domain.User, error)
	GetUserRepo(id string) (domain.UserResponse, error)
	SendPasswordRecovery(id, email, code string) (bool, error)
	VerifyCodeRepository(email string) (string, error)
	UpdatePasswordRepository(email, newPassword string) (bool, error)
	EmailIsValid(email string) error
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

func (*userRepo) SendPasswordRecovery(id, email, code string) (bool, error) {
	c := persistence.Connect()
	stmt, err := c.Prepare(`INSERT INTO recoveryPassword (id, email, code) VALUES (?,?,?)`)
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(id, email, code)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (*userRepo) VerifyCodeRepository(email string) (string, error) {
	c := persistence.Connect()

	res, err := c.Query("SELECT * FROM recoveryPassword WHERE email = ?", email)
	if err != nil {
		return "", err
	}

	var recovery domain.RecoveryPassword

	for res.Next() {
		err := res.Scan(
			&recovery.Id,
			&recovery.Email,
			&recovery.Code,
		)
		if err != nil {
			return "", err
		}
	}
	return recovery.Code, nil
}

func (*userRepo) UpdatePasswordRepository(email, newPassword string) (bool, error) {
	c := persistence.Connect()

	_, err := c.Prepare(`UPDATE user SET password = ? WHERE email = ?`)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (*userRepo) EmailIsValid(email string) error {
	c := persistence.Connect()

	res, err := c.Query(`SELECT * FROM user WHERE email = ?`, email)
	if err != nil {
		return errors.New(`{"error": "Email is invalid"`)
	}

	fmt.Println("email is valid \n", res)
	return nil
}
