package api

import (
	"context"

	"github.com/Solwery-Veronika/auth/pkg/auth"
	"github.com/Solwery-Veronika/gateway/internal/model"
)

type SrvI interface {
	SignupUsecase(ctx context.Context, data model.SignupData) (*auth.SignupResponse, error)
	LoginUsecase(ctx context.Context, data model.LoginData) (*auth.LoginOut, error)
}
