package routes

import (
	"firstProject/app/Http/Controller"
	"firstProject/sproute"
)

func WebList(r sproute.Route) sproute.Route{
	r.GET("/:word", Controller.IndexWeb1)

	return r
}
