package sproute

import (
	"net/http"
)

type Middleware struct {
	NextMiddleware MiddlewareInterface
	Controller     ControllerFunc
}

func (m Middleware) Next(req *http.Request, rp H) Res {
	if m.Controller == nil{
		return m.NextMiddleware.Next(req, rp)
	}

	return m.Controller(req, rp)
}

func (m Middleware) SetController(Controller ControllerFunc) MiddlewareInterface {
	m.Controller = Controller
	return m
}

func (m Middleware) SetNext(NextMiddleware MiddlewareInterface) MiddlewareInterface {
	m.NextMiddleware = NextMiddleware
	return m
}
