package hackgrp

import (
	"net/http"

	"github.com/shohinsan/SaleSphereAPI/foundation/web"
)

// Routes adds specific routes for this group
// Routes adds specific routes for this group
func Routes(app *web.App) {
	app.Handle(http.MethodGet, "/hack", Hack)
}
