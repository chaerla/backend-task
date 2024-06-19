package usecase

import "go.uber.org/fx"

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(
			fx.Annotate(NewUserService, fx.As(new(UserService))),
		),
		fx.Provide(
			fx.Annotate(NewPostService, fx.As(new(PostService))),
		),
	),
)
