package model

type User struct {
	Base
	Email        string
	Name         string
	PasswordHash string
}
