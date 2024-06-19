package main

import (
	"backend-task/app"
	"backend-task/bootstrap/config"
	"backend-task/bootstrap/tracing"
	"backend-task/internal/client"
	"backend-task/internal/cmd"
	"backend-task/internal/usecase"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"os"
)

func startApp(app app.Application) {
	app.Run()
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	app := fx.New(
		config.Module,
		client.Module,
		usecase.Module,
		cmd.Module,
		app.Module,
		fx.Invoke(tracing.InitTracing, startApp),
		fx.NopLogger,
	)
	app.Run()
}
