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

type PostgresRepository struct{}

func NewPostgresRepository() *PostgresRepository {
	return &PostgresRepository{}
}

func (p *PostgresRepository) Add(ctx context.Context, userData entities.User) (user entities.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresRepository) GetByEmail(ctx context.Context, email types.Email) (user entities.User, err error) {
	//TODO implement me
	panic("implement me")
}
