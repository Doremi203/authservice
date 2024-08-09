package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAlreadyRegistered  = errors.New("already registered")
)
