package main

import (
	"os"

	"github.com/kosalnik/keeper/internal/application"
	"github.com/kosalnik/keeper/internal/config"
	"github.com/kosalnik/keeper/internal/log"
)

var Version = "v0.0.0"

func main() {
	cfg := config.NewClient()
	app := application.NewClientCLI(Version, cfg)
	if err := app.Run(os.Args); err != nil {
		log.Fatal("Fail", log.Err(err))
	}
}
