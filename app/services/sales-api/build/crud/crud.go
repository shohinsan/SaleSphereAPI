// Package crud binds the crud domain set of routes into the specified app.
package crud

import (
	"github.com/shohinsan/SaleSphereAPI/app/services/sales-api/handlers/checkgrp"
	"github.com/shohinsan/SaleSphereAPI/app/services/sales-api/handlers/homegrp"
	"github.com/shohinsan/SaleSphereAPI/app/services/sales-api/handlers/productgrp"
	"github.com/shohinsan/SaleSphereAPI/app/services/sales-api/handlers/trangrp"
	"github.com/shohinsan/SaleSphereAPI/app/services/sales-api/handlers/usergrp"
	"github.com/shohinsan/SaleSphereAPI/business/web/mux"
	"github.com/shohinsan/SaleSphereAPI/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {
	checkgrp.Routes(app, checkgrp.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	homegrp.Routes(app, homegrp.Config{
		Log:      cfg.Log,
		Delegate: cfg.Delegate,
		Auth:     cfg.Auth,
		DB:       cfg.DB,
	})

	productgrp.Routes(app, productgrp.Config{
		Log:      cfg.Log,
		Delegate: cfg.Delegate,
		Auth:     cfg.Auth,
		DB:       cfg.DB,
	})

	trangrp.Routes(app, trangrp.Config{
		Log:      cfg.Log,
		Delegate: cfg.Delegate,
		Auth:     cfg.Auth,
		DB:       cfg.DB,
	})

	usergrp.Routes(app, usergrp.Config{
		Log:      cfg.Log,
		Delegate: cfg.Delegate,
		Auth:     cfg.Auth,
		DB:       cfg.DB,
	})
}
