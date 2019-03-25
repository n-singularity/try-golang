package sproute

import "net/http"

type RouteStruck struct {
	method     string
	path       string
	controller ControllerFunc
	middleware []MiddlewareInterface
}

type MiddlewareStruck struct {
	MiddlewareInterface
	method string
	path   string
	next   ControllerFunc
}

type MiddlewareInterface interface {
	Next(req *http.Request, rp H) Res
	SetController(controllerFunc ControllerFunc) MiddlewareInterface
	SetNext(MiddlewareInterface) MiddlewareInterface
}

func (it RouteStruck) Middleware(middleware MiddlewareInterface) RouteStruck {
	it.middleware = append(it.middleware, middleware)
	return UpdateNode(it)
}
