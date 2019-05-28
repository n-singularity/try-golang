package routes

import (
	"firstProject/app/Http/Middlewares"
	"firstProject/sproute"
	"net/http"
)

func ApiListV1(r sproute.Route) sproute.Route {
	v1:=r.GROUP("/api/v1").Middleware(middleware.FirstMiddleware{Params: sproute.H{"name": "ok"}})

	v1.GET("/:word", func(request *http.Request, params sproute.H) sproute.Res {
		return sproute.ResponseJson(200, sproute.H{"messages": params["word"]})
	}).Middleware(middleware.FirstMiddleware{Params: sproute.H{"name": "ok"}}).Middleware(middleware.AnotherMiddleware{})

	return r
}
