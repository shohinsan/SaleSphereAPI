package handlers

import (
	"github.com/shohinsan/SaleSphereAPI/app/services/sales-api/v1/handlers/hackgrp"
	v1 "github.com/shohinsan/SaleSphereAPI/business/web/v1"
	"github.com/shohinsan/SaleSphereAPI/foundation/web"
)

type Router struct{}

// Add
func (Router) Add(app *web.App, cfg v1.APIMuxConfig) {
	hackgrp.Routes(app)
}
