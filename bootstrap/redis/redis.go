package redis

import (
	"backend-task/bootstrap/config"
	"backend-task/pkg/logger"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"os"
)

type RedisClient interface {
	Set(key string, value interface{}) error
	Get(key string) (string, error)
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(config *config.Config) RedisClient {
	client := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		},
	)

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		logger.Log.Error("Failed to connect to Redis:", err)
		os.Exit(1)
	}

	logger.Log.Info("Connected to Redis successfully")

	return &redisClient{client}
}

func (rc *redisClient) Set(key string, value interface{}) error {
	err := rc.client.Set(rc.client.Context(), key, value, 0).Err()
	if err != nil {
		logger.Log.Warn("Error setting value in Redis:", err)
		return err
	}
	logger.Log.Info("Value set successfully in Redis")
	return nil
}

func (rc *redisClient) Get(key string) (string, error) {
	val, err := rc.client.Get(rc.client.Context(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

var Module = fx.Provide(NewRedisClient)
