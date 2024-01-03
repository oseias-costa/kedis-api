package persistence

import "main/entities"

func GetAll() ([]entities.User, error) {
	con := Connect()
	defer con.Close()
	sql := "SELECT * FROM user"

	r, err := con.Query(sql)
	if err != nil {
		return nil, err
	}

	var users []entities.User

	for r.Next() {
		var user entities.User
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
	defer r.Close()

	return users, nil
}
