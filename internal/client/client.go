package client

import (
	"context"

	"github.com/kosalnik/keeper/internal/config"
)

type Client struct {
	g *GRPCClient
}

func NewClient(
	cfg *config.Client,
) *Client {
	return &Client{
		g: NewGRPCClient(cfg.ServerAddr),
	}
}

func (c *Client) Register(login, password string) error {
	return c.g.Register(context.Background(), login, password)
}
