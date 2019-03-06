package routes

import (
	"firstProject/app/Http/Controllers"
	"github.com/gin-gonic/gin"
)

func ApiList(r *gin.Engine) *gin.Engine {

	api := r.Group("/api")

	api.GET("/ping", Controllers.IndexApi)

	return r
}
