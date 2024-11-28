package api

import (
	"context"
	"github.com/Solwery-Veronika/gateway/internal/model"
	"github.com/Solwery-Veronika/gateway/pkg/auth"
)

type SrvI interface {
	SignupUsecase(ctx context.Context, data model.SignupData) (*auth.LoginOut, error)
}
