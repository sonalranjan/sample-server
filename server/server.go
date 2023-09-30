package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/sample-server/src/handler"
	"github.com/sample-server/src/service"

	"github.com/gorilla/mux"
	"github.com/sample-server/src/client"
	"github.com/sample-server/src/config"
	"go.uber.org/fx"
)

type Server struct {
	Client     *client.Client
	Config     *config.Config
	Handler    *handler.Handler
	Service    *service.Service
	Router     *mux.Router
	HTTPServer *http.Server
}

func New(config *config.Config, client *client.Client, handler *handler.Handler, service *service.Service) (*Server, error) {
	router := mux.NewRouter()
	return &Server{
		Client:   client,
		Config:   config,
		Handler:  handler,
		Router:   router,
		HTTPServer: &http.Server{
			Addr:    config.Server.HTTP.Addr,
			Handler: router,
		},
		Service: service,
	}, nil
}

func Register(lifecycle fx.Lifecycle, s *Server) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				s.Router.HandleFunc("/health", s.Handler.Health())
				go func(s *Server) {
					s.Client.Logger.Infow(fmt.Sprintf("http: listen at %s", s.Config.Server.HTTP.Addr))
					if err := s.HTTPServer.ListenAndServe(); err != nil {
						s.Client.Logger.Error(err)
					}
				}
				return nil
			},
			OnStop: func(ctx context.Context) error {
				if err := s.Client.Logger.Sync(); err != nil {
					s.Client.Logger.Errorw("logger sync failed...", "err", err)
				}
				return nil
			},
		},
	)
}
