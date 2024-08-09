package token

import (
	"authservice/internal/domain/types"
	"authservice/internal/domain/user"
	"authservice/pkg/utils"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateUserToken(ctx context.Context, user user.User) (types.Token, error)
}

type JWTService struct {
	config       Config
	timeProvider utils.TimeProvider
}

func NewJWTService(config Config, timeProvider utils.TimeProvider) *JWTService {
	return &JWTService{
		config:       config,
		timeProvider: timeProvider,
	}
}

func (s *JWTService) GenerateUserToken(ctx context.Context, user user.User) (types.Token, error) {
	op := "token.JWTService.GenerateUserToken"

	claims := jwt.MapClaims{
		"sub": user.ID,
	}

	t, err := s.generateToken(ctx, claims)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return t, nil
}

func (s *JWTService) generateToken(_ context.Context, extraClaims jwt.MapClaims) (types.Token, error) {
	curTime := s.timeProvider.UTCNow()
	claims := jwt.MapClaims{
		"exp": curTime.Unix(),
	}
	for k, v := range extraClaims {
		claims[k] = v
	}

	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.config.Key)
	if err != nil {
		return "", err
	}

	return types.Token(t), nil
}
