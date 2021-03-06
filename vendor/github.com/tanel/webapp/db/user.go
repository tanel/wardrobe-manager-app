package db

import (
	"database/sql"

	"github.com/juju/errors"
	"github.com/tanel/webapp/model"
)

// SelectUserByEmail selects user from database by email
func SelectUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := SharedInstance.QueryRow(`
		SELECT
			id,
			password_hash,
			created_at,
			picture,
			name
		FROM
			users
		WHERE
			email = $1
	`, email,
	).Scan(
		&user.ID,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.Picture,
		&user.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Annotate(err, "selecting user by email failed")
	}

	return &user, nil
}

// SelectUserByID selects user from database by ID
func SelectUserByID(ID string) (*model.User, error) {
	var user model.User
	if err := SharedInstance.QueryRow(`
		SELECT
			id,
			password_hash,
			created_at,
			picture,
			name
		FROM
			users
		WHERE
			id = $1
	`, ID,
	).Scan(
		&user.ID,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.Picture,
		&user.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Annotate(err, "selecting user by ID failed")
	}

	return &user, nil
}

// InsertUser inserts user into database
func InsertUser(user model.User) error {
	if _, err := SharedInstance.Exec(`
		INSERT INTO users(
			id,
			email,
			password_hash,
			created_at,
			picture,
			name
		) VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
	`,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
		user.Picture,
		user.Name,
	); err != nil {
		return errors.Annotate(err, "inserting user failed")
	}

	return nil
}

// UpdateUser updates user in database
func UpdateUser(user model.User) error {
	if _, err := SharedInstance.Exec(`
		UPDATE users
		SET
			picture = $1,
			name = $2
		WHERE
			id = $3
	`,
		user.Picture,
		user.Name,
		user.ID,
	); err != nil {
		return errors.Annotate(err, "updating user failed")
	}

	return nil
}
