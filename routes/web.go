package routes

import (
	"firstProject/app/Http/Controller"
	"github.com/gin-gonic/gin"
)

func WebList(r *gin.Engine) *gin.Engine {

	r.GET("/ping1", Controller.IndexWeb1)
	r.GET("/ping2", Controller.IndexWeb2)

	return r
}
