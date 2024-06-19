package cmd

import "go.uber.org/fx"

var Module = fx.Module(
	"cmd",
	fx.Provide(NewGetUsersCmd),
	fx.Provide(NewGetPostsCmd),
	fx.Provide(NewKafkaRunnerCmd),
	fx.Provide(NewRedisRunnerCmd),
)
