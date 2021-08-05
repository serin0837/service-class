// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"net/http"
	"os"

	"github.com/ardanlabs/service/business/web/mid"
	"github.com/ardanlabs/service/foundation/web"
	"go.uber.org/zap"
)

// APIMux constructs an http.Handler with all application routes defined.
func APIMux(build string, shutdown chan os.Signal, log *zap.SugaredLogger) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log))

	// Register debug check endpoints.
	cg := checkGroup{
		build: build,
		log:   log,
	}
	app.Handle(http.MethodGet, "/debug/readiness", cg.readiness)
	app.Handle(http.MethodGet, "/debug/liveness", cg.liveness)
	app.Handle(http.MethodGet, "/testerror", cg.testerror)

	return app
}
