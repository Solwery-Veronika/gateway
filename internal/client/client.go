package client

import (
	"context"

	"github.com/Solwery-Veronika/gateway/pkg/auth"
)

type Client struct {
	client auth.AuthServiceClient
}

func New(client auth.AuthServiceClient) *Client {
	return &Client{client: client}
}

func (c *Client) Login(ctx context.Context, in *auth.LoginIn) (*auth.LoginOut, error) {
	return c.client.Login(ctx, in)
}
