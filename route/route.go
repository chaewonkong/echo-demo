package route

import (
	"net/http"

	"github.com/chaewonkong/echo-demo/handler"
	"github.com/labstack/echo/v4"
)

// Route represents a single route in the application
// associated with a specific HTTP method, path, handler function, and optional middlewares.
type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

// RouteGroup represents a group of routes with a common prefix.
type RouteGroup struct {
	Prefix string
	Routes []Route
}

// RouteGroupFunc is a function type that creates a RouteGroup.
type RouteGroupFunc func(handler handler.Handler) RouteGroup

// UserRoutes returns a RouteGroup for user-related routes.
func UserRoutes(handler handler.Handler) RouteGroup {
	routes := []Route{
		{Method: http.MethodPost, Path: "/", Handler: handler.CreateUser},
		{Method: http.MethodDelete, Path: ":id", Handler: handler.DeleteUser},
	}
	return RouteGroup{
		Prefix: "/users",
		Routes: routes,
	}
}

// PostRoutes returns a RouteGroup for post-related routes.
func PostRoutes(handler handler.Handler) RouteGroup {
	routes := []Route{
		{Method: http.MethodPost, Path: "/", Handler: handler.CreatePost},
		{Method: http.MethodDelete, Path: ":id", Handler: handler.DeletePost},
	}
	return RouteGroup{
		Prefix: "/posts",
		Routes: routes,
	}
}
