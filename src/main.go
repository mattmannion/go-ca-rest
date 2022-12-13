package main

import (
	"_/src/api/controller"
	"_/src/api/repo"
	"_/src/api/router"
	"_/src/api/router/gmux"
	"_/src/api/service"
	"_/src/envs"
)

var (
	RepoLayer  = repo.NewRepoLayer()
	Services   = service.NewServiceLayer(service.Deps{RepoLayer: *RepoLayer})
	Contollers = controller.NewControllerLayer(controller.Deps{ServiceLayer: *Services})
	Router     = router.NewRouterLayer(gmux.NewMuxRouter(), *Contollers, envs.Cfg)
)

func main() {
	Router.Serve()
}
