package sproute

import "net/http"

type RouteStruck struct {
	method     string
	path       string
	controller ControllerFunc
	middleware []MiddlewareFunc
}

type MiddlewareFunc interface {
	Before(req *http.Request, rp H)
	After(res Res, rp H)
}

func (r RouteStruck) Middleware(middleware MiddlewareFunc) RouteStruck{
	r.middleware = append(r.middleware, middleware)
	return UpdateNode(r)
}