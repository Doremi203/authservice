package auth

import (
	"authservice/internal/domain/types"
	"context"
)

type RegisterModel struct {
	Email    types.Email
	Password types.Password
}

type LoginModel struct {
	Email    types.Email
	Password types.Password
	AppID    types.AppID
}

type Service interface {
	Register(ctx context.Context, model RegisterModel) (userID types.UserID, err error)

	Login(ctx context.Context, model LoginModel) (token types.Token, err error)
}
