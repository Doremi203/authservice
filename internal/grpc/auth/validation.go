package authgrpc

import (
	"authservice/internal/domain/services/auth"
	"authservice/internal/domain/types"
	ssov1 "authservice/protos/gen/go/sso"
	"errors"
)

func validateRegisterRequest(req *ssov1.RegisterRequest) (auth.RegisterModel, error) {
	model := auth.RegisterModel{
		Email:    types.Email(req.Email),
		Password: types.Password(req.Password),
	}
	errs := []error{
		model.Email.Validate(),
		model.Password.Validate(),
	}

	if err := validationErrors(errs); err != nil {
		return auth.RegisterModel{}, err
	}

	return model, nil
}

func validateLoginRequest(req *ssov1.LoginRequest) (auth.LoginModel, error) {
	model := auth.LoginModel{
		Email:    types.Email(req.Email),
		Password: types.Password(req.Password),
		AppID:    types.AppID(req.AppId),
	}
	errs := []error{
		model.Email.Validate(),
		model.Password.Validate(),
		model.AppID.Validate(),
	}

	if err := validationErrors(errs); err != nil {
		return auth.LoginModel{}, err
	}

	return model, nil
}

func validationErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	return errors.Join(errs...)
}
