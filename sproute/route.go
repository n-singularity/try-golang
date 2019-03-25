package sproute

import (
	"encoding/json"
	"net/http"
	"strings"
)

//global variable
var RouteList = make(map[string]RouteStruck)

//struck
type Route struct{}

type RouteGroup struct {
	RouteList []RouteStruck
	Prefix    string
}

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

func (it H) ToString() string {
	js, _ := json.Marshal(it)
	return string(js)
}

func (it H) Get(key string) string {
	if it[key] != nil {
		return it[key].(string)
	}

	return ""
}

type ControllerFunc func(*http.Request, H) Res

func Build() Route { return Route{} }

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

func AddNodeGroup(prefix string, method string, path string, controller ControllerFunc) RouteStruck {

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

func (it *Route) GET(p string, controller ControllerFunc) RouteStruck {
	return AddNode("GET", p, controller)
}

func (it Route) Listen(port string) {
	http.HandleFunc("/", it.Handler)
	http.ListenAndServe(port, nil)
}

func CompareUrl(routePath string, urlPath string) (bool, H) {
	routeComponents := strings.Split(routePath, "/")
	urlComponents := strings.Split(urlPath, "/")
	params := H{}

	if len(routeComponents) != len(urlComponents) {
		return false, params
	}

	for key, component := range routeComponents {
		if len(component) > 0 && component[0] == ':' { // check if it is a named param.
			params[component[1:len(component)]] = urlComponents[key]
		} else if component != urlComponents[key] {
			return false, params
		}
	}

	return true, params
}
