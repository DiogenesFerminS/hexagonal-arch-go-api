package ports

import (
	"context"
	"go-api/internal/domain"
)

type UserService interface {
	Create(ctx context.Context, user domain.User) (userId int64, err error)
}

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (userId int64, err error)
}
