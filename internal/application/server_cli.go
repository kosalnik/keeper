package application

import (
	"errors"
	"fmt"

	"github.com/kosalnik/keeper/internal/entity"
	"github.com/kosalnik/keeper/internal/hasher"
	"github.com/kosalnik/keeper/internal/log"
	"github.com/kosalnik/keeper/internal/postgres"
	user2 "github.com/kosalnik/keeper/internal/service"
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
					grpcServer := server.NewGRPCServer(cfg.Listen)
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
					passwordHasher := hasher.NewHasher(`marabumba`)
					userRepo := postgres.NewUserRepository(db)
					userService := user2.NewUserService(passwordHasher, userRepo)
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
