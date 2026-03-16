package user

import (
	"context"
	"errors"
	"go-api/internal/domain"
	"time"
)

func (u UserService) Create(ctx context.Context, user domain.User) (userId int64, err error) {

	user.CreatedAt = time.Now().UTC()

	userId, err = u.Repository.Save(ctx, user)

	if err != nil {
		return 0, errors.New(err.Error())
	}

	return userId, nil
}
