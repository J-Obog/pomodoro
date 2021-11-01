package usersvc

type User struct {
	ID       uint   `json:"id"`
	email    string `json:"email"`
	password string `json:"password"`
}