package auth

import (
	"authservice/internal/domain/services/token"
	"authservice/internal/domain/services/user"
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

type DefaultService struct {
	tokenService   token.Service
	userRepository user.Repository
}

func NewDefaultService(tokenService token.Service, userRepository user.Repository) *DefaultService {
	return &DefaultService{
		tokenService:   tokenService,
		userRepository: userRepository,
	}
}

func (s *DefaultService) Register(ctx context.Context, model RegisterModel) (userID types.UserID, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *DefaultService) Login(ctx context.Context, model LoginModel) (token types.Token, err error) {
	//TODO implement me
	panic("implement me")
}
