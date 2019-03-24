package middleware

import (
	"firstProject/sproute"
	"net/http"
)

type ExampleMiddleware struct {
	Params sproute.H
}

func (m ExampleMiddleware) Build(params sproute.H) ExampleMiddleware {
	m.Params = params
	return m
}

func (m ExampleMiddleware) Before(req *http.Request, rp sproute.H)  {

}

func (m ExampleMiddleware) After(res sproute.Res, rp sproute.H)  {

}
