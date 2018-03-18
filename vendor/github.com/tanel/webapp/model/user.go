package model

import (
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

// User represents user
type User struct {
	Base
	Email        string
	Name         *string
	Picture      *string
	PasswordHash *string
}

// CheckPassword checks if password is correct
func (user User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(password)); err != nil {
		return errors.Annotate(err, "password incorrect")
	}

	return nil
}
