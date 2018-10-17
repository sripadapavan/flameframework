package routes

import (
	"reflect"
	. "github.com/flame/controllers"
)

type Route struct {
	Route_type string
	Controller BaseController
	Url string
	MiddlewareList []reflect.Type
}

func Get() *Route {
	route := new(Route)
	route.Route_type = "GET"
	return route
}

func Post() *Route {
	route := new(Route)
	route.Route_type = "POST"
	return route
}

func (r Route) Define(url string, controller BaseController) Route {
	r.Controller = controller
	r.Url = url
	return r
}

func (r Route) Middleware(middleware ...func()(reflect.Type)) Route {
	for _, ty := range middleware {
		midd := ty()
		r.MiddlewareList = append(r.MiddlewareList, midd)
	}
	return r
}