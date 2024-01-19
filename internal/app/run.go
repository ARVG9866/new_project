package app

import "log/slog"

func (a *App) Run(logger *slog.Logger) error {
	logger.Info("HTTP server is running on Port:", a.appConfig.App.PortHTTP)
	return nil
}
