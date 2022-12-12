package routers

import (
	"_/src/api/controller"
	"_/src/constants"
	"_/src/types/router_types"
)

const prefix string = constants.ApiPrefixV1 + "/posts"

type PostRouter struct {
	router router_types.Router
}

func NewPostRouter(router router_types.IRouter, controller controller.ControllerLayer) *PostRouter {
	return &PostRouter{router: router_types.Router{Router: router, Ctrlr: controller}}
}

func (pr *PostRouter) GetPostsRoute() {
	pr.router.Router.Get(prefix, pr.router.Ctrlr.PostController.GetPosts)
}

func (pr *PostRouter) PostPostRouter() {
	pr.router.Router.Post(prefix, pr.router.Ctrlr.PostController.PostPost)
}
