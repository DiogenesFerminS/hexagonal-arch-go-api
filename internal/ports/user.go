package ports

import "go-api/internal/domain"

type UserService interface {
	Create(user domain.User) (userId int64, err error)
}

type UserRepository interface {
	Save(user domain.User) (userId int64, err error)
}
