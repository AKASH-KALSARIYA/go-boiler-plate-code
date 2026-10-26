package model

type User struct {
	Base
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // Hidden from JSON
}
