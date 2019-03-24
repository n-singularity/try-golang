package routes

import (
	"firstProject/app/Http/Controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func WebList(r *gin.Engine) *gin.Engine {
	r.GET("/",func(c *gin.Context){
		fmt.Printf("aasdasfas")
		c.String(200, "Index Page")
	})


	r.GET("/ping1", Controller.IndexWeb1)
	r.GET("/ping2", Controller.IndexWeb2)

	return r
}
