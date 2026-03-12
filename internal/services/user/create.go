package user

import (
	"errors"
	"go-api/internal/domain"
	"time"
)

func (u UserService) Create(user domain.User) (userId int64, err error) {

	user.CreatedAt = time.Now().UTC()

	userId, err = u.Repository.Save(user)

	if err != nil {
		return 0, errors.New(err.Error())
	}

	return userId, nil
}
