package main

import (
	"_/src/api/controller"
	"_/src/api/repo"
	"_/src/api/router"
	"_/src/api/router/gin_mux"
	"_/src/api/service"
	"_/src/envs"
)

var (
	RepoLayer  = repo.NewRepoLayer()
	Services   = service.NewServiceLayer(service.Deps{RepoLayer: *RepoLayer})
	Contollers = controller.NewControllerLayer(controller.Deps{ServiceLayer: *Services})
	Router     = router.NewRouterLayer(gin_mux.NewMuxRouter(), *Contollers, envs.Cfg)
)

func main() {
	Router.ServeRestApi()
}
