package router_types

import (
	"_/src/api/controller"
	"net/http"
)

type IMux interface {
	Mux() http.Handler
	Get(url string, f func(resp http.ResponseWriter, req *http.Request))
	Post(url string, f func(resp http.ResponseWriter, req *http.Request))
}

type IRouter interface {
	Register()
}

type TMux struct {
	Mux   IMux
	Ctrlr controller.ControllerLayer
}
