package main

import (
	"_/src/api/controller"
	"_/src/api/repo"
	"_/src/api/router"
	"_/src/api/router/muxes"
	"_/src/api/service"
	"_/src/clients/http_client"
	"_/src/clients/pg"
	"_/src/constants"
	"_/src/envs"
)

func main() {
	envs.InitCfg(constants.Env)

	pg.InitDb()

	RepoLayer := repo.NewRepoLayer()
	Services := service.NewServiceLayer(service.Deps{RepoLayer: *RepoLayer})
	Contollers := controller.NewControllerLayer(controller.Deps{ServiceLayer: *Services})
	Router := router.NewRouterLayer(router.Deps{
		Router:     muxes.NewGinMux(),
		Controller: *Contollers,
	})

	http_client.ListenAndServeClient(Router.CreateRestApi())
}
