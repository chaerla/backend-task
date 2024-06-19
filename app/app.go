package app

import (
	cmd2 "backend-task/internal/cmd"
	"backend-task/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

type Application struct {
	rootCmd *cobra.Command
}

func NewApplication(
	getUsersCmd cmd2.GetUsersCmd,
	postUsersCmd cmd2.GetPostsCmd,
	kafkaRunnerCmd cmd2.KafkaRunnerCmd,
	redisRunnerCmd cmd2.RedisRunnerCmd,
) Application {
	rootCmd := &cobra.Command{Use: "app"}
	rootCmd.AddCommand(getUsersCmd)
	rootCmd.AddCommand(postUsersCmd)
	rootCmd.AddCommand(kafkaRunnerCmd)
	rootCmd.AddCommand(redisRunnerCmd)
	return Application{
		rootCmd: rootCmd,
	}
}

func (app Application) Run() {
	if err := app.rootCmd.Execute(); err != nil {
		logger.Log.Fatal("failed to run app")
	}
}

var Module = fx.Module(
	"app",
	fx.Options(fx.Provide(NewApplication)),
)
