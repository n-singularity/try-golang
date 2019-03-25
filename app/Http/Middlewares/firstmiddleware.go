package middleware

import (
	"firstProject/sproute"
	"net/http"
)

type FirstMiddleware struct {
	Params sproute.H
	sproute.Middleware
}

func (m FirstMiddleware) Next(req *http.Request, rp sproute.H) sproute.Res {
	println("1111111")

	res := m.NextMiddleware.Next(req, rp)

	return res
}

func (m FirstMiddleware) SetNext(NextMiddleware sproute.MiddlewareInterface) sproute.MiddlewareInterface {
	m.NextMiddleware = NextMiddleware
	return m
}
