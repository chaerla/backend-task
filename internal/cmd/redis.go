package cmd

import (
	"backend-task/bootstrap/config"
	"backend-task/bootstrap/redis"
	"backend-task/pkg/logger"
	"github.com/spf13/cobra"
)

type RedisRunnerCmd *cobra.Command

func NewRedisRunnerCmd(config *config.Config) RedisRunnerCmd {
	cmd := &cobra.Command{
		Use:   "redis",
		Short: "Use this command to interact with Redis",
	}

	cmd.AddCommand(NewRedisSetCmd(config))
	cmd.AddCommand(NewRedisGetCmd(config))

	return cmd
}

func NewRedisSetCmd(config *config.Config) *cobra.Command {
	var key string
	var value string

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set a value in Redis",
		Run: func(cmd *cobra.Command, args []string) {
			client := redis.NewRedisClient(config)

			if key != "" && value != "" {
				err := client.Set(key, value)
				if err != nil {
					logger.Log.Error("Failed to set value in Redis:", err)
					return
				}
				logger.Log.Info("Successfully set value in Redis")
			} else {
				cmd.Help()
			}
		},
	}

	cmd.Flags().StringVarP(&key, "key", "k", "", "Redis key")
	cmd.Flags().StringVarP(&value, "value", "v", "", "Value to set in Redis")

	cmd.MarkFlagRequired("key")
	cmd.MarkFlagRequired("value")

	return cmd
}

func NewRedisGetCmd(config *config.Config) *cobra.Command {
	var key string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a value from Redis",
		Run: func(cmd *cobra.Command, args []string) {
			client := redis.NewRedisClient(config)

			if key != "" {
				val, err := client.Get(key)
				if err != nil {
					logger.Log.Error("Failed to get value from Redis:", err)
					return
				}
				logger.Log.Info("Value retrieved from Redis: ", val)
			} else {
				cmd.Help()
			}
		},
	}

	cmd.Flags().StringVarP(&key, "key", "k", "", "Redis key")
	cmd.MarkFlagRequired("key")

	return cmd
}
