package auth

import (
	"authservice/internal/domain/password"
	"authservice/internal/domain/token"
	"authservice/internal/domain/types"
	user2 "authservice/internal/domain/user"
	"context"
	"errors"
	"fmt"
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
	userRepository user2.Repository
	hashProvider   password.Provider
}

func NewDefaultService(tokenService token.Service, userRepository user2.Repository, hashProvider password.Provider) *DefaultService {
	return &DefaultService{
		tokenService:   tokenService,
		userRepository: userRepository,
		hashProvider:   hashProvider,
	}
}

func (s *DefaultService) Register(ctx context.Context, model RegisterModel) (_ types.UserID, err error) {
	op := "auth.DefaultService.Register"
	defer func() {
		if err != nil {
			if errors.Is(err, user2.ErrAlreadyExists) {
				err = ErrAlreadyRegistered
			} else {
				err = fmt.Errorf("%s: %w", op, err)
			}
		}
	}()

	hashedPassword, err := s.hashProvider.Hash(model.Password)
	if err != nil {
		return "", err
	}

	userData := user2.User{
		Email:          model.Email,
		HashedPassword: hashedPassword,
	}
	u, err := s.userRepository.Add(ctx, userData)
	if err != nil {
		return "", err
	}

	return u.ID, nil
}

func (s *DefaultService) Login(ctx context.Context, model LoginModel) (_ types.Token, err error) {
	op := "auth.DefaultService.Login"
	defer func() {
		if err != nil {
			switch {
			case errors.Is(err, user2.ErrNotFound):
				err = ErrInvalidCredentials
			case errors.Is(err, ErrInvalidCredentials):
				break
			default:
				err = fmt.Errorf("%s: %w", op, err)
			}
		}
	}()

	u, err := s.userRepository.GetByEmail(ctx, model.Email)
	if err != nil {
		return "", err
	}

	if s.hashProvider.Verify(model.Password, u.HashedPassword) {
		return "", ErrInvalidCredentials
	}

	t, err := s.tokenService.GenerateUserToken(ctx, u)
	if err != nil {
		return "", err
	}

	return t, nil
}
