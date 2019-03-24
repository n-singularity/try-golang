package routes

import (
	"firstProject/app/Http/Controller/ApiController/MathController"
	"firstProject/app/Http/Controller/ApiController/ProductController"
	"github.com/gin-gonic/gin"
)

func ApiList(r *gin.Engine) string {

	api := r.Group("/api")

	api.GET("/product", ProductController.ClassProductController().Index)

	api.GET("/math/sum", MathController.ClassMathController().Sum)

	return r
}
