package application

import (
	"errors"
	"fmt"

	"github.com/kosalnik/keeper/internal/auth"
	"github.com/kosalnik/keeper/internal/entity"
	"github.com/kosalnik/keeper/internal/hasher"
	"github.com/kosalnik/keeper/internal/log"
	"github.com/kosalnik/keeper/internal/postgres"
	service "github.com/kosalnik/keeper/internal/service"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"

	"github.com/kosalnik/keeper/internal/config"
	"github.com/kosalnik/keeper/internal/server"
)

func NewServerCLI(version string, cfg *config.Server) *cli.App {
	return &cli.App{
		Name:    "GophKeeper server",
		Version: version,
		Usage:   "Run and use",
		Commands: []*cli.Command{
			{
				Name: "serve",
				Action: func(context *cli.Context) error {
					jwtSvc, err := service.NewJwtService(cfg.JWT)
					if err != nil {
						return err
					}
					authService := auth.NewAuthenticator(jwtSvc)
					// @TODO через конфиг надо
					db, err := postgres.NewConn(`postgres://postgres:postgres@127.0.0.1:5432/keeper`)
					if err != nil {
						log.Error("create user. fail connect db", log.Err(err))
						return err
					}
					passwordHasher := hasher.NewHasher()
					userRepo := postgres.NewUserRepository(db)
					userSvc := service.NewUserService(passwordHasher, userRepo)
					cert, err := cfg.TLS.Certificate()
					if err != nil {
						log.Panic("failed to load key pair: %s", log.Err(err))
					}
					opts := []server.Opt{
						server.WithAuth(authService),
						server.WithTLS(cert),
					}
					grpcServer := server.NewGRPCServer(cfg.Listen, userSvc, jwtSvc, opts...)
					if err := grpcServer.Serve(); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
						return err
					}
					return nil
				},
			},
			{
				Name: "create-user",
				Action: func(ctx *cli.Context) error {
					login := ctx.Args().Get(0)
					password := ctx.Args().Get(1) // @TODO Пароль надо с консоли запрашивать с подтверждением
					db, err := postgres.NewConn(`postgres://postgres:postgres@127.0.0.1:5432/keeper`)
					if err != nil {
						log.Error("create user. fail connect db", log.Err(err))
						return err
					}
					passwordHasher := hasher.NewHasher()
					userRepo := postgres.NewUserRepository(db)
					userService := service.NewUserService(passwordHasher, userRepo)
					var user *entity.User
					if user, err = userService.Create(ctx.Context, login, password); err != nil {
						log.Error("fail create user", log.Err(err))
						return err
					}
					fmt.Printf("User with login %s has been created. ID=%s", user.Login, user.ID)
					return nil
				},
			},
		},
		Before: func(c *cli.Context) error {
			fmt.Printf("GophKeeper %s\n", c.App.Version)
			return nil
		},
		DefaultCommand: "serve",
	}
}
