package service

import (
	"context"

	"github.com/Solwery-Veronika/auth/pkg/auth"
	"github.com/Solwery-Veronika/gateway/internal/model"
)

type Repo interface {
	Save(chislo int)
}

type AuthClientI interface {
	Signup(ctx context.Context, data model.SignupData) (*auth.SignupResponse, error)
	Login(ctx context.Context, data model.LoginData) (*auth.LoginOut, error)
}
