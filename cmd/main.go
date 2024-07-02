package main

import (
	// std
	"net/http"

	// local
	"uber-fx/internal/app"

	// vendor
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		// fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		// 	return &fxevent.ZapLogger{Logger: log}
		// }), // replace fx logger.
		fx.Provide(
			// * key moment: they have dependencies as parameters in 'constructor' calls.
			// * order doesn't matter.

			zap.NewExample, // add logger (zap.NewProduction).

			app.NewHTTPServer, // add HTTP server.

			fx.Annotate(
				app.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			), // add router.

			app.AsRoute(app.NewEchoHandler),
			app.AsRoute(app.NewHelloHandler),
		),
		fx.Invoke(func(*http.Server) {}), // * to request that the HTTP server is always instantiated, even if none of the other components in the application reference it directly.
	).Run() // starts the application.
}
