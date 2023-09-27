package postgres

import (
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
