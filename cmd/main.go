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
				fx.ParamTags(`name:"echo"`, `name:"hello"`),
			), // add router.

			fx.Annotate(
				app.NewEchoHandler,
				fx.As(new(app.Route)), // declare explicit interface apply.
				fx.ResultTags(`name:"echo"`),
			), // add handler.

			fx.Annotate(
				app.NewHelloHandler,
				fx.As(new(app.Route)), // the same interface.
				fx.ResultTags(`name:"hello"`),
			), // add another handler.
		),
		fx.Invoke(func(*http.Server) {}), // * to request that the HTTP server is always instantiated, even if none of the other components in the application reference it directly.
	).Run() // starts the application.
}

// ? Case:
// 2 instances with of the same type are given,
// and there is no way to differentiate them.
// Example: NewEchoHandler and NewHelloHandler
// are both of type Route interface.
// ! => Error
// * => Solution: fx.ResultTags() use.
