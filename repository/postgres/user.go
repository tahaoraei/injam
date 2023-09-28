package postgres

import (
	"database/sql"
	"injam/entity"
	"log/slog"
)

func (db *DB) Register(user entity.User) (entity.User, error) {
	var lastID int

	if err := db.db.QueryRow(`insert into users(phone_number, name, password) values ($1, $2, $3) RETURNING id`, user.PhoneNumber, user.Name, user.Password).Scan(&lastID); err != nil {
		slog.Error("error in insert user: ", err)
		return entity.User{}, err
	}

	return entity.User{
		ID:          lastID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
	}, nil
}

func (db *DB) IsPhoneNumberUnique(number string) (bool, error) {
	var id int
	if err := db.db.QueryRow(`select id from users where phone_number = $1`, number).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
