package vproductgrp

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/shohinsan/SaleSphereAPI/business/core/views/vproduct"
	"github.com/shohinsan/SaleSphereAPI/business/core/views/vproduct/stores/vproductdb"
	"github.com/shohinsan/SaleSphereAPI/business/web/v1/auth"
	"github.com/shohinsan/SaleSphereAPI/business/web/v1/mid"
	"github.com/shohinsan/SaleSphereAPI/foundation/logger"
	"github.com/shohinsan/SaleSphereAPI/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log  *logger.Logger
	Auth *auth.Auth
	DB   *sqlx.DB
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	vPrdCore := vproduct.NewCore(vproductdb.NewStore(cfg.Log, cfg.DB))

	authen := mid.Authenticate(cfg.Auth)
	ruleAdmin := mid.Authorize(cfg.Auth, auth.RuleAdminOnly)

	hdl := new(vPrdCore)
	app.Handle(http.MethodGet, version, "/vproducts", hdl.Query, authen, ruleAdmin)
}
