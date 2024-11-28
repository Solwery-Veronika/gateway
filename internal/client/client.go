package client

import (
	"context"
	"github.com/Solwery-Veronika/gateway/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"github.com/Solwery-Veronika/gateway/pkg/auth"
)

type Client struct {
	client auth.AuthServiceClient
}

func New() *Client {
	conn, err := grpc.NewClient("0.0.0.0:8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	grpcClient := auth.NewAuthServiceClient(conn) // создаем наш сгенерированый клиент
	return &Client{client: grpcClient}
}

func (c *Client) Login(ctx context.Context, data model.SignupData) (*auth.LoginOut, error) {
	return c.client.Login(ctx, &auth.LoginIn{
		Username: data.Username,
		Password: data.Password,
	})
}
