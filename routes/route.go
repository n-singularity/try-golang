package routes

func Route(method MethodHandler, p string, controller ControllerFunc) string {

	method(request, controller)
	return  "aaaaaa"
}
