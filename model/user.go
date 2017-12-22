package model

// User represents user
type User struct {
	Base
	Email        string
	Name         string
	PasswordHash string
}
