package Controllers

import "github.com/gin-gonic/gin"

func IndexWeb1(c *gin.Context) {
	c.HTML(200, "ping/index.tmpl",gin.H{})
}

func IndexWeb2(c *gin.Context) {
	c.HTML(200, "pong/index.tmpl",gin.H{})
}

func IndexApi(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api pong",
	})
}
