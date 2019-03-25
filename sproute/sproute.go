package sproute

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Route struct{}

type ControllerFunc func(*http.Request, H) Res

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

type MethodHandler func(http.Request, ControllerFunc)

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

var RouteList = make(map[string]RouteStruck)

func Build() Route {
	return Route{}
}

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

func (it *Route) GET(p string, controller ControllerFunc) RouteStruck {
	return AddNode("GET", p, controller)
}

func (it Route) Handler(w http.ResponseWriter, r *http.Request) {
	found := false

	for _, value := range RouteList {
		status, params := CompareUrl(value.path, r.URL.Path)

		if !status {
			continue
		}

		if r.Method != value.method {
			continue
		}

		//Set Middleware
		controller := value.controller


		var middlewareStruck Middleware
		var middleware MiddlewareInterface
		middleware = middlewareStruck
		middleware = middleware.SetController(controller)

		for i := len(value.middleware); i > 0; i-- {
			middleware = value.middleware[i-1].SetNext(middleware)
		}

		res := middleware.Next(r, params)

		w.WriteHeader(res.code)
		fmt.Fprintf(w, res.getResponse())

		found = true
		break
	}

	if found == false {
		w.WriteHeader(404)
		fmt.Fprintf(w, "page not found")
	}
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
