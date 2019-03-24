package Controller

import (
	"firstProject/sproute"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func IndexWeb1(request *http.Request, params sproute.H) sproute.Res {
	return sproute.ResponseJson(200, sproute.H{
		"message": params.Get("word"),
	})
}

func IndexWeb2(c *gin.Context) {
	c.HTML(200, "pong/index.tmpl", gin.H{})
}
