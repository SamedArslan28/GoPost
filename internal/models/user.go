package models

type User struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Created  string `json:"created_at"`
}
