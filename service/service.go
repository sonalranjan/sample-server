package service

import (
	"github.com/sample-server/config"
	"github.com/sample-server/client"
)

type Service struct {
	c         *client.Client
	f         *config.Config
}

func New(client *client.Client, config *config.Config) *Service {
	return &Service{
		c:         client,
		f:         config,
	}
}
