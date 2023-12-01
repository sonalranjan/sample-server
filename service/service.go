package service

import (
	"github.com/srnewbie/sample-server/client"
	"github.com/srnewbie/sample-server/config"
)

type Service struct {
	Client *client.Client
	Config *config.Config
}

func New(config *config.Config, client *client.Client) (*Service, error) {
	return &Service{
		Client: client,
		Config: config,
	}, nil
}
