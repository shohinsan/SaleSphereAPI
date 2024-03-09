package v1

import (
	"os"

	"github.com/shohinsan/SaleSphereAPI/business/web/v1/mid"
	"github.com/shohinsan/SaleSphereAPI/foundation/logger"
	"github.com/shohinsan/SaleSphereAPI/foundation/web"
)

// APIMuxConfig encapsulates the configuration for the API Mux.
type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *logger.Logger
}

type RouteAdder interface {
	Add(app *web.App, cfg APIMuxConfig)
}

func APIMux(cfg APIMuxConfig, routeAdder RouteAdder) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log))

	routeAdder.Add(app, cfg)

	return app
}
