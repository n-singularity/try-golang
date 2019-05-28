package routes

import "firstProject/sproute"

func Route(r sproute.Route) sproute.Route{
	r = ApiListV1(r)
	r = WebList(r)

	return r
}
