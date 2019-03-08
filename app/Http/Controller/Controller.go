package Controller

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {}


func IndexWeb1(c *gin.Context) {
	c.HTML(200, "ping/index.tmpl",gin.H{})
}

func IndexWeb2(c *gin.Context) {
	c.HTML(200, "pong/index.tmpl",gin.H{})
}
