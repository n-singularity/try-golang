package ProductController

import (
	"firstProject/app/Http/Controller"
	"firstProject/sproute"
	"net/http"
)

type ProductController struct {
	Controller.Controller
}

func ClassProductController() ProductController {
	var controller ProductController
	return controller
}


func (it ProductController) Index(request *http.Request, params sproute.H) sproute.Res {
	return sproute.ResponseJson(200, sproute.H{
		"message": "test",
	})
}
