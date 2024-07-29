package types

import (
	"errors"
	"net/mail"
)

type Email string

func (e Email) Validate() error {
	if e == "" {
		return errors.New("email is empty")
	}

	if _, err := mail.ParseAddress(string(e)); err != nil {
		return errors.New("email is invalid")
	}

	return nil
}
