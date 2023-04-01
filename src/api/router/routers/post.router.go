package routers

import (
	"_/src/api/controller"
	"_/src/constants"
	"_/src/types/router_types"
)

const url string = constants.ApiPrefixV1 + "/posts"

type PostRouter struct {
	Router router_types.TMux
}

func NewPostRouter(router router_types.IMux, controller controller.ControllerLayer) *PostRouter {
	return &PostRouter{Router: router_types.TMux{Mux: router, Controller: controller}}
}

func (pr PostRouter) Register() {
	pr.Router.Mux.Get(url, pr.Router.Controller.PostController.GetPosts)
	pr.Router.Mux.Post(url, pr.Router.Controller.PostController.PostPost)
}
