package routers

import (
	"_/src/api/controller"
	"_/src/constants"
	"_/src/types/router_types"
)

const url string = constants.ApiPrefixV1 + "/posts"

type PostRouter struct {
	router router_types.TMux
}

func NewPostRouter(router router_types.IMux, controller controller.ControllerLayer) *PostRouter {
	return &PostRouter{router: router_types.TMux{Mux: router, Ctrlr: controller}}
}

func (pr *PostRouter) Register() {
	pr.router.Mux.Get(url, pr.router.Ctrlr.PostController.GetPosts)
	pr.router.Mux.Post(url, pr.router.Ctrlr.PostController.PostPost)
}
