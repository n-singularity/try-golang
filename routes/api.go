package routes

import (
	"firstProject/app/Http/Controller"
	"firstProject/app/Http/Controller/ApiController/MathController"
	"firstProject/app/Http/Middlewares"
	"firstProject/sproute"
	"net/http"
)

func ApiListV1(r sproute.Route) sproute.Route {
	v1:=r.GROUP("/api/v1").Middleware(middleware.FirstMiddleware{Params: sproute.H{"name": "ok"}})

	v1.GET("/word/:word", func(request *http.Request, params sproute.H) sproute.Res {
		return sproute.ResponseJson(200, sproute.H{"messages": params["word"]})
	}).Middleware(middleware.FirstMiddleware{Params: sproute.H{"name": "ok"}}).Middleware(middleware.AnotherMiddleware{})

	v1.GET("/match/sum", MathController.ClassMathController().Sum)

	v1.GET("/curl", Controller.ClassController().CurlGet)

	v1.POST("/curl", Controller.ClassController().CurlPost)

	v1.POST("/add-line-text", Controller.ClassController().AddValueInFileText)

	return r
}
