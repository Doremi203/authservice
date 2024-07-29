package auth

import (
	"authservice/internal/domain/repositories/user"
	"authservice/internal/domain/services/auth"
	"authservice/internal/domain/services/token"
	"authservice/internal/domain/types"
	"context"
)

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

func (s *DefaultService) Register(ctx context.Context, model auth.RegisterModel) (userID types.UserID, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *DefaultService) Login(ctx context.Context, model auth.LoginModel) (token types.Token, err error) {
	//TODO implement me
	panic("implement me")
}
