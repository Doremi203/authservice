package types

import "errors"

type AppID int

func (i AppID) Validate() error {
	if i == 0 {
		return errors.New("appID is empty")
	}

	return nil
}
