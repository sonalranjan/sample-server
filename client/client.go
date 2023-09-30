package client

import (
	"github.com/sirupsen/logrus"
	"github.com/srnewbie/sample-server/client/logger"
	"github.com/srnewbie/sample-server/config"
)

type Client struct {
	Config *config.Config
	Logger *logrus.Logger
}

func New(config *config.Config) (*Client, error) {
	return &Client{
		Config: config,
		Logger: logger.New(config.Client.Logger),
	}, nil
}
