package password

import (
	"authservice/internal/domain/types"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Provider interface {
	Hash(password types.Password) (hash types.HashedPassword, err error)
	Verify(password types.Password, hash types.HashedPassword) bool
}

type BCryptProvider struct{}

func NewBCryptProvider() *BCryptProvider {
	return &BCryptProvider{}
}

func (p *BCryptProvider) Hash(password types.Password) (types.HashedPassword, error) {
	op := "password.BCryptProvider.Hash"
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("%s: password hashing failed: %w", op, err)
	}

	return types.HashedPassword(h), nil
}

func (p *BCryptProvider) Verify(password types.Password, hash types.HashedPassword) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
