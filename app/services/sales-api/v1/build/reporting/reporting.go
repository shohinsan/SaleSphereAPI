// Package reporting binds the reporting domain set of routes into the specified app.
package reporting

import (
	"github.com/shohinsan/SaleSphereAPI/app/services/sales-api/v1/handlers/checkgrp"
	"github.com/shohinsan/SaleSphereAPI/app/services/sales-api/v1/handlers/vproductgrp"
	"github.com/shohinsan/SaleSphereAPI/business/web/v1/mux"
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

	vproductgrp.Routes(app, vproductgrp.Config{
		Log:  cfg.Log,
		Auth: cfg.Auth,
		DB:   cfg.DB,
	})
}
