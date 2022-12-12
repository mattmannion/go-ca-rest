package router

import (
	"_/src/api/controller"
	"_/src/api/router/routers"
	"_/src/types/router_types"
)

type RouterLayer struct {
	PostRouter routers.PostRouter
}

func NewRouterLayer(host string, port string, router router_types.IRouter, ctrlr controller.ControllerLayer) *RouterLayer {
	return &RouterLayer{PostRouter: *routers.NewPostRouter(router, ctrlr)}
}

func (rl *RouterLayer) Serve() {

}
