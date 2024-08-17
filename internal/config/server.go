package config

import "github.com/kosalnik/keeper/internal/ssl"

const DefaultServerListen = ":9876"

type Server struct {
	Listen string
	TLS    ssl.ServerConfig
}

func NewServer() *Server {
	return &Server{
		Listen: DefaultServerListen,
		TLS:    ssl.ServerConfig{},
	}
}
