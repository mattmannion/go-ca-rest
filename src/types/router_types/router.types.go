package router_types

import (
	"_/src/api/controller"
	"net/http"
)

type IRouter interface {
	Get(url string, f func(resp http.ResponseWriter, req *http.Request))
	Post(url string, f func(resp http.ResponseWriter, req *http.Request))
}

type Router struct {
	Router IRouter
	Ctrlr  controller.ControllerLayer
}
