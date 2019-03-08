package routes

import (
	"firstProject/app/Http/Controller"
	"github.com/gin-gonic/gin"
)

func ApiList(r *gin.Engine) *gin.Engine {

	api := r.Group("/api")

	api.GET("/ping", Controller.ClassApiController().Index)

	return r
}
