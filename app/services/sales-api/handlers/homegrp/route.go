package homegrp

import (
	"net/http"

	"github.com/shohinsan/SaleSphereAPI/business/core/crud/delegate"
	"github.com/shohinsan/SaleSphereAPI/business/core/crud/home"
	"github.com/shohinsan/SaleSphereAPI/business/core/crud/home/stores/homedb"
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
	hmeCore := home.NewCore(cfg.Log, usrCore, cfg.Delegate, homedb.NewStore(cfg.Log, cfg.DB))

	authen := mid.Authenticate(cfg.Auth)
	ruleAny := mid.Authorize(cfg.Auth, auth.RuleAny)
	ruleUserOnly := mid.Authorize(cfg.Auth, auth.RuleUserOnly)
	ruleAuthorizeHome := mid.AuthorizeHome(cfg.Auth, hmeCore)

	hdl := new(hmeCore)
	app.Handle(http.MethodGet, version, "/homes", hdl.query, authen, ruleAny)
	app.Handle(http.MethodGet, version, "/homes/{home_id}", hdl.queryByID, authen, ruleAuthorizeHome)
	app.Handle(http.MethodPost, version, "/homes", hdl.create, authen, ruleUserOnly)
	app.Handle(http.MethodPut, version, "/homes/{home_id}", hdl.update, authen, ruleAuthorizeHome)
	app.Handle(http.MethodDelete, version, "/homes/{home_id}", hdl.delete, authen, ruleAuthorizeHome)
}
