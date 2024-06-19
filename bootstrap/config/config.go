package config

import (
	"backend-task/pkg/logger"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"os"
)

type Config struct {
	DummyApiUrl   string `envconfig:"DUMMY_API_URL" required:"true"`
	DummyApiAppID string `envconfig:"DUMMY_API_APP_ID" required:"true"`
	RedisHost     string `envconfig:"REDIS_HOST" required:"true"`
	RedisPort     string `envconfig:"REDIS_PORT" required:"true"`
	KafkaHost     string `envconfig:"KAFKA_HOST" required:"true"`
	KafkaPort     string `envconfig:"KAFKA_PORT" required:"true"`
	JaegerHost    string `envconfig:"JAEGER_HOST" required:"true"`
	JaegerPort    string `envconfig:"JAEGER_PORT" required:"true"`
	ServiceName   string `envconfig:"SERVICE_NAME" required:"true"`
}

func NewConfig() (*Config, error) {
	var config Config

	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	logger.Log.Info(fmt.Sprintf("Loading env from file: %s", filename))

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", &config); err != nil {
			return nil, errors.Wrap(err, "failed to read from env variable")
		}
		return &config, nil
	}

	if err := godotenv.Load(filename); err != nil {
		return nil, errors.Wrap(err, "failed to read from .env file")
	}

	if err := envconfig.Process("", &config); err != nil {
		return nil, errors.Wrap(err, "failed to read from env variable")
	}

	return &config, nil
}

var Module = fx.Options(fx.Provide(NewConfig))
