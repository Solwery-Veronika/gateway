package service

import (
	"context"
	"fmt"

	"github.com/Solwery-Veronika/auth/pkg/auth"
	"github.com/Solwery-Veronika/gateway/internal/model"
)

type Service struct {
	repo       Repo
	authClient AuthClientI
}

func New(rep Repo, aC AuthClientI) *Service {
	return &Service{
		repo:       rep,
		authClient: aC,
	}
}

func (s *Service) SignupUsecase(ctx context.Context, data model.SignupData) (*auth.SignupResponse, error) {
	if data.Password != data.RetryPassword {
		return nil, fmt.Errorf("password is incorrect")
	}

	resp, err := s.authClient.Signup(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Service) LoginUsecase(ctx context.Context, data model.LoginData) (*auth.LoginOut, error) {
	resp, err := s.authClient.Login(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
