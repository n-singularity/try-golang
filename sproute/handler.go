package sproute

import (
	"fmt"
	"net/http"
	"strings"
)

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
