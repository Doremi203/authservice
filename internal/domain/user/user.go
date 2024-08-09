package user

import (
	"authservice/internal/domain/types"
)

type User struct {
	ID             types.UserID
	Email          types.Email
	HashedPassword types.HashedPassword
}
