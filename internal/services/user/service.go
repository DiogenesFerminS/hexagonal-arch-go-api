package user

import "go-api/internal/ports"

type UserService struct {
	Repository ports.UserRepository
}
