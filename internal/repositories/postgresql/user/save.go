package user

import (
	"context"
	"go-api/cmd/api/utils"
	"go-api/internal/domain"
)

func (r Repository) Save(ctx context.Context, user domain.User) (userId int64, err error) {
	var savedUserId int64

	hashedPassword, err := utils.HashPassword(user.Password)

	r.Client.QueryRow(ctx, "INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3) RETURNING id", user.Username, hashedPassword, user.CreatedAt).Scan(&savedUserId)

	return savedUserId, nil
}
