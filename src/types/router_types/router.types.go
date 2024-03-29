package router_types

import (
	"_/src/api/controller"
	"net/http"
)

type IRoute = interface{}

type IMux interface {
	Mux() http.Handler
	Get(url string, f http.HandlerFunc) IRoute
	Post(url string, f http.HandlerFunc) IRoute
	Put(url string, f http.HandlerFunc) IRoute
	Patch(url string, f http.HandlerFunc) IRoute
	Delete(url string, f http.HandlerFunc) IRoute
}

type IRouter interface {
	Register()
}

type TMux struct {
	Mux        IMux
	Controller controller.ControllerLayer
}
