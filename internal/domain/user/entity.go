package user

import "authservice/internal/domain/types"

type Entity struct {
	ID             string `db:"id"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
}

func FromEntity(entity Entity) User {
	return User{
		ID:             types.UserID(entity.ID),
		Email:          types.Email(entity.Email),
		HashedPassword: types.HashedPassword(entity.HashedPassword),
	}
}
