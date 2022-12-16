package routers

import (
	"_/src/api/controller"
	"_/src/constants"
	"_/src/types/router_types"
)

const url string = constants.ApiPrefixV1 + "/posts"

type PostRouter struct {
	Mux router_types.TMux
}

func NewPostRouter(router router_types.IMux, controller controller.ControllerLayer) *PostRouter {
	return &PostRouter{Mux: router_types.TMux{Mux: router, Ctrlr: controller}}
}

func (pr *PostRouter) Register() {
	pr.Mux.Mux.Get(url, pr.Mux.Ctrlr.PostController.GetPosts)
	pr.Mux.Mux.Post(url, pr.Mux.Ctrlr.PostController.PostPost)
}
