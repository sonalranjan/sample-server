package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/srnewbie/sample-server/client"
	"github.com/srnewbie/sample-server/config"
	"github.com/srnewbie/sample-server/handler"
	"github.com/srnewbie/sample-server/server"
	"go.uber.org/fx"
)

func main() {
	f := fx.New(
		fx.Provide(
			config.New(),
			client.New,
			handler.New,
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
