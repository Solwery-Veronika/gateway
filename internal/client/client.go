package client

import (
	"context"
	"fmt"
	"log"

	"github.com/Solwery-Veronika/auth/pkg/auth"
	"github.com/Solwery-Veronika/gateway/internal/config"
	"github.com/Solwery-Veronika/gateway/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client      auth.AuthServiceClient
	secretToken string
}

func New(cfg *config.Config) *Client {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Client.Host, cfg.Client.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	grpcClient := auth.NewAuthServiceClient(conn) // создаем наш сгенерированый клиент
	return &Client{
		client:      grpcClient,
		secretToken: cfg.Client.Port,
	}
}

func (c *Client) Signup(ctx context.Context, data model.SignupData) (*auth.SignupResponse, error) {
	return c.client.Signup(ctx, &auth.SignupRequest{
		Username: data.Username,
		Password: data.Password,
	})
}

func (c *Client) Login(ctx context.Context, data model.LoginData) (*auth.LoginOut, error) {
	return c.client.Login(ctx, &auth.LoginIn{
		Username: data.Username,
		Password: data.Password,
	})
}
