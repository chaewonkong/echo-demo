package main

import (
	"github.com/chaewonkong/echo-demo/handler"
	"github.com/chaewonkong/echo-demo/route"
	"github.com/chaewonkong/echo-demo/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	h := handler.Handler{}
	r := router.NewRouter(e, h)
	r.RegisterRoutes(
		route.UserRoutes,
		route.PostRoutes,
	)
}
