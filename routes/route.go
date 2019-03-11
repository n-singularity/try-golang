package routes

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) *gin.Engine {

	r = ApiList(r)
	r = WebList(r)

	return r
}
