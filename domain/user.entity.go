package domain

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RecoveryPassword struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Code  string `json:"code"`
}
