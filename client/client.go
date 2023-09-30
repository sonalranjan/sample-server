package client

import (
	"github.com/sample-server/client/logger"
	"github.com/sample-server/config"
)

type Client struct {
	Config       *config.Config
	Logger       *logger.Logger
}

func New(config *config.Config) (*Client, error) {
	return &Client{
		Config:       config,
		Logger:       logger.New(config),
	}, nil
}
