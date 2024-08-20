package main

import (
	"os"

	"github.com/kosalnik/keeper/internal/application"
	"github.com/kosalnik/keeper/internal/config"
	"github.com/kosalnik/keeper/internal/log"
)

var Version = "0.0.0"

func main() {
	cfg := config.NewServer()
	app := application.NewServerCLI(Version, cfg)
	if err := app.Run(os.Args); err != nil {
		log.Fatal("App fail", log.Err(err))
	}
}
