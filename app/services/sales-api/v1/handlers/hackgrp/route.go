package hackgrp

import (
	"net/http"

	"github.com/shohinsan/SaleSphereAPI/business/web/v1/auth"
	"github.com/shohinsan/SaleSphereAPI/business/web/v1/mid"
	"github.com/shohinsan/SaleSphereAPI/foundation/web"
)

type Config struct {
	Auth *auth.Auth
}

// Routes adds specific routes for this group
// Routes adds specific routes for this group
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Auth)
	ruleAdmin := mid.Authorize(cfg.Auth, auth.RuleAdminOnly)

	app.Handle(http.MethodGet, "/hack", version, Hack)
	app.Handle(http.MethodGet, "/hackauth", version, Hack, authen, ruleAdmin)
}
