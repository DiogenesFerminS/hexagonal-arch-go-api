package user

import (
	"errors"
	"go-api/cmd/api/utils"
	"go-api/internal/domain"
)

func (r Repository) Save(user domain.User) (userId int64, err error) {
	query := "INSERT INTO users(username, password, created_at) VALUES ($1, $2, $3) RETURNING id"

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return 0, errors.New("Hashed password failed")
	}

	var savedUserId int64
	err = r.Client.QueryRow(query, user.Username, hashedPassword, user.CreatedAt).Scan(&savedUserId)

	if err != nil {
		return 0, errors.New("Query failed")
	}

	return savedUserId, nil
}
