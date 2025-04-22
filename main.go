package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/chaewonkong/echo-demo/handler"
	"github.com/chaewonkong/echo-demo/route"
	"github.com/chaewonkong/echo-demo/router"
	"github.com/labstack/echo/v4"
)

func main() {
	err := run()
	if err != nil {
		slog.Default().Error("failed to run server", "error", err)
		os.Exit(1)
	}
}

func run() error {
	// load env
	port := os.Getenv("PORT")
	if port == "" {
		return errors.New("port not set")
	}

	e := echo.New()
	h := handler.Handler{}
	r := router.NewRouter(e, h)
	r.RegisterRoutes(
		route.UserRoutes,
		route.PostRoutes,
	)

	return nil
}
