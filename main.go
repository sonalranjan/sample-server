package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
	_ "github.com/srnewbie/sample-server/docs"

	"github.com/srnewbie/sample-server/client"
	"github.com/srnewbie/sample-server/config"
	"github.com/srnewbie/sample-server/handler"
	"github.com/srnewbie/sample-server/server"
	"github.com/srnewbie/sample-server/service"
	"go.uber.org/fx"
)

// Run main() function with arguments like:
// ./bin/sample-server --config-files-dir /etc/sample-server-config/
//
//	@title				SampleServer API
//	@version			0.0.1
//	@description			SampleServer API
//	@Schemes			http
//	@query.collection.format	multi
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in				header
//	@name				X-API-KEY
func main() {
	params := config.Params{}
	_ = kong.Parse(&params)

	f := fx.New(
		fx.Provide(
			config.New(params),
			client.New,
			handler.New,
			server.New,
			service.New,
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
