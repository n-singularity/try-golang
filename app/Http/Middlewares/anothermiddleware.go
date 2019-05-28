package middleware

import (
	"firstProject/sproute"
	"net/http"
)

type AnotherMiddleware struct {
	Params sproute.H
	sproute.Middleware
}

func (m AnotherMiddleware) Next(req *http.Request, rp sproute.H) sproute.Res {
	res := m.NextMiddleware.Next(req, rp)

	return res
}

func (m AnotherMiddleware) SetNext(NextMiddleware sproute.MiddlewareInterface) sproute.MiddlewareInterface {
	m.NextMiddleware = NextMiddleware
	return m
}

