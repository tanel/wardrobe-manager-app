package db

import (
	"database/sql"

	"github.com/juju/errors"
	"github.com/tanel/webapp/model"
)

// SelectUserByEmail selects user from database by email
func SelectUserByEmail(db *sql.DB, email string, user *model.User) error {
	if err := db.QueryRow(`
		SELECT
			id,
			password_hash,
			created_at
		FROM
			users
		WHERE
			email = $1
	`, user.Email,
	).Scan(
		&user.ID,
		&user.PasswordHash,
		&user.CreatedAt,
	); err != nil && err != sql.ErrNoRows {
		return errors.Annotate(err, "selecting user by email failed")
	}

	return nil
}

// InsertUser inserts user into database
func InsertUser(db *sql.DB, user model.User) error {
	if _, err := db.Exec(`
		INSERT INTO users(
			id,
			email,
			password_hash,
			created_at
		) VALUES(
			$1,
			$2,
			$3,
			$4
		)
	`,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
	); err != nil {
		return errors.Annotate(err, "inserting user failed")
	}

	return nil
}
