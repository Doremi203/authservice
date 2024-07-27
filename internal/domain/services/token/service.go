package token

import (
	"authservice/internal/domain/types"
	"context"
)

type Service interface {
	GenerateToken(ctx context.Context, extraClaims map[string]any) (types.Token, error)
}
