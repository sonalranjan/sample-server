package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	Prod    Env = "prod"
	Staging Env = "staging"
	Dev     Env = "dev"
	Local   Env = "local"
)

var (
	home   = "../config/yaml"
	logger *zap.SugaredLogger
	p      string
)

type (
	Env string

	Config struct {
		Server *ServerConfig
		Client *ClientConfig
	}

	ServerConfig struct {
		HTTP     *ServerHTTPConfig
	}

	ServerHTTPConfig struct {
		Addr string
	}

	ClientConfig struct {
		Logger       *LoggerConfig
	}

	LoggerConfig struct {
		Level    string
		Encoding string
	}
)

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = l.Sugar()

	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p = path.Join(filepath.Dir(exe), home)
}

func New(env ...Env) func() (*Config, error) {
	return func() (*Config, error) {
		viper.SetConfigName("base")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(p)

		replacer := strings.NewReplacer(".", "_")
		viper.SetEnvKeyReplacer(replacer)
		viper.AutomaticEnv()

		logger.Infow("creating base config")
		err := viper.ReadInConfig()
		if err != nil {
			logger.Errorw("failed to read in config", "err", err)
			return nil, err
		}

		for _, e := range env {
			logger.Infow(fmt.Sprintf("merging '%s' config", e))
			switch e {
			case Prod:
			case Staging:
			case Dev:
			case Local:
			default:
				return nil, errors.New("unsupported env type")
			}
			viper.SetConfigName(string(e))
			err = viper.MergeInConfig()
			if err != nil {
				return nil, err
			}
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

func GetEnvironment() Env {
	env := Env(os.Getenv("CONFIGURATION_ENV"))
	switch env {
	case Prod:
	case Staging:
	case Dev:
	case Local:
	default:
		env = Local
	}
	return env
}
