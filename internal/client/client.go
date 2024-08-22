package client

import (
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
