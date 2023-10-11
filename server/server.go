package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/srnewbie/sample-server/handler"

	"github.com/gorilla/mux"
	"github.com/srnewbie/sample-server/client"
	"github.com/srnewbie/sample-server/config"
	"go.uber.org/fx"
)

type Server struct {
	Client     *client.Client
	Config     *config.Config
	Handler    *handler.Handler
	Router     *mux.Router
	HTTPServer *http.Server
}

func New(config *config.Config, client *client.Client, handler *handler.Handler) (*Server, error) {
	router := mux.NewRouter()
	return &Server{
		Client:  client,
		Config:  config,
		Handler: handler,
		Router:  router,
		HTTPServer: &http.Server{
			Addr:    config.Server.HTTP.Addr,
			Handler: router,
		},
	}, nil
}

func Register(lifecycle fx.Lifecycle, s *Server) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				s.Router.HandleFunc("/health", s.Handler.Health())
				s.Router.HandleFunc("/questions", s.Handler.CreateQuestions()).Methods("POST")
				s.Router.HandleFunc("/questions", s.Handler.ListQuestions()).Methods("GET")
				s.Router.HandleFunc("/upvote", s.Handler.UpVote()).Methods("POST")
				go func(s *Server) {
					s.Client.Logger.Info(fmt.Sprintf("http: listen at %s", s.Config.Server.HTTP.Addr))
					if err := s.HTTPServer.ListenAndServe(); err != nil {
						s.Client.Logger.Error(err)
					}
				}(s)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)
}
