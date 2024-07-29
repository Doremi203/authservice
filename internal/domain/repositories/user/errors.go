package user

import "errors"

var (
	ErrAlreadyExists = errors.New("user with this email already exists")
)
