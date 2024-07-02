package app

import (
	"context"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(
	lc fx.Lifecycle,
	mux *http.ServeMux,
	log *zap.Logger,
) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}

	lc.Append(fx.Hook{
		OnStart: start(srv, log),
		OnStop:  stop(srv, log),
	})

	return srv
}

func start(
	srv *http.Server,
	log *zap.Logger,
) func(context.Context) error {
	return func(ctx context.Context) error {
		ln, err := net.Listen("tcp", srv.Addr)
		if err != nil {
			return err
		}

		log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
		go func() {
			err := srv.Serve(ln)
			if err != nil {
				log.Error("Failed to start HTTP server", zap.Error(err))
			}
		}()

		return nil
	}
}

func stop(
	srv *http.Server,
	log *zap.Logger,
) func(context.Context) error {
	return func(ctx context.Context) error {
		log.Info("Shutting down HTTP server")
		return srv.Shutdown(ctx)
	}
}
