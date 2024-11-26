package service

import (
	"context"
	"fmt"
	"github.com/Solwery-Veronika/gateway/internal/model"
	"github.com/Solwery-Veronika/gateway/pkg/auth"
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

func (s *Service) SignupUsecase(ctx context.Context, data model.SignupData) (*auth.LoginOut, error) {
	if data.Password != data.RetryPassword {
		return nil, fmt.Errorf("password is incorrect")
	}

	resp, err := s.authClient.Login(ctx, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
