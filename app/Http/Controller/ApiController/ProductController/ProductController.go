package ProductController

import (
	"firstProject/app/Http/Controller"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Controller.Controller
}

func ClassProductController() ProductController {
	var controller ProductController
	return controller
}


func (thisIs ProductController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "product",
	})
}
