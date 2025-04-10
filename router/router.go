package router

import (
	"github.com/chaewonkong/echo-demo/handler"
	"github.com/chaewonkong/echo-demo/route"
	"github.com/labstack/echo/v4"
)

// Router registers routes with the echo instance
type Router struct {
	e       *echo.Echo
	handler handler.Handler
}

// NewRouter creates a new Router instance
func NewRouter(e *echo.Echo, handler handler.Handler) Router {
	return Router{
		e:       e,
		handler: handler,
	}
}

// RegisterRoutes registers the routes with the echo instance
func (r Router) RegisterRoutes(routeGroupFuncs ...route.RouteGroupFunc) {
	for _, f := range routeGroupFuncs {
		routeGroup := f(r.handler)
		for _, route := range routeGroup.Routes {
			r.e.Add(route.Method, routeGroup.Prefix+route.Path, route.Handler, route.Middlewares...)
		}
	}
}
