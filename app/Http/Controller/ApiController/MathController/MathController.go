package MathController

import (
	"firstProject/app/Http/Controller"
	"firstProject/app/Http/Service"
	"github.com/gin-gonic/gin"
)

type MathController struct {
	Controller.Controller
}

func ClassMathController() MathController {
	var controller MathController
	return controller
}


func (it MathController) Sum(c *gin.Context) {
	mathService := Service.ClassMathService()
	mathService.SetX(1)
	mathService.SetY(2)
	mathService.SetZ(3)

	c.JSON(200, gin.H{
		"message": mathService.Sum(),
	})
}
