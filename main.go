package main

import (
	"firstProject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	gin.LoadHTMLGlob("resources/templates/**/*")


	gin = routes.Route(gin)

	gin.Run() // listen and serve on 0.0.0.0:8080
}