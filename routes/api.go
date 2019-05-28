package routes

import (
	"firstProject/app/Http/Middlewares"
	"firstProject/sproute"
	"net/http"
)

func ApiList(r sproute.Route) sproute.Route {
	v1:=r.GROUP("/api/v1").Middleware(middleware.FirstMiddleware{Params: sproute.H{"name": "ok"}})

	v1.GET("/:word", func(request *http.Request, params sproute.H) sproute.Res {
		return sproute.ResponseString(200, params["word"])
	}).Middleware(middleware.FirstMiddleware{Params: sproute.H{"name": "ok"}}).Middleware(middleware.AnotherMiddleware{})

	return r
}
