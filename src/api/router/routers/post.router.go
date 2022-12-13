package routers

import (
	"_/src/api/controller"
	"_/src/constants"
	"_/src/types/router_types"
)

const url string = constants.ApiPrefixV1 + "/posts"

type PostRouter struct {
	router router_types.Router
}

func NewPostRouter(router router_types.IRouter, controller controller.ControllerLayer) *PostRouter {
	return &PostRouter{router: router_types.Router{Router: router, Ctrlr: controller}}
}

func (pr *PostRouter) Register() {
	pr.router.Router.Get(url, pr.router.Ctrlr.PostController.GetPosts)
	pr.router.Router.Post(url, pr.router.Ctrlr.PostController.PostPost)
}
