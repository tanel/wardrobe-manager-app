package model

import (
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

// User represents user
type User struct {
	Base
	Email        string  `json:"email"`
	Name         *string `json:"name,omitempty"`
	Picture      *string `json:"picture,omitempty"`
	PasswordHash *string `json:"-"`
}

// CheckPassword checks if password is correct
func (user User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(password)); err != nil {
		return errors.Annotate(err, "password incorrect")
	}

	return nil
}
