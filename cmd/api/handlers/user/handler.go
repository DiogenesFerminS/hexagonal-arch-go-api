package user

import "go-api/internal/ports"

type Handler struct {
	UserService ports.UserService
}
