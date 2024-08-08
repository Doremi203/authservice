package token

import (
	"authservice/internal/domain/types"
	"authservice/pkg/utils"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(ctx context.Context, extraClaims map[string]any) (types.Token, error)
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

func (s *JWTService) GenerateToken(_ context.Context, extraClaims map[string]any) (types.Token, error) {
	op := "token.JWTService.GenerateToken"
	curTime := s.timeProvider.UTCNow()
	claims := jwt.MapClaims{
		"exp": curTime.Unix(),
	}
	for k, v := range extraClaims {
		claims[k] = v
	}

	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.config.Key)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return types.Token(t), nil
}
