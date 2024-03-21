package productgrp

import (
	"net/http"

	"github.com/shohinsan/SaleSphereAPI/business/core/crud/delegate"
	"github.com/shohinsan/SaleSphereAPI/business/core/crud/product"
	"github.com/shohinsan/SaleSphereAPI/business/core/crud/product/stores/productdb"
	"github.com/shohinsan/SaleSphereAPI/business/core/crud/user"
	"github.com/shohinsan/SaleSphereAPI/business/core/crud/user/stores/usercache"
	"github.com/shohinsan/SaleSphereAPI/business/core/crud/user/stores/userdb"
	"github.com/shohinsan/SaleSphereAPI/business/web/auth"
	"github.com/shohinsan/SaleSphereAPI/business/web/mid"
	"github.com/shohinsan/SaleSphereAPI/foundation/logger"
	"github.com/shohinsan/SaleSphereAPI/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log      *logger.Logger
	Delegate *delegate.Delegate
	Auth     *auth.Auth
	DB       *sqlx.DB
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	usrCore := user.NewCore(cfg.Log, cfg.Delegate, usercache.NewStore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB)))
	prdCore := product.NewCore(cfg.Log, usrCore, cfg.Delegate, productdb.NewStore(cfg.Log, cfg.DB))

	authen := mid.Authenticate(cfg.Auth)
	ruleAny := mid.Authorize(cfg.Auth, auth.RuleAny)
	ruleUserOnly := mid.Authorize(cfg.Auth, auth.RuleUserOnly)
	ruleAuthorizeProduct := mid.AuthorizeProduct(cfg.Auth, prdCore)

	hdl := new(prdCore, usrCore)
	app.Handle(http.MethodGet, version, "/products", hdl.query, authen, ruleAny)
	app.Handle(http.MethodGet, version, "/products/{product_id}", hdl.queryByID, authen, ruleAuthorizeProduct)
	app.Handle(http.MethodPost, version, "/products", hdl.create, authen, ruleUserOnly)
	app.Handle(http.MethodPut, version, "/products/{product_id}", hdl.update, authen, ruleAuthorizeProduct)
	app.Handle(http.MethodDelete, version, "/products/{product_id}", hdl.delete, authen, ruleAuthorizeProduct)
}
