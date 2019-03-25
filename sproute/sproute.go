package sproute

import (
	"encoding/json"
	"net/http"
)
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

func (it Route) Listen(port string) {
	http.HandleFunc("/", it.Handler)
	http.ListenAndServe(port, nil)
}
