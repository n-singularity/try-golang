package routes

import (
	"firstProject/app/Http/Controllers"
	"github.com/gin-gonic/gin"
)

func WebList(r *gin.Engine) *gin.Engine {

	r.GET("/ping1", Controllers.IndexWeb1)
	r.GET("/ping2", Controllers.IndexWeb2)

	return r
}
