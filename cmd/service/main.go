package main

import (
	"net/http"

	"github.com/Solwery-Veronika/gateway/internal/api"
	"github.com/Solwery-Veronika/gateway/internal/client"
	"github.com/Solwery-Veronika/gateway/internal/config"
	"github.com/Solwery-Veronika/gateway/internal/repository"
	"github.com/Solwery-Veronika/gateway/internal/service"
)

// компилируемые (go, c, c++, java)
// интерпритируемые (python, js)

func main() {
	cfg := config.MustLoad()
	repo := repository.New()

	// клиент
	c := client.New(cfg)

	srv := service.New(repo, c)

	handler := api.New(srv)

	http.HandleFunc("/signup", handler.Signup)
	http.HandleFunc("/login", handler.Login)

	err := http.ListenAndServe(":"+cfg.Service.Port, nil)
	if err != nil {
		panic(err)
	}
}
