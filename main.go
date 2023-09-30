package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sample-server/handler"
	"github.com/sample-server/service"
	"github.com/sample-server/client"
	"github.com/sample-server/config"
	"github.com/sample-server/server"
	"go.uber.org/fx"
)

func main() {
	f := fx.New(
		fx.Provide(
			config.New(config.GetEnvironment()),
			client.New,
			handler.New,
			service.New,
			server.New,
		),
		fx.Invoke(server.Register),
	)

	terminate := make(chan os.Signal, 1)
	signal.Notify(
		terminate,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	go func() {
		<-terminate
		_ = f.Stop(context.Background())
	}()

	f.Run()
}
