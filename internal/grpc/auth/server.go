package authgrpc

import (
	"authservice/internal/domain/auth"
	"authservice/internal/domain/user"
	ssov1 "authservice/protos/gen/go/sso"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	authService auth.Service
	log         *slog.Logger
}

func Register(gRPC *grpc.Server, authService auth.Service, log *slog.Logger) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{authService: authService, log: log})
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	op := "authgrpc.ServerAPI.Register"
	log := s.log.With(slog.String("operation", op))

	model, err := validateRegisterRequest(req)
	if err != nil {
		return nil, gRPCValidationError(err)
	}

	id, err := s.authService.Register(ctx, model)
	if err != nil {
		if errors.Is(err, user.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "user with this email already exists")
		}

		log.Error("internal server error", "error", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &ssov1.RegisterResponse{UserId: string(id)}, nil
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	op := "authgrpc.ServerAPI.Login"
	log := s.log.With(slog.String("operation", op))

	model, err := validateLoginRequest(req)
	if err != nil {
		return nil, gRPCValidationError(err)
	}

	token, err := s.authService.Login(ctx, model)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, status.Error(codes.Unauthenticated, "invalid email or password")
		}

		log.Error("internal server error", "error", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &ssov1.LoginResponse{Token: string(token)}, nil
}

func gRPCValidationError(err error) error {
	return status.Error(codes.InvalidArgument, fmt.Sprintf("validation errors: %v", err))
}
