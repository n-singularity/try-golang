package sproute

//global variable
var RouteList = make(map[string]RouteStruck)

func AddNode(method string, path string, controller ControllerFunc) RouteStruck {

	if path[0] != '/' {
		panic("Path has to start with a /.")
	}
	route := RouteStruck{}

	route.method = method
	route.path = path
	route.controller = controller

	RouteList[path] = route

	return route
}

func UpdateNode(route RouteStruck) RouteStruck {
	RouteList[route.path] = route
	return route
}

//common route
type Route struct{}

func (it *Route) GET(p string, controller ControllerFunc) RouteStruck {
	return AddNode("GET", p, controller)
}

func (it *Route) GROUP(p string) RouteGroup {
	return RouteGroup{prefix:p}
}

//common route
type RouteGroup struct {
	middleware []MiddlewareInterface
	prefix     string
}

func (it RouteGroup) Middleware(middleware MiddlewareInterface) RouteGroup {
	it.middleware = append(it.middleware, middleware)
	return it
}

func (it *RouteGroup) GET(p string, controller ControllerFunc) {
	r := AddNode("GET", it.prefix+p, controller)
	for _, mdw := range it.middleware {r.Middleware(mdw)}
}
