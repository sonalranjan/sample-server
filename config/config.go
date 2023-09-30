package config

import (
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

type (
	Config struct {
		Server *ServerConfig
		Client *ClientConfig
	}

	ServerConfig struct {
		HTTP *ServerHTTPConfig
	}

	ServerHTTPConfig struct {
		Addr string
	}

	ClientConfig struct {
		Logger *LoggerConfig
	}

	LoggerConfig struct {
		Level    string
		Encoding string
	}
)

func New() func() (*Config, error) {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = l.Sugar()

	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return func() (*Config, error) {
		viper.SetConfigName("base")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(path.Join(filepath.Dir(exe), "../config/yaml"))
		logger.Infow("creating base config")
		err := viper.ReadInConfig()
		if err != nil {
			logger.Errorw("failed to read in config", "err", err)
			return nil, err
		}

		config := &Config{}
		err = viper.Unmarshal(config)
		if err != nil {
			logger.Errorw("failed to unmarshal config", "err", err)
			return nil, err
		}

		logger.Infow("", "config", config)
		return config, nil
	}
}
