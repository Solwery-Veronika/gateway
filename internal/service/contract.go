package service

import (
	"context"
	"github.com/Solwery-Veronika/gateway/internal/model"
	"github.com/Solwery-Veronika/gateway/pkg/auth"
)

type Repo interface {
	Save(chislo int)
}

type AuthClientI interface {
	Login(ctx context.Context, data model.SignupData) (*auth.LoginOut, error)
}
