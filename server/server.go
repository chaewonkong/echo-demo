package server

import (
	"log/slog"
	"os"

	"github.com/chaewonkong/echo-demo/handler"
	"github.com/chaewonkong/echo-demo/route"
	"github.com/chaewonkong/echo-demo/router"
	"github.com/labstack/echo/v4"
)

const (
	exitZero = iota
	exitNonZero
)

// Run initializes and runs echo server application
func Run() int {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	port := os.Getenv("PORT")
	if port == "" {
		logger.Error("port not set")
		return exitNonZero
	}

	e := echo.New()
	h := handler.Handler{}
	r := router.NewRouter(e, h)
	r.RegisterRoutes(
		route.UserRoutes,
		route.PostRoutes,
	)

	err := e.Start(":" + port)
	if err != nil {
		logger.Error("failed to start server", "error", err)
		return exitNonZero
	}

	return exitZero
}
