package Controller

import (
	"github.com/gin-gonic/gin"
)

type ApiController struct {
	Controller
}

func ClassApiController() ApiController {
	extend := Controller{}
	return ApiController{extend}
}


func (thisIs ApiController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api pong",
	})
}
