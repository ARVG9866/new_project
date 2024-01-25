package app

import (
	"net/http"
)

func (a *App) Run(handler http.Handler) error {
	a.server.logger.Info("HTTP server is running", "Port", a.appConfig.Port)

	return http.ListenAndServe(a.appConfig.Port, handler)
}
