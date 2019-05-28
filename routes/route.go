package routes

import "firstProject/sproute"

func Route(r sproute.Route) sproute.Route{
	r = ApiList(r)
	r = WebList(r)

	return r
}
