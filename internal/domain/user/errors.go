package user

import "errors"

var (
	ErrAlreadyExists = errors.New("user with this email already exists")
	ErrNotFound      = errors.New("user not found")
)
