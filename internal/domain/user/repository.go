package user

import (
	"authservice/internal/domain/types"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Add(ctx context.Context, userData User) (user User, err error)
	GetByEmail(ctx context.Context, email types.Email) (user User, err error)
}

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Add(ctx context.Context, userData User) (User, error) {
	op := "user.PostgresRepository.Add"
	e := Entity{}

	err := r.db.Get(
		&e,
		`INSERT INTO users (email, hashed_password)
				VALUES ($1, $2)
				RETURNING id`,
		userData.Email, userData.HashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrAlreadyExists
		}
		return User{}, fmt.Errorf("%s: %w", op, err)
	}

	return User{}, nil
}

func (r *PostgresRepository) GetByEmail(ctx context.Context, email types.Email) (User, error) {
	op := "user.PostgresRepository.GetByEmail"
	e := Entity{}

	err := r.db.Get(&e, `SELECT id, email, hashed_password FROM users WHERE email = $1`, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrNotFound
		}
		return User{}, fmt.Errorf("%s: %w", op, err)
	}

	return FromEntity(e), nil
}
