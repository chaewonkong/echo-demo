# Echo/v4 Demonstration

A clean and extensible demonstration of routing architecture using [echo/v4](https://github.com/labstack/echo) for Go.

This repository showcases a **declarative, domain-driven routing design** intended for scalable and maintainable web applications.

---

## ✨ Highlights

### 📦 Declarative Routing

Routes are defined as data, not code. Instead of calling `e.GET()` or `e.POST()` directly for every route, all route information is **declared in structured types** and registered centrally.

```go
type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}
```

Each domain exposes a function returning a RouteGroup:

```go
type RouteGroup struct {
	Prefix string
	Routes []Route
}
```

This allows the routing logic to stay simple, testable, and easy to maintain.

### 🧩 Domain-Based Modularity

Each logical module (domain) defines its own routes in isolation, using a consistent and pluggable structure:

```go
type RouteGroupFunc func(handler.Handler) RouteGroup
```

In main.go (or wherever routes are wired), all domain modules can be registered like this:

```go
router.RegisterRoutes(
    TicketRoutes,
    BackfillRoutes,
)
```

### ⚙ Centralized Registration

All routes are ultimately registered via a central Router that encapsulates the Echo instance and handler dependencies:

```go
r.e.Add(route.Method, routeGroup.Prefix+route.Path, route.Handler, route.Middlewares...)
```

## 📁 Project Structure

```text
.
├── handler/            # Business logic handlers
├── route/              # Route definitions per domain
├── router/             # Router orchestrator
├── main.go             # Entry point
```

## ✅ Benefits

- Separation of concerns between handler logic and routing structure
- Easily testable route declarations
- Declarative and extendable design
- Minimal boilerplate when adding new endpoints
- Ready for middleware injection, versioning, and Swagger integration

## 🔧 Requirements
- Go 1.18+
- Echo v4
