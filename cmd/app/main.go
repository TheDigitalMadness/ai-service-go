package main

import (
	"context"
	"fmt"
	"time"

	net_http "net/http"

	openai_client "github.com/TheDigitalMadness/ai-service-go/internal/client/openai"
	"github.com/TheDigitalMadness/ai-service-go/internal/config"
	"github.com/TheDigitalMadness/ai-service-go/internal/controller/http"
	"github.com/TheDigitalMadness/ai-service-go/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// NewOption creates and returns built options for fx DI
func NewOption() fx.Option {
	return fx.Options(
		fx.Provide(config.MustLoad),
		fx.Provide(openai_client.New),
		fx.Provide(service.New),
		fx.Provide(http.New),
		fx.Provide(http.NewRouter),
		fx.Invoke(run),
	)
}

// buildAddress builds address by config and returns it
func buildAddress(cfg *config.Config) string {
	return fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
}

// run starts serving and initializes graceful shutdown on stop
func run(cfg *config.Config, lc fx.Lifecycle, router *gin.Engine) {
	addr := buildAddress(cfg)
	server := &net_http.Server{
		Addr:    addr,
		Handler: router,
	}

	onStart := func(ctx context.Context) error {
		errChan := make(chan error, 1)

		startServer := func(ch chan error) {
			if err := server.ListenAndServe(); err != nil &&
				err != net_http.ErrServerClosed {
				errChan <- err
			}
		}

		go startServer(errChan)

		select {
		case err := <-errChan:
			return err
		case <-time.After(time.Duration(cfg.HTTP.Timeout) * time.Millisecond):
			return nil
		}
	}

	onStop := func(ctx context.Context) error {
		return server.Shutdown(ctx)
	}

	lc.Append(
		fx.Hook{
			OnStart: onStart,
			OnStop:  onStop,
		},
	)
}

func main() {
	options := NewOption()

	fx.New(options).Run()
}
