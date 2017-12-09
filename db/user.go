package db

import (
	"database/sql"

	"github.com/juju/errors"
	"github.com/tanel/wardrobe-manager-app/model"
)

func SelectUserByEmail(email string, user *model.User) error {
	err := db.QueryRow("SELECT id, password_hash FROM users WHERE email = $1", user.Email).Scan(&user.ID, &user.PasswordHash)
	if err != nil && err != sql.ErrNoRows {
		return errors.Annotate(err, "selecting user by email failed")
	}

	return nil
}

func InsertUser(user model.User) error {
	_, err := db.Exec("INSERT INTO users(id, email, password_hash) VALUES($1, $2, $3)", user.ID, user.Email, user.PasswordHash)
	if err != nil {
		return errors.Annotate(err, "inserting user failed")
	}

	return nil
}
