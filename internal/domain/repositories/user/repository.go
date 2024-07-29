package user

import (
	"authservice/internal/domain/entities"
	"authservice/internal/domain/types"
	"context"
)

type Repository interface {
	Add(ctx context.Context, userData entities.User) (user entities.User, err error)
	GetByEmail(ctx context.Context, email types.Email) (user entities.User, err error)
}
