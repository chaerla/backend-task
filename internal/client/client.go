package client

import "go.uber.org/fx"

var Module = fx.Module(
	"client",
	fx.Options(
		fx.Provide(
			fx.Annotate(NewDummyApi, fx.As(new(DummyApi))),
		),
	),
)
