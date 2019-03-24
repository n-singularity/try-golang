package ProductController

import (
	"firstProject/app/Http/Controller"
	"firstProject/sproute"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Controller.Controller
}

func ClassProductController() ProductController {
	var controller ProductController
	return controller
}


func (it ProductController) Index(c *gin.Context) {
	c.JSON(200, sproute.H{
		"message": 123,
	})
}
