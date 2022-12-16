package routers

import (
	"_/src/api/controller"
	"_/src/constants"
	"_/src/types/router_types"
)

const url string = constants.ApiPrefixV1 + "/posts"

type PostRouter struct {
	mux router_types.TMux
}

func NewPostRouter(router router_types.IMux, controller controller.ControllerLayer) *PostRouter {
	return &PostRouter{mux: router_types.TMux{Mux: router, Ctrlr: controller}}
}

func (pr *PostRouter) Register() {
	pr.mux.Mux.Get(url, pr.mux.Ctrlr.PostController.GetPosts)
	pr.mux.Mux.Post(url, pr.mux.Ctrlr.PostController.PostPost)
}
