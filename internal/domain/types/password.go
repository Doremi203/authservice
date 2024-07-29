package types

import "errors"

type Password string

func (p Password) Validate() error {
	if p == "" {
		return errors.New("password is empty")
	}

	return nil
}
