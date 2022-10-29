package main

import (
	"_/src/api/controller"
	"_/src/api/repo"
	"_/src/api/router"
	"_/src/api/service"
	"_/src/constants"
)

var (
	RepoLayer  = repo.NewRepoLayer()
	Services   = service.NewServiceLayer(service.Deps{RepoLayer: *RepoLayer})
	Contollers = controller.NewControllerLayer(controller.Deps{ServiceLayer: *Services})
	Router     = router.NewMuxRouter()
)

func main() {
	Router.Get(constants.ApiPrefixV1+"/posts", Contollers.PostController.GetPosts)
	Router.Post(constants.ApiPrefixV1+"/posts", Contollers.PostController.PostPost)

	Router.Serve(":7890")
}
