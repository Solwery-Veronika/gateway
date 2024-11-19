package api

import "github.com/Solwery-Veronika/gateway/internal/model"

type SrvI interface {
	Summator(data *model.Data)
}
