package checkgrp

import (
	"net/http"

	"github.com/ServiceWeaver/weaver"
	"github.com/jmoiron/sqlx"

	"github.com/vikaskumar1187/publisher_saasv2/services/publisher/foundation/logger"
	"github.com/vikaskumar1187/publisher_saasv2/services/publisher/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	UsingWeaver bool
	Build       string
	Log         *logger.Logger
	DB          *sqlx.DB
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	checkgrp := New(cfg.Build, cfg.Log, cfg.DB)
	app.HandleNoMiddleware(http.MethodGet, version, "/readiness", checkgrp.Readiness)
	app.HandleNoMiddleware(http.MethodGet, version, "/liveness", checkgrp.Liveness)

	if cfg.UsingWeaver {
		app.HandleNoMiddleware(http.MethodGet, "" /*group*/, weaver.HealthzURL, checkgrp.Readiness)
	}
}
