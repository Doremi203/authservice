package entities

import "authservice/internal/domain/types"

type User struct {
	ID       types.UserID
	Email    types.Email
	Password types.Password
}
