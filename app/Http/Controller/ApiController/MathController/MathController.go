package MathController

import (
	"firstProject/app/Http/Controller"
	"firstProject/app/Http/Service"
	"firstProject/sproute"
	"net/http"
)

type MathController struct {
	Controller.Controller
}

func ClassMathController() MathController {
	var controller MathController
	return controller
}

func (it MathController) Sum(request *http.Request, params sproute.H) sproute.Res {
	mathService := Service.ClassMathService()
	mathService.SetX(1)
	mathService.SetY(2)
	mathService.SetZ(3)

	return sproute.ResponseJson(200, sproute.H{
		"message": mathService.Sum(),
	})
}
